package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

type Calculator struct {
	providerID     uint64
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

func (p *Calculator) OnInit(cfg *config.Config, moduleLogger *zap.Logger, providerID uint64) error {
	p.providerID = providerID
	p.moduleLogger = moduleLogger

	return p.actionExecuter.init(p.desktop, moduleLogger)
}

func (p *Calculator) OnStart() {
}

func (p *Calculator) MakeRootListCtrl() api.RootListCtrl {
	return newCalcRootListCtrl(api.ProviderID(p.providerID), p.actionExecuter, p.moduleLogger)
}

func (p *Calculator) GetRoot() ([]*pb.CardItem, error) {
	return []*pb.CardItem{}, nil
}

func (p *Calculator) GetActions(cardID uint64) ([]*pb.ActionItem, error) {
	return []*pb.ActionItem{}, nil
}

func (p *Calculator) Execute(cardID uint64, actionID uint32) (*pb.Result, error) {
	return nil, errors.New("not found item by cardID")
}
