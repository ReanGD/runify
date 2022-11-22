package desktop

import (
	"errors"

	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"go.uber.org/zap"
)

type shortcutData struct {
	hotkey *shortcut.Hotkey
	action *shortcut.Action
}

type handler struct {
	mCtx             *moduleCtx
	ds               module.DisplayServer
	moduleLogger     *zap.Logger
	shortcutByAction map[shortcut.ActionId]*shortcutData
	shortcutByHotkey map[shortcut.HotkeyId]*shortcutData
}

func newHandler() *handler {
	return &handler{
		mCtx:             nil,
		ds:               nil,
		moduleLogger:     nil,
		shortcutByAction: make(map[shortcut.ActionId]*shortcutData),
		shortcutByHotkey: make(map[shortcut.HotkeyId]*shortcutData),
	}
}

func (h *handler) init(mCtx *moduleCtx) error {
	h.mCtx = mCtx
	h.ds = mCtx.ds
	h.moduleLogger = mCtx.moduleLogger

	return nil
}

func (h *handler) start() {
	go func() {
		subsToClipboardRes1 := module.NewChanBoolResult()
		h.ds.SubscribeToClipboard(true, h.mCtx.primaryCh, subsToClipboardRes1)
		subsToClipboardRes2 := module.NewChanBoolResult()
		h.ds.SubscribeToClipboard(false, h.mCtx.clipboardCh, subsToClipboardRes2)
		subsToHotheysRes := module.NewChanBoolResult()
		h.ds.SubscribeToHotkeys(h.mCtx.hotkeyCh, subsToHotheysRes)

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
	// TODO: save clipboard for proivder
}

func (h *handler) onHotkeyMsg(hotkey *shortcut.Hotkey) {
	if shortcutData, ok := h.shortcutByHotkey[hotkey.Id()]; ok {
		h.mCtx.provider.Activate(shortcutData.action)
	}
}

func (h *handler) writeToClipboard(cmd *writeToClipboardCmd) {
	h.ds.WriteToClipboard(cmd.isPrimary, cmd.data, cmd.result)
}

func (h *handler) bindHotkey(shortcutData *shortcutData, cmdResult module.ErrorCodeResult) {
	action := shortcutData.action
	hotkey := shortcutData.hotkey

	bindResult := module.NewChanErrorCodeResult()
	h.ds.BindHotkey(hotkey, bindResult)

	stopCtx := h.mCtx.stopCtx
	moduleLogger := h.moduleLogger
	fields := []zap.Field{hotkey.ZapField(), action.ZapField()}
	h.moduleLogger.Debug("Start add hotkey for shortcut", fields...)
	root := h.mCtx.root
	go func() {
		select {
		case <-stopCtx.Done():
		case res := <-bindResult.GetChannel():
			if res == global.Success {
				moduleLogger.Debug("Finish add hotkey for shortcut success", fields...)
			} else {
				moduleLogger.Warn("Finish add hotkey for shortcut failed", append(fields, res.ZapField())...)
				root.removeShortcutWithoutCheck(action, hotkey)
			}
			cmdResult.SetResult(res)
		}
	}()
}

func (h *handler) unbindHotkey(shortcutData *shortcutData, cmdResult module.VoidResult) {
	unbindResult := module.NewChanBoolResult()
	h.ds.UnbindHotkey(shortcutData.hotkey, unbindResult)

	stopCtx := h.mCtx.stopCtx
	moduleLogger := h.moduleLogger
	fields := []zap.Field{shortcutData.hotkey.ZapField(), shortcutData.action.ZapField()}
	h.moduleLogger.Debug("Start remove hotkey for shortcut",
		shortcutData.hotkey.ZapField(),
		shortcutData.action.ZapField(),
	)
	go func() {
		select {
		case <-stopCtx.Done():
		case res := <-unbindResult.GetChannel():
			if res {
				moduleLogger.Debug("Finish remove hotkey for shortcut success", fields...)
			} else {
				moduleLogger.Warn("Finish remove hotkey for shortcut failed", fields...)
			}
			cmdResult.SetResult()
		}
	}()
}

func (h *handler) addShortcut(cmd *addShortcutCmd) {
	hotkeyId := cmd.hotkey.Id()
	actionId := cmd.action.Id()

	// check if hotkey already exists
	if shortcutData, ok := h.shortcutByHotkey[hotkeyId]; ok {
		if shortcutData.action.Id() == actionId {
			cmd.result.SetResult(global.Success)
		} else {
			cmd.result.SetResult(global.HotkeyUsesByRunify)
		}
		return
	}

	// check if action already exists
	if shortcutData, ok := h.shortcutByAction[actionId]; ok {
		delete(h.shortcutByAction, actionId)
		delete(h.shortcutByHotkey, shortcutData.hotkey.Id())
		h.moduleLogger.Debug("Remove previous shortcut",
			shortcutData.hotkey.ZapField(),
			shortcutData.action.ZapField(),
		)

		h.unbindHotkey(shortcutData, module.NewFuncVoidResult(func() {}))
	}

	shortcutData := &shortcutData{
		hotkey: cmd.hotkey,
		action: cmd.action,
	}
	h.shortcutByAction[cmd.action.Id()] = shortcutData
	h.shortcutByHotkey[cmd.hotkey.Id()] = shortcutData
	h.bindHotkey(shortcutData, cmd.result)
}

func (h *handler) removeShortcut(cmd *removeShortcutCmd) {
	actionId := cmd.action.Id()
	shortcutData, ok := h.shortcutByAction[actionId]
	if !ok {
		cmd.result.SetResult()
		return
	}

	delete(h.shortcutByAction, actionId)
	delete(h.shortcutByHotkey, shortcutData.hotkey.Id())
	h.moduleLogger.Debug("Remove shortcut",
		shortcutData.hotkey.ZapField(),
		shortcutData.action.ZapField(),
	)

	h.unbindHotkey(shortcutData, cmd.result)
}

func (h *handler) removeShortcutWithoutCheck(cmd *removeShortcutWithoutCheckCmd) {
	shortcutData, ok := h.shortcutByAction[cmd.action.Id()]
	if !ok || shortcutData.hotkey.Id() != cmd.hotkey.Id() || shortcutData.action.Id() != cmd.action.Id() {
		return
	}

	delete(h.shortcutByAction, cmd.action.Id())
	delete(h.shortcutByHotkey, cmd.hotkey.Id())
	h.moduleLogger.Debug("Remove shortcut after failed to bind hotkey",
		shortcutData.hotkey.ZapField(),
		shortcutData.action.ZapField(),
	)
}

func (h *handler) stop() {
}
