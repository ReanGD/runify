package portal

import (
	"context"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
)

type Portal struct {
	provider api.Provider
	handler  *handler

	module.Module
}

func New() (*Portal, string) {
	return &Portal{
		provider: nil,
		handler:  newHandler(),
	}, "xdg_desktop_portal"
}

func (m *Portal) SetDeps(provider api.Provider) {
	m.provider = provider
}

func (m *Portal) OnInit() (uint32, error) {
	cfg := m.GetConfig().XDGDesktopPortal
	return cfg.ModuleChLen, m.handler.init(m, cfg)
}

func (m *Portal) OnStart(ctx context.Context) []*types.HandledChannel {
	return m.handler.start()
}

func (m *Portal) OnFinish() {
	m.handler.stop()
}

func (m *Portal) OnRequest(request api.ModuleMsgImpl) (bool, error) {
	// switch r := request.(type) {
	// case *subscribeToClipboardCmd:
	// 	m.handler.subscribeToClipboard(r)
	// case *writeToClipboardCmd:
	// 	m.handler.writeToClipboard(r)
	// case *subscribeToHotkeysCmd:
	// 	m.handler.subscribeToHotkeys(r)
	// case *bindHotkeyCmd:
	// 	m.handler.bindHotkey(r)
	// case *unbindHotkeyCmd:
	// 	m.handler.unbindHotkey(r)

	// default:
	// 	return m.OnRequestUnknownMsg(request)
	// }

	return false, nil
}

// func (m *Portal) SubscribeToClipboard(isPrimary bool, ch chan<- *mime.Data, result api.BoolResult) {
// 	m.AddToChannel(&subscribeToClipboardCmd{
// 		isPrimary: isPrimary,
// 		ch:        ch,
// 		result:    result,
// 	})
// }

// func (m *Portal) WriteToClipboard(isPrimary bool, data *mime.Data, result api.BoolResult) {
// 	m.AddToChannel(&writeToClipboardCmd{
// 		isPrimary: isPrimary,
// 		data:      data,
// 		result:    result,
// 	})
// }

// func (m *Portal) SubscribeToHotkeys(ch chan<- *shortcut.Hotkey, result api.BoolResult) {
// 	m.AddToChannel(&subscribeToHotkeysCmd{
// 		ch:     ch,
// 		result: result,
// 	})
// }

// func (m *Portal) BindHotkey(hotkey *shortcut.Hotkey, result api.ErrorCodeResult) {
// 	m.AddToChannel(&bindHotkeyCmd{
// 		hotkey: hotkey,
// 		result: result,
// 	})
// }

// func (m *Portal) UnbindHotkey(hotkey *shortcut.Hotkey, result api.BoolResult) {
// 	m.AddToChannel(&unbindHotkeyCmd{
// 		hotkey: hotkey,
// 		result: result,
// 	})
// }
