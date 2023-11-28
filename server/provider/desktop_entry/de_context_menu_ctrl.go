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
	formID         api.FormID
	client         api.RpcClient
	id             api.RootListRowID
	actionExecuter *deActionExecuter
	moduleLogger   *zap.Logger
}

func newDEContextMenuCtrl(id api.RootListRowID, actionExecuter *deActionExecuter, moduleLogger *zap.Logger) *DEContextMenuCtrl {
	return &DEContextMenuCtrl{
		formID:         0,
		client:         nil,
		id:             id,
		actionExecuter: actionExecuter,
		moduleLogger:   moduleLogger,
	}
}

func (c *DEContextMenuCtrl) OnOpen(formID api.FormID, client api.RpcClient) []*api.ContextMenuRow {
	c.formID = formID
	c.client = client

	return []*api.ContextMenuRow{
		api.NewContextMenuRow(api.ContextMenuRowID(actionOpen), "Open", "Open\nОткрыть"),
		api.NewContextMenuRow(api.ContextMenuRowID(actionCopyName), "Copy name", "Copy name\nКопировать имя"),
		api.NewContextMenuRow(api.ContextMenuRowID(actionCopyPath), "Copy path", "Copy path\nКопировать путь"),
	}
}

func (c *DEContextMenuCtrl) OnRowActivate(rowID api.ContextMenuRowID) {
	switch uint32(rowID) {
	case actionOpen:
		c.actionExecuter.open(c.client, c.id)
	case actionCopyName:
		c.actionExecuter.copyName(c.client, c.id)
	case actionCopyPath:
		c.actionExecuter.copyPath(c.client, c.id)
	default:
		err := errors.New("unknown menu id")
		c.moduleLogger.Warn("Failed execute menu item",
			c.id.ZapField(),
			rowID.ZapField(),
			zap.Error(err),
		)

		c.client.HideUI(err)
	}
}
