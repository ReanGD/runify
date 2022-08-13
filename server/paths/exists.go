package paths

import (
	"os"
	"path/filepath"
	"syscall"
)

func existsType(name string, targetModeType uint32) (bool, error) {
	modeType, err := lStatMode(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	// is symlink and found not symlink
	if modeType == syscall.S_IFLNK && modeType != targetModeType {
		// resolve symlink
		name, err = filepath.EvalSymlinks(name)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}

			return false, err
		}

		if len(name) == 0 {
			return false, nil
		}

		if modeType, err = lStatMode(name); err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}

			return false, err
		}
	}

	return modeType == targetModeType || targetModeType == syscall.S_IFMT, nil
}

func ExistsDir(name string) (bool, error) {
	return existsType(name, syscall.S_IFDIR)
}

func ExistsFile(name string) (bool, error) {
	return existsType(name, syscall.S_IFREG)
}

func ExistsSymlink(name string) (bool, error) {
	return existsType(name, syscall.S_IFLNK)
}

func Exists(name string) (bool, error) {
	return existsType(name, syscall.S_IFMT)
}
