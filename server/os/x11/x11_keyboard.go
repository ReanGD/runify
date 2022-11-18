package x11

import (
	"errors"

	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

var (
	zapInitKeyboard    = zap.String("Method", "x11.keyboard::init")
	zapStopKeyboard    = zap.String("Method", "x11.keyboard::stop")
	zapBindKeyboard    = zap.String("Method", "x11.keyboard::bind")
	zapUnbindKeyboard  = zap.String("Method", "x11.keyboard::unbind")
	zapOnKeyRelease    = zap.String("Method", "x11.keyboard::onKeyRelease")
	zapOnMappingNotify = zap.String("Method", "x11.keyboard::onMappingNotify")
)

type keyboard struct {
	ignoreMods       []uint16
	keymap           *xproto.GetKeyboardMappingReply
	modmap           *xproto.GetModifierMappingReply
	indexByX11Hotkey map[x11Hotkey]*bindData
	indexByHotkeyId  map[shortcut.HotkeyId]*bindData
	subscriptions    []chan<- *shortcut.Hotkey
	conn             *connection
	window           *window
	errorCtx         *module.ErrorCtx
	moduleLogger     *zap.Logger
	minKeycode       xproto.Keycode
	maxKeycode       xproto.Keycode
}

func newKeyboard() *keyboard {
	ignoreMods := []uint16{
		0,
		xproto.ModMaskLock,                   // Caps lock
		xproto.ModMask2,                      // Num lock
		xproto.ModMaskLock | xproto.ModMask2, // Caps and Num lock
	}

	return &keyboard{
		ignoreMods:       ignoreMods,
		keymap:           nil,
		modmap:           nil,
		indexByX11Hotkey: make(map[x11Hotkey]*bindData, 8),
		indexByHotkeyId:  make(map[shortcut.HotkeyId]*bindData, 8),
		subscriptions:    make([]chan<- *shortcut.Hotkey, 0, 1),
		conn:             nil,
		window:           nil,
		errorCtx:         nil,
		moduleLogger:     nil,
		minKeycode:       0,
		maxKeycode:       0,
	}
}

func (k *keyboard) init(conn *connection, window *window, errorCtx *module.ErrorCtx, moduleLogger *zap.Logger) bool {
	k.conn = conn
	setupInfo := k.conn.getSetupInfo()
	k.minKeycode = setupInfo.MinKeycode
	k.maxKeycode = setupInfo.MaxKeycode
	k.window = window
	k.errorCtx = errorCtx
	k.moduleLogger = moduleLogger

	if !k.updateMaps(zapInitKeyboard) {
		k.moduleLogger.Error("Failed init x11 keyboard", zap.String("Reason", "failed update key and mod map"))
		return false
	}

	return true
}

func (k *keyboard) close() {
	for key, bindData := range k.indexByX11Hotkey {
		_ = k.ungrabKey(key, zapStopKeyboard, bindData.hotkey.ZapField())
	}
	k.indexByX11Hotkey = make(map[x11Hotkey]*bindData)
	k.indexByHotkeyId = make(map[shortcut.HotkeyId]*bindData)
	k.subscriptions = []chan<- *shortcut.Hotkey{}
}

func (k *keyboard) updateMaps(fields ...zap.Field) bool {
	var ok bool

	if k.keymap, ok = k.conn.getKeyboardMapping(fields...); !ok {
		return false
	}
	if k.modmap, ok = k.conn.getModifierMapping(fields...); !ok {
		return false
	}

	return true
}

func (k *keyboard) grabKey(key x11Hotkey, fields ...zap.Field) global.Error {
	for _, m := range k.ignoreMods {
		if errCode := k.window.grabKey(key.mods|m, key.keycode, fields...); errCode != global.Success {
			return errCode
		}
	}

	return global.Success
}

func (k *keyboard) ungrabKey(key x11Hotkey, fields ...zap.Field) bool {
	res := true
	for _, m := range k.ignoreMods {
		if !k.window.ungrabKey(key.mods|m, key.keycode, fields...) {
			res = false
		}
	}

	return res
}

func (k *keyboard) keycodesByKeysym(keysym xproto.Keysym) []xproto.Keycode {
	keycodes := make([]xproto.Keycode, 0)
	set := make(map[xproto.Keycode]bool, 0)

	minKc := int(k.minKeycode)
	keysymsPerKeycode := int(k.keymap.KeysymsPerKeycode)

	for kc := minKc; kc <= int(k.maxKeycode); kc++ {
		keycode := xproto.Keycode(kc)
		offset := (kc - minKc) * keysymsPerKeycode
		for ind := offset; ind < offset+keysymsPerKeycode; ind++ {
			if keysym == k.keymap.Keysyms[ind] && !set[keycode] {
				keycodes = append(keycodes, keycode)
				set[keycode] = true
			}
		}
	}
	return keycodes
}

func (k *keyboard) parseHotkey(hotkey *shortcut.Hotkey, fields ...zap.Field) ([]x11Hotkey, global.Error) {
	mods := uint16(0)
	sMods := hotkey.Mods()
	if sMods&shortcut.ModShift != 0 {
		mods |= xproto.ModMaskShift
		sMods &^= shortcut.ModShift
	}
	if sMods&shortcut.ModControl != 0 {
		mods |= xproto.ModMaskControl
		sMods &^= shortcut.ModControl
	}
	if sMods&shortcut.ModAlt != 0 {
		mods |= xproto.ModMask1
		sMods &^= shortcut.ModAlt
	}
	if sMods&shortcut.ModSuper != 0 {
		mods |= xproto.ModMask4
		sMods &^= shortcut.ModSuper
	}

	if sMods != 0 {
		k.moduleLogger.Warn("Unsupported hotkey modifier", append(fields, sMods.ZapField())...)
		return []x11Hotkey{}, global.HotkeyParseFailed
	}

	sKey := hotkey.Key()
	keysym, ok := gKeysyms[sKey]
	if !ok {
		k.moduleLogger.Warn("Unsupported hotkey key", append(fields, sKey.ZapField())...)
		return []x11Hotkey{}, global.HotkeyParseFailed
	}

	keycodes := k.keycodesByKeysym(keysym)
	res := make([]x11Hotkey, 0, len(keycodes))
	for _, keycode := range keycodes {
		res = append(res, x11Hotkey{mods: mods, keycode: keycode})
	}

	return res, global.Success
}

func (k *keyboard) subscribeToHotkeys(ch chan<- *shortcut.Hotkey) bool {
	k.subscriptions = append(k.subscriptions, ch)
	return true
}

func (k *keyboard) bind(hotkey *shortcut.Hotkey, fields ...zap.Field) global.Error {
	zapFields := append(fields, zapBindKeyboard, hotkey.ZapField())

	id := hotkey.Id()
	if _, ok := k.indexByHotkeyId[id]; ok {
		k.moduleLogger.Warn("Hotkey already bound", zapFields...)
		return global.HotkeyUsesByRunify
	}

	x11Hotkeys, errCode := k.parseHotkey(hotkey, zapFields...)
	if errCode != global.Success {
		return errCode
	}

	for _, x11Hotkey := range x11Hotkeys {
		if _, ok := k.indexByX11Hotkey[x11Hotkey]; ok {
			k.moduleLogger.Info("Hotkey already binds by runify", zapFields...)
			return global.HotkeyUsesByRunify
		}
	}

	for _, bindKey := range x11Hotkeys {
		if errCode := k.grabKey(bindKey, zapFields...); errCode != global.Success {
			return errCode
		}
	}

	bindData := &bindData{id: id, x11Hotkeys: x11Hotkeys, hotkey: hotkey}
	k.indexByHotkeyId[id] = bindData
	for _, x11Hotkey := range x11Hotkeys {
		k.indexByX11Hotkey[x11Hotkey] = bindData
	}

	return global.Success
}

func (k *keyboard) unbind(hotkey *shortcut.Hotkey) bool {
	zapFields := []zap.Field{zapUnbindKeyboard, hotkey.ZapField()}

	id := hotkey.Id()
	if bindData, ok := k.indexByHotkeyId[id]; ok {
		delete(k.indexByHotkeyId, id)
		res := true
		for _, key := range bindData.x11Hotkeys {
			delete(k.indexByX11Hotkey, key)
			if !k.ungrabKey(key, zapFields...) {
				res = false
			}
		}

		return res
	}

	return true
}

func (k *keyboard) onKeyRelease(event xproto.KeyReleaseEvent) {
	mods, keycode := event.State, event.Detail

	for _, m := range k.ignoreMods {
		mods &= ^m
	}
	mods &= ^uint16(xproto.ModMask5)

	x11Hotkey := x11Hotkey{
		mods:    mods,
		keycode: keycode,
	}
	if bindData, ok := k.indexByX11Hotkey[x11Hotkey]; ok {
		hotkey := bindData.hotkey
		k.moduleLogger.Debug("onKeyRelease", zapOnKeyRelease, hotkey.ZapField())

		for _, ch := range k.subscriptions {
			select {
			case ch <- hotkey:
			default:
				k.moduleLogger.Warn("Failed send hotkey data to subscription channel, channel is full",
					zapReadFinish, hotkey.ZapField())
			}
		}
	}
}

func (k *keyboard) onMappingNotify(event xproto.MappingNotifyEvent) {
	if !k.updateMaps(zapOnMappingNotify) {
		k.moduleLogger.Error("Failed update keyboard maps", zapOnMappingNotify)
		k.errorCtx.SendError(errors.New("failed update keyboard mapping"))
		return
	}
	if event.Request == xproto.MappingKeyboard {
		hotkeys := k.indexByHotkeyId
		for key, bindData := range k.indexByX11Hotkey {
			_ = k.ungrabKey(key, zapOnMappingNotify, bindData.hotkey.ZapField())
		}
		k.indexByX11Hotkey = make(map[x11Hotkey]*bindData, 8)
		k.indexByHotkeyId = make(map[shortcut.HotkeyId]*bindData, 8)
		for _, bindData := range hotkeys {
			_ = k.bind(bindData.hotkey, zapOnMappingNotify)
		}
	}
}
