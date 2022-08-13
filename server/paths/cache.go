package paths

import (
	"errors"
)

type cachePaths struct {
	sysTmp   string
	userHome string
}

var (
	cache = cachePaths{}
)

func Init() error {
	var ok bool

	if cache.sysTmp, ok = getenv("TMPDIR"); !ok {
		cache.sysTmp = "/tmp"
	}

	if cache.userHome, ok = getenv("HOME"); !ok {
		return errors.New("Env $HOME is not defined")
	}

	return nil
}
