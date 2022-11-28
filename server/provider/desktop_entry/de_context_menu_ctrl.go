package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
)

const (
	actionOpen uint32 = iota
	actionCopyName
	actionCopyPath
	actionLast
)

type DEContextMenuCtrl struct {
	id       api.RootListRowID
	executer *deExecuter
}

func newDEContextMenuCtrl(id api.RootListRowID, executer *deExecuter) *DEContextMenuCtrl {
	return &DEContextMenuCtrl{
		id:       id,
		executer: executer,
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
		c.executer.open(c.id, result)
	case actionCopyName:
		c.executer.copyName(c.id, result)
	case actionCopyPath:
		c.executer.copyPath(c.id, result)
	default:
		err := errors.New("unknown action")
		result.SetResult(err)
	}
}
