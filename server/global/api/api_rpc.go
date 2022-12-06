package api

type RpcClient interface {
	RootListOpen(formID uint32, ctrl RootListCtrl, rows []*RootListRow)
	RootListAddRows(formID uint32, rows ...*RootListRow)
	RootListChangeRows(formID uint32, rows ...*RootListRow)
	RootListRemoveRows(formID uint32, rows ...RootListRowGlobalID)
	ContextMenuOpen(formID uint32, ctrl ContextMenuCtrl, rows ...*ContextMenuRow)
	CloseAll(msg error)
	CloseOne(formID uint32, msg error)
	ShowMessage(msg error)
}
