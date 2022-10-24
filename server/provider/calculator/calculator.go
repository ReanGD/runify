package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

type Calculator struct {
	providerID   uint64
	moduleLogger *zap.Logger
}

func New() *Calculator {
	return &Calculator{
		providerID:   0,
		moduleLogger: nil,
	}
}

func (p *Calculator) GetName() string {
	return "Calculator"
}

func (p *Calculator) OnInit(cfg *config.Config, moduleLogger *zap.Logger, providerID uint64) error {
	p.providerID = providerID
	p.moduleLogger = moduleLogger
	return nil
}

func (p *Calculator) OnStart() {
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
