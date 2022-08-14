package paths

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ReanGD/runify/server/logger"
)

type WalkFunc func(path string)

func readDir(dirPath string) ([]fs.DirEntry, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	dirs, err := f.ReadDir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return dirs, nil
}

func walkFiles(dirPath string, fn WalkFunc) {
	children, err := readDir(dirPath)
	if err != nil {
		logger.Write("Failed read dir items for path %s, error: %s", dirPath, err)
		return
	}

	for _, child := range children {
		childPath := filepath.Join(dirPath, child.Name())
		isDir := child.IsDir()
		if (child.Type() & os.ModeSymlink) != 0 {
			if isDir, err = ExistsDir(childPath); err != nil {
				logger.Write("Failed call ExistsDir for path %s, error: %s", dirPath, err)
				continue
			}
		}

		if isDir {
			walkFiles(childPath, fn)
		} else {
			fn(childPath)
		}
	}
}

func WalkFiles(dirPath string, fn WalkFunc) {
	dirPath = ExpandUser(dirPath)
	if isDir, err := ExistsDir(dirPath); err != nil {
		logger.Write("Failed call ExistsDir for path %s, error: %s", dirPath, err)
	} else if isDir {
		walkFiles(dirPath, fn)
	} else {
		fn(dirPath)
	}
}
