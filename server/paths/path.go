package paths

import (
	"path/filepath"

	"github.com/ReanGD/runify/server/gtk"
)

func GetSysTmp() string {
	return cache.sysTmp
}

func GetUserHome() string {
	return cache.userHome
}

func GetXDGDataDirs() []string {
	return cache.xdgDataDirs
}

func GetXDGAppDirs() []string {
	return cache.xdgAppDirs
}

func GetIconPath(name string, size int) string {
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

func ExpandUser(path string) string {
	if len(path) >= 1 && path[0] == '~' {
		// TODO: need perf fix
		return filepath.Join(cache.userHome, path[1:])
	}

	return path
}
