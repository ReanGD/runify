package links

import (
	"errors"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/widget"
	"github.com/ReanGD/runify/server/jdb"
	"github.com/ReanGD/runify/server/paths"
	"go.uber.org/zap"
)

const (
	createRowID    = api.RootListRowID(1)
	providerDBName = "settings"
)

type DataModel struct {
	Name    string `json:"name"`
	Aliases string `json:"aliases"`
	Link    string `json:"link"`
}

func (d *DataModel) normalize() {
	d.Name = strings.TrimSpace(d.Name)
	d.Link = strings.TrimSpace(d.Link)

	aliases := map[string]struct{}{
		d.Name: {},
	}
	for _, alias := range strings.Split(d.Aliases, ";") {
		alias = strings.TrimSpace(alias)
		if alias != "" {
			aliases[strings.TrimSpace(alias)] = struct{}{}
		}
	}
	delete(aliases, d.Name)

	var aliasesStr string
	for alias := range aliases {
		if aliasesStr != "" {
			aliasesStr += ";"
		}
		aliasesStr += alias
	}
	d.Aliases = aliasesStr
}

func (d *DataModel) getAliases() string {
	aliases := d.Name
	for _, alias := range strings.Split(d.Aliases, ";") {
		aliases += "\n" + alias
	}

	return aliases
}

type item struct {
	id   api.RootListRowID
	data *DataModel
}

type model struct {
	providerID   api.ProviderID
	db           *jdb.JDB
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
		db:           nil,
		formWidget:   nil,
		moduleLogger: nil,
		dataMutex:    sync.RWMutex{},
		nextID:       createRowID + 1,
		nameIndex:    make(map[string]*item),
		idIndex:      make(map[api.RootListRowID]*item),
		rowsCache:    []*api.RootListRow{},
	}
}

func (m *model) init(providerID api.ProviderID, providerName string, moduleLogger *zap.Logger) error {
	m.providerID = providerID
	m.moduleLogger = moduleLogger

	var err error
	dbDir := filepath.Join(paths.GetAppConfig(), providerName)
	if m.db, err = jdb.New(dbDir, providerDBName, moduleLogger); err != nil {
		return err
	}

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
					Data: "Aliases",
				},
				&widget.TextField{
					Bind: bind(&fields.Aliases).MaxLength(150).ServerSide(),
				},
				&widget.Text{
					Data: "If multiple aliases are required, separate them with a ';' character",
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

	storage := []*DataModel{}
	if err = m.db.Read(&storage); err != nil {
		return err
	}

	for _, item := range storage {
		item.normalize()
		if err = m.addItem(item, false); err != nil {
			return err
		}
	}

	return nil
}

func (m *model) start() {
	_ = m.updateCache(false)
}

func (m *model) updateCache(saveToDB bool) error {
	m.rowsCache = make([]*api.RootListRow, 0, len(m.idIndex)+1)
	m.rowsCache = append(m.rowsCache, api.NewRootListRow(
		api.RowType_Command,
		api.MinPriority,
		m.providerID,
		createRowID,
		"",
		"Create link",
		"Create link"))

	for _, item := range m.idIndex {
		m.rowsCache = append(m.rowsCache, api.NewRootListRow(
			api.RowType_Link,
			api.MinPriority,
			m.providerID,
			item.id,
			"",
			item.data.Name,
			item.data.getAliases(),
		))
	}

	if saveToDB {
		storage := make([]*DataModel, 0, len(m.idIndex))
		for _, item := range m.idIndex {
			storage = append(storage, item.data)
		}

		if err := m.db.Write(&storage); err != nil {
			return err
		}
	}

	return nil
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
	saveToDB := true
	data.normalize()
	if id <= createRowID {
		return m.addItem(data, saveToDB)
	}

	return m.updateItem(id, data, saveToDB)
}

func (m *model) addItem(data *DataModel, saveToDB bool) error {
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

	return m.updateCache(saveToDB)
}

func (m *model) updateItem(id api.RootListRowID, data *DataModel, saveToDB bool) error {
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

	return m.updateCache(saveToDB)
}

func (m *model) removeItem(id api.RootListRowID, saveToDB bool) error {
	m.dataMutex.RLock()
	defer m.dataMutex.RUnlock()

	item, ok := m.idIndex[id]
	if !ok {
		return nil
	}

	delete(m.idIndex, id)
	delete(m.nameIndex, item.data.Name)

	return m.updateCache(saveToDB)
}
