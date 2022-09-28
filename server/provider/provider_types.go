package provider

import (
	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

const (
	desktopEntryID = uint64(1) << 32
	providerIDMask = uint64(0xFFFF) << 32
)

type dataProviderHandler interface {
	GetName() string
	OnInit(cfg *config.Config, moduleLogger *zap.Logger, providerID uint64) error
	OnStart()
	GetRoot() ([]*pb.CardItem, error)
	GetActions(cardID uint64) ([]*pb.ActionItem, error)
	Execute(cardID uint64, actionID uint32) (*pb.Result, error)
}

type getRootCmd struct {
	result chan<- []*pb.CardItem
}

func (c *getRootCmd) onRequestDefault(logger *zap.Logger, reason string) {
	c.result <- []*pb.CardItem{}
	logger.Warn("Process message finished with error",
		zap.String("Request", "GetRoot"),
		zap.String("Reason", reason),
		zap.String("Action", "return empty result"))
}

type getActionsCmd struct {
	cardID uint64
	result chan<- *pb.Actions
}

func (c *getActionsCmd) onRequestDefault(logger *zap.Logger, reason string) {
	c.result <- &pb.Actions{
		Items: []*pb.ActionItem{},
	}
	logger.Warn("Process message finished with error",
		zap.String("Request", "GetActions"),
		zap.Uint64("CardID", c.cardID),
		zap.String("Reason", reason),
		zap.String("Action", "return empty result"))
}

type executeCmd struct {
	cardID   uint64
	actionID uint32
	result   chan<- *pb.Result
}

func (c *executeCmd) onRequestDefault(logger *zap.Logger, reason string) {
	c.result <- &pb.Result{
		Payload: &pb.Result_Empty{},
	}
	logger.Warn("Process message finished with error",
		zap.String("Request", "Execute"),
		zap.Uint64("CardID", c.cardID),
		zap.Uint32("ActionID", c.actionID),
		zap.String("Reason", reason),
		zap.String("Action", "return empty result"))
}

func (c *executeCmd) executeError(logger *zap.Logger, reason string) {
	logger.Warn("Failed execute desktop entry",
		zap.String("Request", "Execute"),
		zap.Uint64("CardID", c.cardID),
		zap.Uint32("ActionID", c.actionID),
		zap.String("Reason", reason),
		zap.String("Action", "return empty result"))
}
