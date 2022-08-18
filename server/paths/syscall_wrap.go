package paths

import "syscall"

func getenv(key string) (string, bool) {
	return syscall.Getenv(key)
}

func lStatMode(path string) (uint32, error) {
	var err error
	var stat syscall.Stat_t
	for {
		err = syscall.Lstat(path, &stat)
		if err != syscall.EINTR {
			break
		}
	}

	// see /usr/lib/go/src/os/stat_linux.go
	return stat.Mode & syscall.S_IFMT, err
}

func statMode(path string) (uint32, error) {
	var err error
	var stat syscall.Stat_t
	for {
		err = syscall.Stat(path, &stat)
		if err != syscall.EINTR {
			break
		}
	}

	// see /usr/lib/go/src/os/stat_linux.go
	return stat.Mode & syscall.S_IFMT, err
}
