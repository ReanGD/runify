package provider

import (
	"context"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/provider/calculator"
	"github.com/ReanGD/runify/server/provider/desktop_entry"
	"github.com/ReanGD/runify/server/provider/root_list"
	"go.uber.org/zap"
)

type providerHandler struct {
	dataProviders  map[api.ProviderID]*dataProvider
	rpc            api.Rpc
	desktop        api.Desktop
	moduleLogger   *zap.Logger
	rootListLogger *zap.Logger
}

func newProviderHandler() *providerHandler {
	return &providerHandler{
		dataProviders:  make(map[api.ProviderID]*dataProvider),
		rpc:            nil,
		moduleLogger:   nil,
		rootListLogger: nil,
	}
}

func (h *providerHandler) onInit(cfg *config.Config, desktop api.Desktop, rpc api.Rpc, moduleLogger *zap.Logger, rootListLogger *zap.Logger) error {
	h.rpc = rpc
	h.desktop = desktop
	h.moduleLogger = moduleLogger
	h.rootListLogger = rootListLogger
	h.dataProviders[desktopEntryID] = newDataProvider(desktopEntryID, desktop_entry.New(desktop))
	h.dataProviders[calculatorID] = newDataProvider(calculatorID, calculator.New(desktop))

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
	result := api.NewFuncErrorCodeResult(func(result global.Error) {})
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

func (h *providerHandler) openRootList() {
	ch := make(chan api.RootListCtrl, len(h.dataProviders))
	for _, dp := range h.dataProviders {
		dp.makeRootListCtrl(ch)
	}

	ctrls := make(map[api.ProviderID]api.RootListCtrl, len(h.dataProviders))
	for id := range h.dataProviders {
		ctrl := <-ch
		ctrls[api.ProviderID(id)] = ctrl
	}

	ctrl := root_list.NewRLRootListCtrl(ctrls, h.rootListLogger)
	h.rpc.OpenRootList(ctrl)
}

func (h *providerHandler) activate(cmd *activateCmd) {
	h.openRootList()
}
