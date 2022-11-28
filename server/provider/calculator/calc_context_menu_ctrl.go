package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

const menuRowID = api.ContextMenuRowID(1)

type CalcContextMenuCtrl struct {
	value          string
	actionExecuter *calcActionExecuter
	moduleLogger   *zap.Logger
}

func newCalcContextMenuCtrl(value string, actionExecuter *calcActionExecuter, moduleLogger *zap.Logger) *CalcContextMenuCtrl {
	return &CalcContextMenuCtrl{
		value:          value,
		actionExecuter: actionExecuter,
		moduleLogger:   moduleLogger,
	}
}

func (c *CalcContextMenuCtrl) GetRows() *api.ContextMenuRows {
	data := api.NewContextMenuRows()
	data.Create = []*api.ContextMenuRow{
		api.NewContextMenuRow(menuRowID, "Copy"),
	}

	return data
}

func (c *CalcContextMenuCtrl) OnRowActivate(id api.ContextMenuRowID, result api.ErrorResult) {
	switch id {
	case menuRowID:
		c.actionExecuter.copyResult(c.value, result)
	default:
		err := errors.New("unknown menu id")
		c.moduleLogger.Warn("Failed execute menu item",
			id.ZapField(),
			zap.Error(err),
		)
		result.SetResult(err)
	}
}
