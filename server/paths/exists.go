package paths

import (
	"os"
	"path/filepath"
	"syscall"
)

func lStat(name string) (*syscall.Stat_t, error) {
	var err error
	var stat syscall.Stat_t
	for {
		err = syscall.Lstat(name, &stat)
		if err != syscall.EINTR {
			break
		}
	}

	return &stat, err
}

func existsType(name string, targetModeType uint32) (bool, error) {
	stat, err := lStat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	// see /usr/lib/go/src/os/stat_linux.go
	modeType := stat.Mode & syscall.S_IFMT

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

		stat, err = lStat(name)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}

			return false, err
		}

		modeType = stat.Mode & syscall.S_IFMT
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
