package paths

import (
	"errors"
	"path/filepath"
	"strings"
)

type cachePaths struct {
	sysTmp      string
	userHome    string
	xdgDataDirs []string
	xdgAppDirs  []string
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

	if cache.sysTmp, ok = getenv("TMPDIR"); !ok {
		cache.sysTmp = "/tmp"
	}

	if cache.userHome, ok = getenv("HOME"); !ok {
		return errors.New("Env $HOME is not defined")
	}

	cache.xdgDataDirs = getXDGDataDirs()
	cache.xdgAppDirs = getXDGAppDirs(cache.xdgDataDirs)

	return nil
}
