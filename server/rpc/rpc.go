package rpc

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

func (m *Rpc) OnInit(cfg *config.Config, provider api.Provider, rootLogger *zap.Logger) <-chan error {
	ch := make(chan error)

	go func() {
		m.Init(rootLogger, ModuleName, cfg.Get().Rpc.ChannelLen)
		uiLogger := rootLogger.With(zap.String("module", "UI"))
		ch <- m.handler.onInit(cfg.Get(), m, provider, uiLogger, m.ModuleLogger)
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
	switch r := request.(type) {
	case *uiClientConnectedCmd:
		m.handler.uiClientConnected(r.pClient)
	case *uiClientDisconnectedCmd:
		m.handler.uiClientDisconnected()
	case *showUICmd:
		m.handler.showUI()
	case *openRootListCmd:
		m.handler.openRootList(r.ctrl)

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
	case *uiClientConnectedCmd:
		m.ModuleLogger.Debug("Message is wrong",
			zap.String("RequestType", "UiClientConnected"),
			zap.String("Reason", reason),
			zap.String("Action", "skip request"))

	case *uiClientDisconnectedCmd:
		m.ModuleLogger.Debug("Message is wrong",
			zap.String("RequestType", "UiClientDisconnected"),
			zap.String("Reason", reason),
			zap.String("Action", "skip request"))

	case *showUICmd:
		m.ModuleLogger.Debug("Message is wrong",
			zap.String("RequestType", "ShowUI"),
			zap.String("Reason", reason),
			zap.String("Action", "skip request"))

	case *openRootListCmd:
		m.ModuleLogger.Debug("Message is wrong",
			zap.String("RequestType", "OpenRootList"),
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

// Inner functions
func (m *Rpc) uiClientConnected(pClient *protoClient) {
	m.AddToChannel(&uiClientConnectedCmd{pClient: pClient})
}

func (m *Rpc) uiClientDisconnected() {
	m.AddToChannel(&uiClientDisconnectedCmd{})
}

// Interface
func (m *Rpc) ShowUI() {
	m.AddToChannel(&showUICmd{})
}

func (m *Rpc) OpenRootList(ctrl api.RootListCtrl) {
	m.AddToChannel(&openRootListCmd{ctrl: ctrl})
}
