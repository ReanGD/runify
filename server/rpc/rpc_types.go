package rpc

import (
	"context"
	"sync"

	"github.com/ReanGD/runify/server/global/api"
)

type showUICmd struct{}

type openRootListCmd struct {
	ctrl api.RootListCtrl
}

type showUIMultiplier struct {
	nextID uint32
	subs   map[uint32]chan struct{}
	mutex  sync.RWMutex
}

func newShowUIMultiplier() *showUIMultiplier {
	return &showUIMultiplier{
		nextID: 0,
		subs:   make(map[uint32]chan struct{}),
		mutex:  sync.RWMutex{},
	}
}

func (m *showUIMultiplier) subscribe() (uint32, <-chan struct{}) {
	ch := make(chan struct{}, 1)
	m.mutex.Lock()
	id := m.nextID
	m.nextID++
	m.subs[id] = ch
	m.mutex.Unlock()
	return id, ch
}

func (m *showUIMultiplier) unsubscribe(id uint32) {
	m.mutex.Lock()
	delete(m.subs, id)
	m.mutex.Unlock()
}

func (m *showUIMultiplier) sendToAll() bool {
	msg := struct{}{}
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	for _, ch := range m.subs {
		ch <- msg
	}

	return len(m.subs) > 0
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
