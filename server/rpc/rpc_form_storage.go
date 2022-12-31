package rpc

import (
	"sync"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

type formHandler interface {
	filterChanged(msg *pb.FilterData) error
	rootListRowActivated(msg *pb.RootListRowGlobalID) error
	rootListMenuActivated(msg *pb.RootListRowGlobalID) error
	contextMenuRowActivated(msg *pb.ContextMenuRowID) error
}

type formStorage struct {
	m            sync.RWMutex
	forms        map[api.FormID]formHandler
	moduleLogger *zap.Logger
	lastID       api.FormID
}

func newFormStorage(moduleLogger *zap.Logger) *formStorage {
	return &formStorage{
		m:            sync.RWMutex{},
		forms:        make(map[api.FormID]formHandler),
		moduleLogger: moduleLogger,
		lastID:       0,
	}
}

func (s *formStorage) addForm(ctrl api.FormCtrl) api.FormID {
	s.m.Lock()
	s.lastID++
	id := s.lastID
	s.forms[s.lastID] = newFormWrapper(ctrl, s.moduleLogger.With(zap.Uint32("FormID", uint32(id))))
	s.m.Unlock()

	return id
}

func (s *formStorage) addRootList(ctrl api.RootListCtrl) api.FormID {
	s.m.Lock()
	s.lastID++
	id := s.lastID
	s.forms[s.lastID] = newRootListWrapper(ctrl, s.moduleLogger.With(zap.Uint32("FormID", uint32(id))))
	s.m.Unlock()

	return id
}

func (s *formStorage) addContextMenu(ctrl api.ContextMenuCtrl) api.FormID {
	s.m.Lock()
	s.lastID++
	id := s.lastID
	s.forms[s.lastID] = newContextMenuWrapper(ctrl, s.moduleLogger.With(zap.Uint32("FormID", uint32(id))))
	s.m.Unlock()

	return id
}

func (s *formStorage) remove(formID api.FormID) bool {
	s.m.Lock()
	_, ok := s.forms[formID]
	delete(s.forms, formID)
	s.m.Unlock()

	return ok
}

func (s *formStorage) removeAll() {
	s.m.Lock()
	s.forms = make(map[api.FormID]formHandler)
	s.m.Unlock()
}

func (s *formStorage) isExists(formID api.FormID) bool {
	s.m.RLock()
	_, ok := s.forms[formID]
	s.m.RUnlock()

	return ok
}

func (s *formStorage) getForHandle(formID api.FormID, msgName string) (formHandler, bool) {
	s.m.RLock()
	handler, ok := s.forms[formID]
	s.m.RUnlock()
	if !ok {
		s.moduleLogger.Debug("Grpc message for unknown form",
			zap.Uint32("FormID", uint32(formID)),
			zap.String("Message", msgName),
		)
		return nil, false
	}

	return handler, ok
}

func (s *formStorage) filterChanged(formID api.FormID, msg *pb.FilterData) error {
	if handler, ok := s.getForHandle(formID, "FilterChanged"); ok {
		return handler.filterChanged(msg)
	}

	return nil
}

func (s *formStorage) rootListRowActivated(formID api.FormID, msg *pb.RootListRowGlobalID) error {
	if handler, ok := s.getForHandle(formID, "RootListRowActivated"); ok {
		return handler.rootListRowActivated(msg)
	}

	return nil
}

func (s *formStorage) rootListMenuActivated(formID api.FormID, msg *pb.RootListRowGlobalID) error {
	if handler, ok := s.getForHandle(formID, "RootListMenuActivated"); ok {
		return handler.rootListMenuActivated(msg)
	}

	return nil
}

func (s *formStorage) contextMenuRowActivated(formID api.FormID, msg *pb.ContextMenuRowID) error {
	if handler, ok := s.getForHandle(formID, "ContextMenuRowActivated"); ok {
		return handler.contextMenuRowActivated(msg)
	}

	return nil
}

func (s *formStorage) formClosed(formID api.FormID) error {
	s.remove(formID)

	return nil
}
