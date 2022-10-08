package paths

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"go.uber.org/zap"
)

const (
	runePathSeparator = '/'
	strPathSeparator  = "/"
	linksWalkedMax    = 255
)

type WalkFunc func(path string, mode PathMode)

// func openDir(name string) (*os.File, error) {
// 	var fd int
// 	var err error
// 	for {
// 		fd, err = syscall.Open(name, syscall.O_CLOEXEC, 0)
// 		if err != syscall.EINTR {
// 			break
// 		}
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	f := os.newFile(uintptr(fd), name, os.kindOpenFile)
// 	f.appendMode = false
// 	return f, nil
// }

func readDir(dirPath string) ([]fs.DirEntry, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	dirs, err := f.ReadDir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return dirs, nil
}

// see /usr/lib/go/src/path/filepath/symlink.go
func resolve(path string, linksWalked int, startPos int, logger *zap.Logger) (string, int, uint32, bool) {
	volLen := 0
	if len(path) > 0 && path[0] == runePathSeparator {
		volLen++
	}
	vol := path[:volLen]

	ind := volLen
	allowDots := true
	if startPos > 0 {
		ind = startPos
		allowDots = false
	}
	dest := path[:ind]
	var err error
	var modeType uint32

	for start, end := ind, ind; start < len(path); start = end {
		for start < len(path) && path[start] == runePathSeparator {
			start++
		}
		end = start
		for end < len(path) && path[end] != runePathSeparator {
			end++
		}

		// The next path component is in path[start:end].
		if end == start {
			// No more path components.
			break
		}

		modeType = 0
		if path[start:end] == "." {
			// Ignore path component ".".
			continue
		} else if path[start:end] == ".." {
			// Back up to previous component if possible.
			// Note that volLen includes any leading slash.

			// Set r to the index of the last slash in dest,
			// after the volume.
			var r int
			for r = len(dest) - 1; r >= volLen; r-- {
				if dest[r] == runePathSeparator {
					break
				}
			}
			if r < volLen || dest[r+1:] == ".." {
				// Either path has no slashes
				// it's empty or it ends in a ".." we had to keep.
				// Either way, keep this "..".
				if len(dest) > volLen {
					dest += strPathSeparator
				}
				dest += ".."
			} else {
				// Discard everything since the last slash.
				dest = dest[:r]
			}
			continue
		}

		// Ordinary path component. Add it to result.
		if len(dest) > 0 && dest[len(dest)-1] != runePathSeparator {
			dest += strPathSeparator
		}

		dest += path[start:end]

		// Check is symlink
		if modeType, err = lStatMode(dest); err != nil {
			if os.IsNotExist(err) {
				// Path {dest} is not exists, it is not error
			} else {
				logger.Warn("Failed get file stat", zap.String("path", dest), zap.Error(err))
			}

			return "", linksWalked, modeType, false
		}

		if modeType != syscall.S_IFLNK {
			if modeType != syscall.S_IFDIR && end < len(path) {
				// found not dir inside path
				logger.Warn("Failed resolve path", zap.String("path", path), zap.Error(syscall.ENOTDIR))
				return "", linksWalked, modeType, false
			}

			continue
		}

		// Resolve symlink.
		modeType = 0
		linksWalked++
		if linksWalked > linksWalkedMax {
			logger.Info("Failed resolve path", zap.String("path", path), zap.Error(errors.New("too many links")))
			return "", linksWalked, modeType, false

		}

		resolvedPath, err := os.Readlink(dest)
		if err != nil {
			if os.IsNotExist(err) {
				// Path {dest} is not exists, it is not error
			} else {
				logger.Warn("Failed get file stat", zap.String("path", dest), zap.Error(err))
			}

			return "", linksWalked, modeType, false
		}
		if len(resolvedPath) == 0 {
			// Path {resolvedPath} is not exists, it is not error
			return "", linksWalked, modeType, false
		}

		if !allowDots && (resolvedPath == "." || resolvedPath == "..") {
			logger.Info("Failed resolve path", zap.String("path", dest), zap.Error(errors.New("circle links")))
			return "", linksWalked, modeType, false
		}
		allowDots = true

		resolvedPath = ExpandUser(resolvedPath)
		path = resolvedPath + path[end:]

		// Symlink to absolute path
		if resolvedPath[0] == runePathSeparator {
			dest = resolvedPath[:1]
			end = 1
		} else {
			// Symlink to relative path; replace last
			// path component in dest.
			var r int
			for r = len(dest) - 1; r >= volLen; r-- {
				if dest[r] == runePathSeparator {
					break
				}
			}
			if r < volLen {
				dest = vol
			} else {
				dest = dest[:r]
			}
			end = 0
		}
	}

	if modeType == 0 {
		if modeType, err = lStatMode(dest); err != nil {
			logger.Warn("Failed get file stat", zap.String("path", dest), zap.Error(err))
			return "", linksWalked, modeType, false
		}
	}
	return filepath.Clean(dest), linksWalked, modeType, true
}

func join(first string, last string) string {
	e1Len := len(first)
	fullLen := e1Len + len(last)
	if first[e1Len-1] != runePathSeparator && last[0] != runePathSeparator {
		fullLen++
	}
	buf := make([]byte, fullLen)
	index := copy(buf, first[:])
	if first[e1Len-1] != runePathSeparator && last[0] != runePathSeparator {
		buf[index] = byte(runePathSeparator)
		index++
	}
	copy(buf[index:], last[:])

	return *(*string)(unsafe.Pointer(&buf))
}

func walkDir(linkPath string, realPath string, linksWalked int, logger *zap.Logger, fn WalkFunc) {
	children, err := readDir(realPath)
	if err != nil {
		logger.Info("Failed get dir items", zap.String("path", realPath), zap.Error(err))
		return
	}

	for _, child := range children {
		childRealPath := join(realPath, child.Name())
		childLinkPath := join(linkPath, child.Name())
		if (child.Type() & os.ModeSymlink) != 0 {
			resolvedPath, linksWalkedChild, modeType, ok := resolve(childRealPath, linksWalked, len(realPath), logger)
			if !ok {
				continue
			}

			fn(childLinkPath, statModeToPathMode(modeType))
			if modeType == syscall.S_IFDIR {
				walkDir(childLinkPath, resolvedPath, linksWalkedChild, logger, fn)
			}
		} else {
			fn(childLinkPath, dirEntryModeToPathMode(child.Type()))
			if child.IsDir() {
				walkDir(childLinkPath, childRealPath, linksWalked, logger, fn)
			}
		}
	}
}

func Walk(dirPath string, logger *zap.Logger, fn WalkFunc) {
	fullDirPath, err := filepath.Abs(ExpandUser(dirPath))
	if err != nil {
		logger.Warn("Failed get abs path", zap.String("path", dirPath), zap.Error(err))
		return
	}

	modeType, err := lStatMode(fullDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Path {fullDirPath} is not exists, it is not error
		} else {
			logger.Warn("Failed get file stat", zap.String("path", fullDirPath), zap.Error(err))
		}

		return
	}

	if modeType == syscall.S_IFLNK {
		resolvedPath, linksWalked, modeType, ok := resolve(fullDirPath, 0, 0, logger)
		if !ok {
			return
		}

		fn(fullDirPath, statModeToPathMode(modeType))
		if modeType == syscall.S_IFDIR {
			walkDir(fullDirPath, resolvedPath, linksWalked, logger, fn)
		}
	} else {
		fn(fullDirPath, statModeToPathMode(modeType))
		if modeType == syscall.S_IFDIR {
			walkDir(fullDirPath, fullDirPath, 0, logger, fn)
		}
	}
}
