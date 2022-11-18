package x11

import (
	"github.com/ReanGD/runify/server/global/module"
	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

var (
	zapInitX11 = zap.String("Method", "x11.X11::init")
)

type x11Handler struct {
	conn      *connection
	atoms     *atomStorage
	clipboard *clipboard
	keyboard  *keyboard

	moduleLogger *zap.Logger
}

func newX11Handler() *x11Handler {
	return &x11Handler{
		conn:         newConnection(),
		atoms:        newAtomStorage(),
		clipboard:    newClipboard(),
		keyboard:     newKeyboard(),
		moduleLogger: nil,
	}
}

func (h *x11Handler) init(eventsCh chan<- interface{}, errorCtx *module.ErrorCtx, moduleLogger *zap.Logger) error {
	h.moduleLogger = moduleLogger

	display := ""
	if !h.conn.init(display, h.atoms, eventsCh, errorCtx, moduleLogger) {
		return initErr
	}

	if !h.atoms.init(h.conn, moduleLogger) {
		return initErr
	}

	screen := h.conn.getDefaultScreen()
	rootWindow := h.conn.newWindow(screen.Root)
	dummyWindow, ok := h.conn.createWindow(screen.Root, screen, zapInitX11)
	if !ok {
		moduleLogger.Warn("Failed create X11 dummy window")
		return initErr
	}

	if !h.clipboard.init(h.conn, h.atoms, dummyWindow, moduleLogger) {
		return initErr
	}

	if !h.keyboard.init(h.conn, rootWindow, errorCtx, moduleLogger) {
		return initErr
	}

	return nil
}

func (h *x11Handler) start() {
	h.conn.start()
}

func (h *x11Handler) onX11Event(event interface{}) {
	switch e := event.(type) {
	case xproto.MappingNotifyEvent:
		h.keyboard.onMappingNotify(e)
	case xproto.KeyReleaseEvent:
		h.keyboard.onKeyRelease(e)
	case xfixes.SelectionNotifyEvent:
		h.clipboard.onSelectionChange(e)
	case xproto.SelectionNotifyEvent:
		h.clipboard.onSelectionNotify(e)
	case xproto.PropertyNotifyEvent:
		h.clipboard.onPropertyNotify(e)
	case xproto.SelectionRequestEvent:
		h.clipboard.onSelectionRequest(e)
	case xproto.SelectionClearEvent:
		h.clipboard.onSelectionClear(e)
	}
}

// func (h *x11Handler) bindShortcut(shortcut string) (bindID, global.Error) {
// 	return h.keyboard.bind(shortcut)
// }

func (h *x11Handler) subscribeToClipboard(cmd *subscribeToClipboardCmd) {
	cmd.result <- h.clipboard.subscribeToClipboard(cmd.isPrimary, cmd.ch)
}

func (h *x11Handler) writeToClipboard(cmd *writeToClipboardCmd) {
	cmd.result <- h.clipboard.writeToClipboard(cmd.isPrimary, cmd.data)
}

func (h *x11Handler) subscribeToHotkeys(cmd *subscribeToHotkeysCmd) {
	cmd.result <- h.keyboard.subscribeToHotkeys(cmd.ch)
}

func (h *x11Handler) stop() {
	h.conn.stop()
	h.keyboard.close()
}
