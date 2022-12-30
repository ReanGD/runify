package widget

import (
	"reflect"
)

type Widget interface {
	check() error
	buildModel(fields reflect.Value) []*Model
	MarshalJSON() ([]byte, error)
}

func Int(value int) *int {
	return &value
}

func Bool(value bool) *bool {
	return &value
}

func String(value string) *string {
	return &value
}
