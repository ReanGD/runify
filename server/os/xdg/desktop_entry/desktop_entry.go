package desktop_entry

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
	"github.com/ReanGD/runify/server/logger"
	"go.uber.org/zap"
)

const ModuleName = "xdg_desktop_entry"

type DesktopEntries struct {
	handler *handler

	module.Module
}

func New() *DesktopEntries {
	return &DesktopEntries{
		handler: newHandler(),
	}
}

func (d *DesktopEntries) OnInit(
	cfg *config.Config, rootLogger *zap.Logger,
) <-chan error {
	ch := make(chan error)

	go func() {
		deCfg := cfg.Get().XDGDesktopEntry
		d.Init(rootLogger, ModuleName, deCfg.ModuleChLen)
		ch <- d.handler.init(d.ModuleLogger)
	}()

	return ch
}

func (d *DesktopEntries) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		d.handler.update()
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

func (d *DesktopEntries) safeRequestLoop(ctx context.Context) (resultIsFinish bool, resultErr error) {
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
		case request = <-messageCh:
			d.MessageWasRead()
			if resultIsFinish, resultErr = d.onRequest(request); resultIsFinish {
				return
			}
		}
	}
}

func (d *DesktopEntries) onRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *updateCmd:
		d.handler.update()
	case *subscribeCmd:
		d.handler.subscribe(r)

	default:
		d.ModuleLogger.Warn("Unknown message received2",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (d *DesktopEntries) onRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *updateCmd:
		r.onRequestDefault(d.ModuleLogger, reason)
	case *subscribeCmd:
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

func (d *DesktopEntries) Update() {
	d.AddToChannel(&updateCmd{})
}

func (d *DesktopEntries) Subscribe(ch chan<- types.DesktopEntries, result api.BoolResult) {
	d.AddToChannel(&subscribeCmd{
		ch:     ch,
		result: result,
	})
}
