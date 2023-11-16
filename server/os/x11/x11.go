package x11

import (
	"context"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
)

const ModuleName = "x11"

type X11 struct {
	handler     *x11Handler
	x11EventsCh chan interface{}

	module.Module
}

func New() (*X11, string) {
	return &X11{
		handler:     newX11Handler(),
		x11EventsCh: nil,
	}, ModuleName
}

func (m *X11) OnInit(cfg *config.Config) <-chan error {
	ch := make(chan error)

	go func() {
		x11Cfg := cfg.Get().DsX11
		m.Init(x11Cfg.ModuleChLen)
		m.x11EventsCh = make(chan interface{}, x11Cfg.X11EventChLen)
		ch <- m.handler.init(m.x11EventsCh, m.ErrorCtx, m.GetModuleLogger())
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
		return m.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (m *X11) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *subscribeToClipboardCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *writeToClipboardCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *subscribeToHotkeysCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *bindHotkeyCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *unbindHotkeyCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)

	default:
		return m.OnRequestDefaultUnknownMsg(request, reason)
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
