package root_list

import (
	"errors"
	"sort"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

type RLRootListCtrl struct {
	ctrls        map[api.ProviderID]api.RootListCtrl
	moduleLogger *zap.Logger
}

func newRLRootListCtrl(ctrls map[api.ProviderID]api.RootListCtrl, moduleLogger *zap.Logger) *RLRootListCtrl {
	return &RLRootListCtrl{
		ctrls:        ctrls,
		moduleLogger: moduleLogger,
	}
}

func (c *RLRootListCtrl) GetRows(out chan *api.RootListRowsUpdate) []*api.RootListRow {
	wg := &sync.WaitGroup{}
	wg.Add(len(c.ctrls))

	resMutex := sync.Mutex{}
	allData := []*api.RootListRow{}
	for _, ctrl := range c.ctrls {
		go func(ctrl api.RootListCtrl) {
			defer wg.Done()
			data := ctrl.GetRows(out)
			resMutex.Lock()
			allData = append(allData, data...)
			resMutex.Unlock()
		}(ctrl)
	}

	// Wait for all controllers send initial data
	// Ans sort by priority, then by name
	wg.Wait()
	sort.SliceStable(allData, func(i, j int) bool {
		if allData[i].Priority == allData[j].Priority {
			return allData[i].Value < allData[j].Value
		}
		return allData[i].Priority > allData[j].Priority
	})

	return allData
}

func (c *RLRootListCtrl) OnFilterChange(value string) {
	for _, ctrl := range c.ctrls {
		ctrl.OnFilterChange(value)
	}
}

func (c *RLRootListCtrl) OnRowActivate(providerID api.ProviderID, rowID api.RootListRowID, result api.ErrorResult) {
	ctrl, ok := c.ctrls[providerID]
	if !ok {
		err := errors.New("provider not found")
		c.moduleLogger.Warn("Failed execute default action",
			providerID.ZapField(),
			rowID.ZapField(),
			zap.Error(err),
		)

		result.SetResult(err)
	} else {
		ctrl.OnRowActivate(providerID, rowID, result)
	}
}

func (c *RLRootListCtrl) OnMenuActivate(providerID api.ProviderID, rowID api.RootListRowID, result api.ContexMenuCtrlOrErrorResult) {
	ctrl, ok := c.ctrls[providerID]
	if !ok {
		err := errors.New("provider not found")
		c.moduleLogger.Warn("Failed open context menu",
			providerID.ZapField(),
			rowID.ZapField(),
			zap.Error(err),
		)

		result.SetResult(api.ContextMenuCtrlOrError{
			Error: err,
		})
	} else {
		ctrl.OnMenuActivate(providerID, rowID, result)
	}
}
