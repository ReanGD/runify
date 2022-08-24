package provider

import (
	"context"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

type providerHandler struct {
	dataProviders []*dataProvider
	moduleLogger  *zap.Logger
}

func newProviderHandler() *providerHandler {
	return &providerHandler{
		dataProviders: []*dataProvider{},
		moduleLogger:  nil,
	}
}

func (h *providerHandler) onInit(cfg *config.Config, moduleLogger *zap.Logger) error {
	h.moduleLogger = moduleLogger
	h.dataProviders = append(h.dataProviders, newDataProvider(desktopEntryID, newDesktopEntry()))

	dpChans := make([]<-chan error, 0, len(h.dataProviders))
	for _, dp := range h.dataProviders {
		dpChans = append(dpChans, dp.onInit(cfg, moduleLogger))
	}
	for _, dpChan := range dpChans {
		if err := <-dpChan; err != nil {
			return err
		}
	}

	return nil
}

func (h *providerHandler) onStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	errCh := make(chan error, 1)

	for _, dp := range h.dataProviders {
		dp.onStart(ctx, wg, errCh)
	}

	return errCh
}

func (h *providerHandler) getRoot() []*pb.Command {
	resultCh := make(chan []*pb.Command, len(h.dataProviders))
	for _, dp := range h.dataProviders {
		dp.getRoot(resultCh)
	}

	result := []*pb.Command{}
	for i := 0; i != len(h.dataProviders); i++ {
		result = append(result, <-resultCh...)
	}

	return result
}
