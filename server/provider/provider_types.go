package provider

import (
	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

const (
	desktopEntryID api.ProviderID = 1
	calculatorID   api.ProviderID = 2
	linksID        api.ProviderID = 3
)

type dependences struct {
	desktop api.Desktop
	de      api.XDGDesktopEntry
	rpc     api.Rpc
}

type dataProviderHandler interface {
	GetName() string
	OnInit(cfg *config.Configuration, moduleLogger *zap.Logger, providerID api.ProviderID) error
	OnStart(errorCtx *module.ErrorCtx) []*types.HandledChannel
	MakeRootListCtrl() api.RootListCtrl
}

type makeRootListCtrlCmd struct {
	result chan<- api.RootListCtrl
}

func (c *makeRootListCtrlCmd) onRequestDefault(logger *zap.Logger, reason string) {
	c.result <- nil
	logger.Warn("Process message finished with error",
		zap.String("Request", "makeRootListCtrl"),
		zap.String("Reason", reason),
		zap.String("Action", "return nil"))
}

type activateCmd struct {
	action *shortcut.Action
}

func (c *activateCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "activate"),
		zap.String("Reason", reason),
		c.action.ZapField(),
		zap.String("Action", "do nothing"),
	)
}
