package rpc

import (
	"context"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
)

type Rpc struct {
	wg      *sync.WaitGroup
	handler *rpcHandler

	module.Module
}

func New() (*Rpc, string) {
	return &Rpc{
		wg:      &sync.WaitGroup{},
		handler: newRpcHandler(),
	}, "rpc"
}

func (m *Rpc) SetDeps() {
}

func (m *Rpc) OnInit() (uint32, error) {
	return m.GetConfig().Rpc.ChannelLen, m.handler.onInit(m)
}

func (m *Rpc) OnStart(ctx context.Context) []*types.HandledChannel {
	m.handler.onStart(ctx, m.wg, m.GetErrorCtx())

	return []*types.HandledChannel{}
}

func (m *Rpc) OnFinish() {
	m.handler.onStop()
	m.wg.Wait()
}

func (m *Rpc) OnRequest(request api.ModuleMsgImpl) (bool, error) {
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
