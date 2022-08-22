package paths

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ReanGD/runify/server/gtk"
	"github.com/ReanGD/runify/server/logger"
)

const (
	appName = "runify"
)

type iconKey struct {
	size int
	name string
}

func newIconKey(size int, name string) iconKey {
	return iconKey{
		size: size,
		name: name,
	}
}

func newIconKeyFromPath(fullpath string) (iconKey, error) {
	_, filename := filepath.Split(fullpath)
	extIndex := strings.LastIndexByte(filename, '.')
	if extIndex == -1 {
		return iconKey{}, fmt.Errorf("Wrong icon cache filename %s, error: not found extension", fullpath)
	}
	filename = filename[:extIndex]

	sepIndex := strings.IndexByte(filename, '_')
	if sepIndex == -1 {
		return iconKey{}, fmt.Errorf("Wrong icon cache filename %s, error: not found separator", fullpath)
	}

	size, err := strconv.Atoi(filename[:sepIndex])
	if err != nil {
		return iconKey{}, fmt.Errorf("Wrong icon cache filename %s, error: can't parse size: %s", fullpath, err)
	}

	return iconKey{
		size: size,
		name: filename[sepIndex+1:],
	}, nil
}

func (k iconKey) toFullPath() string {
	return filepath.Join(GetAppIconCache(), fmt.Sprintf("%d_%s.png", k.size, k.name))
}

type cachePaths struct {
	defaultIconTheme *gtk.IconTheme
	sysTmp           string
	userHome         string
	userCache        string
	appCache         string
	appIconCache     string
	iconPathCache    map[iconKey]string
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

func scanIcons(dirPath string) {
	Walk(dirPath, func(fullpath string, mode PathMode) {
		if mode == ModeRegFile {
			if key, err := newIconKeyFromPath(fullpath); err != nil {
				logger.Write(err.Error())
			} else {
				cache.iconPathCache[key] = fullpath
			}
		}
	})
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
	cache.iconPathCache = make(map[iconKey]string)
	if ok, err = ExistsDir(cache.appIconCache); err != nil {
		return fmt.Errorf("Getting info about appIconCache dir (%s) ended with error: %s", cache.appIconCache, err)
	} else if !ok {
		if err = os.MkdirAll(cache.appIconCache, 0700); err != nil {
			return fmt.Errorf("Creating appIconCache dir (%s) ended with error: %s", cache.appIconCache, err)
		}
	} else {
		defer scanIcons(cache.appIconCache)
	}

	cache.xdgDataDirs = getXDGDataDirs()
	cache.xdgAppDirs = getXDGAppDirs(cache.xdgDataDirs)

	return nil
}
