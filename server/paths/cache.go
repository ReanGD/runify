package paths

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ReanGD/runify/server/gtk"
	"go.uber.org/zap"
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
	iconPathCache    map[iconKey]string
	sysTmp           string
	userHome         string
	userConfig       string
	userCache        string
	appConfig        string
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

func scanIcons(dirPath string, logger *zap.Logger) {
	Walk(dirPath, logger, func(fullpath string, mode PathMode) {
		if mode == ModeRegFile {
			if key, err := newIconKeyFromPath(fullpath); err != nil {
				logger.Info("Failed scan icons", zap.Error(err))
			} else {
				cache.iconPathCache[key] = fullpath
			}
		}
	})
}

func createDir(dirPath string) (existed bool, err error) {
	if existed, err = ExistsDir(dirPath); err != nil {
		return existed, fmt.Errorf("Getting info about dir (%s) ended with error: %s", dirPath, err)
	}

	if !existed {
		if err = os.MkdirAll(dirPath, 0700); err != nil {
			return existed, fmt.Errorf("Creating dir (%s) ended with error: %s", dirPath, err)
		}
	}

	return existed, err
}

func New(logger *zap.Logger) error {
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
	cache.iconPathCache = make(map[iconKey]string)

	if cache.sysTmp, ok = getenv("TMPDIR"); !ok {
		cache.sysTmp = "/tmp"
	}

	if cache.userHome, ok = getenv("HOME"); !ok {
		return errors.New("Env $HOME is not defined")
	}

	cache.userConfig = ExpandUser(getenvDef("XDG_CONFIG_HOM", filepath.Join(cache.userHome, ".config")))
	cache.userCache = ExpandUser(getenvDef("XDG_CACHE_HOME", filepath.Join(cache.userHome, ".cache")))

	cache.appConfig = filepath.Join(cache.userConfig, appName)
	if _, err := createDir(cache.appConfig); err != nil {
		return err
	}
	cache.appCache = filepath.Join(cache.userCache, appName)
	if _, err := createDir(cache.appCache); err != nil {
		return err
	}

	cache.appIconCache = filepath.Join(cache.appCache, "icon")
	if existed, err := createDir(cache.appIconCache); err != nil {
		return err
	} else if existed {
		defer scanIcons(cache.appIconCache, logger)
	}

	cache.xdgDataDirs = getXDGDataDirs()
	cache.xdgAppDirs = getXDGAppDirs(cache.xdgDataDirs)

	return nil
}
