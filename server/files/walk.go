package files

import (
	"io/fs"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/ReanGD/runify/server/logger"
)

type WalkFunc func(filename string)

func expand(filename string) (string, error) {
	if len(filename) == 0 || filename[0] != '~' {
		return filename, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, filename[1:]), nil
}

func readDir(dirname string) ([]fs.DirEntry, error) {
	f, err := os.Open(dirname)
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

func walkSymlink(filename string, fn WalkFunc) {
	resolvedPath, err := filepath.EvalSymlinks(filename)
	if err != nil {
		logger.Write("Symbolic link resolution error %s: %s", filename, err)
		return
	}

	stat, err := os.Lstat(resolvedPath)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Write("The file referenced by the symbolic link was not found %s", resolvedPath)
		} else {
			logger.Write("Error getting FileInfo for a resolved symlink %s: %s", resolvedPath, err)
		}

		return
	}

	if stat.IsDir() {
		WalkDir(resolvedPath, fn)
	} else if (stat.Mode() & os.ModeSymlink) != 0 {
		walkSymlink(resolvedPath, fn)
	} else {
		fn(resolvedPath)
	}
}

func WalkDir(dirname string, fn WalkFunc) {
	children, err := readDir(dirname)
	if err != nil {
		logger.Write("Error read dir for path %s: %s", dirname, err)
		return
	}

	for _, child := range children {
		fullPath := path.Join(dirname, child.Name())
		if child.IsDir() {
			WalkDir(fullPath, fn)
		} else if (child.Type() & os.ModeSymlink) != 0 {
			walkSymlink(fullPath, fn)
		} else {
			fn(fullPath)
		}
	}
}

func Walk(root string, fn WalkFunc) {
	filename, err := expand(root)
	if err != nil {
		logger.Write("File path expand error %s: %s", root, err)
		return
	}

	stat, err := os.Lstat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Write("File path for walk is not exists %s", filename)
		} else {
			logger.Write("Error getting FileInfo for a file path for walk %s:%s", filename, err)
		}

		return
	}

	if stat.IsDir() {
		WalkDir(filename, fn)
	} else if (stat.Mode() & os.ModeSymlink) != 0 {
		walkSymlink(filename, fn)
	} else {
		fn(filename)
	}
}
