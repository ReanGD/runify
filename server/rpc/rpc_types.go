package rpc

import (
	"sync"

	"github.com/ReanGD/runify/server/pb"
)

type showUICmd struct {
}

type showUIMultiplier struct {
	nextID uint32
	subs   map[uint32]chan *pb.ShowWindow
	mutex  sync.RWMutex
}

func newShowUIMultiplier() *showUIMultiplier {
	return &showUIMultiplier{
		nextID: 0,
		subs:   make(map[uint32]chan *pb.ShowWindow),
		mutex:  sync.RWMutex{},
	}
}

func (m *showUIMultiplier) subscribe() (uint32, <-chan *pb.ShowWindow) {
	ch := make(chan *pb.ShowWindow, 1)
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

func (m *showUIMultiplier) sendToAll() {
	msg := &pb.ShowWindow{}
	m.mutex.RLock()
	for _, ch := range m.subs {
		ch <- msg
	}
	m.mutex.RUnlock()
}
