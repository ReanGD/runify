package links

import (
	"github.com/goccy/go-json"
	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/widget"
)

type LinksFormCtrl struct {
	formID         api.FormID
	rowID          api.RootListRowID
	client         api.RpcClient
	model          *model
	dataForm       *widget.DataForm
	rootListformID api.FormID
	moduleLogger   *zap.Logger
}

func newLinksFormCtrl(rowID api.RootListRowID, model *model, rootListformID api.FormID, moduleLogger *zap.Logger) (*LinksFormCtrl, error) {
	dataForm, err := model.createDataForm(rowID)
	if err != nil {
		moduleLogger.Error("Failed create markup for form", zap.Error(err))
		return nil, err
	}

	return &LinksFormCtrl{
		formID:         0,
		rowID:          rowID,
		client:         nil,
		model:          model,
		dataForm:       dataForm,
		rootListformID: rootListformID,
		moduleLogger:   moduleLogger,
	}, nil
}

func (c *LinksFormCtrl) OnOpen(formID api.FormID, client api.RpcClient) *widget.DataForm {
	c.formID = formID
	c.client = client

	return c.dataForm
}

func (c *LinksFormCtrl) OnFieldCheckRequest(requestID uint32, fieldName string, jsonBody string) {
	data := &DataModel{}
	err := json.Unmarshal([]byte(jsonBody), data)
	if err != nil {
		c.moduleLogger.Warn("Failed unmarshal json for check", zap.Error(err))
		c.client.FieldCheckResponse(c.formID, requestID, false, "internal error, failed parse request")
	} else if !c.model.checkItem(c.rowID, data.Name) {
		c.client.FieldCheckResponse(c.formID, requestID, false, "name already exists")
	} else {
		c.client.FieldCheckResponse(c.formID, requestID, true, "")
	}
}

func (c *LinksFormCtrl) OnSubmit(jsonBody string) {
	data := &DataModel{}
	err := json.Unmarshal([]byte(jsonBody), data)
	if err != nil {
		c.client.UserMessage("Failed save link")
		c.moduleLogger.Warn("Failed unmarshal json for save link", zap.Error(err))
	}

	id := c.rowID
	var row *api.RootListRow
	if id <= createRowID {
		row, err = c.model.addItem(data, true)
		if err == nil && row != nil {
			c.client.RootListAddRows(c.rootListformID, row)
		}
	} else {
		row, err = c.model.updateItem(id, data, true)
		if err == nil && row != nil {
			c.client.RootListChangeRows(c.rootListformID, row)
		}
	}

	if err != nil {
		c.client.UserMessage("Failed save link")
		c.moduleLogger.Warn("Failed save link", zap.Error(err))
		return
	}

	c.client.CloseForm(c.formID)
}
