package api

import (
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

type ContextMenuCtrl interface {
	OnOpen(formID FormID, client RpcClient) []*ContextMenuRow
	OnRowActivate(rowID ContextMenuRowID)
}

type RootListCtrl interface {
	OnOpen(formID FormID, client RpcClient) []*RootListRow
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

const (
	RowType_Calculator  RowType = 0
	RowType_Application RowType = 1
	RowType_Unknown     RowType = 2
)

type RowType int16

func (t RowType) ToProtobuf() pb.RootListRowType {
	switch t {
	case RowType_Calculator:
		return pb.RootListRowType_CALCULATOR
	case RowType_Application:
		return pb.RootListRowType_APPLICATION
	default:
		return pb.RootListRowType_UNKNOWN
	}
}

func (t RowType) String() string {
	switch t {
	case RowType_Calculator:
		return "Calculator"
	case RowType_Application:
		return "Application"
	default:
		return "Unknown"
	}
}

func (t RowType) ZapField() zap.Field {
	return zap.String("RowType", t.String())
}

func (t RowType) ZapFieldPrefix(prefix string) zap.Field {
	return zap.String(prefix+"RowType", t.String())
}

const (
	MinPriority uint16 = 0
	MaxPriority uint16 = 0xFFFF
)

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
	rowType    RowType
	Priority   uint16
	ProviderID ProviderID
	ID         RootListRowID
	Icon       string
	Value      string
}

func NewRootListRow(
	rowType RowType, priority uint16, providerID ProviderID, id RootListRowID, icon string, value string,
) *RootListRow {
	return &RootListRow{
		rowType:    rowType,
		Priority:   priority,
		ProviderID: providerID,
		ID:         id,
		Icon:       icon,
		Value:      value,
	}
}

func (r *RootListRow) ToProtobuf() *pb.RootListRow {
	return &pb.RootListRow{
		RowType:    r.rowType.ToProtobuf(),
		ProviderID: uint32(r.ProviderID),
		RowID:      uint32(r.ID),
		Priority:   uint32(r.Priority),
		Icon:       r.Icon,
		Value:      r.Value,
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
