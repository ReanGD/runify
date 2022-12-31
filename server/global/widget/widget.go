package widget

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/goccy/go-json"
)

type DataForm struct {
	form  *Form
	model string
}

func NewDataForm(form *Form, data any) (*DataForm, error) {
	model, err := form.buildModel(data)
	if err != nil {
		return nil, err
	}

	return &DataForm{
		form:  form,
		model: string(model),
	}, nil
}

func (f *DataForm) GetModel() string {
	return f.model
}

func (f *DataForm) GetMarkup() string {
	return f.form.markup
}

type Form struct {
	child  Widget
	markup string
}

func NewForm[TData any](fn func(bind Bind, fields *TData) Widget) (*Form, error) {
	fields := new(TData)

	binder := newBinder(reflect.TypeOf(fields), uintptr(unsafe.Pointer(fields)))
	child := fn(binder.Bind, fields)
	markup, err := json.Marshal(child)
	if err != nil {
		return nil, err
	}

	return &Form{
		child:  child,
		markup: string(markup),
	}, nil
}

func (f *Form) buildMarkup() ([]byte, error) {
	return json.Marshal(f.child)
}

func (f *Form) buildModel(data any) ([]byte, error) {
	fields := reflect.ValueOf(data).Elem()
	models := f.child.buildModel(fields)
	return json.Marshal(models)
}

type WidgetTypeColumn struct{}

func (WidgetTypeColumn) MarshalJSON() ([]byte, error) {
	return []byte(`"Column"`), nil
}

type Column struct {
	WidgetTypeColumn `json:"type"`
	Children         []Widget `json:"children,omitempty"`
}

func (w *Column) buildModel(fields reflect.Value) []*Model {
	res := []*Model{}
	for _, child := range w.Children {
		if arr := child.buildModel(fields); arr != nil {
			res = append(res, arr...)
		}
	}

	return res
}

func (w *Column) MarshalJSON() ([]byte, error) {
	type local *Column
	return json.Marshal(local(w))
}

type WidgetTypeText struct{}

func (WidgetTypeText) MarshalJSON() ([]byte, error) {
	return []byte(`"Text"`), nil
}

type Text struct {
	WidgetTypeText `json:"type"`
	Data           string `json:"data"`
}

func (w *Text) buildModel(fields reflect.Value) []*Model {
	return nil
}

func (w *Text) MarshalJSON() ([]byte, error) {
	type local *Text
	return json.Marshal(local(w))
}

type WidgetTypeTextField struct{}

func (WidgetTypeTextField) MarshalJSON() ([]byte, error) {
	return []byte(`"TextField"`), nil
}

type TextField struct {
	WidgetTypeTextField `json:"type"`
	Bind                *BindField `json:"bind"`
	ObscureText         bool       `json:"obscureText,omitempty"`
	ReadOnly            bool       `json:"readOnly,omitempty"`
	MaxLines            *int       `json:"maxLines,omitempty"`
}

func (w *TextField) buildModel(fields reflect.Value) []*Model {
	return []*Model{
		newModel(w.Bind, fields),
	}
}

func (w *TextField) MarshalJSON() ([]byte, error) {
	if w.Bind != nil {
		return nil, errors.New("bind field for widget TextField is required")
	}

	type local *TextField
	return json.Marshal(local(w))
}
