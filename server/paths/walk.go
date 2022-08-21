package paths

import (
	"io/fs"
	"os"
	"path/filepath"
	"syscall"

	"github.com/ReanGD/runify/server/logger"
)

const (
	runePathSeparator = '/'
	strPathSeparator  = "/"
	linksWalkedMax    = 255
)

type WalkFunc func(path string, mode PathMode)

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
func resolve(path string, linksWalked int, startPos int) (string, int, bool) {
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
		} else if path[start:end] == "." {
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
		modeType, err := lStatMode(dest)
		if err != nil {
			if os.IsNotExist(err) {
				// Path {dest} is not exists, it is not error
			} else {
				logger.Write("Failed read file stat for path %s, error: %s", dest, err)
			}

			return "", linksWalked, false
		}

		if modeType != syscall.S_IFLNK {
			if modeType != syscall.S_IFDIR && end < len(path) {
				// found not dir inside path
				logger.Write("Failed resolve path %s, error: %s", path, syscall.ENOTDIR)
				return "", linksWalked, false
			}

			continue
		}

		// Resolve symlink.
		linksWalked++
		if linksWalked > linksWalkedMax {
			logger.Write("Failed resolve path %s, error: too many links", path, syscall.ENOTDIR)
			return "", linksWalked, false

		}

		resolvedPath, err := os.Readlink(dest)
		if err != nil {
			if os.IsNotExist(err) {
				// Path {dest} is not exists, it is not error
			} else {
				logger.Write("Failed read file stat for path %s, error: %s", dest, err)
			}

			return "", linksWalked, false
		}
		if len(resolvedPath) == 0 {
			// Path {resolvedPath} is not exists, it is not error
			return "", linksWalked, false
		}

		if !allowDots && (resolvedPath == "." || resolvedPath == "..") {
			logger.Write("Path %s is circle links", dest)
			return "", linksWalked, false
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

	return filepath.Clean(dest), linksWalked, true
}

func walkDir(linkPath string, realPath string, linksWalked int, fn WalkFunc) {
	children, err := readDir(realPath)
	if err != nil {
		logger.Write("Failed read dir items for path %s, error: %s", realPath, err)
		return
	}

	for _, child := range children {
		childRealPath := filepath.Join(realPath, child.Name())
		childLinkPath := filepath.Join(linkPath, child.Name())
		if (child.Type() & os.ModeSymlink) != 0 {
			resolvedPath, linksWalkedChild, ok := resolve(childRealPath, linksWalked, len(realPath))
			if !ok {
				continue
			}

			modeType, err := lStatMode(resolvedPath)
			if err != nil {
				logger.Write("Unexpected error for get stat for path %s, error: %s", resolvedPath, err)
				continue
			}

			fn(childLinkPath, statModeToPathMode(modeType))
			if modeType == syscall.S_IFDIR {
				walkDir(childLinkPath, resolvedPath, linksWalkedChild, fn)
			}
		} else {
			fn(childLinkPath, dirEntryModeToPathMode(child.Type()))
			if child.IsDir() {
				walkDir(childLinkPath, childRealPath, linksWalked, fn)
			}
		}
	}
}

func Walk(dirPath string, fn WalkFunc) {
	fullDirPath, err := filepath.Abs(ExpandUser(dirPath))
	if err != nil {
		logger.Write("Failed find abs path for %s, error: %s", dirPath, err)
		return
	}

	modeType, err := lStatMode(fullDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			// // Path {fullDirPath} is not exists, it is not error
		} else {
			logger.Write("Failed read file stat for path %s, error: %s", fullDirPath, err)
		}

		return
	}

	if modeType == syscall.S_IFLNK {
		resolvedPath, linksWalked, ok := resolve(fullDirPath, 0, 0)
		if !ok {
			return
		}

		modeType, err := lStatMode(resolvedPath)
		if err != nil {
			logger.Write("Unexpected error for get stat for path %s, error: %s", resolvedPath, err)
			return
		}

		fn(fullDirPath, statModeToPathMode(modeType))
		if modeType == syscall.S_IFDIR {
			walkDir(fullDirPath, resolvedPath, linksWalked, fn)
		}
	} else {
		fn(fullDirPath, statModeToPathMode(modeType))
		if modeType == syscall.S_IFDIR {
			walkDir(fullDirPath, fullDirPath, 0, fn)
		}
	}
}
