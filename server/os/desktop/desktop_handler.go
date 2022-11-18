package desktop

import (
	"github.com/ReanGD/runify/server/global/module"
	"go.uber.org/zap"
)

type handler struct {
	moduleLogger *zap.Logger
}

func newHandler() *handler {
	return &handler{
		moduleLogger: nil,
	}
}

func (h *handler) init(errorCtx *module.ErrorCtx, moduleLogger *zap.Logger) error {
	h.moduleLogger = moduleLogger

	return nil
}

func (h *handler) start() {
}

func (h *handler) writeToClipboard(cmd *writeToClipboardCmd) {
}

func (h *handler) setHotkey(cmd *setHotkeyCmd) {
}

func (h *handler) removeHotkey(cmd *removeHotkeyCmd) {
}

func (h *handler) stop() {
}
