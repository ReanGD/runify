package rpc

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

type formWrapper struct {
	ctrl         api.FormCtrl
	moduleLogger *zap.Logger
}

func newFormWrapper(ctrl api.FormCtrl, moduleLogger *zap.Logger) *formWrapper {
	return &formWrapper{
		ctrl:         ctrl,
		moduleLogger: moduleLogger,
	}
}

func (w *formWrapper) filterChanged(msg *pb.FilterData) error {
	w.moduleLogger.Error("Unexpected grpc message for form handler", zap.String("Message", "FilterChanged"))
	return errors.New("Unexpected grpc message")
}

func (w *formWrapper) rootListRowActivated(msg *pb.RootListRowGlobalID) error {
	w.moduleLogger.Error("Unexpected grpc message for form handler", zap.String("Message", "RootListRowActivated"))
	return errors.New("Unexpected grpc message")
}

func (w *formWrapper) rootListMenuActivated(msg *pb.RootListRowGlobalID) error {
	w.moduleLogger.Error("Unexpected grpc message for form handler", zap.String("Message", "RootListMenuActivated"))
	return errors.New("Unexpected grpc message")
}

func (w *formWrapper) contextMenuRowActivated(msg *pb.ContextMenuRowID) error {
	w.moduleLogger.Error("Unexpected grpc message for form handler", zap.String("Message", "ContextMenuRowActivated"))
	return errors.New("Unexpected grpc message")
}

type rootListWrapper struct {
	ctrl         api.RootListCtrl
	moduleLogger *zap.Logger
}

func newRootListWrapper(ctrl api.RootListCtrl, moduleLogger *zap.Logger) *rootListWrapper {
	return &rootListWrapper{
		ctrl:         ctrl,
		moduleLogger: moduleLogger,
	}
}

func (w *rootListWrapper) filterChanged(msg *pb.FilterData) error {
	w.ctrl.OnFilterChange(msg.Value)
	return nil
}

func (w *rootListWrapper) rootListRowActivated(msg *pb.RootListRowGlobalID) error {
	w.ctrl.OnRowActivate(api.ProviderID(msg.ProviderID), api.RootListRowID(msg.RowID))
	return nil
}

func (w *rootListWrapper) rootListMenuActivated(msg *pb.RootListRowGlobalID) error {
	w.ctrl.OnMenuActivate(api.ProviderID(msg.ProviderID), api.RootListRowID(msg.RowID))
	return nil
}

func (w *rootListWrapper) contextMenuRowActivated(msg *pb.ContextMenuRowID) error {
	w.moduleLogger.Error("Unexpected grpc message for root list handler", zap.String("Message", "ContextMenuRowActivated"))
	return errors.New("Unexpected grpc message")
}

type contextMenuWrapper struct {
	ctrl         api.ContextMenuCtrl
	moduleLogger *zap.Logger
}

func newContextMenuWrapper(ctrl api.ContextMenuCtrl, moduleLogger *zap.Logger) *contextMenuWrapper {
	return &contextMenuWrapper{
		ctrl:         ctrl,
		moduleLogger: moduleLogger,
	}
}

func (w *contextMenuWrapper) filterChanged(msg *pb.FilterData) error {
	w.moduleLogger.Error("Unexpected grpc message for context menu handler", zap.String("Message", "FilterChanged"))
	return errors.New("Unexpected grpc message")
}

func (w *contextMenuWrapper) rootListRowActivated(msg *pb.RootListRowGlobalID) error {
	w.moduleLogger.Error("Unexpected grpc message for context menu handler", zap.String("Message", "RootListRowActivated"))
	return errors.New("Unexpected grpc message")
}

func (w *contextMenuWrapper) rootListMenuActivated(msg *pb.RootListRowGlobalID) error {
	w.moduleLogger.Error("Unexpected grpc message for context menu handler", zap.String("Message", "RootListMenuActivated"))
	return errors.New("Unexpected grpc message")
}

func (w *contextMenuWrapper) contextMenuRowActivated(msg *pb.ContextMenuRowID) error {
	w.ctrl.OnRowActivate(api.ContextMenuRowID(msg.RowID))
	return nil
}
