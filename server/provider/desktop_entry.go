package provider

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/paths"
	"github.com/ReanGD/runify/server/pb"
	"github.com/rkoesters/xdg/desktop"
	"go.uber.org/zap"
)

type entry struct {
	path  string
	props *desktop.Entry
}

type desktopEntry struct {
	providerID   uint64
	entries      []*entry
	cache        []*pb.CardItem
	moduleLogger *zap.Logger
}

func newDesktopEntry() *desktopEntry {
	return &desktopEntry{
		providerID:   0,
		entries:      []*entry{},
		cache:        []*pb.CardItem{},
		moduleLogger: nil,
	}
}

func (p *desktopEntry) getName() string {
	return "desktopEntry"
}

func (p *desktopEntry) onInit(cfg *config.Config, moduleLogger *zap.Logger, providerID uint64) error {
	p.providerID = providerID
	p.moduleLogger = moduleLogger
	return nil
}

func (p *desktopEntry) onStart() {
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

func (p *desktopEntry) walkXDGDesktopEntries(fn func(fullpath string, props *desktop.Entry)) {
	exists := make(map[string]struct{})
	for _, dirname := range paths.GetXDGAppDirs() {
		paths.Walk(dirname, func(fullpath string, mode paths.PathMode) {
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
			props.Icon = paths.GetNonSvgIconPath(props.Icon, 48)

			fn(fullpath, props)
		})
	}
}

func (p *desktopEntry) getRoot() ([]*pb.CardItem, error) {
	return p.cache, nil
}

func (p *desktopEntry) getActions(cardID uint64) ([]*pb.ActionItem, error) {
	itemID := int(cardID & cardIDMask)
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

func (p *desktopEntry) execute(cardID uint64, actionID uint32) (*pb.Result, error) {
	itemID := int(cardID & cardIDMask)
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
