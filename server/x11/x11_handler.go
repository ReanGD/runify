package x11

import (
	"errors"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/rpc"
	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xproto"
	"github.com/jezek/xgbutil"
	"github.com/jezek/xgbutil/xevent"
	"go.uber.org/zap"
)

type x11Handler struct {
	xConnection *xgbutil.XUtil
	x11EventsCh chan interface{}
	keybind     *x11Keybind
	clipboard   *x11Clipboard

	moduleLogger *zap.Logger
}

func newX11Handler() *x11Handler {
	return &x11Handler{
		xConnection:  nil,
		x11EventsCh:  nil,
		keybind:      newX11Keybind(),
		clipboard:    newX11Clipboard(),
		moduleLogger: nil,
	}
}

func (h *x11Handler) getX11EventsCh() <-chan interface{} {
	return h.x11EventsCh
}

func (h *x11Handler) getHotkeysCh() <-chan hotkeyID {
	return h.keybind.hotkeysCh
}

func (h *x11Handler) onInit(cfg *config.Config, rpc *rpc.Rpc, moduleLogger *zap.Logger) error {
	h.moduleLogger = moduleLogger

	var err error
	h.xConnection, err = xgbutil.NewConn()
	if err != nil {
		h.moduleLogger.Warn("Failed connect to x server", zap.Error(err))
		return errors.New("Failed connect to x server")
	}

	x11EventChannelLen := cfg.Get().X11.X11EventChannelLen
	h.x11EventsCh = make(chan interface{}, x11EventChannelLen)

	err = h.keybind.onInit(cfg, h.xConnection, rpc, moduleLogger)
	if err != nil {
		return err
	}

	atoms, ok := newAtomStorage(h.xConnection.Conn(), moduleLogger)
	if !ok {
		return errInitX11
	}

	err = h.clipboard.onInit(atoms, h.xConnection.Conn(), h.xConnection.Dummy(), moduleLogger)
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
		xevent.Main(h.xConnection)
		wg.Done()
	}()

	startWG.Wait()
	xevent.HookFun(h.hookX11Event).Connect(h.xConnection)
	h.keybind.onStart()
	h.clipboard.onStart()
}

// called in external goroutine
func (h *x11Handler) hookX11Event(xu *xgbutil.XUtil, event interface{}) bool {
	switch e := event.(type) {
	case xfixes.SelectionNotifyEvent:
		xu.TimeSet(e.Timestamp)
		h.x11EventsCh <- event
		return false
	case xproto.SelectionNotifyEvent:
		xu.TimeSet(e.Time)
		h.x11EventsCh <- event
		return false
	case xproto.SelectionRequestEvent:
		xu.TimeSet(e.Time)
		h.x11EventsCh <- event
		return false
	case xproto.SelectionClearEvent:
		xu.TimeSet(e.Time)
		h.x11EventsCh <- event
		return false
	}

	return true
}

func (h *x11Handler) onX11Event(event interface{}) {
	switch e := event.(type) {
	case xfixes.SelectionNotifyEvent:
		h.clipboard.onSelectionChange(e)
	case xproto.SelectionNotifyEvent:
		h.clipboard.onSelectionNotify(e)
	case xproto.SelectionRequestEvent:
		h.clipboard.onSelectionRequest(e)
	case xproto.SelectionClearEvent:
		h.clipboard.onSelectionClear(e)
	}
}

func (h *x11Handler) onHotkey(id hotkeyID) {
	h.keybind.onHotkey(id)
}

func (h *x11Handler) onStop() {
	h.keybind.onStop()
	h.clipboard.onStop()
	h.xConnection.Quit = true
	h.xConnection = nil
}
