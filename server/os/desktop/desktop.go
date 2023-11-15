package desktop

import (
	"context"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

const ModuleName = "desktop"

type Desktop struct {
	handler *handler
	mCtx    *moduleCtx

	module.Module
}

func New() (*Desktop, string) {
	return &Desktop{
		handler: newHandler(),
		mCtx:    nil,
	}, ModuleName
}

func (d *Desktop) OnInit(
	cfg *config.Config, ds api.DisplayServer, provider api.Provider, rootLogger *zap.Logger,
) <-chan error {
	ch := make(chan error)

	go func() {
		desktopCfg := cfg.Get().Desktop
		d.Init(rootLogger, desktopCfg.ModuleChLen)
		d.mCtx = newModuleCtx(d, desktopCfg, ds, provider, d.ErrorCtx, d.ModuleLogger)
		ch <- d.handler.init(d.mCtx)
	}()

	return ch
}

func (d *Desktop) OnStart(ctx context.Context) []*types.HandledChannel {
	d.mCtx.setStopCtx(ctx)
	d.handler.start()

	hChErr := types.NewHandledChannel(d.ErrorCtx.GetChannel(), d.onError)
	primaryCh := types.NewHandledChannel(d.mCtx.primaryCh, d.onPrimary)
	clipboardCh := types.NewHandledChannel(d.mCtx.clipboardCh, d.onClipboard)
	hotkeyCh := types.NewHandledChannel(d.mCtx.hotkeyCh, d.onHotkey)

	return []*types.HandledChannel{hChErr, primaryCh, clipboardCh, hotkeyCh}
}

func (d *Desktop) OnFinish() {
	d.handler.stop()
}

func (d *Desktop) onError(request interface{}) (bool, error) {
	return true, request.(error)
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
		r.onRequestDefault(d.ModuleLogger, reason)
	case *addShortcutCmd:
		r.onRequestDefault(d.ModuleLogger, reason)
	case *removeShortcutCmd:
		r.onRequestDefault(d.ModuleLogger, reason)
	case *removeShortcutWithoutCheckCmd:
		r.onRequestDefault(d.ModuleLogger, reason)

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
