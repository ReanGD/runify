package api

import (
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

type ContextMenuCtrl interface {
	OnOpen(formID uint32, client RpcClient)
	OnRowActivate(rowID ContextMenuRowID)
}

type RootListCtrl interface {
	OnOpen(formID uint32, client RpcClient) []*RootListRow
	OnFilterChange(value string)
	OnRowActivate(providerID ProviderID, rowID RootListRowID)
	OnMenuActivate(providerID ProviderID, rowID RootListRowID)
}

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

const MaxPriority uint16 = 0xFFFF

type RootListRowGlobalID struct {
	ProviderID ProviderID
	ID         RootListRowID
}

func NewRootListRowGlobalID(providerID ProviderID, id RootListRowID) RootListRowGlobalID {
	return RootListRowGlobalID{
		ProviderID: providerID,
		ID:         id,
	}
}

func (r *RootListRowGlobalID) ToProtobuf() *pb.RootListRowGlobalID {
	return &pb.RootListRowGlobalID{
		ProviderID: uint32(r.ProviderID),
		RowID:      uint32(r.ID),
	}
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

func (r *RootListRow) ToProtobuf() *pb.RootListRow {
	return &pb.RootListRow{
		ProviderID: uint32(r.ProviderID),
		RowID:      uint32(r.ID),
		Icon:       r.Icon,
		Value:      r.Value,
		Priority:   uint32(r.Priority),
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

func (r *ContextMenuRow) ToProtobuf() *pb.ContextMenuRow {
	return &pb.ContextMenuRow{
		RowID: uint32(r.ID),
		Value: r.Value,
	}
}
