package desktop

import (
	"context"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
)

type Desktop struct {
	handler *handler
	deps    *dependences
	mCtx    *moduleCtx

	module.Module
}

func New() (*Desktop, string) {
	return &Desktop{
		handler: newHandler(),
		deps:    nil,
		mCtx:    nil,
	}, "desktop"
}

func (m *Desktop) SetDeps(ds api.DisplayServer, provider api.Provider) {
	m.deps = &dependences{
		ds:       ds,
		provider: provider,
	}
}

func (m *Desktop) OnInit() (uint32, error) {
	cfg := m.GetConfig().Desktop
	m.mCtx = newModuleCtx(m, cfg, m.deps)

	return cfg.ModuleChLen, m.handler.init(m.mCtx)
}

func (m *Desktop) OnStart(ctx context.Context) []*types.HandledChannel {
	m.mCtx.setStopCtx(ctx)
	m.handler.start()

	return []*types.HandledChannel{
		types.NewHandledChannel(m.mCtx.primaryCh, m.onPrimary),
		types.NewHandledChannel(m.mCtx.clipboardCh, m.onClipboard),
		types.NewHandledChannel(m.mCtx.hotkeyCh, m.onHotkey),
	}
}

func (m *Desktop) OnFinish() {
	m.handler.stop()
}

func (m *Desktop) onPrimary(request interface{}) (bool, error) {
	m.handler.onClipboardMsg(true, request.(*mime.Data))

	return false, nil
}

func (m *Desktop) onClipboard(request interface{}) (bool, error) {
	m.handler.onClipboardMsg(false, request.(*mime.Data))

	return false, nil
}

func (m *Desktop) onHotkey(request interface{}) (bool, error) {
	m.handler.onHotkeyMsg(request.(*shortcut.Hotkey))

	return false, nil
}

func (m *Desktop) OnRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *writeToClipboardCmd:
		m.handler.writeToClipboard(r)
	case *addShortcutCmd:
		m.handler.addShortcut(r)
	case *removeShortcutCmd:
		m.handler.removeShortcut(r)
	case *removeShortcutWithoutCheckCmd:
		m.handler.removeShortcutWithoutCheck(r)

	default:
		return m.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (m *Desktop) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *writeToClipboardCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *addShortcutCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *removeShortcutCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *removeShortcutWithoutCheckCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)

	default:
		return m.OnRequestDefaultUnknownMsg(request, reason)
	}

	return false, nil
}

func (m *Desktop) WriteToClipboard(isPrimary bool, data *mime.Data, result api.BoolResult) {
	m.AddToChannel(&writeToClipboardCmd{
		isPrimary: isPrimary,
		data:      data,
		result:    result,
	})
}

func (m *Desktop) AddShortcut(action *shortcut.Action, hotkey *shortcut.Hotkey, result api.ErrorCodeResult) {
	m.AddToChannel(&addShortcutCmd{
		action: action,
		hotkey: hotkey,
		result: result,
	})
}

func (m *Desktop) RemoveShortcut(action *shortcut.Action, result api.VoidResult) {
	m.AddToChannel(&removeShortcutCmd{
		action: action,
		result: result,
	})
}

func (m *Desktop) removeShortcutWithoutCheck(action *shortcut.Action, hotkey *shortcut.Hotkey) {
	m.AddToChannel(&removeShortcutWithoutCheckCmd{
		action: action,
		hotkey: hotkey,
	})
}
