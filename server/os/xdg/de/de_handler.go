package de

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rkoesters/xdg/desktop"

	"github.com/ReanGD/runify/server/global/types"
	"github.com/ReanGD/runify/server/paths"
	"go.uber.org/zap"
)

type handler struct {
	iconCache     *iconCache
	dfileCache    []*desktopFile
	subscriptions []chan<- types.DesktopFiles

	moduleLogger *zap.Logger
}

func newHandler() *handler {
	return &handler{
		iconCache:     nil,
		dfileCache:    []*desktopFile{},
		subscriptions: []chan<- types.DesktopFiles{},
		moduleLogger:  nil,
	}
}

func (h *handler) init(moduleLogger *zap.Logger) error {
	h.moduleLogger = moduleLogger

	var err error
	h.iconCache, err = newIconCache(moduleLogger)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) dfileCacheSend(ch chan<- types.DesktopFiles) {
	copy := make(types.DesktopFiles, len(h.dfileCache))
	for ind, val := range h.dfileCache {
		copy[ind] = val
	}

	select {
	case ch <- copy:
	default:
		h.moduleLogger.Warn("Failed send desktop entry file data array to subscription channel, channel is full")
	}
}

func (h *handler) update() {
	dfileCache := make([]*desktopFile, 0, len(h.dfileCache))
	h.walkXDGDesktopFiles(func(dfile *desktopFile) {
		dfileCache = append(dfileCache, dfile)
	})

	h.dfileCache = dfileCache
	for _, ch := range h.subscriptions {
		h.dfileCacheSend(ch)
	}
}

func (h *handler) subscribe(cmd *subscribeCmd) {
	h.subscriptions = append(h.subscriptions, cmd.ch)
	h.dfileCacheSend(cmd.ch)
	cmd.result.SetResult(true)
}

func (h *handler) stop() {
}

func (h *handler) walkXDGDesktopFiles(fn func(dfile *desktopFile)) {
	exists := make(map[string]struct{})
	for _, dirname := range paths.GetXDGAppDirs() {
		idStart := len(dirname) + 1
		idEnd := len(".desktop")
		paths.Walk(dirname, h.moduleLogger, func(filePath string, mode paths.PathMode) {
			if mode != paths.ModeRegFile || filepath.Ext(filePath) != ".desktop" {
				return
			}

			id := strings.ReplaceAll(filePath[idStart:len(filePath)-idEnd], "/", "_")
			if _, ok := exists[id]; ok {
				return
			}

			f, err := os.Open(filePath)
			if err != nil {
				h.moduleLogger.Info("Error open desktop entry file", zap.String("path", filePath), zap.Error(err))
				return
			}

			props, err := desktop.New(f)
			f.Close()
			if err != nil {
				h.moduleLogger.Info("Error parse desktop entry file", zap.String("path", filePath), zap.Error(err))
				return
			}

			if props.NoDisplay || props.Hidden {
				return
			}

			exists[id] = struct{}{}
			fn(newDesktopFile(
				id,
				filePath,
				h.iconCache.getNonSvgIconPath(props.Icon, 48),
				props,
			))
		})
	}
}
