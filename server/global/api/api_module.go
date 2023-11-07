package api

import (
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/global/types"
)

type Provider interface {
	Activate(action *shortcut.Action)
}

type Rpc interface {
	OpenRootList(ctrl RootListCtrl)
}

type DisplayServer interface {
	SubscribeToClipboard(isPrimary bool, ch chan<- *mime.Data, result BoolResult)
	WriteToClipboard(isPrimary bool, data *mime.Data, result BoolResult)
	SubscribeToHotkeys(ch chan<- *shortcut.Hotkey, result BoolResult)
	BindHotkey(hotkey *shortcut.Hotkey, result ErrorCodeResult)
	UnbindHotkey(hotkey *shortcut.Hotkey, result BoolResult)
}

type XDGDesktopEntry interface {
	Update()
	Subscribe(ch chan<- types.DesktopEntries, result BoolResult)
}

type Desktop interface {
	WriteToClipboard(isPrimary bool, data *mime.Data, result BoolResult)
	AddShortcut(action *shortcut.Action, hotkey *shortcut.Hotkey, result ErrorCodeResult)
	RemoveShortcut(action *shortcut.Action, result VoidResult)
}
