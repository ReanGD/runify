package providers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ReanGD/runify/server/files"
	"github.com/ReanGD/runify/server/logger"
	"github.com/rkoesters/xdg/desktop"
)

func getXDGDataDirs() []string {
	var xdgDataDirs []string

	if xdgDataHome, ok := os.LookupEnv("XDG_DATA_HOME"); ok {
		xdgDataDirs = append(xdgDataDirs, xdgDataHome)
	} else {
		xdgDataDirs = append(xdgDataDirs, path.Join("~", ".local", "share"))
	}

	if str, ok := os.LookupEnv("XDG_DATA_DIRS"); ok {
		xdgDataDirs = append(xdgDataDirs, strings.Split(str, ":")...)
	} else {
		xdgDataDirs = append(xdgDataDirs, "/usr/local/share", "/usr/share")
	}

	return xdgDataDirs
}

func getXDGAppDirs() []string {
	xdgDataDirs := getXDGDataDirs()
	xdgAppDirs := []string{}
	for _, dirname := range xdgDataDirs {
		fullpath := path.Join(dirname, "applications")
		if _, err := os.Lstat(fullpath); os.IsNotExist(err) {
			continue
		}

		xdgAppDirs = append(xdgAppDirs, fullpath)
	}

	return xdgAppDirs
}

func walkXDGDesktopEntries(iconTheme *IconTheme, fn func(fullpath string, entry *desktop.Entry)) {
	exists := make(map[string]struct{})
	for _, dirname := range getXDGAppDirs() {
		files.Walk(dirname, func(fullpath string) {
			if filepath.Ext(fullpath) != ".desktop" {
				return
			}

			_, filename := filepath.Split(fullpath)
			if _, ok := exists[filename]; ok {
				return
			}
			exists[filename] = struct{}{}

			f, err := os.Open(fullpath)
			if err != nil {
				logger.Write("Error open file %s: %s", fullpath, err)
				return
			}

			entry, err := desktop.New(f)
			f.Close()
			if err != nil {
				logger.Write("Error parse desktop entry %s: %s", fullpath, err)
				return
			}

			if entry.NoDisplay || entry.Hidden {
				return
			}

			if len(entry.Icon) != 0 {
				if _, err = os.Stat(entry.Icon); err != nil {
					entry.Icon = iconTheme.LookupIcon(entry.Icon, 48, 0)
				}
			}

			fn(fullpath, entry)
		})
	}
}

func Get() {
	Init()
	iconTheme, err := IconThemeGetDefault()
	if err != nil {
		logger.Write("Error get default icon theme: %s", err)
		return
	}

	walkXDGDesktopEntries(iconTheme, func(fullpath string, entry *desktop.Entry) {
		fmt.Println(fullpath)
	})
}
