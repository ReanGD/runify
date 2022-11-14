package x11

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ReanGD/runify/server/system"
	"github.com/ReanGD/runify/server/system/module"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

const (
	dbgShowShortcut = true
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
	ignoreMods     []uint16
	modIdByName    map[string]uint16
	modNameById    map[uint16]string
	keysymIdByName map[string]xproto.Keysym
	keysymNameById map[xproto.Keysym]string
	keymap         *xproto.GetKeyboardMappingReply
	modmap         *xproto.GetModifierMappingReply
	bindById       map[bindID]*bindData
	bindByKey      map[bindKey]*bindData
	conn           *connection
	window         *window
	errorCtx       *module.ErrorCtx
	moduleLogger   *zap.Logger
	nextBindID     bindID
	minKeycode     xproto.Keycode
	maxKeycode     xproto.Keycode
}

func newKeyboard() *keyboard {
	ignoreMods := []uint16{
		0,
		xproto.ModMaskLock,                   // Caps lock
		xproto.ModMask2,                      // Num lock
		xproto.ModMaskLock | xproto.ModMask2, // Caps and Num lock
	}

	modIdByName := map[string]uint16{
		"Shift":    xproto.ModMaskShift,
		"CapsLock": xproto.ModMaskLock,
		"Control":  xproto.ModMaskControl,
		"Alt":      xproto.ModMask1,
		"NumLock":  xproto.ModMask2,
		"Mod3":     xproto.ModMask3,
		"Super":    xproto.ModMask4,
		"Mod5":     xproto.ModMask5,
	}

	modNameById := make(map[uint16]string, len(modIdByName))
	for name, id := range modIdByName {
		modNameById[id] = name
	}

	keysymNameById := make(map[xproto.Keysym]string, len(gKeysyms))
	for name, id := range gKeysyms {
		keysymNameById[id] = name
	}

	return &keyboard{
		ignoreMods:     ignoreMods,
		modIdByName:    modIdByName,
		modNameById:    modNameById,
		keysymIdByName: gKeysyms,
		keysymNameById: keysymNameById,
		keymap:         nil,
		modmap:         nil,
		bindById:       make(map[bindID]*bindData, 8),
		bindByKey:      make(map[bindKey]*bindData, 8),
		conn:           nil,
		window:         nil,
		errorCtx:       nil,
		moduleLogger:   nil,
		nextBindID:     1,
		minKeycode:     0,
		maxKeycode:     0,
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
	for key, bindData := range k.bindByKey {
		_ = k.ungrabKey(
			key, zapStopKeyboard, zap.String("Shortcut", bindData.shortcut), zap.Uint16("BindID", uint16(bindData.id)))
	}
	k.bindByKey = make(map[bindKey]*bindData)
	k.bindById = make(map[bindID]*bindData)
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

func (k *keyboard) grabKey(key bindKey, fields ...zap.Field) system.Error {
	for _, m := range k.ignoreMods {
		if errCode := k.window.grabKey(key.mods|m, key.keycode, fields...); errCode != system.Success {
			return errCode
		}
	}

	return system.Success
}

func (k *keyboard) ungrabKey(key bindKey, fields ...zap.Field) bool {
	res := true
	for _, m := range k.ignoreMods {
		if !k.window.ungrabKey(key.mods|m, key.keycode, fields...) {
			res = false
		}
	}

	return res
}

func (k *keyboard) modsToStr(mods uint16, joinStr string) string {
	res := ""
	for mask, name := range k.modNameById {
		if mods&uint16(mask) != 0 {
			if res != "" {
				res += joinStr
			}
			res += name
			mods &= ^uint16(mask)
		}
	}

	if mods != 0 {
		if res != "" {
			res += joinStr
		}
		res += fmt.Sprintf("0x%04X", mods)
	}

	return res
}

func (k *keyboard) keysymsByKeycode(keycode xproto.Keycode) []xproto.Keysym {
	if keycode < k.minKeycode || keycode > k.maxKeycode {
		k.moduleLogger.Warn("Invalid keycode",
			zap.Uint8("Keycode", uint8(keycode)),
			zap.Uint8("MinKeycode", uint8(k.minKeycode)),
			zap.Uint8("MaxKeycode", uint8(k.maxKeycode)))
		return []xproto.Keysym{}
	}

	cnt := int(k.keymap.KeysymsPerKeycode)
	res := make([]xproto.Keysym, 0, cnt)
	offset := int(keycode-k.minKeycode) * cnt
	for column := 0; column != cnt; column++ {
		unique := true
		keysym := k.keymap.Keysyms[offset+column]
		for i := 0; i < len(res); i++ {
			if keysym == res[i] {
				unique = false
				break
			}
		}
		if unique {
			res = append(res, keysym)
		}
	}

	return res
}

func (k *keyboard) keycodeToStr(keycode xproto.Keycode, joinStr string) string {
	keysyms := k.keysymsByKeycode(keycode)
	if len(keysyms) == 0 {
		return ""
	}

	res := ""
	for _, keysym := range keysyms {
		if res != "" {
			res += joinStr
		}
		if name, ok := k.keysymNameById[keysym]; ok {
			res += name
		} else {
			res += fmt.Sprintf("0x%08X", keysym)
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

func (k *keyboard) keycodesByKeysymStr(str string) []xproto.Keycode {
	keysym, ok := k.keysymIdByName[str]
	if !ok {
		keysym, ok = k.keysymIdByName[strings.Title(str)]
	}
	if !ok {
		keysym, ok = k.keysymIdByName[strings.ToLower(str)]
	}
	if !ok {
		keysym, ok = k.keysymIdByName[strings.ToUpper(str)]
	}
	if !ok {
		return []xproto.Keycode{}
	}

	return k.keycodesByKeysym(keysym)
}

func (k *keyboard) parseShortcut(shortcut string, fields ...zap.Field) ([]bindKey, system.Error) {
	mods := uint16(0)
	keycodes := []xproto.Keycode{}

	for _, part := range strings.Split(shortcut, "+") {
		switch strings.ToLower(part) {
		case "shift":
			mods |= xproto.ModMaskShift
		case "control":
			mods |= xproto.ModMaskControl
		case "alt":
			mods |= xproto.ModMask1
		case "super":
			mods |= xproto.ModMask4
		default:
			if len(keycodes) == 0 {
				keycodes = k.keycodesByKeysymStr(part)
			}
		}
	}

	if len(keycodes) == 0 {
		k.moduleLogger.Info("Failed parse shortcut", fields...)
		return []bindKey{}, system.ShortcutParseFailed
	}

	res := make([]bindKey, 0, len(keycodes))
	for _, keycode := range keycodes {
		res = append(res, bindKey{mods: mods, keycode: keycode})
	}

	return res, system.Success
}

func (k *keyboard) bindImpl(shortcut string, bindID bindID, fields ...zap.Field) system.Error {
	fields = append(fields, zap.String("Shortcut", shortcut))

	bindKeys, errCode := k.parseShortcut(shortcut, fields...)
	if errCode != system.Success {
		return errCode
	}

	for _, bindKey := range bindKeys {
		if _, ok := k.bindByKey[bindKey]; ok {
			k.moduleLogger.Info("Shortcut already binds by runify", fields...)
			return system.ShortcutUsesByRunify
		}
	}

	bindData := &bindData{id: bindID, keys: bindKeys, shortcut: shortcut}

	for _, bindKey := range bindKeys {
		if errCode := k.grabKey(bindKey, fields...); errCode != system.Success {
			return errCode
		}
	}

	k.bindById[bindID] = bindData
	for _, bindKey := range bindKeys {
		k.bindByKey[bindKey] = bindData
	}

	return system.Success
}

func (k *keyboard) bind(shortcut string) (bindID, system.Error) {
	if errCode := k.bindImpl(shortcut, k.nextBindID, zapBindKeyboard); errCode != system.Success {
		return 0, errCode
	}

	k.nextBindID++
	return k.nextBindID - 1, system.Success
}

func (k *keyboard) unbind(id bindID) bool {
	if bindData, ok := k.bindById[id]; ok {
		zapFields := []zap.Field{
			zapUnbindKeyboard,
			zap.String("Shortcut", bindData.shortcut),
			zap.Uint16("BindID", uint16(bindData.id)),
		}

		delete(k.bindById, id)
		res := true
		for _, key := range bindData.keys {
			delete(k.bindByKey, key)
			if !k.ungrabKey(key, zapFields...) {
				res = false
			}
		}

		return res
	}

	k.moduleLogger.Info("Failed unbind shortcut",
		zapUnbindKeyboard, zap.Uint16("BindID", uint16(id)), zap.Error(errors.New("bind not found")))
	return false
}

func (k *keyboard) onKeyRelease(event xproto.KeyReleaseEvent) {
	mods, keycode := event.State, event.Detail

	modsStr := k.modsToStr(mods, "+")
	keycodeStr := k.keycodeToStr(keycode, "|")
	shortcut := fmt.Sprintf("%s [%s]", modsStr, keycodeStr)
	k.moduleLogger.Debug("onKeyRelease", zapOnKeyRelease, zap.String("Shortcut", shortcut))

	for _, m := range k.ignoreMods {
		mods &= ^m
	}
	mods &= ^uint16(xproto.ModMask5)

	bindKey := bindKey{
		mods:    mods,
		keycode: keycode,
	}
	if bindData, ok := k.bindByKey[bindKey]; ok {
		// TODO: send bindData.id
		_ = bindData
	}
}

func (k *keyboard) onMappingNotify(event xproto.MappingNotifyEvent) {
	if !k.updateMaps(zapOnMappingNotify) {
		k.moduleLogger.Error("Failed update keyboard maps", zapOnMappingNotify)
		k.errorCtx.SendError(errors.New("failed update keyboard mapping"))
		return
	}
	if event.Request == xproto.MappingKeyboard {
		bindById := k.bindById
		for key, bindData := range k.bindByKey {
			_ = k.ungrabKey(
				key, zapOnMappingNotify, zap.String("Shortcut", bindData.shortcut), zap.Uint16("BindID", uint16(bindData.id)))
		}
		k.bindByKey = make(map[bindKey]*bindData, 8)
		k.bindById = make(map[bindID]*bindData, 8)
		for id, bindData := range bindById {
			_ = k.bindImpl(bindData.shortcut, id, zapOnMappingNotify)
		}
	}
}
