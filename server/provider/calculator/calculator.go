package calculator

import (
	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

type Calculator struct {
	providerID     api.ProviderID
	desktop        api.Desktop
	actionExecuter *calcActionExecuter
	moduleLogger   *zap.Logger
}

func New(desktop api.Desktop) *Calculator {
	return &Calculator{
		providerID:     0,
		desktop:        desktop,
		actionExecuter: newCalcActionExecuter(),
		moduleLogger:   nil,
	}
}

func (p *Calculator) GetName() string {
	return "Calculator"
}

func (p *Calculator) OnInit(cfg *config.Configuration, moduleLogger *zap.Logger, providerID api.ProviderID) error {
	p.providerID = providerID
	p.moduleLogger = moduleLogger

	return p.actionExecuter.init(p.desktop, moduleLogger)
}

func (p *Calculator) OnStart(errorCtx *module.ErrorCtx) []*types.HandledChannel {
	return []*types.HandledChannel{}
}

func (p *Calculator) MakeRootListCtrl() api.RootListCtrl {
	return newCalcRootListCtrl(p.providerID, p.actionExecuter, p.moduleLogger)
}
