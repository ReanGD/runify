package provider

import (
	"context"
	"errors"
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
	deps           *dependences
	dpChans        []<-chan error
	dataProviders  map[api.ProviderID]*dataProvider
	wg             *sync.WaitGroup
	showUIHotkey   *shortcut.Hotkey
	doneErrWaitCh  chan struct{}
	moduleLogger   *zap.Logger
	rootListLogger *zap.Logger
}

func newProviderHandler() *providerHandler {
	return &providerHandler{
		cfg:            nil,
		deps:           nil,
		dpChans:        make([]<-chan error, 0, 3),
		dataProviders:  make(map[api.ProviderID]*dataProvider),
		wg:             &sync.WaitGroup{},
		showUIHotkey:   nil,
		doneErrWaitCh:  make(chan struct{}),
		moduleLogger:   nil,
		rootListLogger: nil,
	}
}

func (h *providerHandler) addProvider(providerID api.ProviderID, handler dataProviderHandler) {
	dp := newDataProvider(providerID, handler)
	h.dataProviders[providerID] = dp
	h.dpChans = append(h.dpChans, dp.Init(dp, handler.GetName(), module.SUB_MODULE, h.cfg, h.moduleLogger))
}

func (h *providerHandler) onInit(root *Provider, deps *dependences) error {
	h.cfg = root.GetConfig()
	h.deps = deps
	h.moduleLogger = root.GetModuleLogger()
	h.rootListLogger = root.NewSubmoduleLogger("RootList")
	desktop := deps.desktop

	h.addProvider(desktopEntryID, desktop_entry.New(desktop, deps.de))
	h.addProvider(calculatorID, calculator.New(desktop))
	h.addProvider(linksID, links.New(desktop))

	for _, dpChan := range h.dpChans {
		if err := <-dpChan; err != nil {
			return err
		}
	}

	var err error
	if h.showUIHotkey, err = shortcut.NewHotkey(h.cfg.Shortcuts.Root); err != nil {
		return err
	}

	return nil
}

func (h *providerHandler) onStart(ctx context.Context, errorCtx *module.ErrorCtx) {
	cases := make([]reflect.SelectCase, len(h.dataProviders)+1)

	result := api.NewFuncErrorCodeResult(func(result global.Error) {
		if result != global.Success {
			errorCtx.SendError(errors.New(result.String()))
		}
	})
	h.deps.desktop.AddShortcut(shortcut.NewAction("Show UI"), h.showUIHotkey, result)

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
	h.deps.rpc.OpenRootList(ctrl)
}

func (h *providerHandler) activate(cmd *activateCmd) {
	h.openRootList()
}
