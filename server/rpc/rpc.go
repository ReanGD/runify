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

func (m *Rpc) OnInit(cfg *config.Config, rootLogger *zap.Logger) <-chan error {
	ch := make(chan error)

	go func() {
		m.Init(rootLogger, ModuleName, cfg.Get().Rpc.ChannelLen)
		uiLogger := rootLogger.With(zap.String("module", "UI"))
		ch <- m.handler.onInit(cfg.Get(), m, uiLogger, m.ModuleLogger)
	}()

	return ch
}

func (m *Rpc) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		errCh := m.handler.onStart(ctx, wg)
		m.ModuleLogger.Info("Start")

		hChErr := module.NewHandledChannel(errCh, m.onError)
		for {
			if isFinish, err := m.SafeRequestLoop(
				ctx, m.onRequest, m.onRequestDefault, []*module.HandledChannel{hChErr}); isFinish {
				m.handler.onStop()
				ch <- err
				wg.Done()
				return
			}
		}
	}()

	return ch
}

func (m *Rpc) onError(request interface{}) (bool, error) {
	return true, request.(error)
}

func (m *Rpc) onRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *serverStartedCmd:
		m.handler.serverStarted()
	case *uiClientConnectedCmd:
		m.handler.uiClientConnected(r.pClient)
	case *uiClientDisconnectedCmd:
		m.handler.uiClientDisconnected(r.id)
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
	switch r := request.(type) {
	case *serverStartedCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *uiClientConnectedCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *uiClientDisconnectedCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *openRootListCmd:
		r.onRequestDefault(m.ModuleLogger, reason)

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
func (m *Rpc) serverStarted() {
	m.AddToChannel(&serverStartedCmd{})
}

func (m *Rpc) uiClientConnected(pClient *protoClient) {
	m.AddToChannel(&uiClientConnectedCmd{pClient: pClient})
}

func (m *Rpc) uiClientDisconnected(id uint32) {
	m.AddToChannel(&uiClientDisconnectedCmd{
		id: id,
	})
}

// Interface functions
func (m *Rpc) OpenRootList(ctrl api.RootListCtrl) {
	m.AddToChannel(&openRootListCmd{ctrl: ctrl})
}
