package providers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ReanGD/runify/server/logger"
	"github.com/ReanGD/runify/server/paths"
	"github.com/rkoesters/xdg/desktop"
)

func walkXDGDesktopEntries(fn func(fullpath string, entry *desktop.Entry)) {
	exists := make(map[string]struct{})
	for _, dirname := range paths.GetXDGAppDirs() {
		paths.Walk(dirname, func(fullpath string, mode paths.PathMode) {
			if mode != paths.ModeRegFile || filepath.Ext(fullpath) != ".desktop" {
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
			entry.Icon = paths.GetNonSvgIconPath(entry.Icon, 48)

			fn(fullpath, entry)
		})
	}
}

func Get() {
	cnt := 0
	walkXDGDesktopEntries(func(fullpath string, entry *desktop.Entry) {
		if filepath.Ext(entry.Icon) == ".svg" {
			cnt++
		}
		fmt.Println(fullpath, entry.Icon)
	})
	fmt.Println(cnt)
}
