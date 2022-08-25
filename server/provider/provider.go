package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/logger"
	"github.com/ReanGD/runify/server/pb"
	"github.com/ReanGD/runify/server/system/module"
	"go.uber.org/zap"
)

const moduleName = "provider"

type Provider struct {
	handler *providerHandler

	module.Module
}

func New() *Provider {
	return &Provider{
		handler: newProviderHandler(),
	}
}

func (p *Provider) OnInit(cfg *config.Config, rootLogger *zap.Logger) <-chan error {
	ch := make(chan error)

	go func() {
		channelLen := cfg.GetConfiguration().Provider.ChannelLen
		p.Init(rootLogger, moduleName, channelLen)

		ch <- p.handler.onInit(cfg, p.ModuleLogger)
	}()

	return ch
}

func (p *Provider) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		p.ModuleLogger.Info("Start")
		errCh := p.handler.onStart(ctx, wg)

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
	case *getRootCmd:
		r.result <- p.handler.getRoot()

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
	case *getRootCmd:
		r.result <- []*pb.Command{}
		p.ModuleLogger.Debug("Message is wrong",
			zap.String("RequestType", "getRoot"),
			zap.String("Reason", reason),
			zap.String("Action", "skip request"))

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

func (p *Provider) GetRoot() <-chan []*pb.Command {
	ch := make(chan []*pb.Command, 1)
	p.AddToChannel(&getRootCmd{
		result: ch,
	})

	return ch
}