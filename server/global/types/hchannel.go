package types

import "reflect"

type HandledChannel struct {
	ch      reflect.Value
	handler func(interface{}) (bool, error)
}

func NewHandledChannel[T any](ch <-chan T, handler func(interface{}) (bool, error)) *HandledChannel {
	return &HandledChannel{
		ch:      reflect.ValueOf(ch),
		handler: handler,
	}
}

func (h *HandledChannel) Channel() reflect.Value {
	return h.ch
}

func (h *HandledChannel) Handle(request interface{}) (bool, error) {
	return h.handler(request)
}
