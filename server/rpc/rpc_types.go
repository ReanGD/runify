package rpc

import (
	"sync"
)

type showUICmd struct {
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
