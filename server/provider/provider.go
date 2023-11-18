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
	deps    *dependences

	module.Module
}

func New() (*Provider, string) {
	return &Provider{
		handler: newProviderHandler(),
		deps:    nil,
	}, ModuleName
}

func (p *Provider) SetDeps(desktop api.Desktop, de api.XDGDesktopEntry, rpc api.Rpc) {
	p.deps = &dependences{
		desktop: desktop,
		de:      de,
		rpc:     rpc,
	}
}

func (p *Provider) OnInit() (uint32, error) {
	chLen := p.GetConfig().Provider.ChannelLen
	return chLen, p.handler.onInit(p, p.deps)
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
