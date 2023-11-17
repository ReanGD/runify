package rpc

import (
	"context"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

const ModuleName = "rpc"

type Rpc struct {
	wg      *sync.WaitGroup
	handler *rpcHandler

	module.Module
}

func New() (*Rpc, string) {
	return &Rpc{
		wg:      &sync.WaitGroup{},
		handler: newRpcHandler(),
	}, ModuleName
}

func (m *Rpc) OnInit() <-chan error {
	ch := make(chan error)

	go func() {
		cfg := m.GetConfig()
		m.Init(cfg.Rpc.ChannelLen)
		uiLogger := m.GetRootLogger().With(zap.String("module", "UI"))
		ch <- m.handler.onInit(cfg, m, uiLogger, m.GetModuleLogger())
	}()

	return ch
}

func (m *Rpc) OnStart(ctx context.Context) []*types.HandledChannel {
	m.handler.onStart(ctx, m.wg, m.ErrorCtx)

	return []*types.HandledChannel{}
}

func (m *Rpc) OnFinish() {
	m.handler.onStop()
	m.wg.Wait()
}

func (m *Rpc) OnRequest(request interface{}) (bool, error) {
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
		return m.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (m *Rpc) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *serverStartedCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *uiClientConnectedCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *uiClientDisconnectedCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)
	case *openRootListCmd:
		r.onRequestDefault(m.GetModuleLogger(), reason)

	default:
		return m.OnRequestDefaultUnknownMsg(request, reason)
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
