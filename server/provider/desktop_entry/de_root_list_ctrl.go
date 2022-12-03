package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

type DERootListCtrl struct {
	model          *deModel
	actionExecuter *deActionExecuter
	sender         api.RootListRowsUpdateSender
	moduleLogger   *zap.Logger
}

func newDERootListCtrl(model *deModel, actionExecuter *deActionExecuter, moduleLogger *zap.Logger) *DERootListCtrl {
	return &DERootListCtrl{
		model:          model,
		actionExecuter: actionExecuter,
		sender:         nil,
		moduleLogger:   moduleLogger,
	}
}

func (c *DERootListCtrl) OnOpen(sender api.RootListRowsUpdateSender) []*api.RootListRow {
	c.sender = sender
	return c.model.getRows()
}

func (c *DERootListCtrl) OnFilterChange(value string) {
	// pass
}

func (c *DERootListCtrl) OnRowActivate(providerID api.ProviderID, rowID api.RootListRowID, result api.ErrorResult) {
	c.actionExecuter.open(rowID, result)
}

func (c *DERootListCtrl) OnMenuActivate(providerID api.ProviderID, rowID api.RootListRowID, result api.ContexMenuCtrlOrErrorResult) {
	_, ok := c.model.getEntry(rowID)
	if !ok {
		err := errors.New("row data not found")
		c.moduleLogger.Warn("Failed open context menu",
			rowID.ZapField(),
			zap.Error(err),
		)

		result.SetResult(api.ContextMenuCtrlOrError{
			Error: err,
		})
	} else {
		result.SetResult(api.ContextMenuCtrlOrError{
			Ctrl: newDEContextMenuCtrl(rowID, c.actionExecuter, c.moduleLogger),
		})
	}
}
