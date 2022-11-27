package desktop

import (
	"context"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"go.uber.org/zap"
)

type moduleCtx struct {
	root         *Desktop
	ds           api.DisplayServer
	provider     api.Provider
	primaryCh    chan *mime.Data
	clipboardCh  chan *mime.Data
	hotkeyCh     chan *shortcut.Hotkey
	errorCtx     *module.ErrorCtx
	stopCtx      context.Context
	moduleLogger *zap.Logger
}

func newModuleCtx(
	root *Desktop,
	cfg *config.DesktopCfg,
	ds api.DisplayServer,
	provider api.Provider,
	ErrorCtx *module.ErrorCtx,
	moduleLogger *zap.Logger) *moduleCtx {

	return &moduleCtx{
		root:         root,
		ds:           ds,
		provider:     provider,
		primaryCh:    make(chan *mime.Data, cfg.PrimarySubscriptionChLen),
		clipboardCh:  make(chan *mime.Data, cfg.ClipboardSubscriptionChLen),
		hotkeyCh:     make(chan *shortcut.Hotkey, cfg.HotkeySubscriptionChLen),
		errorCtx:     ErrorCtx,
		stopCtx:      nil,
		moduleLogger: moduleLogger,
	}
}

func (c *moduleCtx) setStopCtx(stopCtx context.Context) {
	c.stopCtx = stopCtx
}

type writeToClipboardCmd struct {
	isPrimary bool
	data      *mime.Data
	result    api.BoolResult
}

func (c *writeToClipboardCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "writeToClipboard"),
		zap.Bool("IsPrimary", c.isPrimary),
		zap.String("Reason", reason),
		zap.String("Action", "return error"))
	c.result.SetResult(false)
}

type addShortcutCmd struct {
	action *shortcut.Action
	hotkey *shortcut.Hotkey
	result api.ErrorCodeResult
}

func (c *addShortcutCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "addShortcut"),
		c.action.ZapField(),
		c.hotkey.ZapField(),
		zap.String("Reason", reason),
		zap.String("Action", "return error"))
	c.result.SetResult(global.HotkeyBindError)
}

type removeShortcutCmd struct {
	action *shortcut.Action
	result api.VoidResult
}

func (c *removeShortcutCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "removeShortcut"),
		c.action.ZapField(),
		zap.String("Reason", reason),
		zap.String("Action", "Do nothing, just log error and return"))
	c.result.SetResult()
}

type removeShortcutWithoutCheckCmd struct {
	action *shortcut.Action
	hotkey *shortcut.Hotkey
}

func (c *removeShortcutWithoutCheckCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "removeShortcutWithoutCheck"),
		c.action.ZapField(),
		zap.String("Reason", reason),
		zap.String("Action", "Do nothing"))
}
