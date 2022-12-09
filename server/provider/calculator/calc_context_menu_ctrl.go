package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

const menuRowID = api.ContextMenuRowID(1)

type CalcContextMenuCtrl struct {
	formID         api.FormID
	client         api.RpcClient
	value          string
	actionExecuter *calcActionExecuter
	moduleLogger   *zap.Logger
}

func newCalcContextMenuCtrl(value string, actionExecuter *calcActionExecuter, moduleLogger *zap.Logger) *CalcContextMenuCtrl {
	return &CalcContextMenuCtrl{
		formID:         0,
		client:         nil,
		value:          value,
		actionExecuter: actionExecuter,
		moduleLogger:   moduleLogger,
	}
}

func (c *CalcContextMenuCtrl) OnOpen(formID api.FormID, client api.RpcClient) []*api.ContextMenuRow {
	c.formID = formID
	c.client = client
	return []*api.ContextMenuRow{api.NewContextMenuRow(menuRowID, "Copy")}
}

func (c *CalcContextMenuCtrl) OnRowActivate(rowID api.ContextMenuRowID) {
	switch rowID {
	case menuRowID:
		c.actionExecuter.copyResult(c.client, c.value)
	default:
		err := errors.New("unknown menu id")
		c.moduleLogger.Warn("Failed execute menu item",
			rowID.ZapField(),
			zap.String("Value", c.value),
			zap.Error(err),
		)
		c.client.CloseAll(err)
	}
}
