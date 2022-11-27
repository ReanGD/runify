package api

type ItemMenuCtrl interface {
	GetRows() <-chan ItemMenuRows
	OnRowActivate(id ItemMenuRowID) VoidResult
}

type ItemMenuCtrlOrError struct {
	Ctrl  ItemMenuCtrl
	Error error
}

type ItemMenuCtrlOrErrorResult TResult[ItemMenuCtrlOrError]

type RootMenuCtrl interface {
	GetRowsCh() <-chan RootMenuRows
	OnFilterChange(value string)
	OnRowActivate(id RootMenuRowID) VoidResult
	OnMenuActivate(id RootMenuRowID) ItemMenuCtrlOrErrorResult
}
