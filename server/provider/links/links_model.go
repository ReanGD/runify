package links

import (
	"errors"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/widget"
	"go.uber.org/zap"
)

const createRowID = api.RootListRowID(1)

type DataModel struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type item struct {
	id   api.RootListRowID
	data *DataModel
}

type model struct {
	providerID   api.ProviderID
	formWidget   *widget.Form
	moduleLogger *zap.Logger

	dataMutex sync.RWMutex
	nextID    api.RootListRowID
	nameIndex map[string]*item
	idIndex   map[api.RootListRowID]*item
	rowsCache []*api.RootListRow
}

func newModel() *model {
	return &model{
		providerID:   0,
		formWidget:   nil,
		moduleLogger: nil,
		dataMutex:    sync.RWMutex{},
		nextID:       createRowID + 1,
		nameIndex:    make(map[string]*item),
		idIndex:      make(map[api.RootListRowID]*item),
		rowsCache:    []*api.RootListRow{},
	}
}

func (m *model) init(providerID api.ProviderID, moduleLogger *zap.Logger) error {
	m.providerID = providerID
	m.moduleLogger = moduleLogger

	var err error
	m.formWidget, err = widget.NewForm(func(bind widget.Bind, fields *DataModel) widget.Widget {
		return &widget.Column{
			Children: []widget.Widget{
				&widget.Text{
					Data: "Name",
				},
				&widget.TextField{
					Bind: bind(&fields.Name).Required().MaxLength(25).ServerSide(),
				},
				&widget.Text{
					Data: "Link",
				},
				&widget.TextField{
					Bind: bind(&fields.Link).Required(),
				},
			},
		}
	})

	if err != nil {
		m.moduleLogger.Error("Failed create widgets", zap.Error(err))
		return err
	}

	return nil
}

func (m *model) start() {
	m.updateCache()
}

func (m *model) updateCache() {
	m.rowsCache = make([]*api.RootListRow, 0, len(m.idIndex)+1)
	m.rowsCache = append(m.rowsCache, api.NewRootListRow(
		api.RowType_Command, api.MinPriority, m.providerID, createRowID, "", "Create link"))

	for _, item := range m.idIndex {
		m.rowsCache = append(m.rowsCache, api.NewRootListRow(
			api.RowType_Link, api.MinPriority, m.providerID, item.id, "", item.data.Name))
	}
}

func (m *model) getRows() []*api.RootListRow {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()

	return m.rowsCache
}

func (m *model) createDataForm(rowID api.RootListRowID) (*widget.DataForm, error) {
	if rowID <= createRowID {
		return widget.NewDataForm(m.formWidget, &DataModel{})
	}

	item, ok := m.getItem(rowID)
	if !ok {
		return nil, errors.New("item ID not found")
	}

	return widget.NewDataForm(m.formWidget, item.data)
}

func (m *model) getItem(id api.RootListRowID) (*item, bool) {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()
	res, ok := m.idIndex[id]
	return res, ok
}

func (m *model) checkItem(id api.RootListRowID, name string) bool {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()

	if nameItem, ok := m.nameIndex[name]; ok && nameItem.id != id {
		return false
	}

	return true
}

func (m *model) saveItem(id api.RootListRowID, data *DataModel) error {
	if id <= createRowID {
		return m.addItem(data)
	}

	return m.updateItem(id, data)
}

func (m *model) addItem(data *DataModel) error {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()

	_, ok := m.nameIndex[data.Name]
	if ok {
		return errors.New("item with this name already exists")
	}

	item := &item{
		id:   m.nextID,
		data: data,
	}
	m.nextID++
	m.idIndex[item.id] = item
	m.nameIndex[data.Name] = item

	m.updateCache()

	return nil
}

func (m *model) updateItem(id api.RootListRowID, data *DataModel) error {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()

	item, ok := m.idIndex[id]
	if !ok {
		return errors.New("item not found")
	}

	if nameItem, ok := m.nameIndex[data.Name]; ok && nameItem.id != id {
		return errors.New("item with this name already exists")
	}

	delete(m.nameIndex, data.Name)
	item.data = data
	m.nameIndex[data.Name] = item

	m.updateCache()

	return nil
}

func (m *model) removeItem(id api.RootListRowID) {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()

	item, ok := m.idIndex[id]
	if !ok {
		return
	}

	delete(m.idIndex, id)
	delete(m.nameIndex, item.data.Name)

	m.updateCache()
}
