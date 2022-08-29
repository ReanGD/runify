package x11

import (
	"errors"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/rpc"
	"github.com/jezek/xgbutil"
	"github.com/jezek/xgbutil/keybind"
	"github.com/jezek/xgbutil/xevent"
	"go.uber.org/zap"
)

type hotkeyID uint32

func (h hotkeyID) ZapField() zap.Field {
	return zap.Uint32("HotkeyID", uint32(h))
}

type x11Keybind struct {
	hotkeysCh   chan hotkeyID
	xConnection *xgbutil.XUtil
	rpc         *rpc.Rpc

	moduleLogger *zap.Logger
}

func newX11Keybind() *x11Keybind {
	return &x11Keybind{
		hotkeysCh:    nil,
		xConnection:  nil,
		rpc:          nil,
		moduleLogger: nil,
	}
}

func (h *x11Keybind) onInit(cfg *config.Config, xConnection *xgbutil.XUtil, rpc *rpc.Rpc, moduleLogger *zap.Logger) error {
	channelLen := cfg.Get().X11.HotkeysChannelLen
	h.hotkeysCh = make(chan hotkeyID, channelLen)
	h.xConnection = xConnection
	h.rpc = rpc
	h.moduleLogger = moduleLogger
	keybind.Initialize(h.xConnection)
	h.register(hotkeyID(0), "Mod4-r")

	return nil
}

func (h *x11Keybind) onStart() {
}

func (h *x11Keybind) onHotkey(id hotkeyID) {
	h.rpc.ShowUI()
}

func (h *x11Keybind) onStop() {
	keybind.Detach(h.xConnection, h.xConnection.RootWin())
}

func (h *x11Keybind) register(id hotkeyID, keyStr string) error {
	if keyStr == "" {
		return nil
	}

	xu := h.xConnection
	_, _, err := keybind.ParseString(xu, keyStr)
	if err != nil {
		h.moduleLogger.Warn("Failed parse hotkey", zap.String("Hotkey", keyStr), id.ZapField(), zap.Error(err))
		return errors.New("Failed parse hotkey")
	}

	err = keybind.KeyReleaseFun(
		func(X *xgbutil.XUtil, e xevent.KeyReleaseEvent) {
			h.hotkeysCh <- id
		}).Connect(xu, xu.RootWin(), keyStr, true)
	if err != nil {
		h.moduleLogger.Warn("Failed register hotkey", zap.String("Hotkey", keyStr), id.ZapField(), zap.Error(err))
		return errors.New("Failed register hotkey")
	}

	return err
}
