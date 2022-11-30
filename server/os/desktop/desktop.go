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

		for {
			if isFinish, err := d.safeRequestLoop(ctx); isFinish {
				d.handler.stop()
				ch <- err
				wg.Done()
				return
			}
		}
	}()

	return ch
}

func (d *Desktop) safeRequestLoop(ctx context.Context) (resultIsFinish bool, resultErr error) {
	var request interface{}
	defer func() {
		if recoverResult := recover(); recoverResult != nil {
			if request != nil {
				reason := d.RecoverLog(recoverResult, request)
				resultIsFinish, resultErr = d.onRequestDefault(request, reason)
			} else {
				_ = d.RecoverLog(recoverResult, "unknown request")
			}
		}
	}()

	done := ctx.Done()
	errorCh := d.ErrorCtx.GetChannel()
	messageCh := d.GetReadChannel()
	primaryCh := d.mCtx.primaryCh
	clipboardCh := d.mCtx.clipboardCh
	hotkeyCh := d.mCtx.hotkeyCh

	for {
		request = nil
		select {
		case <-done:
			resultIsFinish = true
			resultErr = nil
			return
		case err := <-errorCh:
			resultIsFinish = true
			resultErr = err
			return
		case msg := <-primaryCh:
			d.handler.onClipboardMsg(true, msg)
		case msg := <-clipboardCh:
			d.handler.onClipboardMsg(false, msg)
		case msg := <-hotkeyCh:
			d.handler.onHotkeyMsg(msg)
		case request = <-messageCh:
			d.MessageWasRead()
			if resultIsFinish, resultErr = d.onRequest(request); resultIsFinish {
				return
			}
		}
	}
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
