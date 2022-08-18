package paths

import (
	"os"
	"syscall"
)

func ExistsDir(path string) (bool, error) {
	modeType, err := statMode(ExpandUser(path))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return modeType == syscall.S_IFDIR, nil
}

func ExistsFile(path string) (bool, error) {
	modeType, err := statMode(ExpandUser(path))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return modeType == syscall.S_IFREG, nil
}

func ExistsSymlink(path string) (bool, error) {
	modeType, err := lStatMode(ExpandUser(path))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return modeType == syscall.S_IFLNK, nil
}

func Exists(path string) (bool, error) {
	_, err := statMode(ExpandUser(path))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
