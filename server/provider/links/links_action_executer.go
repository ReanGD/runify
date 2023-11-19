package links

import (
	"errors"
	"os/exec"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/mime"
	"go.uber.org/zap"
)

type actionExecuter struct {
	desktop      api.Desktop
	moduleLogger *zap.Logger
}

func newActionExecuter() *actionExecuter {
	return &actionExecuter{
		desktop:      nil,
		moduleLogger: nil,
	}
}

func (e *actionExecuter) init(desktop api.Desktop, moduleLogger *zap.Logger) error {
	e.desktop = desktop
	e.moduleLogger = moduleLogger

	return nil
}

func (e *actionExecuter) openLink(client api.RpcClient, itemData *DataModel) {
	cmd := exec.Command("xdg-open", itemData.Link)
	if err := cmd.Start(); err != nil {
		e.moduleLogger.Error("Failed open link",
			zap.String("link", itemData.Link),
			zap.Error(err))
		client.HideUI(err)
		return
	}

	e.moduleLogger.Debug("Open link success", zap.String("link", itemData.Link))
	go cmd.Wait()

	client.HideUI(nil)
}

func (e *actionExecuter) copyValue(client api.RpcClient, text string) {
	copyResult := api.NewChanBoolResult()
	e.desktop.WriteToClipboard(false, mime.NewTextData(text), copyResult)
	res := <-copyResult.GetChannel()
	if !res {
		e.moduleLogger.Warn("Failed copy value to clipboard")
		client.HideUI(errors.New("Failed copy calculator result"))
	} else {
		client.HideUI(nil)
	}
}
