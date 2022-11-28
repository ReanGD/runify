package calculator

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"go.uber.org/zap"
)

type calcActionExecuter struct {
	desktop      api.Desktop
	moduleLogger *zap.Logger
}

func newCalcActionExecuter() *calcActionExecuter {
	return &calcActionExecuter{
		desktop:      nil,
		moduleLogger: nil,
	}
}

func (e *calcActionExecuter) init(desktop api.Desktop, moduleLogger *zap.Logger) error {
	e.desktop = desktop
	e.moduleLogger = moduleLogger

	return nil
}

func (e *calcActionExecuter) copyResult(text string, result api.ErrorResult) {
	copyResult := api.NewFuncBoolResult(func(ok bool) {
		var err error
		if !ok {
			err = errors.New("clipboard copy failed")
			e.moduleLogger.Warn("Failed copy result",
				rootRowID.ZapField(),
			)
		}
		result.SetResult(err)
	})
	e.desktop.WriteToClipboard(false, mime.NewTextData(text), copyResult)
}
