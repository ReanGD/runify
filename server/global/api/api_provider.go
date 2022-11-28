package api

type ContextMenuCtrl interface {
	GetRows() []*ContextMenuRow
	OnRowActivate(rowID ContextMenuRowID, result ErrorResult)
}

type RootListCtrl interface {
	GetRows(out chan<- *RootListRowsUpdate) []*RootListRow
	OnFilterChange(value string)
	OnRowActivate(providerID ProviderID, rowID RootListRowID, result ErrorResult)
	OnMenuActivate(providerID ProviderID, rowID RootListRowID, result ContexMenuCtrlOrErrorResult)
}
