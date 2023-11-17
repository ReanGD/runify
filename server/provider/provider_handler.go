package provider

import (
	"context"
	"reflect"
	"sync"

	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/provider/calculator"
	"github.com/ReanGD/runify/server/provider/desktop_entry"
	"github.com/ReanGD/runify/server/provider/links"
	"github.com/ReanGD/runify/server/provider/root_list"
)

type providerHandler struct {
	cfg            *config.Configuration
	dataProviders  map[api.ProviderID]*dataProvider
	rpc            api.Rpc
	desktop        api.Desktop
	wg             *sync.WaitGroup
	doneErrWaitCh  chan struct{}
	moduleLogger   *zap.Logger
	rootListLogger *zap.Logger
}

func newProviderHandler() *providerHandler {
	return &providerHandler{
		cfg:            nil,
		dataProviders:  make(map[api.ProviderID]*dataProvider),
		rpc:            nil,
		desktop:        nil,
		wg:             &sync.WaitGroup{},
		doneErrWaitCh:  make(chan struct{}),
		moduleLogger:   nil,
		rootListLogger: nil,
	}
}

func (h *providerHandler) addProvider(providerID api.ProviderID, handler dataProviderHandler) {
	dp := newDataProvider(providerID, handler)
	h.dataProviders[providerID] = dp
	dp.Create(dp, handler.GetName(), module.SUB_MODULE, h.cfg, h.moduleLogger)
}

func (h *providerHandler) onInit(
	cfg *config.Configuration,
	desktop api.Desktop,
	de api.XDGDesktopEntry,
	rpc api.Rpc,
	moduleLogger *zap.Logger,
	rootListLogger *zap.Logger,
) error {
	h.cfg = cfg
	h.rpc = rpc
	h.desktop = desktop
	h.moduleLogger = moduleLogger
	h.rootListLogger = rootListLogger

	h.addProvider(desktopEntryID, desktop_entry.New(desktop, de))
	h.addProvider(calculatorID, calculator.New(desktop))
	h.addProvider(linksID, links.New(desktop))

	dpChans := make([]<-chan error, 0, len(h.dataProviders))
	for _, dp := range h.dataProviders {
		dpChans = append(dpChans, dp.onInit())
	}
	for _, dpChan := range dpChans {
		if err := <-dpChan; err != nil {
			return err
		}
	}
	hotkey, err := shortcut.NewHotkey(cfg.Shortcuts.Root)
	if err != nil {
		return err
	}
	result := api.NewFuncErrorCodeResult(func(result global.Error) {})
	h.desktop.AddShortcut(shortcut.NewAction("Show UI"), hotkey, result)

	return nil
}

func (h *providerHandler) onStart(ctx context.Context, errorCtx *module.ErrorCtx) {
	cases := make([]reflect.SelectCase, len(h.dataProviders)+1)

	caseNum := 0
	for _, dp := range h.dataProviders {
		ch := dp.Start(ctx, h.wg)
		cases[caseNum] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		caseNum++
	}

	cases[caseNum] = reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(h.doneErrWaitCh),
	}

	go func() {
		h.wg.Add(1)
		defer h.wg.Done()

		if _, recv, recvOk := reflect.Select(cases); !recvOk {
			errorCtx.SendError(recv.Interface().(error))
		}
	}()
}

func (h *providerHandler) onFinish() {
	if h.doneErrWaitCh != nil {
		close(h.doneErrWaitCh)
		h.doneErrWaitCh = nil
	}
	h.wg.Wait()
}

func (h *providerHandler) openRootList() {
	chans := make(map[api.ProviderID]chan api.RootListCtrl, len(h.dataProviders))
	for id, dp := range h.dataProviders {
		ch := make(chan api.RootListCtrl, 1)
		chans[id] = ch
		dp.makeRootListCtrl(ch)
	}

	ctrls := make(map[api.ProviderID]api.RootListCtrl, len(h.dataProviders))
	for id, ch := range chans {
		ctrl := <-ch
		ctrls[id] = ctrl
	}

	ctrl := root_list.NewRLRootListCtrl(ctrls, h.rootListLogger)
	h.rpc.OpenRootList(ctrl)
}

func (h *providerHandler) activate(cmd *activateCmd) {
	h.openRootList()
}
