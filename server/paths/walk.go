package paths

import (
	"io/fs"
	"os"
)

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
