package module

import (
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

type Module struct {
	ErrorCtx     *ErrorCtx
	RootLogger   *zap.Logger
	ModuleLogger *zap.Logger

	Channel
}

func (m *Module) Init(rootLogger *zap.Logger, moduleName string, channelLen uint32) {
	m.ErrorCtx = newErrorCtx()
	m.RootLogger = rootLogger
	m.ModuleLogger = rootLogger.With(zap.String("module", moduleName))
	m.Channel.Init(channelLen)
}

func (m *Module) InitSubmodule(rootLogger *zap.Logger, submoduleName string, channelLen uint32) {
	m.RootLogger = rootLogger
	m.ModuleLogger = rootLogger.With(zap.String("module", submoduleName))
	m.Channel.Init(channelLen)
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
