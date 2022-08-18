package paths

import (
	"path/filepath"
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

func ExpandUser(path string) string {
	if len(path) >= 1 && path[0] == '~' {
		// TODO: need perf fix
		return filepath.Join(cache.userHome, path[1:])
	}

	return path
}
