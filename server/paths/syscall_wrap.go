package paths

import "syscall"

func getenv(key string) (string, bool) {
	return syscall.Getenv(key)
}

func lStatMode(name string) (uint32, error) {
	var err error
	var stat syscall.Stat_t
	for {
		err = syscall.Lstat(name, &stat)
		if err != syscall.EINTR {
			break
		}
	}

	// see /usr/lib/go/src/os/stat_linux.go
	return stat.Mode & syscall.S_IFMT, err
}
