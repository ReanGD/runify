package desktop

import (
	"errors"

	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
)

type handler struct {
	mCtx *moduleCtx
}

func newHandler() *handler {
	return &handler{
		mCtx: nil,
	}
}

func (h *handler) init(mCtx *moduleCtx) error {
	h.mCtx = mCtx

	return nil
}

func (h *handler) start() {
	go func() {
		subsToClipboardRes1 := module.NewChanBoolResult()
		h.mCtx.ds.SubscribeToClipboard(true, h.mCtx.primaryCh, subsToClipboardRes1)
		subsToClipboardRes2 := module.NewChanBoolResult()
		h.mCtx.ds.SubscribeToClipboard(false, h.mCtx.clipboardCh, subsToClipboardRes2)
		subsToHotheysRes := module.NewChanBoolResult()
		h.mCtx.ds.SubscribeToHotkeys(h.mCtx.hotkeyCh, subsToHotheysRes)

		if res := <-subsToClipboardRes1.GetChannel(); !res {
			h.mCtx.errorCtx.SendError(errors.New("subscribe to primary clipboard failed"))
			return
		}

		if res := <-subsToClipboardRes2.GetChannel(); !res {
			h.mCtx.errorCtx.SendError(errors.New("subscribe to clipboard failed"))
			return
		}

		if res := <-subsToHotheysRes.GetChannel(); !res {
			h.mCtx.errorCtx.SendError(errors.New("subscribe to hotkeys failed"))
			return
		}
	}()
}

func (h *handler) onClipboardMsg(isPrimary bool, data *mime.Data) {
}

func (h *handler) onHotkeyMsg(hotkey *shortcut.Hotkey) {
}

func (h *handler) writeToClipboard(cmd *writeToClipboardCmd) {
}

func (h *handler) setHotkey(cmd *setHotkeyCmd) {
}

func (h *handler) removeHotkey(cmd *removeHotkeyCmd) {
}

func (h *handler) stop() {
}
