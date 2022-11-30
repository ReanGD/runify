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
	sysTmp       string
	userHome     string
	userConfig   string
	userData     string
	userCache    string
	appConfig    string
	appData      string
	appCache     string
	appIconCache string
	xdgDataDirs  []string
	xdgAppDirs   []string
}

var cache = cachePaths{}

func getXDGDataDirs() []string {
	var xdgDataDirs []string

	xdgDataHome := getenvDef("XDG_DATA_HOME", "~/.local/share")
	if ok, _ := ExistsDir(xdgDataHome); ok {
		xdgDataDirs = append(xdgDataDirs, xdgDataHome)
	}

	if str, ok := getenv("XDG_DATA_DIRS"); ok {
		for _, dirPath := range strings.Split(str, ":") {
			if ok, _ := ExistsDir(dirPath); ok {
				xdgDataDirs = append(xdgDataDirs, dirPath)
			}
		}
	} else {
		if ok, _ := ExistsDir("/usr/local/share"); ok {
			xdgDataDirs = append(xdgDataDirs, "/usr/local/share")
		}
		if ok, _ := ExistsDir("/usr/share"); ok {
			xdgDataDirs = append(xdgDataDirs, "/usr/share")
		}
	}

	return xdgDataDirs
}

func getXDGAppDirs(xdgDataDirs []string) []string {
	var xdgAppDirs []string
	for _, xdgDataDir := range xdgDataDirs {
		xdgAppDir := filepath.Join(xdgDataDir, "applications")
		if ok, _ := ExistsDir(xdgAppDir); ok {
			xdgAppDirs = append(xdgAppDirs, xdgAppDir)
		}
	}

	return xdgAppDirs
}

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

	if cache.sysTmp, ok = getenv("TMPDIR"); !ok {
		cache.sysTmp = "/tmp"
	}

	if cache.userHome, ok = getenv("HOME"); !ok {
		return errors.New("Env $HOME is not defined")
	}

	cache.userConfig = ExpandUser(getenvDef("XDG_CONFIG_HOME", filepath.Join(cache.userHome, ".config")))
	cache.userData = ExpandUser(getenvDef("XDG_DATA_HOME", filepath.Join(cache.userHome, ".local", "share")))
	cache.userCache = ExpandUser(getenvDef("XDG_CACHE_HOME", filepath.Join(cache.userHome, ".cache")))

	cache.appConfig = filepath.Join(cache.userConfig, appName)
	if _, err := createDir(cache.appConfig); err != nil {
		return err
	}
	setenv("RUNIFY_CONFIG_DIR", cache.appConfig)

	cache.appData = filepath.Join(cache.userData, appName)
	if _, err := createDir(cache.appData); err != nil {
		return err
	}
	setenv("RUNIFY_DATA_DIR", cache.appData)

	cache.appCache = filepath.Join(cache.userCache, appName)
	if _, err := createDir(cache.appCache); err != nil {
		return err
	}
	setenv("RUNIFY_CACHE_DIR", cache.appCache)

	cache.appIconCache = filepath.Join(cache.appCache, "icon")
	if _, err := createDir(cache.appIconCache); err != nil {
		return err
	}

	cache.xdgDataDirs = getXDGDataDirs()
	cache.xdgAppDirs = getXDGAppDirs(cache.xdgDataDirs)

	return nil
}
