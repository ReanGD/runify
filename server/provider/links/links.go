package links

import (
	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
)

type Links struct {
	providerID     api.ProviderID
	desktop        api.Desktop
	model          *model
	actionExecuter *actionExecuter
	moduleLogger   *zap.Logger
}

func New(desktop api.Desktop) *Links {
	return &Links{
		providerID:     0,
		desktop:        desktop,
		model:          newModel(),
		actionExecuter: newActionExecuter(),
		moduleLogger:   nil,
	}
}

func (p *Links) GetName() string {
	return "Bookmark"
}

func (p *Links) OnInit(cfg *config.Configuration, moduleLogger *zap.Logger, providerID api.ProviderID) error {
	p.providerID = providerID
	p.moduleLogger = moduleLogger
	if err := p.model.init(providerID, p.GetName(), moduleLogger); err != nil {
		return err
	}

	return p.actionExecuter.init(p.desktop, moduleLogger)
}

func (p *Links) OnStart(errorCtx *module.ErrorCtx) []*types.HandledChannel {
	p.model.start()

	return []*types.HandledChannel{}
}

func (p *Links) MakeRootListCtrl() api.RootListCtrl {
	return newLinksRootListCtrl(p.providerID, p.model, p.actionExecuter, p.moduleLogger)
}
