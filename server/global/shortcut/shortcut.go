package shortcut

import (
	"errors"
	"strings"

	"go.uber.org/zap"
)

var (
	ErrDuplicateModifier    = errors.New("duplicate modifier")
	ErrSecondKeyUnsupported = errors.New("second key unsupported")
	ErrNoKeySpecified       = errors.New("no key specified")
)

type Hotkey struct {
	mods ModId
	key  KeyId
	text string
}

func NewHotkey(s string) (*Hotkey, error) {
	var ok bool
	var key KeyId
	var mods ModId
	var textMods string
	var textKey string

	for _, part := range strings.Split(s, "+") {
		switch strings.ToLower(part) {
		case "shift":
			if mods&ModShift != 0 {
				return nil, ErrDuplicateModifier
			}
			mods |= ModShift
			textMods += "Shift+"

		case "ctrl":
			fallthrough

		case "control":
			if mods&ModControl != 0 {
				return nil, ErrDuplicateModifier
			}
			mods |= ModControl
			textMods += "Control+"

		case "alt":
			if mods&ModAlt != 0 {
				return nil, ErrDuplicateModifier
			}
			mods |= ModAlt
			textMods += "Alt+"

		case "super":
			if mods&ModSuper != 0 {
				return nil, ErrDuplicateModifier
			}
			mods |= ModSuper
			textMods += "Super+"

		default:
			if key != 0 {
				return nil, ErrSecondKeyUnsupported
			}

			textKey = part
			key, ok = KeyStrToId[textKey]
			if !ok {
				textKey = strings.Title(part)
				key, ok = KeyStrToId[textKey]
			}
			if !ok {
				textKey = strings.ToLower(part)
				key, ok = KeyStrToId[textKey]
			}
			if !ok {
				textKey = strings.ToUpper(part)
				key, ok = KeyStrToId[textKey]
			}
			if !ok {
				return nil, errors.New("unknown key or modifier: " + part)
			}
		}
	}

	if key == 0 {
		return nil, ErrNoKeySpecified
	}

	return &Hotkey{
		mods: mods,
		key:  key,
		text: textMods + textKey,
	}, nil
}

func (h *Hotkey) Mods() ModId {
	return h.mods
}

func (h *Hotkey) Key() KeyId {
	return h.key
}

func (h *Hotkey) Id() HotkeyId {
	return HotkeyId(uint32(h.mods)<<16 | uint32(h.key))
}

func (h *Hotkey) String() string {
	return h.text
}

func (h *Hotkey) GoString() string {
	return h.String()
}

func (h *Hotkey) ZapField() zap.Field {
	return zap.String("Hotkey", h.String())
}

func (h *Hotkey) ZapFieldPrefix(prefix string) zap.Field {
	return zap.String(prefix+"Hotkey", h.String())
}
