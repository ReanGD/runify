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
	rowsCh         chan *api.RootListRows
	moduleLogger   *zap.Logger
}

func newCalcRootListCtrl(providerID api.ProviderID, actionExecuter *calcActionExecuter, moduleLogger *zap.Logger) *CalcRootListCtrl {
	return &CalcRootListCtrl{
		actualValue:    "",
		providerID:     providerID,
		executer:       NewExecuter(),
		actionExecuter: actionExecuter,
		rowsCh:         make(chan *api.RootListRows, 10),
		moduleLogger:   moduleLogger,
	}
}

func (c *CalcRootListCtrl) GetRowsCh() <-chan *api.RootListRows {
	c.rowsCh <- api.NewRootListRows()
	return c.rowsCh
}

func (c *CalcRootListCtrl) OnFilterChange(text string) {
	calcResult := c.executer.Execute(text)
	if calcResult.ParserErr != nil {
		if len(c.actualValue) == 0 {
			data := api.NewRootListRows()
			data.Remove = append(data.Remove, api.NewRootListRowRemove(c.providerID, rootRowID))
		}
		return
	}

	value := calcResult.Value.Value()
	strValue := value.String()
	if strValue == c.actualValue {
		return
	}

	data := api.NewRootListRows()
	rows := []*api.RootListRow{api.NewRootListRow(c.providerID, rootRowID, "", strValue, api.MaxPriority)}
	if len(c.actualValue) == 0 {
		data.Change = rows
	} else {
		data.Create = rows
	}
	c.actualValue = strValue
}

func (c *CalcRootListCtrl) OnRowActivate(id api.RootListRowID, result api.ErrorResult) {
	c.actionExecuter.copyResult(c.actualValue, result)
}

func (c *CalcRootListCtrl) OnMenuActivate(id api.RootListRowID, result api.ContexMenuCtrlOrErrorResult) {
	if id != rootRowID {
		err := errors.New("row data not found")
		c.moduleLogger.Warn("Failed open menu",
			id.ZapField(),
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
