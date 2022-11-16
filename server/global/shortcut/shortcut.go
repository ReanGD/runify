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

type Shortcut struct {
	mods ModId
	key  KeyId
	text string
}

func NewShortcut(s string) (*Shortcut, error) {
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
			textMods += "Ctrl+"

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

	return &Shortcut{
		mods: mods,
		key:  key,
		text: textMods + textKey,
	}, nil
}

func (s *Shortcut) Mods() ModId {
	return s.mods
}

func (s *Shortcut) Key() KeyId {
	return s.key
}

func (s *Shortcut) IsEqual(other *Shortcut) bool {
	return s.mods == other.mods && s.key == other.key
}

func (s *Shortcut) String() string {
	return s.text
}

func (s *Shortcut) GoString() string {
	return s.String()
}

func (s *Shortcut) ZapField() zap.Field {
	return zap.String("Shortcut", s.String())
}

func (s *Shortcut) ZapFieldPrefix(prefix string) zap.Field {
	return zap.String(prefix+"Shortcut", s.String())
}
