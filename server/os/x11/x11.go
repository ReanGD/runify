package x11

import (
	"context"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
)

type X11 struct {
	handler     *x11Handler
	x11EventsCh chan interface{}

	module.Module
}

func New() (*X11, string) {
	return &X11{
		handler:     newX11Handler(),
		x11EventsCh: nil,
	}, "x11"
}

func (m *X11) SetDeps() {
}

func (m *X11) OnInit() (uint32, error) {
	cfg := m.GetConfig().DsX11
	m.x11EventsCh = make(chan interface{}, cfg.X11EventChLen)
	return cfg.ModuleChLen, m.handler.init(m, m.x11EventsCh)
}

func (m *X11) OnStart(ctx context.Context) []*types.HandledChannel {
	m.handler.start()

	return []*types.HandledChannel{
		types.NewHandledChannel(m.x11EventsCh, m.onX11Events),
	}
}

func (m *X11) OnFinish() {
	m.handler.stop()
}

func (m *X11) onX11Events(event interface{}) (bool, error) {
	m.handler.onX11Event(event)

	return false, nil
}

func (m *X11) OnRequest(request api.ModuleMsgImpl) (bool, error) {
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
