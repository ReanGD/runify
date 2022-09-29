package paths

import (
	"os"
	"path/filepath"
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

func ExpandUser(path string) string {
	if len(path) >= 1 && path[0] == '~' {
		// TODO: need perf fix
		return filepath.Join(cache.userHome, path[1:])
	}

	return path
}

func ExpandAll(path string) string {
	return os.ExpandEnv(ExpandUser(path))
}
