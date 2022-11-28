package api

import "go.uber.org/zap"

type ProviderID uint32

func (id ProviderID) ZapField() zap.Field {
	return zap.Uint32("ProviderID", uint32(id))
}

func (id ProviderID) ZapFieldPrefix(prefix string) zap.Field {
	return zap.Uint32(prefix+"ProviderID", uint32(id))
}

type RootListRowID uint32

func (id RootListRowID) ZapField() zap.Field {
	return zap.Uint32("RootListRowID", uint32(id))
}

func (id RootListRowID) ZapFieldPrefix(prefix string) zap.Field {
	return zap.Uint32(prefix+"RootListRowID", uint32(id))
}

type ContextMenuRowID uint32

func (id ContextMenuRowID) ZapField() zap.Field {
	return zap.Uint32("ContextMenuRowID", uint32(id))
}

func (id ContextMenuRowID) ZapFieldPrefix(prefix string) zap.Field {
	return zap.Uint32(prefix+"ContextMenuRowID", uint32(id))
}

type RootListRow struct {
	ProviderID ProviderID
	ID         RootListRowID
	Icon       string
	Value      string
	Priority   uint16
}

func NewRootListRow(providerID ProviderID, id RootListRowID, icon string, value string, priority uint16) *RootListRow {
	return &RootListRow{
		ProviderID: providerID,
		ID:         id,
		Icon:       icon,
		Value:      value,
		Priority:   priority,
	}
}

type RootListRows struct {
	Create []*RootListRow
	Change []*RootListRow
	Remove []*RootListRow
}

func NewRootListRows() *RootListRows {
	return &RootListRows{
		Create: []*RootListRow{},
		Change: []*RootListRow{},
		Remove: []*RootListRow{},
	}
}

type ContextMenuRow struct {
	ID    ContextMenuRowID
	Value string
}

func NewContextMenuRow(id ContextMenuRowID, value string) *ContextMenuRow {
	return &ContextMenuRow{
		ID:    id,
		Value: value,
	}
}

type ContextMenuRows struct {
	Create []*ContextMenuRow
}

func NewContextMenuRows() *ContextMenuRows {
	return &ContextMenuRows{
		Create: []*ContextMenuRow{},
	}
}

type ContextMenuCtrlOrError struct {
	Ctrl  ContextMenuCtrl
	Error error
}

type ContexMenuCtrlOrErrorResult TResult[ContextMenuCtrlOrError]
