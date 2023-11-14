package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/types"
	"go.uber.org/zap"
)

type deActionExecuter struct {
	terminal     string
	desktop      api.Desktop
	model        *deModel
	moduleLogger *zap.Logger
}

func newDEActionExecuter() *deActionExecuter {
	return &deActionExecuter{
		terminal:     "",
		desktop:      nil,
		model:        nil,
		moduleLogger: nil,
	}
}

func (e *deActionExecuter) init(cfg *config.Config, desktop api.Desktop, model *deModel, moduleLogger *zap.Logger) error {
	e.terminal = cfg.Get().System.Terminal
	e.desktop = desktop
	e.model = model
	e.moduleLogger = moduleLogger

	return nil
}

func (e *deActionExecuter) getEntry(id api.RootListRowID, logMsg string) (*types.DesktopEntry, error) {
	entry, ok := e.model.getEntry(id)
	if !ok {
		err := errors.New("row data not found")
		e.moduleLogger.Warn(logMsg,
			id.ZapField(),
			zap.Error(err),
		)

		return nil, err
	}

	return entry, nil
}

func (e *deActionExecuter) open(client api.RpcClient, id api.RootListRowID) {
	logMsg := "Failed execute desktop entry"
	entry, err := e.getEntry(id, logMsg)
	if err != nil {
		client.HideUI(err)
		return
	}

	err = execCmd(entry.Exec(), entry.InTerminal(), e.terminal)
	if err != nil {
		e.moduleLogger.Warn(logMsg,
			id.ZapField(),
			zap.String("EntryPath", entry.FilePath()),
			zap.Error(err),
		)
	}

	client.HideUI(err)
}

func (e *deActionExecuter) copy(id api.RootListRowID, entry *types.DesktopEntry, data *mime.Data, logMsg string) bool {
	copyResult := api.NewChanBoolResult()
	e.desktop.WriteToClipboard(false, data, copyResult)
	res := <-copyResult.GetChannel()
	if !res {
		e.moduleLogger.Warn(logMsg,
			id.ZapField(),
			zap.String("EntryPath", entry.FilePath()),
		)

		return false
	}

	return true
}

func (e *deActionExecuter) copyName(client api.RpcClient, id api.RootListRowID) {
	logMsg := "Failed copy name of desktop entry"
	entry, err := e.getEntry(id, logMsg)
	if err != nil {
		client.HideUI(err)
		return
	}

	if e.copy(id, entry, mime.NewTextData(entry.Name()), logMsg) {
		client.HideUI(nil)
	}

	client.HideUI(errors.New("Failed copy desktop entry name"))
}

func (e *deActionExecuter) copyPath(client api.RpcClient, id api.RootListRowID) {
	logMsg := "Failed copy path of desktop entry"
	entry, err := e.getEntry(id, logMsg)
	if err != nil {
		client.HideUI(err)
		return
	}

	if e.copy(id, entry, mime.NewTextData(entry.FilePath()), logMsg) {
		client.HideUI(nil)
	}

	client.HideUI(errors.New("Failed copy desktop entry path"))
}
