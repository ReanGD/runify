package module

import (
	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/pb"
)

type Provider interface {
	GetRoot() <-chan []*pb.CardItem
	GetActions(cardID uint64) <-chan *pb.Actions
	Execute(cardID uint64, actionID uint32) <-chan *pb.Result
}

type Rpc interface {
	ShowUI()
}

type X11 interface {
	WriteToClipboard(isPrimary bool, data *mime.Data)
}

type DisplayServer interface {
	SubscribeToClipboard(isPrimary bool, ch chan<- *mime.Data, result BoolResult)
	WriteToClipboard(isPrimary bool, data *mime.Data, result BoolResult)
	SubscribeToHotkeys(ch chan<- *shortcut.Hotkey, result BoolResult)
	BindHotkey(hotkey *shortcut.Hotkey, result ErrorCodeResult)
	UnbindHotkey(hotkey *shortcut.Hotkey, result BoolResult)
}

type Desktop interface {
	WriteToClipboard(isPrimary bool, data *mime.Data) <-chan bool
	SetHotkey(action *shortcut.Action, hotkey *shortcut.Hotkey) <-chan global.Error
	RemoveHotkey(action *shortcut.Action) <-chan bool
}
