package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

const rootRowID = api.RootListRowID(1)

type CalcRootListCtrl struct {
	visible        bool
	actualValue    string
	formID         api.FormID
	providerID     api.ProviderID
	executer       *Executer
	actionExecuter *calcActionExecuter
	client         api.RpcClient
	moduleLogger   *zap.Logger
}

func newCalcRootListCtrl(providerID api.ProviderID, actionExecuter *calcActionExecuter, moduleLogger *zap.Logger) *CalcRootListCtrl {
	return &CalcRootListCtrl{
		visible:        false,
		actualValue:    "",
		formID:         0,
		providerID:     providerID,
		executer:       NewExecuter(),
		actionExecuter: actionExecuter,
		client:         nil,
		moduleLogger:   moduleLogger,
	}
}

func (c *CalcRootListCtrl) OnOpen(formID api.FormID, client api.RpcClient) []*api.RootListRow {
	c.formID = formID
	c.client = client
	return []*api.RootListRow{}
}

func (c *CalcRootListCtrl) OnFilterChange(text string) {
	calcResult := c.executer.Execute(text)
	if calcResult.ParserErr != nil {
		if c.visible {
			c.client.RootListRemoveRows(c.formID, api.NewRootListRowGlobalID(c.providerID, rootRowID))
		}
		c.visible = false
		c.actualValue = ""
		return
	}

	value := calcResult.Value.Value()
	strValue := value.String()
	if strValue == c.actualValue {
		return
	}

	row := api.NewRootListRow(c.providerID, rootRowID, "", "|"+text+" = "+strValue, api.MaxPriority)
	if c.visible {
		c.client.RootListChangeRows(c.formID, row)
	} else {
		c.client.RootListAddRows(c.formID, row)
	}

	c.visible = true
	c.actualValue = strValue
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

		c.client.HideUI(err)
	} else {
		menuCtrl := newCalcContextMenuCtrl(c.actualValue, c.actionExecuter, c.moduleLogger)
		c.client.AddContextMenu(menuCtrl)
	}
}
