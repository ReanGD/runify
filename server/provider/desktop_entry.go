package provider

import (
	"os"
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
	commands     []*pb.Command
	moduleLogger *zap.Logger
}

func newDesktopEntry() desktopEntry {
	return desktopEntry{
		providerID:   0,
		entries:      []*entry{},
		commands:     []*pb.Command{},
		moduleLogger: nil,
	}
}

func (p *desktopEntry) getName() string {
	return "desktopEntry"
}

func (p *desktopEntry) init(cfg *config.Config, moduleLogger *zap.Logger, providerID uint64) error {
	p.providerID = providerID
	p.moduleLogger = moduleLogger
	return nil
}

func (p *desktopEntry) start() {
	id := p.providerID
	entries := p.entries
	commands := p.commands
	p.walkXDGDesktopEntries(func(path string, props *desktop.Entry) {
		entries = append(entries, &entry{
			path:  path,
			props: props,
		})
		commands = append(commands, &pb.Command{
			Id:   id,
			Name: props.Name,
			Icon: props.Icon,
		})
		id++
	})

	p.entries = entries
	p.commands = commands
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

func (p *desktopEntry) getRoot() []*pb.Command {
	return p.commands
}
