package desktop_entry

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ReanGD/runify/server/gtk"
	"github.com/ReanGD/runify/server/paths"
	"go.uber.org/zap"
)

type iconKey struct {
	size int
	name string
}

func newIconKey(size int, name string) iconKey {
	return iconKey{
		size: size,
		name: name,
	}
}

func newIconKeyFromPath(fullpath string) (iconKey, error) {
	_, filename := filepath.Split(fullpath)
	extIndex := strings.LastIndexByte(filename, '.')
	if extIndex == -1 {
		return iconKey{}, fmt.Errorf("Wrong icon cache filename %s, error: not found extension", fullpath)
	}
	filename = filename[:extIndex]

	sepIndex := strings.IndexByte(filename, '_')
	if sepIndex == -1 {
		return iconKey{}, fmt.Errorf("Wrong icon cache filename %s, error: not found separator", fullpath)
	}

	size, err := strconv.Atoi(filename[:sepIndex])
	if err != nil {
		return iconKey{}, fmt.Errorf("Wrong icon cache filename %s, error: can't parse size: %s", fullpath, err)
	}

	return iconKey{
		size: size,
		name: filename[sepIndex+1:],
	}, nil
}

func (k iconKey) toFullPath() string {
	return filepath.Join(paths.GetAppIconCache(), fmt.Sprintf("%d_%s.png", k.size, k.name))
}

type iconCache struct {
	defaultIconTheme *gtk.IconTheme
	iconPathCache    map[iconKey]string
	logger           *zap.Logger
}

func newIconCache(logger *zap.Logger) (*iconCache, error) {
	gtk.Init()

	defaultIconTheme, err := gtk.IconThemeGetDefault()
	if err != nil {
		return nil, fmt.Errorf("Getting default theme for icons ended with error: %s", err)
	}

	iconPathCache := make(map[iconKey]string)
	paths.Walk(paths.GetAppIconCache(), logger, func(fullpath string, mode paths.PathMode) {
		if mode == paths.ModeRegFile {
			if key, err := newIconKeyFromPath(fullpath); err != nil {
				logger.Info("Failed scan icons", zap.Error(err))
			} else {
				iconPathCache[key] = fullpath
			}
		}
	})

	return &iconCache{
		defaultIconTheme: defaultIconTheme,
		iconPathCache:    iconPathCache,
		logger:           logger,
	}, nil
}

func (c *iconCache) resolveIcon(name string, size int) string {
	if len(name) == 0 {
		return ""
	}

	if filepath.IsAbs(name) {
		if ok, _ := paths.ExistsFile(name); ok {
			return name
		}
	}

	iconPath := c.defaultIconTheme.LookupIcon(name, size, gtk.ICON_LOOKUP_NO_SVG)
	if len(iconPath) == 0 {
		iconPath = c.defaultIconTheme.LookupIcon(name, size, 0)
	}

	return iconPath
}

func (c *iconCache) getNonSvgIconPath(name string, size int) string {
	key := newIconKey(size, name)
	if path, ok := c.iconPathCache[key]; ok {
		return path
	}

	path := c.resolveIcon(name, size)
	if len(path) == 0 {
		c.iconPathCache[key] = ""
		return ""
	}

	if filepath.Ext(path) != ".svg" && !strings.HasPrefix(path, "/org") {
		// TODO: resize icon
		c.iconPathCache[key] = path
		return path
	}

	pBuf, err := c.defaultIconTheme.LoadIcon(name, size, 0)
	if err != nil {
		c.logger.Info("Failed load icon", zap.String("name", name), zap.Error(err))
		return ""
	}

	path = key.toFullPath()
	err = pBuf.SavePNG(path, 0)
	if err != nil {
		c.logger.Info("Failed save icon", zap.String("name", name), zap.String("path", path), zap.Error(err))
		return ""
	}

	c.iconPathCache[key] = path
	return path
}
