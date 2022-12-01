package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
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

func (e *deActionExecuter) start() {
}

func (e *deActionExecuter) getEntry(id api.RootListRowID, result api.ErrorResult, logMsg string) (*entry, bool) {
	entry, ok := e.model.getEntry(id)
	if !ok {
		err := errors.New("row data not found")
		e.moduleLogger.Warn(logMsg,
			id.ZapField(),
			zap.Error(err),
		)

		result.SetResult(err)
	}

	return entry, ok
}

func (e *deActionExecuter) open(id api.RootListRowID, result api.ErrorResult) {
	logMsg := "Failed execute desktop entry"
	entry, ok := e.getEntry(id, result, logMsg)
	if !ok {
		return
	}

	err := execCmd(entry.props.Exec, entry.props.Terminal, e.terminal)
	if err != nil {
		e.moduleLogger.Warn(logMsg,
			id.ZapField(),
			zap.String("EntryPath", entry.path),
			zap.Error(err),
		)
	}

	result.SetResult(err)
}

func (e *deActionExecuter) copy(id api.RootListRowID, entry *entry, data *mime.Data, logMsg string, result api.ErrorResult) {
	copyResult := api.NewFuncBoolResult(func(ok bool) {
		var err error
		if !ok {
			err = errors.New("clipboard copy failed")
			e.moduleLogger.Warn(logMsg,
				id.ZapField(),
				zap.String("EntryPath", entry.path),
			)
		}
		result.SetResult(err)
	})
	e.desktop.WriteToClipboard(false, data, copyResult)
}

func (e *deActionExecuter) copyName(id api.RootListRowID, result api.ErrorResult) {
	logMsg := "Failed copy name of desktop entry"
	entry, ok := e.getEntry(id, result, logMsg)
	if !ok {
		return
	}

	e.copy(id, entry, mime.NewTextData(entry.props.Name), logMsg, result)
}

func (e *deActionExecuter) copyPath(id api.RootListRowID, result api.ErrorResult) {
	logMsg := "Failed copy path of desktop entry"
	entry, ok := e.getEntry(id, result, logMsg)
	if !ok {
		return
	}

	e.copy(id, entry, mime.NewTextData(entry.path), logMsg, result)
}
