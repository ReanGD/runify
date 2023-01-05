package links

import (
	"github.com/ReanGD/runify/server/global/api"

	"go.uber.org/zap"
)

type LinksRootListCtrl struct {
	formID         api.FormID
	providerID     api.ProviderID
	model          *model
	actionExecuter *actionExecuter
	client         api.RpcClient
	moduleLogger   *zap.Logger
}

func newLinksRootListCtrl(
	providerID api.ProviderID, model *model, actionExecuter *actionExecuter, moduleLogger *zap.Logger,
) *LinksRootListCtrl {
	return &LinksRootListCtrl{
		formID:         0,
		providerID:     providerID,
		model:          model,
		actionExecuter: actionExecuter,
		client:         nil,
		moduleLogger:   moduleLogger,
	}
}

func (c *LinksRootListCtrl) OnOpen(formID api.FormID, client api.RpcClient) []*api.RootListRow {
	c.formID = formID
	c.client = client
	return c.model.getRows()
}

func (c *LinksRootListCtrl) OnFilterChange(text string) {
	// pass
}

func (c *LinksRootListCtrl) OnRowActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	if rowID == createRowID {
		formCtrl, err := newLinksFormCtrl(rowID, c.model, c.moduleLogger)
		if err != nil {
			c.client.UserMessage(err.Error())
		} else {
			c.client.AddForm(formCtrl)
		}
		return
	}

	item, ok := c.model.getItem(rowID)
	if !ok {
		c.client.UserMessage("Item not found")
		return
	}
	c.actionExecuter.openLink(c.client, item.data)
}

func (c *LinksRootListCtrl) OnMenuActivate(providerID api.ProviderID, rowID api.RootListRowID) {
	menuCtrl, err := newLinksContextMenuCtrl(rowID, c.model, c.actionExecuter, c.moduleLogger)
	if err != nil {
		c.client.UserMessage(err.Error())
		c.moduleLogger.Warn("Failed open menu",
			rowID.ZapField(),
			zap.Error(err),
		)
	} else {
		c.client.AddContextMenu(menuCtrl)
	}
}
