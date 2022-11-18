package desktop

import (
	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/shortcut"
	"go.uber.org/zap"
)

type writeToClipboardCmd struct {
	isPrimary bool
	data      *mime.Data
	result    chan<- bool
}

func (c *writeToClipboardCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "writeToClipboard"),
		zap.Bool("IsPrimary", c.isPrimary),
		zap.String("Reason", reason),
		zap.String("Action", "return error"))
	c.result <- false
}

type setHotkeyCmd struct {
	action *shortcut.Action
	hotkey *shortcut.Hotkey
	result chan<- global.Error
}

func (c *setHotkeyCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "setHotkey"),
		c.action.ZapField(),
		c.hotkey.ZapField(),
		zap.String("Reason", reason),
		zap.String("Action", "return error"))
	c.result <- global.HotkeyBindError
}

type removeHotkeyCmd struct {
	action *shortcut.Action
	result chan<- bool
}

func (c *removeHotkeyCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "removeHotkey"),
		c.action.ZapField(),
		zap.String("Reason", reason),
		zap.String("Action", "return error"))
	c.result <- false
}
