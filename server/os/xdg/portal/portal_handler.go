package portal

import (
	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

type handler struct {
	dbusClient      *dbusClient
	globalShortcuts *globalShortcuts
	errorCtx        *module.ErrorCtx
	moduleLogger    *zap.Logger
}

func newHandler() *handler {
	return &handler{
		dbusClient:      nil,
		globalShortcuts: nil,
		errorCtx:        nil,
		moduleLogger:    nil,
	}
}

func (h *handler) init(root *Portal, cfg *config.XDGDesktopPortalCfg) error {
	h.errorCtx = root.GetErrorCtx()
	logger := root.GetModuleLogger()
	h.moduleLogger = logger

	var err error
	if h.dbusClient, err = newDBusClient(h.errorCtx, cfg, h.moduleLogger); err != nil {
		return err
	}

	if h.globalShortcuts, err = newGlobalShortcuts(
		h.dbusClient, root.provider, h.errorCtx, h.moduleLogger); err != nil {
		return err
	}

	return nil
}

func (h *handler) start() []*types.HandledChannel {
	ch := h.dbusClient.start()
	chs := []*types.HandledChannel{ch}

	if err := h.globalShortcuts.createSession(); err != nil {
		h.errorCtx.SendError(err)
		return chs
	}

	// shortcuts := []globalShortcutDefinition{
	// 	newGlobalShortcutDefinition(shortcutOpen, "Open runify (preferred: MOD4+r)"),
	// }

	// if err := h.dHandler.globalShortcutsBind(shortcuts); err != nil {
	// 	h.moduleLogger.Error("Failed to bind global shortcuts", zap.Error(err))
	// 	h.errorCtx.SendError(err)
	// 	return
	// }

	return chs
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
	if h.globalShortcuts != nil {
		h.globalShortcuts.close()
		h.globalShortcuts = nil
	}

	if h.dbusClient != nil {
		h.dbusClient.close()
		h.dbusClient = nil
	}
}
