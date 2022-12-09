package api

type FormID uint32

type RpcClient interface {
	AddRootList(ctrl RootListCtrl)
	RootListAddRows(formID FormID, rows ...*RootListRow)
	RootListChangeRows(formID FormID, rows ...*RootListRow)
	RootListRemoveRows(formID FormID, rows ...RootListRowGlobalID)
	AddContextMenu(ctrl ContextMenuCtrl)
	CloseAll(msg error)
	CloseOne(formID FormID, msg error)
	ShowMessage(msg error)
}
