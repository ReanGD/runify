package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
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
		p.Init(p, rootLogger, ModuleName, channelLen)

		ch <- p.handler.onInit(cfg, desktop, rpc, p.ModuleLogger, p.NewSubmoduleLogger(p.ModuleLogger, "RootList"))
	}()

	return ch
}

func (p *Provider) OnStart(ctx context.Context) []*types.HandledChannel {
	p.handler.onStart(ctx)

	hChErr := types.NewHandledChannel(p.handler.getErrCh(), p.onError)
	return []*types.HandledChannel{hChErr}
}

func (p *Provider) OnFinish() {
	p.handler.onFinish()
}

func (p *Provider) onError(request interface{}) (bool, error) {
	return true, request.(error)
}

func (p *Provider) OnRequest(request interface{}) (bool, error) {
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

func (p *Provider) OnRequestDefault(request interface{}, reason string) (bool, error) {
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
