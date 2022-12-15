package root_list

import (
	"errors"
	"sort"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

type RLRootListCtrl struct {
	client       api.RpcClient
	ctrls        map[api.ProviderID]api.RootListCtrl
	moduleLogger *zap.Logger
}

func NewRLRootListCtrl(ctrls map[api.ProviderID]api.RootListCtrl, moduleLogger *zap.Logger) *RLRootListCtrl {
	return &RLRootListCtrl{
		ctrls:        ctrls,
		moduleLogger: moduleLogger,
	}
}

func (c *RLRootListCtrl) OnOpen(formID api.FormID, client api.RpcClient) []*api.RootListRow {
	c.client = client
	wg := &sync.WaitGroup{}
	wg.Add(len(c.ctrls))

	resMutex := sync.Mutex{}
	allData := []*api.RootListRow{}
	for _, ctrl := range c.ctrls {
		go func(ctrl api.RootListCtrl) {
			defer wg.Done()
			data := ctrl.OnOpen(formID, client)
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

func (c *RLRootListCtrl) OnRowActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	ctrl, ok := c.ctrls[providerID]
	if !ok {
		err := errors.New("provider not found")
		c.moduleLogger.Warn("Failed execute default action",
			providerID.ZapField(),
			rowID.ZapField(),
			zap.Error(err),
		)

		c.client.HideUI(err)
	} else {
		ctrl.OnRowActivate(providerID, rowID)
	}
}

func (c *RLRootListCtrl) OnMenuActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	ctrl, ok := c.ctrls[providerID]
	if !ok {
		err := errors.New("provider not found")
		c.moduleLogger.Warn("Failed open context menu",
			providerID.ZapField(),
			rowID.ZapField(),
			zap.Error(err),
		)

		c.client.HideUI(err)
	} else {
		ctrl.OnMenuActivate(providerID, rowID)
	}
}
