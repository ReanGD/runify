package desktop

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/logger"
	"go.uber.org/zap"
)

const ModuleName = "desktop"

type Desktop struct {
	handler *handler
	mCtx    *moduleCtx

	module.Module
}

func New() *Desktop {
	return &Desktop{
		handler: newHandler(),
		mCtx:    nil,
	}
}

func (d *Desktop) OnInit(
	cfg *config.Config, ds api.DisplayServer, provider api.Provider, rootLogger *zap.Logger,
) <-chan error {
	ch := make(chan error)

	go func() {
		desktopCfg := cfg.Get().Desktop
		d.Init(rootLogger, ModuleName, desktopCfg.ModuleChLen)
		d.mCtx = newModuleCtx(d, desktopCfg, ds, provider, d.ErrorCtx, d.ModuleLogger)
		ch <- d.handler.init(d.mCtx)
	}()

	return ch
}

func (d *Desktop) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		d.mCtx.setStopCtx(ctx)
		d.handler.start()
		d.ModuleLogger.Info("Start")

		hChErr := module.NewHandledChannel(d.ErrorCtx.GetChannel(), d.onError)
		primaryCh := module.NewHandledChannel(d.mCtx.primaryCh, d.onPrimary)
		clipboardCh := module.NewHandledChannel(d.mCtx.clipboardCh, d.onClipboard)
		hotkeyCh := module.NewHandledChannel(d.mCtx.hotkeyCh, d.onHotkey)

		hChs := []*module.HandledChannel{hChErr, primaryCh, clipboardCh, hotkeyCh}
		for {
			if isFinish, err := d.SafeRequestLoop(
				ctx, d.onRequest, d.onRequestDefault, hChs); isFinish {
				d.handler.stop()
				ch <- err
				wg.Done()
				return
			}
		}
	}()

	return ch
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

func (d *Desktop) onRequest(request interface{}) (bool, error) {
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
		d.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (d *Desktop) onRequestDefault(request interface{}, reason string) (bool, error) {
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
		d.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.String("Reason", reason),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
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
