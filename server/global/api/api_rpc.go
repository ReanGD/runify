package api

type FormID uint32

type RpcClient interface {
	AddForm(ctrl FormCtrl)
	AddRootList(ctrl RootListCtrl)
	RootListAddRows(formID FormID, rows ...*RootListRow)
	RootListChangeRows(formID FormID, rows ...*RootListRow)
	RootListRemoveRows(formID FormID, rows ...RootListRowGlobalID)
	AddContextMenu(ctrl ContextMenuCtrl)
	FieldCheckResponse(formID FormID, requestID uint32, result bool, msg string)
	UserMessage(msg string)
	CloseForm(formID FormID)
	HideUI(msg error)
	CloseUI()
}
