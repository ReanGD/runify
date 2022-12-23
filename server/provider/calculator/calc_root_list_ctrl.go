package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/interpreter"

	"go.uber.org/zap"
)

const rootRowID = api.RootListRowID(1)

type CalcRootListCtrl struct {
	visible        bool
	lastResult     string
	formID         api.FormID
	providerID     api.ProviderID
	interpreter    *interpreter.Interpreter
	actionExecuter *calcActionExecuter
	client         api.RpcClient
	moduleLogger   *zap.Logger
}

func newCalcRootListCtrl(providerID api.ProviderID, actionExecuter *calcActionExecuter, moduleLogger *zap.Logger) *CalcRootListCtrl {
	return &CalcRootListCtrl{
		visible:        false,
		lastResult:     "",
		formID:         0,
		providerID:     providerID,
		interpreter:    interpreter.New(),
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
	res := c.interpreter.Execute(text)
	if !res.IsExprValid() {
		if c.visible {
			c.client.RootListRemoveRows(c.formID, api.NewRootListRowGlobalID(c.providerID, rootRowID))
		}
		c.visible = false
		c.lastResult = ""
		return
	}

	userResult := res.UserResult()
	if userResult == c.lastResult {
		return
	}

	row := api.NewRootListRow(api.RowType_Calculator, api.MaxPriority, c.providerID, rootRowID, "", text+"\n"+userResult)
	if c.visible {
		c.client.RootListChangeRows(c.formID, row)
	} else {
		c.client.RootListAddRows(c.formID, row)
	}

	c.visible = true
	c.lastResult = userResult
}

func (c *CalcRootListCtrl) OnRowActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	c.actionExecuter.copyResult(c.client, c.lastResult)
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
		menuCtrl := newCalcContextMenuCtrl(c.lastResult, c.actionExecuter, c.moduleLogger)
		c.client.AddContextMenu(menuCtrl)
	}
}
