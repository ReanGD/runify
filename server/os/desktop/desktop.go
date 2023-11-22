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

func (d *Desktop) SetDeps(ds api.DisplayServer, provider api.Provider) {
	d.deps = &dependences{
		ds:       ds,
		provider: provider,
	}
}

func (d *Desktop) OnInit() (uint32, error) {
	cfg := d.GetConfig().Desktop
	d.mCtx = newModuleCtx(d, cfg, d.deps)

	return cfg.ModuleChLen, d.handler.init(d.mCtx)
}

func (d *Desktop) OnStart(ctx context.Context) []*types.HandledChannel {
	d.mCtx.setStopCtx(ctx)
	d.handler.start()

	return []*types.HandledChannel{
		types.NewHandledChannel(d.mCtx.primaryCh, d.onPrimary),
		types.NewHandledChannel(d.mCtx.clipboardCh, d.onClipboard),
		types.NewHandledChannel(d.mCtx.hotkeyCh, d.onHotkey),
	}
}

func (d *Desktop) OnFinish() {
	d.handler.stop()
}

func (d *Desktop) onPrimary(request interface{}) (bool, error) {
	d.handler.onClipboardMsg(true, request.(*mime.Data))

	return false, nil
}

func (d *Desktop) onClipboard(request interface{}) (bool, error) {
	d.handler.onClipboardMsg(false, request.(*mime.Data))

	return false, nil
}

func (d *Desktop) onHotkey(request interface{}) (bool, error) {
	d.handler.onHotkeyMsg(request.(*shortcut.Hotkey))

	return false, nil
}

func (d *Desktop) OnRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *writeToClipboardCmd:
		d.handler.writeToClipboard(r)
	case *addShortcutCmd:
		d.handler.addShortcut(r)
	case *removeShortcutCmd:
		d.handler.removeShortcut(r)
	case *removeShortcutWithoutCheckCmd:
		d.handler.removeShortcutWithoutCheck(r)

	default:
		return d.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (d *Desktop) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *writeToClipboardCmd:
		r.onRequestDefault(d.GetModuleLogger(), reason)
	case *addShortcutCmd:
		r.onRequestDefault(d.GetModuleLogger(), reason)
	case *removeShortcutCmd:
		r.onRequestDefault(d.GetModuleLogger(), reason)
	case *removeShortcutWithoutCheckCmd:
		r.onRequestDefault(d.GetModuleLogger(), reason)

	default:
		return d.OnRequestDefaultUnknownMsg(request, reason)
	}

	return false, nil
}

func (d *Desktop) WriteToClipboard(isPrimary bool, data *mime.Data, result api.BoolResult) {
	d.AddToChannel(&writeToClipboardCmd{
		isPrimary: isPrimary,
		data:      data,
		result:    result,
	})
}

func (d *Desktop) AddShortcut(action *shortcut.Action, hotkey *shortcut.Hotkey, result api.ErrorCodeResult) {
	d.AddToChannel(&addShortcutCmd{
		action: action,
		hotkey: hotkey,
		result: result,
	})
}

func (d *Desktop) RemoveShortcut(action *shortcut.Action, result api.VoidResult) {
	d.AddToChannel(&removeShortcutCmd{
		action: action,
		result: result,
	})
}

func (d *Desktop) removeShortcutWithoutCheck(action *shortcut.Action, hotkey *shortcut.Hotkey) {
	d.AddToChannel(&removeShortcutWithoutCheckCmd{
		action: action,
		hotkey: hotkey,
	})
}
