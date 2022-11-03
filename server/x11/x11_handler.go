package x11

import (
	"sync"

	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/system"
	"github.com/ReanGD/runify/server/system/module"
)

type x11Handler struct {
	cfg         *config.Configuration
	xConnection *x11Connection
	errorCh     chan error
	x11EventsCh chan interface{}
	shortcutCh  chan bindID
	keyboard    *x11Keyboard
	clipboard   *x11Clipboard

	moduleLogger *zap.Logger
}

func newX11Handler() *x11Handler {
	return &x11Handler{
		cfg:          nil,
		xConnection:  newX11Connection(),
		errorCh:      make(chan error),
		x11EventsCh:  nil,
		shortcutCh:   nil,
		keyboard:     newX11Keyboard(),
		clipboard:    newX11Clipboard(),
		moduleLogger: nil,
	}
}

func (h *x11Handler) getErrorCh() <-chan error {
	return h.errorCh
}

func (h *x11Handler) getX11EventsCh() <-chan interface{} {
	return h.x11EventsCh
}

func (h *x11Handler) getShortcutCh() <-chan bindID {
	return h.shortcutCh
}

func (h *x11Handler) onInit(cfg *config.Configuration, rpc module.Rpc, moduleLogger *zap.Logger) error {
	h.cfg = cfg
	h.moduleLogger = moduleLogger

	var err error
	display := ""
	if err = h.xConnection.onInit(display, moduleLogger); err != nil {
		return err
	}

	x11EventChannelLen := h.cfg.X11.X11EventChannelLen
	h.x11EventsCh = make(chan interface{}, x11EventChannelLen)
	h.shortcutCh = make(chan bindID, h.cfg.X11.HotkeysChannelLen)
	err = h.keyboard.onInit(h.xConnection.connection, h.xConnection.rootWindowID, h.errorCh, h.shortcutCh, moduleLogger)
	if err != nil {
		return err
	}

	atoms, ok := newAtomStorage(h.xConnection.connection, moduleLogger)
	if !ok {
		return errInitX11
	}

	err = h.clipboard.onInit(atoms, h.xConnection.connection, h.xConnection.dummyWindowID, moduleLogger)
	if err != nil {
		return err
	}

	return nil
}

func (h *x11Handler) onStart(wg *sync.WaitGroup) {
	wg.Add(1)

	startWG := &sync.WaitGroup{}
	startWG.Add(1)
	go func() {
		startWG.Done()
		if err := h.xConnection.onStart(h.x11EventsCh); err != nil {
			h.errorCh <- err
		}
		wg.Done()
	}()

	startWG.Wait()
	h.keyboard.onStart()
	h.clipboard.onStart()
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

func (h *x11Handler) bindShortcut(shortcut string) (bindID, system.Error) {
	return h.keyboard.bind(shortcut)
}

func (h *x11Handler) writeToClipboard(cmd *writeToClipboardCmd) {
	h.clipboard.writeToClipboard(cmd.isPrimary, cmd.data)
}

func (h *x11Handler) onStop() {
	h.keyboard.onStop()
	h.clipboard.onStop()
	h.xConnection.quit = true
	h.xConnection = nil
}
