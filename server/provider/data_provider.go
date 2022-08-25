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

type dataProviderHandler interface {
	getName() string
	onInit(cfg *config.Config, moduleLogger *zap.Logger, providerID uint64) error
	onStart()
	getRoot() []*pb.Command
}

type dataProvider struct {
	providerID uint64
	handler    dataProviderHandler

	module.Module
}

func newDataProvider(providerID uint64, handler dataProviderHandler) *dataProvider {
	return &dataProvider{
		providerID: providerID,
		handler:    handler,
	}
}

func (p *dataProvider) onInit(cfg *config.Config, rootProviderLogger *zap.Logger) <-chan error {
	ch := make(chan error)
	go func() {
		channelLen := cfg.Get().Provider.SubModuleChannelLen
		p.InitSubmodule(rootProviderLogger, p.handler.getName(), channelLen)

		ch <- p.handler.onInit(cfg, p.ModuleLogger, p.providerID)
	}()

	return ch
}

func (p *dataProvider) onStart(ctx context.Context, wg *sync.WaitGroup, errCh chan<- error) {
	wg.Add(1)
	go func() {
		p.handler.onStart()

		for {
			if isFinish, err := p.safeRequestLoop(ctx); isFinish {
				if err != nil {
					errCh <- err
				}
				wg.Done()
				return
			}
		}
	}()
}

func (p *dataProvider) safeRequestLoop(ctx context.Context) (resultIsFinish bool, resultErr error) {
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
		case request = <-messageCh:
			p.MessageWasRead()
			if resultIsFinish, resultErr = p.onRequest(request); resultIsFinish {
				return
			}
		}
	}
}

func (p *dataProvider) onRequest(request interface{}) (bool, error) {
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

func (p *dataProvider) onRequestDefault(request interface{}, reason string) (bool, error) {
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

func (p *dataProvider) getRoot(result chan<- []*pb.Command) {
	p.AddToChannel(&getRootCmd{
		result: result,
	})
}
