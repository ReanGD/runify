package module

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/global/types"
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

type Module struct {
	ErrorCtx     *ErrorCtx
	impl         api.ModuleImpl
	cfg          *config.Configuration
	rootLogger   *zap.Logger
	moduleLogger *zap.Logger
	name         string

	Channel
}

func (m *Module) Create(impl api.ModuleImpl, name string, cfg *config.Configuration, rootLogger *zap.Logger) {
	m.ErrorCtx = newErrorCtx()
	m.impl = impl
	m.cfg = cfg
	m.rootLogger = rootLogger
	m.moduleLogger = nil
	m.name = name
}

func (m *Module) Init(channelLen uint32) {
	m.moduleLogger = m.rootLogger.With(zap.String("Module", m.name))
	m.Channel.Init(channelLen)
}

func (m *Module) InitSubmodule(channelLen uint32) {
	m.moduleLogger = m.NewSubmoduleLogger(m.name)
	m.Channel.Init(channelLen)
}

func (m *Module) NewSubmoduleLogger(submoduleName string) *zap.Logger {
	return m.rootLogger.With(zap.String("SubModule", submoduleName))
}

func (m *Module) GetName() string {
	return m.name
}

func (m *Module) GetConfig() *config.Configuration {
	return m.cfg
}

func (m *Module) GetRootLogger() *zap.Logger {
	return m.rootLogger
}

func (m *Module) GetModuleLogger() *zap.Logger {
	return m.moduleLogger
}

func (m *Module) recoverLog(recoverResult interface{}, request interface{}) string {
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
		m.moduleLogger.Error("Panic during message processing",
			zap.String("Request", fmt.Sprintf("%v", request)),
			zap.Stringer("RequestType", reflect.TypeOf(request)),
			zap.Error(err),
			logger.LogicalError)
	} else {
		m.moduleLogger.Error("Panic during message processing",
			zap.String("Request", "unknown"),
			zap.String("RequestType", "unknown"),
			zap.Error(err),
			logger.LogicalError)
	}

	return "panic during message processing: " + err.Error()
}

func (m *Module) safeRequestLoop(ctx context.Context, hChannels []*types.HandledChannel) (resultIsFinish bool, resultErr error) {
	var request interface{}
	defer func() {
		if recoverResult := recover(); recoverResult != nil {
			if request != nil {
				reason := m.recoverLog(recoverResult, request)
				resultIsFinish, resultErr = m.impl.OnRequestDefault(request, reason)
			} else {
				_ = m.recoverLog(recoverResult, "unknown request")
			}
		}
	}()

	onRequest := m.impl.OnRequest
	done := ctx.Done()
	messageCh := m.GetReadChannel()
	cntCh := len(hChannels) + 2
	doneChIdx := cntCh - 1
	messageChIdx := cntCh - 2

	cases := make([]reflect.SelectCase, cntCh)
	for i, hc := range hChannels {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: hc.Channel()}
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
			m.moduleLogger.Error("Error during message processing",
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
			if resultIsFinish, resultErr = hChannels[chosen].Handle(recv.Interface()); resultIsFinish {
				return
			}
		}
	}
}

func (m *Module) Start(ctx context.Context, wg *sync.WaitGroup) <-chan error {
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		defer wg.Done()
		m.moduleLogger.Info("Start")
		hChs := m.impl.OnStart(ctx)

		for {
			if isFinish, err := m.safeRequestLoop(ctx, hChs); isFinish {
				m.impl.OnFinish()
				ch <- err
				return
			}
		}
	}()

	return ch
}

func (m *Module) OnRequestUnknownMsg(request interface{}) (bool, error) {
	m.moduleLogger.Warn("Unknown message received",
		zap.String("Request", fmt.Sprintf("%v", request)),
		zap.Stringer("RequestType", reflect.TypeOf(request)),
		logger.LogicalError)
	return true, errors.New("unknown message received")
}

func (m *Module) OnRequestDefaultUnknownMsg(request interface{}, reason string) (bool, error) {
	m.moduleLogger.Warn("Unknown message received",
		zap.String("Request", fmt.Sprintf("%v", request)),
		zap.String("Reason", reason),
		zap.Stringer("RequestType", reflect.TypeOf(request)),
		logger.LogicalError)
	return true, errors.New("unknown message received")
}
