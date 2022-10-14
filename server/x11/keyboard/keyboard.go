package keyboard

import (
	"fmt"
	"strings"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"github.com/jezek/xgbutil"
	"github.com/jezek/xgbutil/xevent"
)

type KeyKey struct {
	Evtype int
	Win    xproto.Window
	Mod    uint16
	Code   xproto.Keycode
}

type CallbackKey func()

type KeyString struct {
	Str      string
	Callback CallbackKey
	Evtype   int
	Win      xproto.Window
	Grab     bool
}

type bindData struct {
	grabCounter int
	callbacks   []CallbackKey
}

func newBindData() *bindData {
	return &bindData{
		grabCounter: 0,
		callbacks:   make([]CallbackKey, 0),
	}
}

type Keyboard struct {
	connection *xgb.Conn
	setup      *xproto.SetupInfo
	keymap     *xproto.GetKeyboardMappingReply
	modmap     *xproto.GetModifierMappingReply
	binds      map[KeyKey]*bindData
	keystrings []KeyString
}

func NewKeyboard(connection *xgb.Conn) *Keyboard {
	return &Keyboard{
		connection: connection,
		setup:      xproto.Setup(connection),
		keymap:     nil,
		modmap:     nil,
		binds:      make(map[KeyKey]*bindData, 10),
		keystrings: make([]KeyString, 0, 10),
	}
}

func (k *Keyboard) updateMaps() {
	min := k.setup.MinKeycode
	max := k.setup.MaxKeycode
	var err error
	if k.keymap, err = xproto.GetKeyboardMapping(k.connection, min, byte(max-min+1)).Reply(); err != nil {
		panic(fmt.Sprintf("COULD NOT GET KEYBOARD MAPPING: %v\n"+
			"THIS IS AN UNRECOVERABLE ERROR.\n",
			err))
	}
	if k.modmap, err = xproto.GetModifierMapping(k.connection).Reply(); err != nil {
		panic(fmt.Sprintf("COULD NOT GET MODIFIER MAPPING: %v\n"+
			"THIS IS AN UNRECOVERABLE ERROR.\n",
			err))
	}
	for i := 0; i != len(k.modmap.Keycodes)/int(k.modmap.KeycodesPerModifier); i++ {
		arr := []string{}
		for j := 0; j != int(k.modmap.KeycodesPerModifier); j++ {
			arr = append(arr, fmt.Sprintf("0x%x", k.modmap.Keycodes[i*int(k.modmap.KeycodesPerModifier)+j]))
		}
		fmt.Println(arr)
	}

}

func (k *Keyboard) Init() {
	k.updateMaps()
}

func (k *Keyboard) isGrabbed(evtype int, win xproto.Window, mods uint16, keycode xproto.Keycode) bool {
	key := KeyKey{Evtype: evtype, Win: win, Mod: mods, Code: keycode}
	if data, ok := k.binds[key]; ok && data.grabCounter > 0 {
		return true
	}

	return false
}

func (k *Keyboard) attachKeyBindCallback(evtype int, win xproto.Window, mods uint16, keycode xproto.Keycode, fn CallbackKey) {
	key := KeyKey{Evtype: evtype, Win: win, Mod: mods, Code: keycode}

	data, ok := k.binds[key]
	if !ok {
		data = newBindData()
		k.binds[key] = data
	}

	data.callbacks = append(data.callbacks, fn)
	data.grabCounter++
}

func (k *Keyboard) addKeyString(callback CallbackKey, evtype int, win xproto.Window, keyStr string, grab bool) {
	val := KeyString{
		Str:      keyStr,
		Callback: callback,
		Evtype:   evtype,
		Win:      win,
		Grab:     grab,
	}
	k.keystrings = append(k.keystrings, val)
}

func (k *Keyboard) runKeyBindCallbacks(event interface{}, evtype int, win xproto.Window, mods uint16, keycode xproto.Keycode) {
	key := KeyKey{Evtype: evtype, Win: win, Mod: mods, Code: keycode}
	if data, ok := k.binds[key]; ok {
		fns := make([]CallbackKey, len(data.callbacks))
		copy(fns, data.callbacks)
		for _, fn := range fns {
			fn()
		}
	}
}

func (k *Keyboard) keysymGetWithMap(keycode xproto.Keycode, column byte) xproto.Keysym {
	min := k.setup.MinKeycode
	i := (int(keycode)-int(min))*int(k.keymap.KeysymsPerKeycode) + int(column)

	return k.keymap.Keysyms[i]
}

func (k *Keyboard) keysymGet(keycode xproto.Keycode, column byte) xproto.Keysym {
	return k.keysymGetWithMap(keycode, column)
}

func (k *Keyboard) keycodesGet(keysym xproto.Keysym) []xproto.Keycode {
	var c byte
	var keycode xproto.Keycode
	keycodes := make([]xproto.Keycode, 0)
	set := make(map[xproto.Keycode]bool, 0)

	for kc := int(k.setup.MinKeycode); kc <= int(k.setup.MaxKeycode); kc++ {
		keycode = xproto.Keycode(kc)
		for c = 0; c < k.keymap.KeysymsPerKeycode; c++ {
			if keysym == k.keysymGet(keycode, c) && !set[keycode] {
				keycodes = append(keycodes, keycode)
				set[keycode] = true
			}
		}
	}
	return keycodes
}

func (k *Keyboard) strToKeycodes(str string) []xproto.Keycode {
	// Do some fancy case stuff before we give up.
	sym, ok := keysyms[str]
	if !ok {
		sym, ok = keysyms[strings.Title(str)]
	}
	if !ok {
		sym, ok = keysyms[strings.ToLower(str)]
	}
	if !ok {
		sym, ok = keysyms[strings.ToUpper(str)]
	}

	// If we don't know what 'str' is, return 0.
	// There will probably be a bad access. We should do better than that...
	if !ok {
		return []xproto.Keycode{}
	}
	return k.keycodesGet(sym)
}

func (k *Keyboard) ParseString(s string) (uint16, []xproto.Keycode, error) {
	mods, kcs := uint16(0), []xproto.Keycode{}
	for _, part := range strings.Split(s, "-") {
		switch strings.ToLower(part) {
		case "shift":
			mods |= xproto.ModMaskShift
		case "lock":
			mods |= xproto.ModMaskLock
		case "control":
			mods |= xproto.ModMaskControl
		case "mod1":
			mods |= xproto.ModMask1
		case "mod2":
			mods |= xproto.ModMask2
		case "mod3":
			mods |= xproto.ModMask3
		case "mod4":
			mods |= xproto.ModMask4
		case "mod5":
			mods |= xproto.ModMask5
		case "any":
			mods |= xproto.ModMaskAny
		default: // a key code!
			if len(kcs) == 0 { // only accept the first keycode we see
				kcs = k.strToKeycodes(part)
			}
		}
	}

	if len(kcs) == 0 {
		return 0, nil, fmt.Errorf("Could not find a valid keycode in the "+
			"string '%s'. Key binding failed.", s)
	}

	return mods, kcs, nil
}

func (k *Keyboard) GrabChecked(win xproto.Window, mods uint16, key xproto.Keycode) error {
	var err error
	for _, m := range ignoreMods {
		err = xproto.GrabKeyChecked(k.connection, true, win, mods|m, key, xproto.GrabModeAsync, xproto.GrabModeAsync).Check()
		if err != nil {
			return err
		}
	}
	return nil
}

// Ungrab undoes Grab. It will handle all combinations od modifiers found
// in ignoreMods.
func (k *Keyboard) Ungrab(win xproto.Window, mods uint16, key xproto.Keycode) {
	for _, m := range ignoreMods {
		xproto.UngrabKeyChecked(k.connection, key, win, mods|m).Check()
	}
}

// func (k *Keyboard) connectedKeyBind(evtype int, win xproto.Window) bool {
// 	// Since we can't create a full key, loop through all key binds
// 	// and check if evtype and window match.
// 	for key := range k.keybinds {
// 		if key.Evtype == evtype && key.Win == win {
// 			return true
// 		}
// 	}
// 	return false
// }

func (k *Keyboard) Connect(callback CallbackKey, evtype int, win xproto.Window, keyStr string, grab, reconnect bool) error {
	// Get the mods/key first
	mods, keycodes, err := k.ParseString(keyStr)
	if err != nil {
		return err
	}

	// Only do the grab if we haven't yet on this window.
	for _, keycode := range keycodes {
		if grab && !k.isGrabbed(evtype, win, mods, keycode) {
			if err := k.GrabChecked(win, mods, keycode); err != nil {
				// If a bad access, let's be nice and give a good error message.
				switch err.(type) {
				case xproto.AccessError:
					return fmt.Errorf("Got a bad access error when trying to "+
						"bind '%s'. This usually means another client has "+
						"already grabbed this keybinding.", keyStr)
				default:
					return fmt.Errorf("Could not bind '%s' because: %s",
						keyStr, err)
				}
			}
		}

		// If we've never grabbed anything on this window before, we need to
		// make sure we can respond to it in the main event loop.
		// Never do this if we're reconnecting.
		if !reconnect {
			// var allCb xgbutil.Callback
			// if evtype == xevent.KeyPress {
			// 	allCb = xevent.KeyPressFun(runKeyPressCallbacks)
			// } else {
			// 	allCb = xevent.KeyReleaseFun(runKeyReleaseCallbacks)
			// }

			// // If this is the first Key{Press|Release}Event on this window,
			// // then we need to listen to Key{Press|Release} events in the main
			// // loop.
			// if !k.connectedKeyBind(evtype, win) {
			// 	allCb.Connect(xu, win)
			// }
		}

		// Finally, attach the callback.
		k.attachKeyBindCallback(evtype, win, mods, keycode, callback)
	}

	// Keep track of all unique key connections.
	if !reconnect {
		k.addKeyString(callback, evtype, win, keyStr, grab)
	}

	return nil
}

// DeduceKeyInfo AND's the "ignored modifiers" out of the state returned by
// a Key{Press,Release} event. This is useful to connect a (state, keycode)
// tuple from an event with a tuple specified by the user.
func DeduceKeyInfo(state uint16, detail xproto.Keycode) (uint16, xproto.Keycode) {
	mods, kc := state, detail
	for _, m := range ignoreMods {
		mods &= ^m
	}
	return mods, kc
}

func (k *Keyboard) OnKeyRelease(event xproto.KeyReleaseEvent) {
	mods, kc := DeduceKeyInfo(event.State, event.Detail)

	k.runKeyBindCallbacks(event, xevent.KeyRelease, event.Event, mods, kc)
}

func (k *Keyboard) OnMappingNotify(event xproto.MappingNotifyEvent) {
	k.updateMaps()
	if event.Request == xproto.MappingKeyboard {
		k.DetachAll()
		for _, ks := range k.keystrings {
			err := k.Connect(ks.Callback, ks.Evtype, ks.Win, ks.Str, ks.Grab, true)
			if err != nil {
				xgbutil.Logger.Println(err)
			}
		}
	}
}

func (k *Keyboard) keyKeys() []KeyKey {
	keys := make([]KeyKey, len(k.binds))
	i := 0
	for key := range k.binds {
		keys[i] = key
		i++
	}

	return keys
}

func (k *Keyboard) DetachKey(key KeyKey) {
	if _, ok := k.binds[key]; ok {
		delete(k.binds, key)
		if k.isGrabbed(key.Evtype, key.Win, key.Mod, key.Code) {
			k.Ungrab(key.Win, key.Mod, key.Code)
		}
	}
}

func (k *Keyboard) DetachAll() {
	mkeys := k.keyKeys()
	k.binds = make(map[KeyKey]*bindData, 10)
	for _, key := range mkeys {
		if k.isGrabbed(key.Evtype, key.Win, key.Mod, key.Code) {
			k.Ungrab(key.Win, key.Mod, key.Code)
		}
	}
}
