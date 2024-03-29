package links

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

const (
	menuOpenCmdRowID api.ContextMenuRowID = iota
	menuOpenLinkRowID
	menuEditLinkRowID
	menuCopyLinkRowID
	menuCopyNameRowID
	menuRemoveLinkRowID
)

type LinksContextMenuCtrl struct {
	formID         api.FormID
	itemRowID      api.RootListRowID
	client         api.RpcClient
	itemData       *DataModel
	model          *model
	actionExecuter *actionExecuter
	providerID     api.ProviderID
	rootListformID api.FormID
	moduleLogger   *zap.Logger
}

func newLinksContextMenuCtrl(
	itemRowID api.RootListRowID,
	model *model,
	actionExecuter *actionExecuter,
	providerID api.ProviderID,
	rootListformID api.FormID,
	moduleLogger *zap.Logger,
) (*LinksContextMenuCtrl, error) {
	var itemData *DataModel
	if itemRowID > createRowID {
		item, ok := model.getItem(itemRowID)
		if !ok {
			return nil, errors.New("item not found")
		}

		itemData = item.data
	}

	return &LinksContextMenuCtrl{
		formID:         0,
		itemRowID:      itemRowID,
		client:         nil,
		itemData:       itemData,
		model:          model,
		actionExecuter: actionExecuter,
		providerID:     providerID,
		rootListformID: rootListformID,
		moduleLogger:   moduleLogger,
	}, nil
}

func (c *LinksContextMenuCtrl) OnOpen(formID api.FormID, client api.RpcClient) []*api.ContextMenuRow {
	c.formID = formID
	c.client = client

	if c.itemRowID <= createRowID {
		return []*api.ContextMenuRow{
			api.NewContextMenuRow(menuOpenCmdRowID, "Open command", "Open command\nОткрыть команду"),
		}
	}

	return []*api.ContextMenuRow{
		api.NewContextMenuRow(menuOpenLinkRowID, "Open", "Open\nОткрыть"),
		api.NewContextMenuRow(menuEditLinkRowID, "Edit", "Edit\nРедактировать"),
		api.NewContextMenuRow(menuCopyLinkRowID, "Copy link", "Copy link\nКопировать ссылку"),
		api.NewContextMenuRow(menuCopyNameRowID, "Copy name", "Copy name\nКопировать имя"),
		api.NewContextMenuRow(menuRemoveLinkRowID, "Remove", "Remove\nУдалить"),
	}
}

func (c *LinksContextMenuCtrl) OnRowActivate(menuRowID api.ContextMenuRowID) {
	if c.itemRowID <= createRowID {
		if menuRowID == menuOpenCmdRowID {
			formCtrl, err := newLinksFormCtrl(c.itemRowID, c.model, c.rootListformID, c.moduleLogger)
			if err != nil {
				c.client.UserMessage(err.Error())
			} else {
				c.client.AddForm(formCtrl)
			}
		} else {
			c.client.UserMessage("Unknown menu item ID")
			c.moduleLogger.Warn("Unknown menu item ID for create row", menuRowID.ZapField())
		}

		return
	}

	switch menuRowID {
	case menuOpenLinkRowID:
		c.actionExecuter.openLink(c.client, c.itemData)
	case menuEditLinkRowID:
		formCtrl, err := newLinksFormCtrl(c.itemRowID, c.model, c.rootListformID, c.moduleLogger)
		c.client.CloseForm(c.formID)
		if err != nil {
			c.client.UserMessage(err.Error())
		} else {
			c.client.AddForm(formCtrl)
		}
	case menuCopyLinkRowID:
		c.actionExecuter.copyValue(c.client, c.itemData.Link)
	case menuCopyNameRowID:
		c.actionExecuter.copyValue(c.client, c.itemData.Name)
	case menuRemoveLinkRowID:
		if err := c.model.removeItem(c.itemRowID, true); err != nil {
			c.client.UserMessage("Failed remove link")
		} else {
			c.client.RootListRemoveRows(c.rootListformID, api.NewRootListRowGlobalID(c.providerID, c.itemRowID))
		}
		c.client.CloseForm(c.formID)
	default:
		err := errors.New("unknown menu id")
		c.moduleLogger.Warn("Failed execute menu item",
			menuRowID.ZapField(),
			zap.Error(err),
		)
		c.client.HideUI(err)
	}
}
