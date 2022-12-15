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

func (e *calcActionExecuter) copyResult(client api.RpcClient, text string) {
	copyResult := api.NewChanBoolResult()
	e.desktop.WriteToClipboard(false, mime.NewTextData(text), copyResult)
	res := <-copyResult.GetChannel()
	if !res {
		e.moduleLogger.Warn("Failed copy calculator result",
			rootRowID.ZapField(),
		)
		client.HideUI(errors.New("Failed copy calculator result"))
	}

	client.HideUI(nil)
}
