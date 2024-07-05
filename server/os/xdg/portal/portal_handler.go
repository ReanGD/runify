package portal

import (
	"github.com/ReanGD/runify/server/global/module"
	"github.com/godbus/dbus/v5"
	"go.uber.org/zap"
)

// var zapInitX11 = zap.String("Method", "x11.X11::init")

type handler struct {
	conn         *dbus.Conn
	dHandler     *dbusHandler
	errorCtx     *module.ErrorCtx
	moduleLogger *zap.Logger
}

func newHandler() *handler {
	return &handler{
		conn:         nil,
		dHandler:     nil,
		errorCtx:     nil,
		moduleLogger: nil,
	}
}

func (h *handler) init(root *Portal) error {
	h.errorCtx = root.GetErrorCtx()
	logger := root.GetModuleLogger()
	h.moduleLogger = logger

	var err error
	if h.conn, err = dbus.SessionBus(); err != nil {
		h.moduleLogger.Error("Failed to connect to session dbus", zap.Error(err))
		return initErr
	}

	opts := []dbus.MatchOption{
		dbus.WithMatchInterface(portalRequestName),
		dbus.WithMatchMember(portalRequestMemberResponse),
	}
	if err := h.conn.AddMatchSignal(opts...); err != nil {
		h.moduleLogger.Error("Failed to add match dbus signal", zap.Error(err))
		return initErr
	}

	h.dHandler = newDbusHandler(h.conn, root.provider, h.errorCtx, h.moduleLogger)

	return nil
}

func (h *handler) start(signalsCh chan *dbus.Signal) {
	h.conn.Signal(signalsCh)

	if err := h.dHandler.globalShortcutsCreateSession(); err != nil {
		h.moduleLogger.Error("Failed to create global shortcuts session", zap.Error(err))
		h.errorCtx.SendError(err)
		return
	}

	// shortcuts := []globalShortcutDefinition{
	// 	newGlobalShortcutDefinition(shortcutOpen, "Open runify (preferred: MOD4+r)"),
	// }

	// if err := h.dHandler.globalShortcutsBind(shortcuts); err != nil {
	// 	h.moduleLogger.Error("Failed to bind global shortcuts", zap.Error(err))
	// 	h.errorCtx.SendError(err)
	// 	return
	// }
}

func (h *handler) onSignal(event interface{}) {
	signal, ok := event.(*dbus.Signal)
	if !ok {
		h.moduleLogger.Warn("Failed to cast dbus signal")
		return
	}

	h.dHandler.onSignal(signal)
}

// func (h *handler) subscribeToClipboard(cmd *subscribeToClipboardCmd) {
// 	cmd.result.SetResult(h.clipboard.subscribeToClipboard(cmd.isPrimary, cmd.ch))
// }

// func (h *handler) writeToClipboard(cmd *writeToClipboardCmd) {
// 	cmd.result.SetResult(h.clipboard.writeToClipboard(cmd.isPrimary, cmd.data))
// }

// func (h *handler) subscribeToHotkeys(cmd *subscribeToHotkeysCmd) {
// 	cmd.result.SetResult(h.keyboard.subscribeToHotkeys(cmd.ch))
// }

// func (h *handler) bindHotkey(cmd *bindHotkeyCmd) {
// 	cmd.result.SetResult(h.keyboard.bind(cmd.hotkey))
// }

// func (h *handler) unbindHotkey(cmd *unbindHotkeyCmd) {
// 	cmd.result.SetResult(h.keyboard.unbind(cmd.hotkey))
// }

func (h *handler) stop() {
	if h.conn != nil {
		h.dHandler.close()
		if err := h.conn.Close(); err != nil {
			h.moduleLogger.Warn("Failed to close dbus connection", zap.Error(err))
		}
		h.conn = nil
	}
}
