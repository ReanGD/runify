package widget

import "reflect"

type Model struct {
	Name       string                 `json:"name"`
	Value      interface{}            `json:"value"`
	Validators map[string]interface{} `json:"validators,omitempty"`
}

func newModel(bind *BindField, fields reflect.Value) *Model {
	return &Model{
		Name:       bind.jsonName,
		Value:      fields.Field(bind.fieldNum).Interface(),
		Validators: bind.validators,
	}
}
