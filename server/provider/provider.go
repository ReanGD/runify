package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/logger"
	"go.uber.org/zap"
)

const ModuleName = "provider"

type Provider struct {
	handler *providerHandler

	module.Module
}

func New() *Provider {
	return &Provider{
		handler: newProviderHandler(),
	}
}

func (p *Provider) OnInit(cfg *config.Config, desktop api.Desktop, rpc api.Rpc, rootLogger *zap.Logger) <-chan error {
	ch := make(chan error)

	go func() {
		channelLen := cfg.Get().Provider.ChannelLen
		p.Init(rootLogger, ModuleName, channelLen)

		ch <- p.handler.onInit(cfg, desktop, rpc, p.ModuleLogger, p.NewSubmoduleLogger(p.ModuleLogger, "RootList"))
	}()

	return ch
}

func (p *Provider) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		errCh := p.handler.onStart(ctx, wg)
		p.ModuleLogger.Info("Start")

		for {
			if isFinish, err := p.safeRequestLoop(ctx, errCh); isFinish {
				ch <- err
				wg.Done()
				return
			}
		}
	}()

	return ch
}

func (p *Provider) safeRequestLoop(ctx context.Context, errCh <-chan error) (resultIsFinish bool, resultErr error) {
	var request interface{}
	defer func() {
		if recoverResult := recover(); recoverResult != nil {
			if request != nil {
				reason := p.RecoverLog(recoverResult, request)
				resultIsFinish, resultErr = p.onRequestDefault(request, reason)
			} else {
				_ = p.RecoverLog(recoverResult, "unknown request")
			}
		}
	}()

	messageCh := p.GetReadChannel()
	done := ctx.Done()

	for {
		request = nil
		select {
		case <-done:
			resultIsFinish = true
			resultErr = nil
			return
		case resultErr = <-errCh:
			resultIsFinish = true
			return
		case request = <-messageCh:
			p.MessageWasRead()
			if resultIsFinish, resultErr = p.onRequest(request); resultIsFinish {
				return
			}
		}
	}
}

func (p *Provider) onRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *activateCmd:
		p.handler.activate(r)

	default:
		p.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (p *Provider) onRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *activateCmd:
		r.onRequestDefault(p.ModuleLogger, reason)

	default:
		p.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.String("Reason", reason),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (p *Provider) Activate(action *shortcut.Action) {
	p.AddToChannel(&activateCmd{
		action: action,
	})
}
