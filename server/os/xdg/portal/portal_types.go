package portal

import (
	"errors"

	"github.com/godbus/dbus/v5"
)

var initErr = errors.New("failed to init xdg_desktop_portal module")

const (
	portalName = `org.freedesktop.portal.Desktop`
	portalPath = dbus.ObjectPath(`/org/freedesktop/portal/desktop`)

	portalRequestName           = `org.freedesktop.portal.Request`
	portalRequestMemberResponse = `Response`
	portalRequestSignalResponse = portalRequestName + `.` + portalRequestMemberResponse

	portalSessionName        = `org.freedesktop.portal.Session`
	portalSessionMethodClose = portalSessionName + `.Close`

	globalShortcutsName = `org.freedesktop.portal.GlobalShortcuts`

	globalShortcutsSignalActivated   = globalShortcutsName + `.Activated`
	globalShortcutsSignalDeactivated = globalShortcutsName + `.Deactivated`

	globalShortcutsMethodCreateSession = globalShortcutsName + `.CreateSession`
	globalShortcutsMethodListShortcuts = globalShortcutsName + `.ListShortcuts`
	globalShortcutsMethodBindShortcuts = globalShortcutsName + `.BindShortcuts`

	shortcutPrefix = `com.runify.`
	shortcutOpen   = shortcutPrefix + `open`
)

type dbusCallback func(*dbus.Signal)

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
