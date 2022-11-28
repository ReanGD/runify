package desktop_entry

import (
	"errors"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"go.uber.org/zap"
)

type deExecuter struct {
	providerID   api.ProviderID
	terminal     string
	desktop      api.Desktop
	model        *deModel
	moduleLogger *zap.Logger
}

func newDEExecuter() *deExecuter {
	return &deExecuter{
		providerID:   0,
		terminal:     "",
		desktop:      nil,
		model:        nil,
		moduleLogger: nil,
	}
}

func (e *deExecuter) init(
	providerID api.ProviderID, cfg *config.Config, desktop api.Desktop, model *deModel, moduleLogger *zap.Logger) error {

	e.providerID = providerID
	e.terminal = cfg.Get().System.Terminal
	e.desktop = desktop
	e.model = model
	e.moduleLogger = moduleLogger

	return nil
}

func (e *deExecuter) getEntry(id api.RootListRowID, result api.ErrorResult, logMsg string) (*entry, bool) {
	entry, ok := e.model.getEntry(id)
	if !ok {
		err := errors.New("entry not found")
		e.moduleLogger.Warn("Failed execute desktop entry",
			e.providerID.ZapField(),
			id.ZapField(),
			zap.Error(err),
		)

		result.SetResult(err)
	}

	return entry, ok
}

func (e *deExecuter) open(id api.RootListRowID, result api.ErrorResult) {
	entry, ok := e.getEntry(id, result, "Failed execute desktop entry")
	if !ok {
		return
	}

	err := execCmd(entry.props.Exec, entry.props.Terminal, e.terminal)
	if err != nil {
		e.moduleLogger.Warn("Failed execute desktop entry",
			e.providerID.ZapField(),
			id.ZapField(),
			zap.String("EntryPath", entry.path),
			zap.Error(err),
		)
	}

	result.SetResult(err)
}

func (e *deExecuter) copyName(id api.RootListRowID, result api.ErrorResult) {
	entry, ok := e.getEntry(id, result, "Failed copy name of desktop entry")
	if !ok {
		return
	}

	copyResult := api.NewFuncBoolResult(func(ok bool) {
		var err error
		if !ok {
			err = errors.New("clipboard copy failed")
			e.moduleLogger.Warn("Failed copy name of desktop entry",
				e.providerID.ZapField(),
				id.ZapField(),
				zap.String("EntryPath", entry.path),
			)
		}
		result.SetResult(err)
	})
	e.desktop.WriteToClipboard(false, mime.NewTextData(entry.props.Name), copyResult)
}

func (e *deExecuter) copyPath(id api.RootListRowID, result api.ErrorResult) {
	entry, ok := e.getEntry(id, result, "Failed copy path of desktop entry")
	if !ok {
		return
	}

	copyResult := api.NewFuncBoolResult(func(ok bool) {
		var err error
		if !ok {
			err = errors.New("clipboard copy failed")
			e.moduleLogger.Warn("Failed copy path of desktop entry",
				e.providerID.ZapField(),
				id.ZapField(),
				zap.String("EntryPath", entry.path),
			)
		}
		result.SetResult(err)
	})

	e.desktop.WriteToClipboard(false, mime.NewTextData(entry.path), copyResult)
}
