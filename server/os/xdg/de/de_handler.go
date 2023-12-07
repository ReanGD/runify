package de

import (
	"path/filepath"
	"strings"

	"github.com/rkoesters/xdg/keyfile"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/types"
	"github.com/ReanGD/runify/server/paths"
	"go.uber.org/zap"
)

type handler struct {
	mimeStorage   *mimeStorage
	iconCache     *iconCache
	dfileCache    []*desktopFile
	subscriptions []chan<- types.DesktopFiles
	mainLocale    keyfile.Locale
	dopLocale     keyfile.Locale

	xdgAppDirs      []string
	xdgMimeAppFiles []string

	moduleLogger *zap.Logger
}

func newHandler() *handler {
	return &handler{
		mimeStorage:     newMimeStorage(),
		iconCache:       nil,
		dfileCache:      []*desktopFile{},
		subscriptions:   []chan<- types.DesktopFiles{},
		mainLocale:      keyfile.Locale{},
		dopLocale:       keyfile.Locale{},
		xdgAppDirs:      getXDGAppDirs(),
		xdgMimeAppFiles: getXDGMimeAppFiles(),
		moduleLogger:    nil,
	}
}

func (h *handler) init(cfg *config.Configuration, moduleLogger *zap.Logger) error {
	h.moduleLogger = moduleLogger

	var err error
	h.iconCache, err = newIconCache(moduleLogger)
	if err != nil {
		return err
	}

	h.mainLocale, err = keyfile.ParseLocale(cfg.System.MainLocale)
	if err != nil {
		h.moduleLogger.Warn("Error parse main locale",
			zap.String("action", "use default locale"),
			zap.String("locale", cfg.System.MainLocale),
			zap.Error(err))
		h.mainLocale = keyfile.DefaultLocale()
	}

	h.dopLocale, err = keyfile.ParseLocale(cfg.System.DopLocale)
	if err != nil {
		h.moduleLogger.Warn("Error parse dop locale",
			zap.String("action", "use default locale"),
			zap.String("locale", cfg.System.DopLocale),
			zap.Error(err))
		h.dopLocale = keyfile.DefaultLocale()
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
	h.readXDGFiles()
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

func (h *handler) readXDGFiles() {
	exists := make(map[string]struct{})
	mimeStorage := newMimeStorage()
	dfileCache := make([]*desktopFile, 0, len(h.dfileCache))
	for _, dirname := range h.xdgAppDirs {
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

			dFile := newDesktopFile(id, filePath, h.mainLocale, h.dopLocale, mimeStorage, h.moduleLogger)
			if dFile == nil {
				return
			}
			dFile.iconPath = h.iconCache.getNonSvgIconPath(dFile.icon, 48)

			exists[id] = struct{}{}
			dfileCache = append(dfileCache, dFile)
		})
	}

	h.dfileCache = dfileCache
	h.mimeStorage = mimeStorage
}
