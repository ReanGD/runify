package x11

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
	"github.com/ReanGD/runify/server/logger"
	"go.uber.org/zap"
)

const ModuleName = "x11"

type X11 struct {
	handler     *x11Handler
	x11EventsCh chan interface{}
	module.Module
}

func New() *X11 {
	return &X11{
		handler:     newX11Handler(),
		x11EventsCh: nil,
	}
}

func (m *X11) OnInit(cfg *config.Config, rootLogger *zap.Logger) <-chan error {
	ch := make(chan error)

	go func() {
		x11Cfg := cfg.Get().DsX11
		m.Init(m, rootLogger, ModuleName, x11Cfg.ModuleChLen)
		m.x11EventsCh = make(chan interface{}, x11Cfg.X11EventChLen)
		ch <- m.handler.init(m.x11EventsCh, m.ErrorCtx, m.ModuleLogger)
	}()

	return ch
}

func (m *X11) OnStart(ctx context.Context) []*types.HandledChannel {
	m.handler.start()

	hChErr := types.NewHandledChannel(m.ErrorCtx.GetChannel(), m.onError)
	hChX11Events := types.NewHandledChannel(m.x11EventsCh, m.onX11Events)

	return []*types.HandledChannel{hChErr, hChX11Events}
}

func (m *X11) OnFinish() {
	m.handler.stop()
}

func (m *X11) onError(request interface{}) (bool, error) {
	return true, request.(error)
}

func (m *X11) onX11Events(event interface{}) (bool, error) {
	m.handler.onX11Event(event)

	return false, nil
}

func (m *X11) OnRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *subscribeToClipboardCmd:
		m.handler.subscribeToClipboard(r)
	case *writeToClipboardCmd:
		m.handler.writeToClipboard(r)
	case *subscribeToHotkeysCmd:
		m.handler.subscribeToHotkeys(r)
	case *bindHotkeyCmd:
		m.handler.bindHotkey(r)
	case *unbindHotkeyCmd:
		m.handler.unbindHotkey(r)

	default:
		m.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (m *X11) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *subscribeToClipboardCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *writeToClipboardCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *subscribeToHotkeysCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *bindHotkeyCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *unbindHotkeyCmd:
		r.onRequestDefault(m.ModuleLogger, reason)

	default:
		m.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.String("Reason", reason),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (m *X11) SubscribeToClipboard(isPrimary bool, ch chan<- *mime.Data, result api.BoolResult) {
	m.AddToChannel(&subscribeToClipboardCmd{
		isPrimary: isPrimary,
		ch:        ch,
		result:    result,
	})
}

func (m *X11) WriteToClipboard(isPrimary bool, data *mime.Data, result api.BoolResult) {
	m.AddToChannel(&writeToClipboardCmd{
		isPrimary: isPrimary,
		data:      data,
		result:    result,
	})
}

func (m *X11) SubscribeToHotkeys(ch chan<- *shortcut.Hotkey, result api.BoolResult) {
	m.AddToChannel(&subscribeToHotkeysCmd{
		ch:     ch,
		result: result,
	})
}

func (m *X11) BindHotkey(hotkey *shortcut.Hotkey, result api.ErrorCodeResult) {
	m.AddToChannel(&bindHotkeyCmd{
		hotkey: hotkey,
		result: result,
	})
}

func (m *X11) UnbindHotkey(hotkey *shortcut.Hotkey, result api.BoolResult) {
	m.AddToChannel(&unbindHotkeyCmd{
		hotkey: hotkey,
		result: result,
	})
}
