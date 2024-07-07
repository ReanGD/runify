package portal

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/godbus/dbus/v5"
	"go.uber.org/zap"
)

const (
	globalShortcutsName = "org.freedesktop.portal.GlobalShortcuts"

	globalShortcutsSignalActivated   = globalShortcutsName + ".Activated"
	globalShortcutsSignalDeactivated = globalShortcutsName + ".Deactivated"

	globalShortcutsMethodCreateSession = globalShortcutsName + ".CreateSession"
	globalShortcutsMethodListShortcuts = globalShortcutsName + ".ListShortcuts"
	globalShortcutsMethodBindShortcuts = globalShortcutsName + ".BindShortcuts"

	portalSessionName        = "org.freedesktop.portal.Session"
	portalSessionMethodClose = portalSessionName + ".Close"

	shortcutPrefix = "com.runify."
	shortcutOpen   = shortcutPrefix + "open"
)

var (
	gsCreateSessionErr = errors.New("GlobalShortcuts failed create session")
	gsListShortcutsErr = errors.New("GlobalShortcuts failed get list of shortcuts")
	gsBindShortcutsErr = errors.New("GlobalShortcuts failed bind shortcuts")
)

type globalShortcuts struct {
	client            *dbusClient
	provider          api.Provider
	sessionObj        dbus.BusObject
	sessionObjectPath dbus.ObjectPath
	shortcutsState    map[string]struct{}

	errorCtx     *module.ErrorCtx
	moduleLogger *zap.Logger
}

func newGlobalShortcuts(
	client *dbusClient,
	provider api.Provider,
	errorCtx *module.ErrorCtx,
	moduleLogger *zap.Logger,
) (*globalShortcuts, error) {
	result := &globalShortcuts{
		client:            client,
		provider:          provider,
		sessionObj:        nil,
		sessionObjectPath: "",
		shortcutsState:    make(map[string]struct{}),
		errorCtx:          errorCtx,
		moduleLogger:      moduleLogger.With(zap.String("api", "GlobalShortcuts")),
	}

	if !client.signalSubscribe(globalShortcutsSignalActivated, result.onActivated) {
		result.moduleLogger.Error("Failed to subscribe to activated signal")
		return nil, initErr
	}

	if !client.signalSubscribe(globalShortcutsSignalDeactivated, result.onDeactivated) {
		result.moduleLogger.Error("Failed to subscribe to deactivated signal")
		return nil, initErr
	}

	return result, nil
}

func (gs *globalShortcuts) onActivated(body []interface{}) {
	if len(body) != 4 {
		gs.moduleLogger.Warn("Failed parsing activated signal body", zap.Any("body", body))
		return
	}

	shortcutName, ok := body[1].(string)
	if !ok {
		gs.moduleLogger.Warn("Failed parsing activated signal shortcut", zap.Any("body", body))
		return
	}

	if _, exists := gs.shortcutsState[shortcutName]; exists {
		return
	}
	gs.shortcutsState[shortcutName] = struct{}{}

	gs.provider.Activate(shortcut.NewAction(shortcutName))
}

func (gs *globalShortcuts) onDeactivated(body []interface{}) {
	if len(body) != 4 {
		gs.moduleLogger.Warn("Failed parsing deactivated signal body", zap.Any("body", body))
		return
	}

	shortcutName, ok := body[1].(string)
	if !ok {
		gs.moduleLogger.Warn("Failed parsing deactivated signal shortcut", zap.Any("body", body))
		return
	}

	if _, exists := gs.shortcutsState[shortcutName]; !exists {
		return
	}

	delete(gs.shortcutsState, shortcutName)
}

func (gs *globalShortcuts) onCreateSessionResponse(body []interface{}) error {
	if len(body) != 2 {
		gs.moduleLogger.Error("Failed parsing create session response", zap.Any("body", body))
		return gsCreateSessionErr
	}

	response, ok := body[1].(map[string]dbus.Variant)
	if !ok {
		gs.moduleLogger.Error("Failed parsing create session response body", zap.Any("body", body))
		return gsCreateSessionErr

	}

	sessionHandle, ok := response["session_handle"]
	if !ok {
		gs.moduleLogger.Error("Failed find 'session_handle' in create session response", zap.Any("body", body))
		return gsCreateSessionErr
	}

	if err := sessionHandle.Store(&gs.sessionObjectPath); err != nil {
		gs.moduleLogger.Error("Failed get session path from GlobalShortcuts create session response", zap.Any("body", sessionHandle))
		return gsCreateSessionErr
	}
	gs.sessionObj = gs.client.newObject(gs.sessionObjectPath)

	shortcuts := []globalShortcutDefinition{
		newGlobalShortcutDefinition(shortcutOpen, "Open runify (preferred: MOD4+r)"),
	}

	if err := gs.bind(shortcuts); err != nil {
		gs.moduleLogger.Error("Failed to bind global shortcuts", zap.Error(err))
		return err
	}

	return nil
}

func (gs *globalShortcuts) createSession() error {
	sessionToken := gs.client.newToken()
	handleToken := gs.client.newToken()

	err := gs.client.request(
		globalShortcutsMethodCreateSession,
		handleToken,
		gs.onCreateSessionResponse,
		map[string]dbus.Variant{
			"handle_token":         dbus.MakeVariant(handleToken),
			"session_handle_token": dbus.MakeVariant(sessionToken),
		})
	if err != nil {
		gs.moduleLogger.Error("Failed to create global shortcuts session", zap.Error(err))
		return gsListShortcutsErr
	}

	return nil
}

// nolint: unused
func (gs *globalShortcuts) onListResponse(body []interface{}) error {
	gs.moduleLogger.Info("List response", zap.Any("body", body))
	return nil
}

// nolint: unused
func (gs *globalShortcuts) list() error {
	handleToken := gs.client.newToken()

	err := gs.client.request(
		globalShortcutsMethodListShortcuts,
		handleToken,
		gs.onListResponse,
		gs.sessionObjectPath,
		map[string]dbus.Variant{
			"handle_token": dbus.MakeVariant(handleToken),
		},
	)
	if err != nil {
		gs.moduleLogger.Error("Failed to list global shortcuts", zap.Error(err))
	}

	return err
}

func (gs *globalShortcuts) onBindResponse(body []interface{}) error {
	gs.moduleLogger.Debug("Bind response", zap.Any("body", body))
	return nil
}

func (gs *globalShortcuts) bind(shortcuts []globalShortcutDefinition) error {
	handleToken := gs.client.newToken()
	parentWindow := ""

	err := gs.client.request(
		globalShortcutsMethodBindShortcuts,
		handleToken,
		gs.onBindResponse,
		gs.sessionObjectPath,
		shortcuts,
		parentWindow,
		map[string]dbus.Variant{
			"handle_token": dbus.MakeVariant(handleToken),
		},
	)
	if err != nil {
		gs.moduleLogger.Error("Failed to bind global shortcuts", zap.Error(err))
		return gsBindShortcutsErr
	}

	return nil
}

func (gs *globalShortcuts) close() {
	if gs.sessionObj == nil {
		return
	}

	if err := gs.sessionObj.Call(portalSessionMethodClose, 0).Err; err != nil {
		gs.moduleLogger.Warn("Failed to close global shortcuts session", zap.Error(err))
	}

	gs.sessionObj = nil
}
