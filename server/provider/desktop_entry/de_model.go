package desktop_entry

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/paths"
	"github.com/rkoesters/xdg/desktop"
	"go.uber.org/zap"
)

type entry struct {
	path  string
	props *desktop.Entry
}

type deModel struct {
	providerID   api.ProviderID
	nextID       api.RootListRowID
	nameIndex    map[string]api.RootListRowID
	entries      map[api.RootListRowID]*entry
	dataMutex    sync.RWMutex
	dataCache    []*api.RootListRow
	iconCache    *iconCache
	moduleLogger *zap.Logger
}

func newDEModel() *deModel {
	return &deModel{
		providerID:   0,
		nextID:       1,
		nameIndex:    make(map[string]api.RootListRowID),
		entries:      make(map[api.RootListRowID]*entry),
		dataMutex:    sync.RWMutex{},
		dataCache:    []*api.RootListRow{},
		iconCache:    nil,
		moduleLogger: nil,
	}
}

func (m *deModel) init(providerID api.ProviderID, moduleLogger *zap.Logger) error {
	m.providerID = providerID
	m.moduleLogger = moduleLogger
	var err error
	m.iconCache, err = newIconCache(moduleLogger)
	return err
}

func (m *deModel) start() {
	m.update()
}

func (m *deModel) update() {
	entries := make(map[api.RootListRowID]*entry)
	dataCache := make([]*api.RootListRow, 0, len(m.dataCache))
	m.walkXDGDesktopEntries(func(path string, props *desktop.Entry) {
		id, ok := m.nameIndex[props.Name]
		if !ok {
			id = m.nextID
			m.nextID++
		}

		entries[id] = &entry{
			path:  path,
			props: props,
		}

		dataCache = append(dataCache, api.NewRootListRow(
			api.RowType_Application, api.MinPriority, m.providerID, id, props.Icon, props.Name))
	})

	m.dataMutex.Lock()
	m.entries = entries
	m.dataCache = dataCache
	m.dataMutex.Unlock()
}

func (m *deModel) walkXDGDesktopEntries(fn func(fullpath string, props *desktop.Entry)) {
	exists := make(map[string]struct{})
	for _, dirname := range paths.GetXDGAppDirs() {
		paths.Walk(dirname, m.moduleLogger, func(fullpath string, mode paths.PathMode) {
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
				m.moduleLogger.Info("Error open desktop entry file", zap.String("path", fullpath), zap.Error(err))
				return
			}

			props, err := desktop.New(f)
			f.Close()
			if err != nil {
				m.moduleLogger.Info("Error parse desktop entry file", zap.String("path", fullpath), zap.Error(err))
				return
			}

			if props.NoDisplay || props.Hidden {
				return
			}
			props.Icon = m.iconCache.getNonSvgIconPath(props.Icon, 48)

			fn(fullpath, props)
		})
	}
}

func (m *deModel) getRows() []*api.RootListRow {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()
	return m.dataCache
}

func (m *deModel) getEntry(id api.RootListRowID) (*entry, bool) {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()
	res, ok := m.entries[id]
	return res, ok
}
