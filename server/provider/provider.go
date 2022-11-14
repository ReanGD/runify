package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/logger"
	"github.com/ReanGD/runify/server/pb"
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

func (p *Provider) OnInit(cfg *config.Config, x11 module.X11, rootLogger *zap.Logger) <-chan error {
	ch := make(chan error)

	go func() {
		channelLen := cfg.Get().Provider.ChannelLen
		p.Init(rootLogger, ModuleName, channelLen)

		ch <- p.handler.onInit(cfg, x11, p.ModuleLogger)
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
	case *getRootCmd:
		p.handler.getRoot(r)
	case *getActionsCmd:
		p.handler.getActions(r)
	case *executeCmd:
		p.handler.execute(r)

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

func (p *Provider) GetRoot() <-chan []*pb.CardItem {
	ch := make(chan []*pb.CardItem, 1)
	p.AddToChannel(&getRootCmd{
		result: ch,
	})

	return ch
}

func (p *Provider) GetActions(cardID uint64) <-chan *pb.Actions {
	ch := make(chan *pb.Actions, 1)
	p.AddToChannel(&getActionsCmd{
		cardID: cardID,
		result: ch,
	})

	return ch
}

func (p *Provider) Execute(cardID uint64, actionID uint32) <-chan *pb.Result {
	ch := make(chan *pb.Result, 1)
	p.AddToChannel(&executeCmd{
		cardID:   cardID,
		actionID: actionID,
		result:   ch,
	})

	return ch
}
