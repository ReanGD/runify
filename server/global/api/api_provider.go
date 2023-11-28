package api

import (
	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/global/widget"
	"github.com/ReanGD/runify/server/pb"
)

type FormCtrl interface {
	OnOpen(formID FormID, client RpcClient) *widget.DataForm
	OnFieldCheckRequest(requestID uint32, fieldName string, jsonBody string)
	OnSubmit(jsonBody string)
}

type RootListCtrl interface {
	OnOpen(formID FormID, client RpcClient) []*RootListRow
	OnFilterChange(value string)
	OnRowActivate(providerID ProviderID, rowID RootListRowID)
	OnMenuActivate(providerID ProviderID, rowID RootListRowID)
}

type ContextMenuCtrl interface {
	OnOpen(formID FormID, client RpcClient) []*ContextMenuRow
	OnRowActivate(rowID ContextMenuRowID)
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
	RowType_Command     RowType = 2
	RowType_Link        RowType = 3
	RowType_Unknown     RowType = 4
)

type RowType int16

func (t RowType) ToProtobuf() pb.RootListRowType {
	switch t {
	case RowType_Calculator:
		return pb.RootListRowType_CALCULATOR
	case RowType_Application:
		return pb.RootListRowType_APPLICATION
	case RowType_Command:
		return pb.RootListRowType_COMMAND
	case RowType_Link:
		return pb.RootListRowType_LINK
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
	case RowType_Command:
		return "Command"
	case RowType_Link:
		return "Link"
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
	rowType     RowType
	Priority    uint16
	ProviderID  ProviderID
	ID          RootListRowID
	Icon        string
	DisplayName string
	SearchNames string
}

func NewRootListRow(
	rowType RowType,
	priority uint16,
	providerID ProviderID,
	id RootListRowID,
	icon string,
	displayName string,
	searchNames string,
) *RootListRow {
	return &RootListRow{
		rowType:     rowType,
		Priority:    priority,
		ProviderID:  providerID,
		ID:          id,
		Icon:        icon,
		DisplayName: displayName,
		SearchNames: searchNames,
	}
}

func (r *RootListRow) ToProtobuf() *pb.RootListRow {
	return &pb.RootListRow{
		RowType:     r.rowType.ToProtobuf(),
		ProviderID:  uint32(r.ProviderID),
		RowID:       uint32(r.ID),
		Priority:    uint32(r.Priority),
		Icon:        r.Icon,
		DisplayName: r.DisplayName,
		SearchNames: r.SearchNames,
	}
}

type ContextMenuRow struct {
	ID          ContextMenuRowID
	DisplayName string
	SearchNames string
}

func NewContextMenuRow(id ContextMenuRowID, displayName, searchNames string) *ContextMenuRow {
	return &ContextMenuRow{
		ID:          id,
		DisplayName: displayName,
		SearchNames: searchNames,
	}
}

func (r *ContextMenuRow) ToProtobuf() *pb.ContextMenuRow {
	return &pb.ContextMenuRow{
		RowID:       uint32(r.ID),
		DisplayName: r.DisplayName,
		SearchNames: r.SearchNames,
	}
}
