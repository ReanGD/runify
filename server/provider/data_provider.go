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

func (p *dataProvider) onInit() <-chan error {
	ch := make(chan error)
	go func() {
		cfg := p.GetConfig()
		p.Init(cfg.Provider.SubModuleChannelLen)

		ch <- p.handler.OnInit(cfg, p.GetModuleLogger(), p.providerID)
	}()

	return ch
}

func (p *dataProvider) OnStart(ctx context.Context) []*types.HandledChannel {
	hChs := []*types.HandledChannel{}

	return append(hChs, p.handler.OnStart(p.ErrorCtx)...)
}

func (p *dataProvider) OnFinish() {
}

func (p *dataProvider) OnRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *makeRootListCtrlCmd:
		r.result <- p.handler.MakeRootListCtrl()

	default:
		return p.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (p *dataProvider) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *makeRootListCtrlCmd:
		r.onRequestDefault(p.GetModuleLogger(), reason)

	default:
		return p.OnRequestDefaultUnknownMsg(request, reason)
	}

	return false, nil
}

func (p *dataProvider) makeRootListCtrl(result chan<- api.RootListCtrl) {
	p.AddToChannel(&makeRootListCtrlCmd{
		result: result,
	})
}
