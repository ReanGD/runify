package paths

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	appName = "runify"
)

type cachePaths struct {
	sysTmp          string
	userHome        string
	dataHome        string
	configHome      string
	cacheHome       string
	dataDirs        []string
	configDirs      []string
	allDataDirs     []string
	allConfigDirs   []string
	appDataDir      string
	appConfigDir    string
	appCacheDir     string
	appIconCacheDir string
}

var cache = cachePaths{}

func createDir(dirPath string) (existed bool, err error) {
	if existed, err = ExistsDir(dirPath); err != nil {
		return existed, fmt.Errorf("Getting info about dir (%s) ended with error: %s", dirPath, err)
	}

	if !existed {
		if err = os.MkdirAll(dirPath, 0o700); err != nil {
			return existed, fmt.Errorf("Creating dir (%s) ended with error: %s", dirPath, err)
		}
	}

	return existed, err
}

func New() error {
	var ok bool

	if len(appName) == 0 {
		return errors.New("Application name is empty")
	}

	if cache.sysTmp, ok = GetEnv("TMPDIR"); !ok {
		cache.sysTmp = "/tmp"
	}

	if cache.userHome, ok = GetEnv("HOME"); !ok {
		return errors.New("Env $HOME is not defined")
	}

	cache.dataHome = ExpandUser(GetEnvDef("XDG_DATA_HOME", "~/.local/share"))
	cache.configHome = ExpandUser(GetEnvDef("XDG_CONFIG_HOME", "~/.config"))
	cache.cacheHome = ExpandUser(GetEnvDef("XDG_CACHE_HOME", "~/.cache"))

	envDataDirs := GetEnvDef("XDG_DATA_DIRS", "/usr/local/share:/usr/share")
	for _, dirPath := range strings.Split(envDataDirs, ":") {
		dirPath = ExpandUser(dirPath)
		if ok, _ := ExistsDir(dirPath); ok {
			cache.dataDirs = append(cache.dataDirs, dirPath)
		}
	}

	envConfigDirs := GetEnvDef("XDG_CONFIG_DIRS", "/etc/xdg")
	for _, dirPath := range strings.Split(envConfigDirs, ":") {
		dirPath = ExpandUser(dirPath)
		if ok, _ := ExistsDir(dirPath); ok {
			cache.configDirs = append(cache.configDirs, dirPath)
		}
	}

	if ok, _ := ExistsDir(cache.dataHome); ok {
		cache.allDataDirs = append(cache.allDataDirs, cache.dataHome)
	}
	cache.allDataDirs = append(cache.allDataDirs, cache.dataDirs...)

	if ok, _ := ExistsDir(cache.configHome); ok {
		cache.allConfigDirs = append(cache.allConfigDirs, cache.configHome)
	}
	cache.allConfigDirs = append(cache.allConfigDirs, cache.configDirs...)

	cache.appDataDir = filepath.Join(cache.dataHome, appName)
	if _, err := createDir(cache.appDataDir); err != nil {
		return err
	}
	setenv("RUNIFY_DATA_DIR", cache.appDataDir)

	cache.appConfigDir = filepath.Join(cache.configHome, appName)
	if _, err := createDir(cache.appConfigDir); err != nil {
		return err
	}
	setenv("RUNIFY_CONFIG_DIR", cache.appConfigDir)

	cache.appCacheDir = filepath.Join(cache.cacheHome, appName)
	if _, err := createDir(cache.appCacheDir); err != nil {
		return err
	}
	setenv("RUNIFY_CACHE_DIR", cache.appCacheDir)

	cache.appIconCacheDir = filepath.Join(cache.appCacheDir, "icon")
	if _, err := createDir(cache.appIconCacheDir); err != nil {
		return err
	}

	return nil
}
