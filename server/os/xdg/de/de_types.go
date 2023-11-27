package de

import (
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/types"
	"github.com/rkoesters/xdg/desktop"
	"go.uber.org/zap"
)

type desktopFile struct {
	// The unique id
	id string

	// The full path to the desktop entry file
	filePath string

	// The full path to the icon file
	iconPath string

	props *desktop.Entry
}

func newDesktopFile(id, filePath, iconPath string, props *desktop.Entry) *desktopFile {
	return &desktopFile{
		id:       id,
		filePath: filePath,
		iconPath: iconPath,
		props:    props,
	}
}

func (f *desktopFile) ID() string {
	return f.id
}

func (f *desktopFile) FilePath() string {
	return f.filePath
}

func (f *desktopFile) IconPath() string {
	return f.iconPath
}

func (f *desktopFile) Name() string {
	return f.props.Name
}

func (f *desktopFile) Exec() string {
	return f.props.Exec
}

func (f *desktopFile) InTerminal() bool {
	return f.props.Terminal
}

type updateCmd struct{}

func (c *updateCmd) OnRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "update"),
		zap.String("Reason", reason),
		zap.String("Action", "do nothing"))
}

type subscribeCmd struct {
	ch     chan<- types.DesktopFiles
	result api.BoolResult
}

func (c *subscribeCmd) OnRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "subscribe"),
		zap.String("Reason", reason),
		zap.String("Action", "subscription not activated, return false"))
	c.result.SetResult(false)
}
