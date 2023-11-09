package links

import (
	"github.com/goccy/go-json"
	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/widget"
)

type LinksFormCtrl struct {
	formID       api.FormID
	rowID        api.RootListRowID
	client       api.RpcClient
	model        *model
	dataForm     *widget.DataForm
	moduleLogger *zap.Logger
}

func newLinksFormCtrl(rowID api.RootListRowID, model *model, moduleLogger *zap.Logger) (*LinksFormCtrl, error) {
	dataForm, err := model.createDataForm(rowID)
	if err != nil {
		moduleLogger.Error("Failed create markup for form", zap.Error(err))
		return nil, err
	}

	return &LinksFormCtrl{
		formID:       0,
		rowID:        rowID,
		client:       nil,
		model:        model,
		dataForm:     dataForm,
		moduleLogger: moduleLogger,
	}, nil
}

func (c *LinksFormCtrl) OnOpen(formID api.FormID, client api.RpcClient) *widget.DataForm {
	c.formID = formID
	c.client = client

	return c.dataForm
}

func (c *LinksFormCtrl) OnFieldCheckRequest(requestID uint32, fieldName string, jsonBody string) {
	result := true

	data := &DataModel{}
	err := json.Unmarshal([]byte(jsonBody), data)
	if err != nil {
		c.moduleLogger.Warn("Failed unmarshal json for check", zap.Error(err))
		result = false
	} else {
		result = c.model.checkItem(c.rowID, data.Name)
	}

	c.client.FieldCheckResponse(c.formID, requestID, result, "name already exists")
}

func (c *LinksFormCtrl) OnSubmit(jsonBody string) {
	data := &DataModel{}
	err := json.Unmarshal([]byte(jsonBody), data)
	if err != nil {
		c.client.UserMessage("Failed save link")
		c.moduleLogger.Warn("Failed unmarshal json for save link", zap.Error(err))
	}

	err = c.model.saveItem(c.rowID, data)
	if err != nil {
		c.client.UserMessage("Failed save link")
		c.moduleLogger.Warn("Failed save link", zap.Error(err))
		return
	}

	c.client.CloseForm(c.formID)
}