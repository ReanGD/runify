package paths

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ReanGD/runify/server/gtk"
)

const (
	appName = "runify"
)

type cachePaths struct {
	defaultIconTheme *gtk.IconTheme
	sysTmp           string
	userHome         string
	userCache        string
	appCache         string
	appIconCache     string
	xdgDataDirs      []string
	xdgAppDirs       []string
}

var (
	cache = cachePaths{}
)

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

func Init() error {
	var ok bool
	var err error

	if len(appName) == 0 {
		return errors.New("Application name is empty")
	}

	gtk.Init()
	cache.defaultIconTheme, err = gtk.IconThemeGetDefault()
	if err != nil {
		return fmt.Errorf("Getting default theme for icons ended with error: %s", err)
	}

	if cache.sysTmp, ok = getenv("TMPDIR"); !ok {
		cache.sysTmp = "/tmp"
	}

	if cache.userHome, ok = getenv("HOME"); !ok {
		return errors.New("Env $HOME is not defined")
	}

	cache.userCache = ExpandUser(getenvDef("XDG_CACHE_HOME", filepath.Join(cache.userHome, ".cache")))

	cache.appCache = filepath.Join(cache.userCache, appName)
	if ok, err = ExistsDir(cache.appCache); err != nil {
		return fmt.Errorf("Getting info about appCache dir (%s) ended with error: %s", cache.appCache, err)
	} else if !ok {
		if err = os.MkdirAll(cache.appCache, 0700); err != nil {
			return fmt.Errorf("Creating appCache dir (%s) ended with error: %s", cache.appCache, err)
		}
	}

	cache.appIconCache = filepath.Join(cache.appCache, "icon")
	if ok, err = ExistsDir(cache.appIconCache); err != nil {
		return fmt.Errorf("Getting info about appIconCache dir (%s) ended with error: %s", cache.appIconCache, err)
	} else if !ok {
		if err = os.MkdirAll(cache.appIconCache, 0700); err != nil {
			return fmt.Errorf("Creating appIconCache dir (%s) ended with error: %s", cache.appIconCache, err)
		}
	}

	cache.xdgDataDirs = getXDGDataDirs()
	cache.xdgAppDirs = getXDGAppDirs(cache.xdgDataDirs)

	return nil
}
