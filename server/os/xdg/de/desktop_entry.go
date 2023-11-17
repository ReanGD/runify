package de

import (
	"context"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
)

const ModuleName = "xdg_desktop_entry"

type XDGDesktopEntry struct {
	handler *handler

	module.Module
}

func New() (*XDGDesktopEntry, string) {
	return &XDGDesktopEntry{
		handler: newHandler(),
	}, ModuleName
}

func (d *XDGDesktopEntry) OnInit() <-chan error {
	ch := make(chan error)

	go func() {
		d.Init(d.GetConfig().XDGDesktopEntry.ModuleChLen)
		ch <- d.handler.init(d.GetModuleLogger())
	}()

	return ch
}

func (d *XDGDesktopEntry) OnStart(ctx context.Context) []*types.HandledChannel {
	d.handler.update()

	hChErr := types.NewHandledChannel(d.ErrorCtx.GetChannel(), d.onError)
	return []*types.HandledChannel{hChErr}
}

func (d *XDGDesktopEntry) OnFinish() {
	d.handler.stop()
}

func (d *XDGDesktopEntry) onError(request interface{}) (bool, error) {
	return true, request.(error)
}

func (d *XDGDesktopEntry) OnRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *updateCmd:
		d.handler.update()
	case *subscribeCmd:
		d.handler.subscribe(r)

	default:
		return d.OnRequestUnknownMsg(request)
	}

	return false, nil
}

func (d *XDGDesktopEntry) OnRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *updateCmd:
		r.onRequestDefault(d.GetModuleLogger(), reason)
	case *subscribeCmd:
		r.onRequestDefault(d.GetModuleLogger(), reason)

	default:
		return d.OnRequestDefaultUnknownMsg(request, reason)
	}

	return false, nil
}

func (d *XDGDesktopEntry) Update() {
	d.AddToChannel(&updateCmd{})
}

func (d *XDGDesktopEntry) Subscribe(ch chan<- types.DesktopEntries, result api.BoolResult) {
	d.AddToChannel(&subscribeCmd{
		ch:     ch,
		result: result,
	})
}
