package desktop_entry

import "github.com/ReanGD/runify/server/global/api"

type DERootListCtrl struct {
	model    *deModel
	executer *deExecuter
}

func newDERootListCtrl(model *deModel, executer *deExecuter) *DERootListCtrl {
	return &DERootListCtrl{
		model:    model,
		executer: executer,
	}
}

func (c *DERootListCtrl) GetRowsCh() <-chan *api.RootListRows {
	data := api.NewRootListRows()
	data.Create = c.model.getRows()

	ch := make(chan *api.RootListRows, 1)
	ch <- data
	return ch
}

func (c *DERootListCtrl) OnFilterChange(value string) {
	// pass
}

func (c *DERootListCtrl) OnRowActivate(id api.RootListRowID, result api.ErrorResult) {
	c.executer.open(id, result)
}

func (c *DERootListCtrl) OnMenuActivate(id api.RootListRowID, result api.ContexMenuCtrlOrErrorResult) {
	// pass
}
