package x11

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ReanGD/runify/server/system"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

const (
	dbgShowShortcut = true
)

var (
	zapInitKeyboard    = zap.String("Method", "x11Keyboard::onInit")
	zapStopKeyboard    = zap.String("Method", "x11Keyboard::onStop")
	zapBindKeyboard    = zap.String("Method", "x11Keyboard::bind")
	zapUnbindKeyboard  = zap.String("Method", "x11Keyboard::unbind")
	zapOnKeyRelease    = zap.String("Method", "x11Keyboard::onKeyRelease")
	zapOnMappingNotify = zap.String("Method", "x11Keyboard::onMappingNotify")
)

type bindID uint16

type bindKey struct {
	mods    uint16
	keycode xproto.Keycode
}

type bindData struct {
	id       bindID
	keys     []bindKey
	shortcut string
}

type x11Keyboard struct {
	ignoreMods     []uint16
	modIdByName    map[string]uint16
	modNameById    map[uint16]string
	keysymIdByName map[string]xproto.Keysym
	keysymNameById map[xproto.Keysym]string
	keymap         *xproto.GetKeyboardMappingReply
	modmap         *xproto.GetModifierMappingReply
	bindById       map[bindID]*bindData
	bindByKey      map[bindKey]*bindData
	errorCh        chan<- error
	shortcutCh     chan<- bindID
	connection     *xgb.Conn
	moduleLogger   *zap.Logger
	window         xproto.Window
	nextBindID     bindID
	minKeycode     xproto.Keycode
	maxKeycode     xproto.Keycode
}

func newX11Keyboard() *x11Keyboard {
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

	return &x11Keyboard{
		ignoreMods:     ignoreMods,
		modIdByName:    modIdByName,
		modNameById:    modNameById,
		keysymIdByName: gKeysyms,
		keysymNameById: keysymNameById,
		keymap:         nil,
		modmap:         nil,
		bindById:       make(map[bindID]*bindData, 8),
		bindByKey:      make(map[bindKey]*bindData, 8),
		errorCh:        nil,
		shortcutCh:     nil,
		connection:     nil,
		moduleLogger:   nil,
		window:         0,
		nextBindID:     1,
		minKeycode:     0,
		maxKeycode:     0,
	}
}

func (k *x11Keyboard) onInit(
	connection *xgb.Conn, window xproto.Window, errorCh chan<- error, shortcutCh chan<- bindID, moduleLogger *zap.Logger) error {

	setupInfo := xproto.Setup(connection)
	k.minKeycode = setupInfo.MinKeycode
	k.maxKeycode = setupInfo.MaxKeycode
	k.shortcutCh = shortcutCh
	k.connection = connection
	k.window = window
	k.moduleLogger = moduleLogger

	if !k.updateMaps(zapInitKeyboard) {
		return errInitX11Keyboard
	}

	return nil
}

func (k *x11Keyboard) onStart() {

}

func (k *x11Keyboard) onStop() {
	for key, bindData := range k.bindByKey {
		_ = k.ungrabKey(
			key, zapStopKeyboard, zap.String("Shortcut", bindData.shortcut), zap.Uint16("BindID", uint16(bindData.id)))
	}
	k.bindByKey = make(map[bindKey]*bindData)
	k.bindById = make(map[bindID]*bindData)
}

func (k *x11Keyboard) updateMaps(fields ...zap.Field) bool {
	var err error

	firstKeycode := k.minKeycode
	count := byte(k.maxKeycode - k.minKeycode + 1)
	if k.keymap, err = xproto.GetKeyboardMapping(k.connection, firstKeycode, count).Reply(); err != nil {
		k.moduleLogger.Error("Failed to get keyboard mapping", append(fields, zap.Error(err))...)
		return false
	}

	if k.modmap, err = xproto.GetModifierMapping(k.connection).Reply(); err != nil {
		k.moduleLogger.Error("Failed to get modifier mapping", append(fields, zap.Error(err))...)
		return false
	}

	return true
}

func (k *x11Keyboard) grabKey(key bindKey, fields ...zap.Field) system.Error {
	var err error
	for _, m := range k.ignoreMods {
		err = xproto.GrabKeyChecked(
			k.connection, true, k.window, key.mods|m, key.keycode, xproto.GrabModeAsync, xproto.GrabModeAsync).Check()
		if err != nil {
			switch err.(type) {
			case xproto.AccessError:
				accessErr := errors.New("keyboard shortcut is already taken by another application")
				k.moduleLogger.Info("Failed call x11 grab key", append(fields, zap.Error(accessErr))...)
				return system.ShortcutUsesByExternalApp
			default:
				k.moduleLogger.Info("Failed call x11 grab key", append(fields, zap.Error(err))...)
				return system.ShortcutBindError
			}
		}
	}

	return system.Success
}

func (k *x11Keyboard) ungrabKey(key bindKey, fields ...zap.Field) bool {
	for _, m := range k.ignoreMods {
		err := xproto.UngrabKeyChecked(k.connection, key.keycode, k.window, key.mods|m).Check()
		if err != nil {
			k.moduleLogger.Warn("Failed call x11 ungrab key", append(fields, zap.Error(err))...)
			return false
		}
	}

	return false
}

func (k *x11Keyboard) modsToStr(mods uint16, joinStr string) string {
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

func (k *x11Keyboard) keysymsByKeycode(keycode xproto.Keycode) []xproto.Keysym {
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

func (k *x11Keyboard) keycodeToStr(keycode xproto.Keycode, joinStr string) string {
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

func (k *x11Keyboard) keycodesByKeysym(keysym xproto.Keysym) []xproto.Keycode {
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

func (k *x11Keyboard) keycodesByKeysymStr(str string) []xproto.Keycode {
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

func (k *x11Keyboard) parseShortcut(shortcut string, fields ...zap.Field) ([]bindKey, system.Error) {
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

func (k *x11Keyboard) bindImpl(shortcut string, bindID bindID, fields ...zap.Field) system.Error {
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

func (k *x11Keyboard) bind(shortcut string) (bindID, system.Error) {
	if errCode := k.bindImpl(shortcut, k.nextBindID, zapBindKeyboard); errCode != system.Success {
		return 0, errCode
	}

	k.nextBindID++
	return k.nextBindID - 1, system.Success
}

func (k *x11Keyboard) unbind(id bindID) bool {
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

func (k *x11Keyboard) onKeyRelease(event xproto.KeyReleaseEvent) {
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
		k.shortcutCh <- bindData.id
	}
}

func (k *x11Keyboard) onMappingNotify(event xproto.MappingNotifyEvent) {
	if !k.updateMaps(zapOnMappingNotify) {
		k.errorCh <- errors.New("failed update keyboard mapping")
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
