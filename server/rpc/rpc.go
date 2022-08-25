package rpc

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/logger"
	"github.com/ReanGD/runify/server/provider"
	"github.com/ReanGD/runify/server/system/module"
	"go.uber.org/zap"
)

const ModuleName = "rpc"

type Rpc struct {
	handler *rpcHandler

	module.Module
}

func New() *Rpc {
	return &Rpc{
		handler: newRpcHandler(),
	}
}

func (m *Rpc) OnInit(cfg *config.Config, rootLogger *zap.Logger, provider *provider.Provider) <-chan error {
	ch := make(chan error)

	go func() {
		rpcCfg := &cfg.Get().Rpc
		m.Init(rootLogger, ModuleName, rpcCfg.ChannelLen)

		ch <- m.handler.onInit(rpcCfg, m.ModuleLogger, provider)
	}()

	return ch
}

func (m *Rpc) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		errCh := m.handler.onStart(ctx, wg)
		m.ModuleLogger.Info("Start")

		for {
			if isFinish, err := m.safeRequestLoop(ctx, errCh); isFinish {
				m.handler.onStop()
				ch <- err
				wg.Done()
				return
			}
		}
	}()

	return ch
}

func (m *Rpc) safeRequestLoop(ctx context.Context, errCh <-chan error) (resultIsFinish bool, resultErr error) {
	var request interface{}
	defer func() {
		if recoverResult := recover(); recoverResult != nil {
			if request != nil {
				reason := m.RecoverLog(recoverResult, request)
				resultIsFinish, resultErr = m.onRequestDefault(request, reason)
			} else {
				_ = m.RecoverLog(recoverResult, "unknown request")
			}
		}
	}()

	messageCh := m.GetReadChannel()
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
			m.MessageWasRead()
			if resultIsFinish, resultErr = m.onRequest(request); resultIsFinish {
				return
			}
		}
	}
}

func (m *Rpc) onRequest(request interface{}) (bool, error) {
	switch request.(type) {
	case *showUICmd:
		m.handler.onShowUI()

	default:
		m.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (m *Rpc) onRequestDefault(request interface{}, reason string) (bool, error) {
	switch request.(type) {
	case *showUICmd:
		m.ModuleLogger.Debug("Message is wrong",
			zap.String("RequestType", "ShowUI"),
			zap.String("Reason", reason),
			zap.String("Action", "skip request"))

	default:
		m.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.String("Reason", reason),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (m *Rpc) ShowUI() {
	m.AddToChannel(&showUICmd{})
}
