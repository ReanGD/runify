package api

type ContextMenuCtrl interface {
	GetRows() *ContextMenuRows
	OnRowActivate(id ContextMenuRowID, result ErrorResult)
}

type RootListCtrl interface {
	GetRowsCh() <-chan *RootListRows
	OnFilterChange(value string)
	OnRowActivate(id RootListRowID, result ErrorResult)
	OnMenuActivate(id RootListRowID, result ContexMenuCtrlOrErrorResult)
}
