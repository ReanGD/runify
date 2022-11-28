package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

const (
	actionOpen uint32 = iota + 1
	actionCopyName
	actionCopyPath
	actionLast
)

type DEContextMenuCtrl struct {
	id             api.RootListRowID
	actionExecuter *deActionExecuter
	moduleLogger   *zap.Logger
}

func newDEContextMenuCtrl(id api.RootListRowID, actionExecuter *deActionExecuter, moduleLogger *zap.Logger) *DEContextMenuCtrl {
	return &DEContextMenuCtrl{
		id:             id,
		actionExecuter: actionExecuter,
		moduleLogger:   moduleLogger,
	}
}

func (c *DEContextMenuCtrl) GetRows() *api.ContextMenuRows {
	data := api.NewContextMenuRows()
	data.Create = []*api.ContextMenuRow{
		api.NewContextMenuRow(api.ContextMenuRowID(actionOpen), "Open"),
		api.NewContextMenuRow(api.ContextMenuRowID(actionCopyName), "Copy name"),
		api.NewContextMenuRow(api.ContextMenuRowID(actionCopyPath), "Copy path"),
	}

	return data
}

func (c *DEContextMenuCtrl) OnRowActivate(id api.ContextMenuRowID, result api.ErrorResult) {
	switch uint32(id) {
	case actionOpen:
		c.actionExecuter.open(c.id, result)
	case actionCopyName:
		c.actionExecuter.copyName(c.id, result)
	case actionCopyPath:
		c.actionExecuter.copyPath(c.id, result)
	default:
		err := errors.New("unknown menu id")
		c.moduleLogger.Warn("Failed execute menu item",
			id.ZapField(),
			zap.Error(err),
		)
		result.SetResult(err)
	}
}
