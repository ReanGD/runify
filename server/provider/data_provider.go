package provider

import (
	"context"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
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

func (m *dataProvider) SetDeps() {
}

func (m *dataProvider) OnInit() (uint32, error) {
	cfg := m.GetConfig()
	chLen := cfg.Provider.ChannelLen

	return chLen, m.handler.OnInit(cfg, m.GetModuleLogger(), m.providerID)
}

func (m *dataProvider) OnStart(ctx context.Context) []*types.HandledChannel {
	hChs := []*types.HandledChannel{}

	return append(hChs, m.handler.OnStart(m.ErrorCtx)...)
}

func (m *dataProvider) OnFinish() {
}

func (m *dataProvider) OnRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *makeRootListCtrlCmd:
		r.result <- m.handler.MakeRootListCtrl()

	default:
		return m.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (m *dataProvider) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *makeRootListCtrlCmd:
		r.OnRequestDefault(m.GetModuleLogger(), reason)

	default:
		return m.OnRequestDefaultUnknownMsg(request, reason)
	}

	return false, nil
}

func (m *dataProvider) makeRootListCtrl(result chan<- api.RootListCtrl) {
	m.AddToChannel(&makeRootListCtrlCmd{
		result: result,
	})
}
