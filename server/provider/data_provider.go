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
		if data, err := p.handler.getRoot(); err != nil {
			r.onRequestDefault(p.ModuleLogger, err.Error())
		} else {
			r.result <- data
		}
	case *getActionsCmd:
		if data, err := p.handler.getActions(r.cardID); err != nil {
			r.onRequestDefault(p.ModuleLogger, err.Error())
		} else {
			r.result <- &pb.Actions{
				Items: data,
			}
		}
	case *executeCmd:
		if data, err := p.handler.execute(r.cardID, r.actionID); err != nil {
			r.onRequestDefault(p.ModuleLogger, err.Error())
		} else {
			r.result <- data
		}

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
		r.onRequestDefault(p.ModuleLogger, reason)
	case *getActionsCmd:
		r.onRequestDefault(p.ModuleLogger, reason)
	case *executeCmd:
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

func (p *dataProvider) getRoot() <-chan []*pb.CardItem {
	ch := make(chan []*pb.CardItem, 1)
	p.AddToChannel(&getRootCmd{
		result: ch,
	})

	return ch
}

func (p *dataProvider) getActions(cardID uint64) <-chan *pb.Actions {
	ch := make(chan *pb.Actions, 1)
	p.AddToChannel(&getActionsCmd{
		cardID: cardID,
		result: ch,
	})

	return ch
}

func (p *dataProvider) execute(cardID uint64, actionID uint32) <-chan *pb.Result {
	ch := make(chan *pb.Result, 1)
	p.AddToChannel(&executeCmd{
		cardID:   cardID,
		actionID: actionID,
		result:   ch,
	})

	return ch
}
