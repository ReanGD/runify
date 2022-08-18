package paths

import "syscall"

func getenv(key string) (string, bool) {
	return syscall.Getenv(key)
}

func getenvDef(key string, defValue string) string {
	if res, ok := syscall.Getenv(key); ok {
		return res
	}

	return defValue
}

// S_IFMT   = 0xf000 0b1111+000000000000 - mask
// S_IFIFO  = 0x1000 0b0001+000000000000 - named pipe
// S_IFCHR  = 0x2000 0b0010+000000000000 - ? character-oriented device file
// S_IFBLK  = 0x6000 0b0110+000000000000 - ? block-oriented device file
// S_IFDIR  = 0x4000 0b0100+000000000000 - directory
// S_IFREG  = 0x8000 0b1000+000000000000 - regular file
// S_IFLNK  = 0xa000 0b1010+000000000000 - Symlink
// S_IFSOCK = 0xc000 0b1100+000000000000 - Socket
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
