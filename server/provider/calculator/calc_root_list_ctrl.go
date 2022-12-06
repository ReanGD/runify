package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

const rootRowID = api.RootListRowID(1)

type CalcRootListCtrl struct {
	actualValue    string
	formID         uint32
	providerID     api.ProviderID
	executer       *Executer
	actionExecuter *calcActionExecuter
	client         api.RpcClient
	moduleLogger   *zap.Logger
}

func newCalcRootListCtrl(providerID api.ProviderID, actionExecuter *calcActionExecuter, moduleLogger *zap.Logger) *CalcRootListCtrl {
	return &CalcRootListCtrl{
		actualValue:    "",
		formID:         0,
		providerID:     providerID,
		executer:       NewExecuter(),
		actionExecuter: actionExecuter,
		client:         nil,
		moduleLogger:   moduleLogger,
	}
}

func (c *CalcRootListCtrl) OnOpen(formID uint32, client api.RpcClient) []*api.RootListRow {
	c.formID = formID
	c.client = client
	return []*api.RootListRow{}
}

func (c *CalcRootListCtrl) OnFilterChange(text string) {
	calcResult := c.executer.Execute(text)
	if calcResult.ParserErr != nil {
		if len(c.actualValue) == 0 {
			c.client.RootListRemoveRows(c.formID, api.NewRootListRowGlobalID(c.providerID, rootRowID))
		}
		return
	}

	value := calcResult.Value.Value()
	strValue := value.String()
	if strValue == c.actualValue {
		return
	}

	c.actualValue = strValue
	row := api.NewRootListRow(c.providerID, rootRowID, "", strValue, api.MaxPriority)
	if len(c.actualValue) == 0 {
		c.client.RootListChangeRows(c.formID, row)
	} else {
		c.client.RootListAddRows(c.formID, row)
	}
}

func (c *CalcRootListCtrl) OnRowActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	c.actionExecuter.copyResult(c.client, c.actualValue)
}

func (c *CalcRootListCtrl) OnMenuActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	if rowID != rootRowID {
		err := errors.New("row data not found")
		c.moduleLogger.Warn("Failed open menu",
			rowID.ZapField(),
			zap.Error(err),
		)

		c.client.CloseAll(err)
	} else {
		menuCtrl := newCalcContextMenuCtrl(c.actualValue, c.actionExecuter, c.moduleLogger)
		menuCtrl.OnOpen(c.formID+1, c.client)
	}
}
