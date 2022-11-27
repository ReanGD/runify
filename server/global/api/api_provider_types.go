package api

type RootMenuRowID uint32
type ItemMenuRowID uint32

type RootMenuRow struct {
	ID       RootMenuRowID
	Icon     string
	Value    string
	Priority uint16
}

func NewRootMenuRow(id RootMenuRowID, icon string, value string, priority uint16) *RootMenuRow {
	return &RootMenuRow{
		ID:       id,
		Icon:     icon,
		Value:    value,
		Priority: priority,
	}
}

type RootMenuRows struct {
	Create []*RootMenuRow
	Change []*RootMenuRow
	Remove []*RootMenuRow
}

func NewRootMenuRows() *RootMenuRows {
	return &RootMenuRows{
		Create: []*RootMenuRow{},
		Change: []*RootMenuRow{},
		Remove: []*RootMenuRow{},
	}
}

type ItemMenuRow struct {
	ID    ItemMenuRowID
	Value string
}

func NewItemMenuRow(id ItemMenuRowID, value string) *ItemMenuRow {
	return &ItemMenuRow{
		ID:    id,
		Value: value,
	}
}

type ItemMenuRows struct {
	rows []*ItemMenuRow
}
