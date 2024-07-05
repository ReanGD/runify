package portal

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/godbus/dbus/v5"
	"go.uber.org/zap"
)

type dbusHandler struct {
	conn      *dbus.Conn
	provider  api.Provider
	busObj    dbus.BusObject
	callbacks map[dbus.ObjectPath]dbusCallback

	gsSessionObj        dbus.BusObject
	gsSessionObjectPath dbus.ObjectPath
	gsShortcutsState    map[string]struct{}

	errorCtx     *module.ErrorCtx
	moduleLogger *zap.Logger
}

func newDbusHandler(conn *dbus.Conn, provider api.Provider, errorCtx *module.ErrorCtx, moduleLogger *zap.Logger) *dbusHandler {
	return &dbusHandler{
		conn:                conn,
		provider:            provider,
		busObj:              conn.Object(portalName, portalPath),
		callbacks:           make(map[dbus.ObjectPath]dbusCallback),
		gsSessionObj:        nil,
		gsSessionObjectPath: "",
		gsShortcutsState:    make(map[string]struct{}),
		errorCtx:            errorCtx,
		moduleLogger:        moduleLogger,
	}
}

func (h *dbusHandler) requestPath(token string) dbus.ObjectPath {
	busName := h.conn.Names()[0]
	busName = strings.ReplaceAll(busName, `:`, ``)
	busName = strings.ReplaceAll(busName, `.`, `_`)
	return dbus.ObjectPath(fmt.Sprintf("/org/freedesktop/portal/desktop/request/%s/%s", busName, token))
}

func (h *dbusHandler) newToken() string {
	return fmt.Sprintf("runify_%d", rand.Int())
}

func (h *dbusHandler) request(method string, token string, callback dbusCallback, args ...any) error {
	requestPath := h.requestPath(token)
	h.callbacks[requestPath] = callback
	return h.busObj.Call(method, 0, args...).Err
}

func (h *dbusHandler) onResponse(signal *dbus.Signal) {
	if callback, ok := h.callbacks[signal.Path]; ok {
		delete(h.callbacks, signal.Path)
		callback(signal)
	}
}

func (h *dbusHandler) globalShortcutsActivated(signal *dbus.Signal) {
	if len(signal.Body) != 4 {
		h.moduleLogger.Warn("failed parsing GlobalShortcuts activated signal body", zap.Any("body", signal.Body))
		return
	}

	shortcutName, ok := signal.Body[1].(string)
	if !ok {
		h.moduleLogger.Warn("failed parsing GlobalShortcuts activated signal shortcut", zap.Any("body", signal.Body))
		return
	}

	if _, exists := h.gsShortcutsState[shortcutName]; exists {
		return
	}

	h.provider.Activate(shortcut.NewAction(shortcutName))

	h.gsShortcutsState[shortcutName] = struct{}{}
}

func (h *dbusHandler) globalShortcutsDeactivated(signal *dbus.Signal) {
	if len(signal.Body) != 4 {
		h.moduleLogger.Warn("failed parsing GlobalShortcuts deactivated signal body", zap.Any("body", signal.Body))
		return
	}

	shortcutName, ok := signal.Body[1].(string)
	if !ok {
		h.moduleLogger.Warn("failed parsing GlobalShortcuts deactivated signal shortcut", zap.Any("body", signal.Body))
		return
	}

	if _, exists := h.gsShortcutsState[shortcutName]; !exists {
		return
	}

	delete(h.gsShortcutsState, shortcutName)
}

func (h *dbusHandler) onSignal(signal *dbus.Signal) {
	switch signal.Name {
	case portalRequestSignalResponse:
		h.onResponse(signal)
	case globalShortcutsSignalActivated:
		h.globalShortcutsActivated(signal)
	case globalShortcutsSignalDeactivated:
		h.globalShortcutsDeactivated(signal)
	default:
		h.moduleLogger.Warn("Unknown dbus signal", zap.String("signal", signal.Name))
	}
}

func (h *dbusHandler) globalShortcutsCreateSessionResponse(signal *dbus.Signal) {
	if len(signal.Body) != 2 {
		h.moduleLogger.Error("failed parsing GlobalShortcuts create session reply", zap.Any("body", signal.Body))
		h.errorCtx.SendError(errors.New("GlobalShortcuts failed create session"))
		return
	}

	body, ok := signal.Body[1].(map[string]dbus.Variant)
	if !ok {
		h.moduleLogger.Error("failed parsing GlobalShortcuts create session reply body", zap.Any("body", signal.Body))
		h.errorCtx.SendError(errors.New("GlobalShortcuts failed create session"))
		return
	}

	sessionHandle, ok := body["session_handle"]
	if !ok {
		h.moduleLogger.Error("failed find 'session_handle' in GlobalShortcuts create session reply", zap.Any("body", signal.Body))
		h.errorCtx.SendError(errors.New("GlobalShortcuts failed create session"))
		return
	}

	if err := sessionHandle.Store(&h.gsSessionObjectPath); err != nil {
		h.moduleLogger.Error("failed get session path from GlobalShortcuts create session reply", zap.Any("body", sessionHandle))
		h.errorCtx.SendError(errors.New("GlobalShortcuts failed create session"))
		return
	}
	h.gsSessionObj = h.conn.Object(portalName, h.gsSessionObjectPath)

	shortcuts := []globalShortcutDefinition{
		newGlobalShortcutDefinition(shortcutOpen, "Open runify (preferred: MOD4+r)"),
	}

	if err := h.globalShortcutsBind(shortcuts); err != nil {
		h.moduleLogger.Error("Failed to bind global shortcuts", zap.Error(err))
		h.errorCtx.SendError(err)
		return
	}
}

func (h *dbusHandler) globalShortcutsCreateSession() error {
	sessionToken := h.newToken()
	handleToken := h.newToken()

	return h.request(
		globalShortcutsMethodCreateSession,
		handleToken,
		h.globalShortcutsCreateSessionResponse,
		map[string]dbus.Variant{
			"handle_token":         dbus.MakeVariant(handleToken),
			"session_handle_token": dbus.MakeVariant(sessionToken),
		})
}

func (h *dbusHandler) globalShortcutsListResponse(signal *dbus.Signal) {
	h.moduleLogger.Info("globalShortcutsListResponse", zap.Any("signal", signal))
}

func (h *dbusHandler) globalShortcutsList() error {
	handleToken := h.newToken()

	return h.request(
		globalShortcutsMethodListShortcuts,
		handleToken,
		h.globalShortcutsListResponse,
		h.gsSessionObjectPath,
		map[string]dbus.Variant{
			"handle_token": dbus.MakeVariant(handleToken),
		},
	)
}

func (h *dbusHandler) globalShortcutsBindResponse(signal *dbus.Signal) {
	h.moduleLogger.Debug("globalShortcutsBindResponse", zap.Any("signal", signal))
}

func (h *dbusHandler) globalShortcutsBind(shortcuts []globalShortcutDefinition) error {
	handleToken := h.newToken()
	parentWindow := ""

	return h.request(
		globalShortcutsMethodBindShortcuts,
		handleToken,
		h.globalShortcutsBindResponse,
		h.gsSessionObjectPath,
		shortcuts,
		parentWindow,
		map[string]dbus.Variant{
			"handle_token": dbus.MakeVariant(handleToken),
		},
	)
}

func (h *dbusHandler) close() error {
	// defer func() {
	// 	if h.conn != nil {
	// 		if err := h.conn.Close(); err != nil {
	// 			h.moduleLogger.Warn("Failed to close dbus connection", zap.Error(err))
	// 		}
	// 		h.conn = nil
	// 	}
	// }()

	if h.gsSessionObj != nil {
		if err := h.gsSessionObj.Call(portalSessionMethodClose, 0).Err; err != nil {
			return err
		}
		h.gsSessionObj = nil
	}

	return nil
}
