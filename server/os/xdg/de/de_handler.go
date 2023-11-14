package de

import (
	"os"
	"path/filepath"

	"github.com/rkoesters/xdg/desktop"

	"github.com/ReanGD/runify/server/global/types"
	"github.com/ReanGD/runify/server/paths"
	"go.uber.org/zap"
)

type handler struct {
	iconCache     *iconCache
	entries       types.DesktopEntries
	subscriptions []chan<- types.DesktopEntries

	moduleLogger *zap.Logger
}

func newHandler() *handler {
	return &handler{
		iconCache:     nil,
		entries:       types.DesktopEntries{},
		subscriptions: []chan<- types.DesktopEntries{},
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

func (h *handler) sendToSubscribers() {
	for _, ch := range h.subscriptions {
		select {
		case ch <- h.entries.Clone():
		default:
			h.moduleLogger.Warn("Failed send desktop entry data array to subscription channel, channel is full")
		}
	}
}

func (h *handler) update() {
	entries := make(types.DesktopEntries, 0, len(h.entries))
	h.walkXDGDesktopEntries(func(entry *types.DesktopEntry) {
		entries = append(entries, entry)
	})

	h.entries = entries
	h.sendToSubscribers()
}

func (h *handler) subscribe(cmd *subscribeCmd) {
	h.subscriptions = append(h.subscriptions, cmd.ch)

	select {
	case cmd.ch <- h.entries.Clone():
	default:
		h.moduleLogger.Warn("Failed send desktop entry data array to subscription channel, channel is full")
	}

	cmd.result.SetResult(true)
}

func (h *handler) stop() {
}

func (h *handler) walkXDGDesktopEntries(fn func(entry *types.DesktopEntry)) {
	exists := make(map[string]struct{})
	for _, dirname := range paths.GetXDGAppDirs() {
		paths.Walk(dirname, h.moduleLogger, func(filePath string, mode paths.PathMode) {
			if mode != paths.ModeRegFile || filepath.Ext(filePath) != ".desktop" {
				return
			}

			_, filename := filepath.Split(filePath)
			if _, ok := exists[filename]; ok {
				return
			}
			exists[filename] = struct{}{}

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

			fn(types.NewDesktopEntry(
				filePath,
				h.iconCache.getNonSvgIconPath(props.Icon, 48),
				props.Name,
				props.Exec,
				props.Terminal,
			))
		})
	}
}
