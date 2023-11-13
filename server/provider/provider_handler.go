package provider

import (
	"context"
	"reflect"
	"sync"

	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/provider/calculator"
	"github.com/ReanGD/runify/server/provider/desktop_entry"
	"github.com/ReanGD/runify/server/provider/links"
	"github.com/ReanGD/runify/server/provider/root_list"
)

type providerHandler struct {
	dataProviders  map[api.ProviderID]*dataProvider
	rpc            api.Rpc
	desktop        api.Desktop
	errCh          chan error
	wg             *sync.WaitGroup
	doneErrWaitCh  chan struct{}
	moduleLogger   *zap.Logger
	rootListLogger *zap.Logger
}

func newProviderHandler() *providerHandler {
	return &providerHandler{
		dataProviders:  make(map[api.ProviderID]*dataProvider),
		rpc:            nil,
		desktop:        nil,
		errCh:          make(chan error, 1),
		wg:             &sync.WaitGroup{},
		doneErrWaitCh:  make(chan struct{}),
		moduleLogger:   nil,
		rootListLogger: nil,
	}
}

func (h *providerHandler) getErrCh() <-chan error {
	return h.errCh
}

func (h *providerHandler) onInit(cfg *config.Config, desktop api.Desktop, rpc api.Rpc, moduleLogger *zap.Logger, rootListLogger *zap.Logger) error {
	h.rpc = rpc
	h.desktop = desktop
	h.moduleLogger = moduleLogger
	h.rootListLogger = rootListLogger
	h.dataProviders[desktopEntryID] = newDataProvider(desktopEntryID, desktop_entry.New(desktop))
	h.dataProviders[calculatorID] = newDataProvider(calculatorID, calculator.New(desktop))
	h.dataProviders[linksID] = newDataProvider(linksID, links.New(desktop))

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

func (h *providerHandler) onStart(ctx context.Context) {
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

		_, recv, recvOk := reflect.Select(cases)
		if !recvOk {
			h.errCh <- recv.Interface().(error)
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
