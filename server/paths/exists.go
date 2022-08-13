package paths

import (
	"os"
	"path/filepath"
	"syscall"
)

func existsType(path string, targetModeType uint32) (bool, error) {
	modeType, err := lStatMode(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	// is symlink and found not symlink
	if modeType == syscall.S_IFLNK && modeType != targetModeType {
		// resolve symlink
		path, err = filepath.EvalSymlinks(path)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}

			return false, err
		}

		if len(path) == 0 {
			return false, nil
		}

		if modeType, err = lStatMode(path); err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}

			return false, err
		}
	}

	return modeType == targetModeType || targetModeType == syscall.S_IFMT, nil
}

func ExistsDir(path string) (bool, error) {
	return existsType(path, syscall.S_IFDIR)
}

func ExistsFile(path string) (bool, error) {
	return existsType(path, syscall.S_IFREG)
}

func ExistsSymlink(path string) (bool, error) {
	return existsType(path, syscall.S_IFLNK)
}

func Exists(path string) (bool, error) {
	return existsType(path, syscall.S_IFMT)
}
