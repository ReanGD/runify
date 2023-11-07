package module

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync/atomic"

	"github.com/ReanGD/runify/server/logger"
	"go.uber.org/zap"
)

type ErrorCtx struct {
	ch chan error
}

func newErrorCtx() *ErrorCtx {
	return &ErrorCtx{
		ch: make(chan error, 1),
	}
}

func (e *ErrorCtx) GetChannel() <-chan error {
	return e.ch
}

func (e *ErrorCtx) SendError(err error) {
	select {
	case e.ch <- err:
	default:
	}
}

type Channel struct {
	messageCh chan interface{}

	queueOverflow int32
}

func (c *Channel) Init(channelLen uint32) {
	c.messageCh = make(chan interface{}, channelLen)
	c.queueOverflow = 0
}

func (c *Channel) GetReadChannel() <-chan interface{} {
	return c.messageCh
}

func (c *Channel) AddToChannel(value interface{}) {
	select {
	case c.messageCh <- value:
		atomic.StoreInt32(&c.queueOverflow, 0)
	default:
		atomic.StoreInt32(&c.queueOverflow, 1)
		c.messageCh <- value
	}
}

func (c *Channel) MessageWasRead() {
}

func (c *Channel) ClearModuleQueue() {
	for {
		select {
		case <-c.messageCh:
			c.MessageWasRead()
		default:
			return
		}
	}
}

func (c *Channel) IsOverflow() bool {
	return atomic.LoadInt32(&c.queueOverflow) == 1
}

type HandledChannel struct {
	ch      reflect.Value
	handler func(interface{}) (bool, error)
}

func NewHandledChannel[T any](ch <-chan T, handler func(interface{}) (bool, error)) *HandledChannel {
	return &HandledChannel{
		ch:      reflect.ValueOf(ch),
		handler: handler,
	}
}

type Module struct {
	ErrorCtx     *ErrorCtx
	RootLogger   *zap.Logger
	ModuleLogger *zap.Logger

	Channel
}

func (m *Module) Init(rootLogger *zap.Logger, moduleName string, channelLen uint32) {
	m.ErrorCtx = newErrorCtx()
	m.RootLogger = rootLogger
	m.ModuleLogger = rootLogger.With(zap.String("Module", moduleName))
	m.Channel.Init(channelLen)
}

func (m *Module) InitSubmodule(rootLogger *zap.Logger, submoduleName string, channelLen uint32) {
	m.RootLogger = rootLogger
	m.ModuleLogger = m.NewSubmoduleLogger(rootLogger, submoduleName)
	m.Channel.Init(channelLen)
}

func (m *Module) NewSubmoduleLogger(rootLogger *zap.Logger, submoduleName string) *zap.Logger {
	return rootLogger.With(zap.String("SubModule", submoduleName))
}

func (m *Module) RecoverLog(recoverResult interface{}, request interface{}) string {
	var err error
	switch val := recoverResult.(type) {
	case string:
		err = errors.New(val)
	case error:
		err = val
	default:
		err = fmt.Errorf("%v", val)
	}

	if request != nil {
		m.ModuleLogger.Error("Panic during message processing",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			zap.Error(err),
			logger.LogicalError)
	} else {
		m.ModuleLogger.Error("Panic during message processing",
			zap.String("Request", "unknown"),
			zap.String("RequestType", "unknown"),
			zap.Error(err),
			logger.LogicalError)
	}

	return "panic during message processing: " + err.Error()
}

func (m *Module) SafeRequestLoop(
	ctx context.Context,
	onRequest func(request interface{}) (bool, error),
	onRequestDefault func(request interface{}, reason string) (bool, error),
	hChannels []*HandledChannel,
) (resultIsFinish bool, resultErr error) {
	var request interface{}
	defer func() {
		if recoverResult := recover(); recoverResult != nil {
			if request != nil {
				reason := m.RecoverLog(recoverResult, request)
				resultIsFinish, resultErr = onRequestDefault(request, reason)
			} else {
				_ = m.RecoverLog(recoverResult, "unknown request")
			}
		}
	}()

	done := ctx.Done()
	messageCh := m.GetReadChannel()
	cntCh := len(hChannels) + 2
	doneChIdx := cntCh - 1
	messageChIdx := cntCh - 2

	cases := make([]reflect.SelectCase, cntCh)
	for i, hc := range hChannels {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: hc.ch}
	}

	cases[messageChIdx] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(messageCh)}
	cases[doneChIdx] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(done)}

	for {
		request = nil
		chosen, recv, recvOk := reflect.Select(cases)
		if chosen == doneChIdx {
			resultIsFinish = true
			resultErr = nil
			return
		}

		if !recvOk {
			resultIsFinish = true
			resultErr = errors.New("channel closed")
			m.ModuleLogger.Error("Error during message processing",
				zap.String("Request", "unknown"),
				zap.String("RequestType", "unknown"),
				zap.Error(resultErr),
				logger.LogicalError)
		}

		if chosen == messageChIdx {
			request = recv.Interface()
			m.MessageWasRead()
			if resultIsFinish, resultErr = onRequest(request); resultIsFinish {
				return
			}
		} else {
			if resultIsFinish, resultErr = hChannels[chosen].handler(recv.Interface()); resultIsFinish {
				return
			}
		}
	}
}
