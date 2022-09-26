package paths

import (
	"path/filepath"
	"strings"

	"github.com/ReanGD/runify/server/gtk"
	"go.uber.org/zap"
)

func GetSysTmp() string {
	return cache.sysTmp
}

func GetUserHome() string {
	return cache.userHome
}

func GetUserConfig() string {
	return cache.userConfig
}

func GetUserCache() string {
	return cache.userCache
}

func GetAppConfig() string {
	return cache.appConfig
}

func GetAppCache() string {
	return cache.appCache
}

func GetAppIconCache() string {
	return cache.appIconCache
}

func GetXDGDataDirs() []string {
	return cache.xdgDataDirs
}

func GetXDGAppDirs() []string {
	return cache.xdgAppDirs
}

func ResolveIcon(name string, size int) string {
	if len(name) == 0 {
		return ""
	}

	if filepath.IsAbs(name) {
		if ok, _ := ExistsFile(name); ok {
			return name
		}
	}

	iconPath := cache.defaultIconTheme.LookupIcon(name, size, gtk.ICON_LOOKUP_NO_SVG)
	if len(iconPath) == 0 {
		iconPath = cache.defaultIconTheme.LookupIcon(name, size, 0)
	}

	return iconPath
}

func GetNonSvgIconPath(name string, size int, logger *zap.Logger) string {
	key := newIconKey(size, name)
	if path, ok := cache.iconPathCache[key]; ok {
		return path
	}

	path := ResolveIcon(name, size)
	if len(path) == 0 {
		cache.iconPathCache[key] = ""
		return ""
	}

	if filepath.Ext(path) != ".svg" && !strings.HasPrefix(path, "/org") {
		// TODO: resize icon
		cache.iconPathCache[key] = path
		return path
	}

	pBuf, err := cache.defaultIconTheme.LoadIcon(name, size, 0)
	if err != nil {
		logger.Info("Failed load icon", zap.String("name", name), zap.Error(err))
		return ""
	}

	path = key.toFullPath()
	err = pBuf.SavePNG(path, 0)
	if err != nil {
		logger.Info("Failed save icon", zap.String("name", name), zap.String("path", path), zap.Error(err))
		return ""
	}

	cache.iconPathCache[key] = path
	return path
}

func ExpandUser(path string) string {
	if len(path) >= 1 && path[0] == '~' {
		// TODO: need perf fix
		return filepath.Join(cache.userHome, path[1:])
	}

	return path
}
