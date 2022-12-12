package rpc

import (
	"context"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"go.uber.org/zap"
)

type uiClientConnectedCmd struct {
	pClient *protoClient
}

func (c *uiClientConnectedCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "UIClientConnected"),
		zap.String("Reason", reason),
		zap.String("Action", "skip request"))
}

type uiClientDisconnectedCmd struct{}

func (c *uiClientDisconnectedCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "UIClientDisconnected"),
		zap.String("Reason", reason),
		zap.String("Action", "skip request"))
}

type openRootListCmd struct {
	ctrl api.RootListCtrl
}

func (c *openRootListCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "OpenRootList"),
		zap.String("Reason", reason),
		zap.String("Action", "skip request"))
}

type streamProcessor struct {
	wg        sync.WaitGroup
	stopMutex sync.Mutex
	doneCh    <-chan struct{}
	cancel    context.CancelFunc
	err       error
}

func newStreamProcessor(streamCtx context.Context) *streamProcessor {
	ctx, cancel := context.WithCancel(streamCtx)
	res := &streamProcessor{
		wg:        sync.WaitGroup{},
		stopMutex: sync.Mutex{},
		doneCh:    ctx.Done(),
		cancel:    cancel,
		err:       nil,
	}
	res.wg.Add(1)

	return res
}

func (p *streamProcessor) stop(err error) {
	p.stopMutex.Lock()
	defer p.stopMutex.Unlock()

	if p.cancel != nil {
		p.err = err
		p.cancel()
		p.cancel = nil
	}
}

func (p *streamProcessor) runRecv(fn func(doneCh <-chan struct{}) error) {
	go func() {
		defer p.wg.Done()
		p.stop(fn(p.doneCh))
	}()
}

func (p *streamProcessor) runSend(fn func(doneCh <-chan struct{}) error) {
	p.stop(fn(p.doneCh))
}

func (p *streamProcessor) wait() error {
	p.wg.Wait()
	return p.err
}
