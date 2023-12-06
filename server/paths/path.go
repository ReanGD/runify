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

func GetDataHome() string {
	return cache.dataHome
}

func GetConfigHome() string {
	return cache.configHome
}

func GetCacheHome() string {
	return cache.cacheHome
}

func GetDataDirs() []string {
	return cache.dataDirs
}

func GetConfigDirs() []string {
	return cache.configDirs
}

// if exists(dataHome) + dataDirs
func GetAllDataDirs() []string {
	return cache.allDataDirs
}

// if exists(configHome) + configDirs
func GetAllConfigDirs() []string {
	return cache.allConfigDirs
}

func GetAppDataDir() string {
	return cache.appDataDir
}

func GetAppConfigDir() string {
	return cache.appConfigDir
}

func GetAppCacheDir() string {
	return cache.appCacheDir
}

func GetAppIconCacheDir() string {
	return cache.appIconCacheDir
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
