package provider

import (
	"context"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
)

type Provider struct {
	handler *providerHandler
	deps    *dependences

	module.Module
}

func New() (*Provider, string) {
	return &Provider{
		handler: newProviderHandler(),
		deps:    nil,
	}, "provider"
}

func (m *Provider) SetDeps(desktop api.Desktop, de api.XDGDesktopEntry, rpc api.Rpc) {
	m.deps = &dependences{
		desktop: desktop,
		de:      de,
		rpc:     rpc,
	}
}

func (m *Provider) OnInit() (uint32, error) {
	chLen := m.GetConfig().Provider.ChannelLen
	return chLen, m.handler.onInit(m, m.deps)
}

func (m *Provider) OnStart(ctx context.Context) []*types.HandledChannel {
	m.handler.onStart(ctx, m.GetErrorCtx())

	return []*types.HandledChannel{}
}

func (m *Provider) OnFinish() {
	m.handler.onFinish()
}

func (m *Provider) OnRequest(request api.ModuleMsgImpl) (bool, error) {
	switch r := request.(type) {
	case *activateCmd:
		m.handler.activate(r)

	default:
		return m.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (m *Provider) Activate(action *shortcut.Action) {
	m.AddToChannel(&activateCmd{
		action: action,
	})
}
