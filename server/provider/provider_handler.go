package provider

import (
	"context"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/pb"
	"github.com/ReanGD/runify/server/provider/calculator"
	"github.com/ReanGD/runify/server/provider/desktop_entry"
	"go.uber.org/zap"
)

type providerHandler struct {
	dataProviders map[uint64]*dataProvider
	rpc           module.Rpc
	desktop       module.Desktop
	moduleLogger  *zap.Logger
}

func newProviderHandler() *providerHandler {
	return &providerHandler{
		dataProviders: make(map[uint64]*dataProvider),
		rpc:           nil,
		moduleLogger:  nil,
	}
}

func (h *providerHandler) onInit(cfg *config.Config, desktop module.Desktop, rpc module.Rpc, moduleLogger *zap.Logger) error {
	h.rpc = rpc
	h.desktop = desktop
	h.moduleLogger = moduleLogger
	h.dataProviders[desktopEntryID] = newDataProvider(desktopEntryID, desktop_entry.New(desktop))
	h.dataProviders[calculatorID] = newDataProvider(calculatorID, calculator.New())

	dpChans := make([]<-chan error, 0, len(h.dataProviders))
	for _, dp := range h.dataProviders {
		dpChans = append(dpChans, dp.onInit(cfg, moduleLogger))
	}
	for _, dpChan := range dpChans {
		if err := <-dpChan; err != nil {
			return err
		}
	}
	hotkey, err := shortcut.NewHotkey(cfg.Get().Shortcuts.Root)
	if err != nil {
		return err
	}
	result := module.NewFuncErrorCodeResult(func(result global.Error) {})
	h.desktop.AddShortcut(shortcut.NewAction("Show UI"), hotkey, result)

	return nil
}

func (h *providerHandler) onStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	errCh := make(chan error, 1)

	for _, dp := range h.dataProviders {
		dp.onStart(ctx, wg, errCh)
	}

	return errCh
}

func (h *providerHandler) getRoot(cmd *getRootCmd) {
	chans := make([]<-chan []*pb.CardItem, 0, len(h.dataProviders))
	for _, dp := range h.dataProviders {
		chans = append(chans, dp.getRoot())
	}

	result := []*pb.CardItem{}
	for _, ch := range chans {
		result = append(result, <-ch...)
	}

	cmd.result <- result
}

func (h *providerHandler) getActions(cmd *getActionsCmd) {
	providerID := cmd.cardID & providerIDMask
	provider, ok := h.dataProviders[providerID]
	if !ok {
		cmd.onRequestDefault(h.moduleLogger, "Not found provider")
	} else {
		data := <-provider.getActions(cmd.cardID)
		cmd.result <- data
	}
}

func (h *providerHandler) execute(cmd *executeCmd) {
	providerID := cmd.cardID & providerIDMask
	provider, ok := h.dataProviders[providerID]
	if !ok {
		cmd.onRequestDefault(h.moduleLogger, "Not found provider")
	} else {
		data := <-provider.execute(cmd.cardID, cmd.actionID)
		cmd.result <- data
	}
}

func (h *providerHandler) activate(cmd *activateCmd) {
	h.rpc.ShowUI()
}
