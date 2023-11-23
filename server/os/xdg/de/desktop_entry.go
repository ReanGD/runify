package de

import (
	"context"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
)

type XDGDesktopEntry struct {
	handler *handler

	module.Module
}

func New() (*XDGDesktopEntry, string) {
	return &XDGDesktopEntry{
		handler: newHandler(),
	}, "xdg_desktop_entry"
}

func (m *XDGDesktopEntry) SetDeps() {
}

func (m *XDGDesktopEntry) OnInit() (uint32, error) {
	chLen := m.GetConfig().XDGDesktopEntry.ModuleChLen

	return chLen, m.handler.init(m.GetModuleLogger())
}

func (m *XDGDesktopEntry) OnStart(ctx context.Context) []*types.HandledChannel {
	m.handler.update()

	return []*types.HandledChannel{}
}

func (m *XDGDesktopEntry) OnFinish() {
	m.handler.stop()
}

func (m *XDGDesktopEntry) OnRequest(request api.ModuleMsgImpl) (bool, error) {
	switch r := request.(type) {
	case *updateCmd:
		m.handler.update()
	case *subscribeCmd:
		m.handler.subscribe(r)

	default:
		return m.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (m *XDGDesktopEntry) Update() {
	m.AddToChannel(&updateCmd{})
}

func (m *XDGDesktopEntry) Subscribe(ch chan<- types.DesktopEntries, result api.BoolResult) {
	m.AddToChannel(&subscribeCmd{
		ch:     ch,
		result: result,
	})
}
