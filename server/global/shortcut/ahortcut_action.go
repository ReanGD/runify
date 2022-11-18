package shortcut

import "go.uber.org/zap"

type Action struct {
	name string
}

func NewAction(name string) *Action {
	return &Action{
		name: name,
	}
}

func (a *Action) String() string {
	return a.name
}

func (a *Action) GoString() string {
	return a.String()
}

func (a *Action) ZapField() zap.Field {
	return zap.String("Action", a.String())
}

func (a *Action) ZapFieldPrefix(prefix string) zap.Field {
	return zap.String(prefix+"Hotkey", a.String())
}
