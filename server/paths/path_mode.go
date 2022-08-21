package paths

import (
	"io/fs"
	"syscall"
)

type PathMode uint8

const (
	ModeNamedPipe   PathMode = 1
	ModeCharDevice  PathMode = 2 // character-oriented device file
	ModeBlockDevice PathMode = 3 // block-oriented device file
	ModeDir         PathMode = 4
	ModeRegFile     PathMode = 5
	ModeSymlink     PathMode = 6
	ModeSocket      PathMode = 7
	ModeUnknown     PathMode = 8
)

func dirEntryModeToPathMode(mode fs.FileMode) PathMode {
	if (mode & fs.ModeNamedPipe) != 0 {
		return ModeNamedPipe
	}

	if (mode & fs.ModeDevice) != 0 {
		if (mode & fs.ModeCharDevice) != 0 {
			return ModeCharDevice
		}
		return ModeBlockDevice
	}

	if (mode & fs.ModeDir) != 0 {
		return ModeDir
	}

	if mode == 0 {
		return ModeRegFile
	}

	if (mode & fs.ModeSymlink) != 0 {
		return ModeSymlink
	}

	if (mode & fs.ModeSocket) != 0 {
		return ModeSocket
	}

	return ModeUnknown
}

func statModeToPathMode(mode uint32) PathMode {
	switch mode {
	case syscall.S_IFIFO:
		return ModeNamedPipe
	case syscall.S_IFCHR:
		return ModeCharDevice
	case syscall.S_IFBLK:
		return ModeBlockDevice
	case syscall.S_IFDIR:
		return ModeDir
	case syscall.S_IFREG:
		return ModeRegFile
	case syscall.S_IFLNK:
		return ModeSymlink
	case syscall.S_IFSOCK:
		return ModeSocket
	default:
		return ModeUnknown
	}
}
