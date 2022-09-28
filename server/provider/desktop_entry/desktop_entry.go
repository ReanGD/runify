package desktop_entry

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/paths"
	"github.com/ReanGD/runify/server/pb"
	"github.com/ReanGD/runify/server/provider/pcommon"
	"github.com/rkoesters/xdg/desktop"
	"go.uber.org/zap"
)

type entry struct {
	path  string
	props *desktop.Entry
}

type DesktopEntry struct {
	providerID   uint64
	iconsCache   *iconCache
	entries      []*entry
	cache        []*pb.CardItem
	moduleLogger *zap.Logger
}

func NewDesktopEntry() *DesktopEntry {
	return &DesktopEntry{
		providerID:   0,
		iconsCache:   nil,
		entries:      []*entry{},
		cache:        []*pb.CardItem{},
		moduleLogger: nil,
	}
}

func (p *DesktopEntry) GetName() string {
	return "desktopEntry"
}

func (p *DesktopEntry) OnInit(cfg *config.Config, moduleLogger *zap.Logger, providerID uint64) error {
	p.providerID = providerID
	p.moduleLogger = moduleLogger
	var err error
	p.iconsCache, err = newIconCache(moduleLogger)
	return err
}

func (p *DesktopEntry) OnStart() {
	id := p.providerID
	entries := p.entries
	cache := p.cache
	p.walkXDGDesktopEntries(func(path string, props *desktop.Entry) {
		entries = append(entries, &entry{
			path:  path,
			props: props,
		})
		cache = append(cache, &pb.CardItem{
			CardID: id,
			Name:   props.Name,
			Icon:   props.Icon,
		})
		id++
	})

	p.entries = entries
	p.cache = cache
}

func (p *DesktopEntry) walkXDGDesktopEntries(fn func(fullpath string, props *desktop.Entry)) {
	exists := make(map[string]struct{})
	for _, dirname := range paths.GetXDGAppDirs() {
		paths.Walk(dirname, p.moduleLogger, func(fullpath string, mode paths.PathMode) {
			if mode != paths.ModeRegFile || filepath.Ext(fullpath) != ".desktop" {
				return
			}

			_, filename := filepath.Split(fullpath)
			if _, ok := exists[filename]; ok {
				return
			}
			exists[filename] = struct{}{}

			f, err := os.Open(fullpath)
			if err != nil {
				p.moduleLogger.Info("Error open desktop entry file", zap.String("path", fullpath), zap.Error(err))
				return
			}

			props, err := desktop.New(f)
			f.Close()
			if err != nil {
				p.moduleLogger.Info("Error parse desktop entry file", zap.String("path", fullpath), zap.Error(err))
				return
			}

			if props.NoDisplay || props.Hidden {
				return
			}
			props.Icon = p.iconsCache.getNonSvgIconPath(props.Icon, 48)

			fn(fullpath, props)
		})
	}
}

func (p *DesktopEntry) GetRoot() ([]*pb.CardItem, error) {
	return p.cache, nil
}

func (p *DesktopEntry) GetActions(cardID uint64) ([]*pb.ActionItem, error) {
	itemID := int(cardID & pcommon.CardIDMask)
	if itemID >= len(p.entries) {
		return nil, errors.New("not found item by cardID")
	}

	return []*pb.ActionItem{{
		ActionID: 0,
		Name:     "Open",
	}, {
		ActionID: 1,
		Name:     "Copy name",
	}, {
		ActionID: 2,
		Name:     "Copy path",
	}}, nil
}

func (p *DesktopEntry) Execute(cardID uint64, actionID uint32) (*pb.Result, error) {
	itemID := int(cardID & pcommon.CardIDMask)
	if itemID >= len(p.entries) {
		return nil, errors.New("not found item by cardID")
	}

	// TODO: copy run from dex
	entry := p.entries[itemID]
	c := exec.Command("dex", entry.path)
	err := c.Run()
	if err != nil {
		p.moduleLogger.Warn("Failed execute desktop entry",
			zap.String("Request", "Execute"),
			zap.Uint64("CardID", cardID),
			zap.Uint32("ActionID", actionID),
			zap.String("EntryPath", entry.path),
			zap.Error(err))

		return &pb.Result{
			Payload: &pb.Result_Hide{
				Hide: &pb.HideWindow{
					Error: "Failed execute desktop entry",
				},
			},
		}, nil
	}

	return &pb.Result{
		Payload: &pb.Result_Hide{
			Hide: &pb.HideWindow{},
		},
	}, nil
}
