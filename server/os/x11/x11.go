package x11

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/ReanGD/runify/server/logger"
	"go.uber.org/zap"
)

const ModuleName = "x11"

type X11 struct {
	handler     *x11Handler
	x11EventsCh chan interface{}
	module.Module
}

func New() *X11 {
	return &X11{
		handler:     newX11Handler(),
		x11EventsCh: nil,
	}
}

func (m *X11) OnInit(cfg *config.Config, rootLogger *zap.Logger) <-chan error {
	ch := make(chan error)

	go func() {
		m.Init(rootLogger, ModuleName, cfg.Get().X11.ChannelLen)
		m.x11EventsCh = make(chan interface{}, cfg.Get().X11.X11EventChannelLen)
		ch <- m.handler.init(m.x11EventsCh, m.ErrorCtx, m.ModuleLogger)
	}()

	return ch
}

func (m *X11) OnStart(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		m.handler.start()
		m.ModuleLogger.Info("Start")

		for {
			if isFinish, err := m.safeRequestLoop(ctx); isFinish {
				m.handler.stop()
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
	errorCh := m.ErrorCtx.GetChannel()
	x11EventsCh := m.x11EventsCh
	messageCh := m.GetReadChannel()

	for {
		request = nil
		select {
		case <-done:
			resultIsFinish = true
			resultErr = nil
			return
		case err := <-errorCh:
			resultIsFinish = true
			resultErr = err
			return
		case event := <-x11EventsCh:
			m.handler.onX11Event(event)
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
	case *subscribeToClipboardCmd:
		m.handler.subscribeToClipboard(r)
	case *writeToClipboardCmd:
		m.handler.writeToClipboard(r)
	case *subscribeToHotkeysCmd:
		m.handler.subscribeToHotkeys(r)
	case *bindHotkeyCmd:
		m.handler.bindHotkey(r)
	case *unbindHotkeyCmd:
		m.handler.unbindHotkey(r)

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
	case *subscribeToClipboardCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *writeToClipboardCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *subscribeToHotkeysCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *bindHotkeyCmd:
		r.onRequestDefault(m.ModuleLogger, reason)
	case *unbindHotkeyCmd:
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

func (m *X11) SubscribeToClipboard(isPrimary bool, ch chan<- *mime.Data) <-chan bool {
	result := make(chan bool, 1)
	m.AddToChannel(&subscribeToClipboardCmd{
		isPrimary: isPrimary,
		ch:        ch,
		result:    result,
	})

	return result
}

func (m *X11) WriteToClipboard(isPrimary bool, data *mime.Data) <-chan bool {
	result := make(chan bool, 1)
	m.AddToChannel(&writeToClipboardCmd{
		isPrimary: isPrimary,
		data:      data,
		result:    result,
	})

	return result
}

func (m *X11) SubscribeToHotkeys(ch chan<- *shortcut.Hotkey) <-chan bool {
	result := make(chan bool, 1)
	m.AddToChannel(&subscribeToHotkeysCmd{
		ch:     ch,
		result: result,
	})

	return result
}

func (m *X11) BindHotkey(hotkey *shortcut.Hotkey) <-chan global.Error {
	result := make(chan global.Error, 1)
	m.AddToChannel(&bindHotkeyCmd{
		hotkey: hotkey,
		result: result,
	})

	return result
}

func (m *X11) UnbindHotkey(hotkey *shortcut.Hotkey) <-chan bool {
	result := make(chan bool, 1)
	m.AddToChannel(&unbindHotkeyCmd{
		hotkey: hotkey,
		result: result,
	})

	return result
}
