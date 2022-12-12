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
	"github.com/ReanGD/runify/server/logger"
	"go.uber.org/zap"
)

type dataProvider struct {
	providerID api.ProviderID
	handler    dataProviderHandler

	module.Module
}

func newDataProvider(providerID api.ProviderID, handler dataProviderHandler) *dataProvider {
	return &dataProvider{
		providerID: providerID,
		handler:    handler,
	}
}

func (p *dataProvider) onInit(cfg *config.Config, rootProviderLogger *zap.Logger) <-chan error {
	ch := make(chan error)
	go func() {
		channelLen := cfg.Get().Provider.SubModuleChannelLen
		p.InitSubmodule(rootProviderLogger, p.handler.GetName(), channelLen)

		ch <- p.handler.OnInit(cfg, p.ModuleLogger, p.providerID)
	}()

	return ch
}

func (p *dataProvider) onStart(ctx context.Context, wg *sync.WaitGroup, errCh chan<- error) {
	wg.Add(1)
	go func() {
		p.handler.OnStart()

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
	case *makeRootListCtrlCmd:
		r.result <- p.handler.MakeRootListCtrl()

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
	case *makeRootListCtrlCmd:
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

func (p *dataProvider) makeRootListCtrl(result chan<- api.RootListCtrl) {
	p.AddToChannel(&makeRootListCtrlCmd{
		result: result,
	})
}
