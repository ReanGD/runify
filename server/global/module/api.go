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
	SubscribeToClipboard(isPrimary bool, ch chan<- *mime.Data) <-chan bool
	WriteToClipboard(isPrimary bool, data *mime.Data) <-chan bool
	SubscribeToHotkeys(ch chan<- *shortcut.Hotkey) <-chan bool
	BindHotkey(hotkey *shortcut.Hotkey) <-chan global.Error
	UnbindHotkey(hotkey *shortcut.Hotkey) <-chan bool
}
