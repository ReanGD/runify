package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

type DesktopEntry struct {
	desktop        api.Desktop
	de             api.XDGDesktopEntry
	cfg            *config.Configuration
	model          *deModel
	actionExecuter *deActionExecuter
	moduleLogger   *zap.Logger
}

func New(desktop api.Desktop, de api.XDGDesktopEntry) *DesktopEntry {
	return &DesktopEntry{
		desktop:        desktop,
		de:             de,
		cfg:            nil,
		model:          newDEModel(),
		actionExecuter: newDEActionExecuter(),
		moduleLogger:   nil,
	}
}

func (p *DesktopEntry) GetName() string {
	return "DesktopEntry"
}

func (p *DesktopEntry) OnInit(cfg *config.Configuration, moduleLogger *zap.Logger, providerID api.ProviderID) error {
	p.cfg = cfg
	p.moduleLogger = moduleLogger
	if err := p.model.init(providerID, moduleLogger); err != nil {
		return err
	}
	if err := p.actionExecuter.init(cfg, p.desktop, p.model, moduleLogger); err != nil {
		return err
	}

	return nil
}

func (p *DesktopEntry) OnStart(errorCtx *module.ErrorCtx) []*types.HandledChannel {
	desktopEntriesCh := make(chan types.DesktopFiles, p.cfg.Provider.DesktopEntry.DesktopEntriesChLen)
	subsToDesktopEntriesRes := api.NewChanBoolResult()
	p.de.Subscribe(desktopEntriesCh, subsToDesktopEntriesRes)

	if res := <-subsToDesktopEntriesRes.GetChannel(); !res {
		errorCtx.SendError(errors.New("subscribe to XDGDesktopEntry failed"))
		return []*types.HandledChannel{}
	}

	return []*types.HandledChannel{
		types.NewHandledChannel(desktopEntriesCh, p.model.onDesktopEntries),
	}
}

func (p *DesktopEntry) MakeRootListCtrl() api.RootListCtrl {
	return newDERootListCtrl(p.model, p.actionExecuter, p.moduleLogger)
}
