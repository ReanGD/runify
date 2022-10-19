package x11

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/logger"
	"github.com/ReanGD/runify/server/system/mime"
	"github.com/ReanGD/runify/server/system/module"
	"go.uber.org/zap"
)

const ModuleName = "x11"

type X11 struct {
	handler *x11Handler
	rpc     module.Rpc
	cfg     *config.Configuration

	module.Module
}

func New() *X11 {
	return &X11{
		handler: newX11Handler(),
		rpc:     nil,
		cfg:     nil,
	}
}

func (m *X11) OnInit(cfg *config.Config, rpc module.Rpc, rootLogger *zap.Logger) <-chan error {
	ch := make(chan error)

	go func() {
		m.rpc = rpc
		m.cfg = cfg.Get()
		channelLen := m.cfg.X11.ChannelLen
		m.Init(rootLogger, ModuleName, channelLen)

		ch <- m.handler.onInit(m.cfg, rpc, m.ModuleLogger)
	}()

	return ch
}

func (m *X11) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		m.handler.onStart(wg)
		m.ModuleLogger.Info("Start")

		for {
			if isFinish, err := m.safeRequestLoop(ctx); isFinish {
				m.handler.onStop()
				ch <- err
				wg.Done()
				return
			}
		}
	}()

	return ch
}

func (m *X11) safeRequestLoop(ctx context.Context) (resultIsFinish bool, resultErr error) {
	var request interface{}
	defer func() {
		if recoverResult := recover(); recoverResult != nil {
			if request != nil {
				reason := m.RecoverLog(recoverResult, request)
				resultIsFinish, resultErr = m.onRequestDefault(request, reason)
			} else {
				_ = m.RecoverLog(recoverResult, "unknown request")
			}
		}
	}()

	done := ctx.Done()
	errorCh := m.handler.getErrorCh()
	x11EventsCh := m.handler.getX11EventsCh()
	shortcutCh := m.handler.getShortcutCh()
	messageCh := m.GetReadChannel()

	bindID, _ := m.handler.bindShortcut(m.cfg.UI.ShowShortcut)

	for {
		request = nil
		select {
		case <-done:
			resultIsFinish = true
			resultErr = nil
			return
		// TODO: move this case
		case err := <-errorCh:
			resultIsFinish = true
			resultErr = err
			return
		case event := <-x11EventsCh:
			m.handler.onX11Event(event)
		case id := <-shortcutCh:
			if id == bindID {
				m.rpc.ShowUI()
			}
		case request = <-messageCh:
			m.MessageWasRead()
			if resultIsFinish, resultErr = m.onRequest(request); resultIsFinish {
				return
			}
		}
	}
}

func (m *X11) onRequest(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *writeToClipboardCmd:
		m.handler.writeToClipboard(r)
	default:
		m.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (m *X11) onRequestDefault(request interface{}, reason string) (bool, error) {
	switch r := request.(type) {
	case *writeToClipboardCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	default:
		m.ModuleLogger.Warn("Unknown message received",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.String("Reason", reason),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			logger.LogicalError)
		return true, errors.New("unknown message received")
	}

	return false, nil
}

func (m *X11) WriteToClipboard(isPrimary bool, data *mime.Data) {
	m.AddToChannel(&writeToClipboardCmd{
		isPrimary: isPrimary,
		data:      data,
	})
}
