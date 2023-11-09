package de

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

type XDGDesktopEntry struct {
	handler *handler

	module.Module
}

func New() *XDGDesktopEntry {
	return &XDGDesktopEntry{
		handler: newHandler(),
	}
}

func (d *XDGDesktopEntry) OnInit(
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

func (d *XDGDesktopEntry) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		d.handler.update()
		d.ModuleLogger.Info("Start")

		hChErr := module.NewHandledChannel(d.ErrorCtx.GetChannel(), d.onError)
		for {
			if isFinish, err := d.SafeRequestLoop(
				ctx, d.onRequest, d.onRequestDefault, []*module.HandledChannel{hChErr}); isFinish {
				d.handler.stop()
				ch <- err
				wg.Done()
				return
			}
		}
	}()

	return ch
}

func (d *XDGDesktopEntry) onError(request interface{}) (bool, error) {
	return true, request.(error)
}

func (d *XDGDesktopEntry) onRequest(request interface{}) (bool, error) {
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

func (d *XDGDesktopEntry) onRequestDefault(request interface{}, reason string) (bool, error) {
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

func (d *XDGDesktopEntry) Update() {
	d.AddToChannel(&updateCmd{})
}

func (d *XDGDesktopEntry) Subscribe(ch chan<- types.DesktopEntries, result api.BoolResult) {
	d.AddToChannel(&subscribeCmd{
		ch:     ch,
		result: result,
	})
}
