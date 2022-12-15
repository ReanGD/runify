package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

type DERootListCtrl struct {
	formID         api.FormID
	model          *deModel
	actionExecuter *deActionExecuter
	client         api.RpcClient
	moduleLogger   *zap.Logger
}

func newDERootListCtrl(model *deModel, actionExecuter *deActionExecuter, moduleLogger *zap.Logger) *DERootListCtrl {
	return &DERootListCtrl{
		formID:         0,
		model:          model,
		actionExecuter: actionExecuter,
		client:         nil,
		moduleLogger:   moduleLogger,
	}
}

func (c *DERootListCtrl) OnOpen(formID api.FormID, client api.RpcClient) []*api.RootListRow {
	c.formID = formID
	c.client = client
	return c.model.getRows()
}

func (c *DERootListCtrl) OnFilterChange(value string) {
	// pass
}

func (c *DERootListCtrl) OnRowActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	c.actionExecuter.open(c.client, rowID)
}

func (c *DERootListCtrl) OnMenuActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	_, ok := c.model.getEntry(rowID)
	if !ok {
		err := errors.New("row data not found")
		c.moduleLogger.Warn("Failed open context menu",
			rowID.ZapField(),
			zap.Error(err),
		)

		c.client.HideUI(err)
	} else {
		menuCtrl := newDEContextMenuCtrl(rowID, c.actionExecuter, c.moduleLogger)
		c.client.AddContextMenu(menuCtrl)
	}
}
