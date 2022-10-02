package x11

import (
	"errors"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/rpc"
	"github.com/jezek/xgbutil"
	"github.com/jezek/xgbutil/xevent"
	"go.uber.org/zap"
)

type x11Handler struct {
	xConnection *xgbutil.XUtil
	keybind     *x11Keybind
	clipboard   *x11Clipboard

	moduleLogger *zap.Logger
}

func newX11Handler() *x11Handler {
	return &x11Handler{
		xConnection:  nil,
		keybind:      newX11Keybind(),
		clipboard:    newX11Clipboard(),
		moduleLogger: nil,
	}
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

	err = h.keybind.onInit(cfg, h.xConnection, rpc, moduleLogger)
	if err != nil {
		return err
	}

	err = h.clipboard.onInit(cfg, h.xConnection, moduleLogger)
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
	h.keybind.onStart()
	h.clipboard.onStart()
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
