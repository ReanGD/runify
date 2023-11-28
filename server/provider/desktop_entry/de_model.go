package desktop_entry

import (
	"fmt"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

type deModel struct {
	providerID   api.ProviderID
	nextID       api.RootListRowID
	nameIndex    map[string]api.RootListRowID
	entriesIndex map[api.RootListRowID]types.DesktopFile
	dataMutex    sync.RWMutex
	dataCache    []*api.RootListRow
	moduleLogger *zap.Logger
}

func newDEModel() *deModel {
	return &deModel{
		providerID:   0,
		nextID:       1,
		nameIndex:    make(map[string]api.RootListRowID),
		entriesIndex: make(map[api.RootListRowID]types.DesktopFile),
		dataMutex:    sync.RWMutex{},
		dataCache:    []*api.RootListRow{},
		moduleLogger: nil,
	}
}

func (m *deModel) init(providerID api.ProviderID, moduleLogger *zap.Logger) error {
	m.providerID = providerID
	m.moduleLogger = moduleLogger

	return nil
}

func (m *deModel) onDesktopEntries(request interface{}) (bool, error) {
	entries, ok := request.(types.DesktopFiles)
	if !ok {
		return true, fmt.Errorf("invalid request type, expected types.DesktopFiles, but got %T", request)
	}
	entriesIndex := make(map[api.RootListRowID]types.DesktopFile)
	dataCache := make([]*api.RootListRow, 0, len(m.dataCache))

	for _, entry := range entries {
		name := entry.Name()
		id, ok := m.nameIndex[name]
		if !ok {
			id = m.nextID
			m.nextID++
			m.nameIndex[name] = id
		}

		entriesIndex[id] = entry

		dataCache = append(dataCache, api.NewRootListRow(
			api.RowType_Application,
			api.MinPriority,
			m.providerID,
			id,
			entry.IconPath(),
			name,
			entry.SearchNames()))
	}

	m.dataMutex.Lock()
	m.entriesIndex = entriesIndex
	m.dataCache = dataCache
	m.dataMutex.Unlock()

	return false, nil
}

func (m *deModel) getRows() []*api.RootListRow {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()
	return m.dataCache
}

func (m *deModel) getEntry(id api.RootListRowID) (types.DesktopFile, bool) {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()
	res, ok := m.entriesIndex[id]
	return res, ok
}
