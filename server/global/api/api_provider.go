package api

type ContextMenuCtrl interface {
	OnOpen() []*ContextMenuRow
	OnRowActivate(rowID ContextMenuRowID, result ErrorResult)
}

type RootListCtrl interface {
	OnOpen(sender RootListRowsUpdateSender) []*RootListRow
	OnFilterChange(value string)
	OnRowActivate(providerID ProviderID, rowID RootListRowID, result ErrorResult)
	OnMenuActivate(providerID ProviderID, rowID RootListRowID, result ContexMenuCtrlOrErrorResult)
}
