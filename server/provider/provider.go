package provider

import (
	"context"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
)

const ModuleName = "provider"

type Provider struct {
	handler *providerHandler

	module.Module
}

func New() (*Provider, string) {
	return &Provider{
		handler: newProviderHandler(),
	}, ModuleName
}

func (p *Provider) OnInit(desktop api.Desktop, de api.XDGDesktopEntry, rpc api.Rpc) <-chan error {
	ch := make(chan error)

	go func() {
		cfg := p.GetConfig()
		p.Init(cfg.Provider.ChannelLen)

		ch <- p.handler.onInit(cfg, desktop, de, rpc, p.GetModuleLogger(), p.NewSubmoduleLogger("RootList"))
	}()

	return ch
}

func (p *Provider) OnStart(ctx context.Context) []*types.HandledChannel {
	p.handler.onStart(ctx, p.ErrorCtx)

	return []*types.HandledChannel{}
}

func (p *Provider) OnFinish() {
	p.handler.onFinish()
}

func (p *Provider) OnRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *activateCmd:
		p.handler.activate(r)

	default:
		return p.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (p *Provider) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *activateCmd:
		r.onRequestDefault(p.GetModuleLogger(), reason)

	default:
		return p.OnRequestDefaultUnknownMsg(request, reason)
	}

	return false, nil
}

func (p *Provider) Activate(action *shortcut.Action) {
	p.AddToChannel(&activateCmd{
		action: action,
	})
}
