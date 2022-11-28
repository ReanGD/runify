package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

type DERootListCtrl struct {
	model          *deModel
	actionExecuter *deActionExecuter
	moduleLogger   *zap.Logger
}

func newDERootListCtrl(model *deModel, actionExecuter *deActionExecuter, moduleLogger *zap.Logger) *DERootListCtrl {
	return &DERootListCtrl{
		model:          model,
		actionExecuter: actionExecuter,
		moduleLogger:   moduleLogger,
	}
}

func (c *DERootListCtrl) GetRowsCh() <-chan *api.RootListRows {
	data := api.NewRootListRows()
	data.Create = c.model.getRows()

	ch := make(chan *api.RootListRows, 1)
	ch <- data
	return ch
}

func (c *DERootListCtrl) OnFilterChange(value string) {
	// pass
}

func (c *DERootListCtrl) OnRowActivate(id api.RootListRowID, result api.ErrorResult) {
	c.actionExecuter.open(id, result)
}

func (c *DERootListCtrl) OnMenuActivate(id api.RootListRowID, result api.ContexMenuCtrlOrErrorResult) {
	_, ok := c.model.getEntry(id)
	if !ok {
		err := errors.New("row data not found")
		c.moduleLogger.Warn("Failed open context menu",
			id.ZapField(),
			zap.Error(err),
		)

		result.SetResult(api.ContextMenuCtrlOrError{
			Error: err,
		})
	} else {
		result.SetResult(api.ContextMenuCtrlOrError{
			Ctrl: newDEContextMenuCtrl(id, c.actionExecuter, c.moduleLogger),
		})
	}
}
