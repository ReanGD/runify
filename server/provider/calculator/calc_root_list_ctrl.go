package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

const rootRowID = api.RootListRowID(1)

type CalcRootListCtrl struct {
	actualValue    string
	providerID     api.ProviderID
	executer       *Executer
	actionExecuter *calcActionExecuter
	sender         api.RootListRowsUpdateSender
	moduleLogger   *zap.Logger
}

func newCalcRootListCtrl(providerID api.ProviderID, actionExecuter *calcActionExecuter, moduleLogger *zap.Logger) *CalcRootListCtrl {
	return &CalcRootListCtrl{
		actualValue:    "",
		providerID:     providerID,
		executer:       NewExecuter(),
		actionExecuter: actionExecuter,
		sender:         nil,
		moduleLogger:   moduleLogger,
	}
}

func (c *CalcRootListCtrl) OnOpen(sender api.RootListRowsUpdateSender) []*api.RootListRow {
	c.sender = sender
	return []*api.RootListRow{}
}

func (c *CalcRootListCtrl) OnFilterChange(text string) {
	calcResult := c.executer.Execute(text)
	if calcResult.ParserErr != nil {
		if len(c.actualValue) == 0 {
			update := api.NewRootListRowsUpdate()
			update.Remove = append(update.Remove, api.NewRootListRowGlobalID(c.providerID, rootRowID))
			c.sender.Send(update)
		}
		return
	}

	value := calcResult.Value.Value()
	strValue := value.String()
	if strValue == c.actualValue {
		return
	}

	data := api.NewRootListRowsUpdate()
	rows := []*api.RootListRow{api.NewRootListRow(c.providerID, rootRowID, "", strValue, api.MaxPriority)}
	if len(c.actualValue) == 0 {
		data.Change = rows
	} else {
		data.Create = rows
	}
	c.actualValue = strValue
	c.sender.Send(data)
}

func (c *CalcRootListCtrl) OnRowActivate(providerID api.ProviderID, rowID api.RootListRowID, result api.ErrorResult) {
	c.actionExecuter.copyResult(c.actualValue, result)
}

func (c *CalcRootListCtrl) OnMenuActivate(providerID api.ProviderID, rowID api.RootListRowID, result api.ContexMenuCtrlOrErrorResult) {
	if rowID != rootRowID {
		err := errors.New("row data not found")
		c.moduleLogger.Warn("Failed open menu",
			rowID.ZapField(),
			zap.Error(err),
		)

		result.SetResult(api.ContextMenuCtrlOrError{
			Error: err,
		})
	} else {
		result.SetResult(api.ContextMenuCtrlOrError{
			Ctrl: newCalcContextMenuCtrl(c.actualValue, c.actionExecuter, c.moduleLogger),
		})
	}
}
