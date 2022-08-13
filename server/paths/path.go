package paths

import (
	"path/filepath"
	"strings"
)

func GetSysTmp() string {
	return cache.sysTmp
}

func GetUserHome() string {
	return cache.userHome
}

func ExpandUser(path string) string {
	if strings.HasPrefix(path, "~") {
		return filepath.Join(cache.userHome, path[1:])
	}

	return path
}
