package portal

import (
	"errors"

	"github.com/godbus/dbus/v5"
)

var (
	initErr  = errors.New("failed to init xdg_desktop_portal module")
	startErr = errors.New("failed to start xdg_desktop_portal module")
)

type (
	dbusResponseCallback func([]interface{}) error
	dbusSignalCallback   func([]interface{})
)

type globalShortcutDefinition struct {
	ID   string
	Data map[string]dbus.Variant
}

func newGlobalShortcutDefinition(id, description string) globalShortcutDefinition {
	return globalShortcutDefinition{
		ID: id,
		Data: map[string]dbus.Variant{
			"description": dbus.MakeVariant(description),
			// "preferred_trigger": dbus.MakeVariant(`MOD4+r`),
			// "trigger_description": dbus.MakeVariant(`MOD4+r`),
		},
	}
}
