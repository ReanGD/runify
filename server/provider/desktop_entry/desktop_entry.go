package desktop_entry

import (
	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

type DesktopEntry struct {
	desktop        api.Desktop
	model          *deModel
	actionExecuter *deActionExecuter
	moduleLogger   *zap.Logger
}

func New(desktop api.Desktop) *DesktopEntry {
	return &DesktopEntry{
		desktop:        desktop,
		model:          newDEModel(),
		actionExecuter: newDEActionExecuter(),
		moduleLogger:   nil,
	}
}

func (p *DesktopEntry) GetName() string {
	return "DesktopEntry"
}

func (p *DesktopEntry) OnInit(cfg *config.Config, moduleLogger *zap.Logger, providerID api.ProviderID) error {
	p.moduleLogger = moduleLogger
	if err := p.model.init(providerID, moduleLogger); err != nil {
		return err
	}
	if err := p.actionExecuter.init(cfg, p.desktop, p.model, moduleLogger); err != nil {
		return err
	}

	return nil
}

func (p *DesktopEntry) OnStart(errorCtx *module.ErrorCtx) []*types.HandledChannel {
	return []*types.HandledChannel{}
}

func (p *DesktopEntry) MakeRootListCtrl() api.RootListCtrl {
	return newDERootListCtrl(p.model, p.actionExecuter, p.moduleLogger)
}
