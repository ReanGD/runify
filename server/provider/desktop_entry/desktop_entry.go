package desktop_entry

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/paths"
	"github.com/ReanGD/runify/server/pb"
	"github.com/ReanGD/runify/server/provider/pcommon"
	"github.com/ReanGD/runify/server/system/mime"
	"github.com/ReanGD/runify/server/system/module"
	"github.com/rkoesters/xdg/desktop"
	"go.uber.org/zap"
)

const (
	actionOpen uint32 = iota
	actionCopyName
	actionCopyPath
	actionLast
)

type entry struct {
	path  string
	props *desktop.Entry
}

type DesktopEntry struct {
	providerID   uint64
	terminal     string
	x11          module.X11
	iconsCache   *iconCache
	entries      []*entry
	cache        []*pb.CardItem
	moduleLogger *zap.Logger
}

func New(x11 module.X11) *DesktopEntry {
	return &DesktopEntry{
		providerID:   0,
		terminal:     "",
		x11:          x11,
		iconsCache:   nil,
		entries:      []*entry{},
		cache:        []*pb.CardItem{},
		moduleLogger: nil,
	}
}

func (p *DesktopEntry) GetName() string {
	return "DesktopEntry"
}

func (p *DesktopEntry) OnInit(cfg *config.Config, moduleLogger *zap.Logger, providerID uint64) error {
	p.providerID = providerID
	p.moduleLogger = moduleLogger
	p.terminal = cfg.Get().Provider.Terminal
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
		ActionID: actionOpen,
		Name:     "Open",
	}, {
		ActionID: actionCopyName,
		Name:     "Copy name",
	}, {
		ActionID: actionCopyPath,
		Name:     "Copy path",
	}}, nil
}

func (p *DesktopEntry) Execute(cardID uint64, actionID uint32) (*pb.Result, error) {
	itemID := int(cardID & pcommon.CardIDMask)
	if itemID >= len(p.entries) {
		return nil, errors.New("not found item by cardID")
	}
	if actionID >= actionLast {
		return nil, errors.New("not found action by actionID")
	}

	entry := p.entries[itemID]
	switch actionID {
	case actionOpen:
		err := execCmd(entry.props.Exec, entry.props.Terminal, p.terminal)
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
	case actionCopyName:
		p.x11.WriteToClipboard(false, mime.NewTextData(entry.props.Name))
	case actionCopyPath:
		p.x11.WriteToClipboard(false, mime.NewTextData(entry.path))
	}

	return &pb.Result{
		Payload: &pb.Result_Hide{
			Hide: &pb.HideWindow{},
		},
	}, nil
}
