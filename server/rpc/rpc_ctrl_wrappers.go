package rpc

import (
	"errors"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
)

type rootListWrapper struct {
	ctrl         api.RootListCtrl
	moduleLogger *zap.Logger
}

func newRootListHandler(ctrl api.RootListCtrl, moduleLogger *zap.Logger) *rootListWrapper {
	return &rootListWrapper{
		ctrl:         ctrl,
		moduleLogger: moduleLogger,
	}
}

func (h *rootListWrapper) filterChanged(msg *pb.FilterData) error {
	h.ctrl.OnFilterChange(msg.Value)
	return nil
}

func (h *rootListWrapper) rootListRowActivated(msg *pb.RootListRowGlobalID) error {
	h.ctrl.OnRowActivate(api.ProviderID(msg.ProviderID), api.RootListRowID(msg.RowID))
	return nil
}

func (h *rootListWrapper) rootListMenuActivated(msg *pb.RootListRowGlobalID) error {
	h.ctrl.OnMenuActivate(api.ProviderID(msg.ProviderID), api.RootListRowID(msg.RowID))
	return nil
}

func (h *rootListWrapper) contextMenuRowActivated(msg *pb.ContextMenuRowID) error {
	h.moduleLogger.Error("Unexpected grpc message for root list handler", zap.String("Message", "ContextMenuRowActivated"))
	return errors.New("Unexpected grpc message")
}

type contextMenuWrapper struct {
	ctrl         api.ContextMenuCtrl
	moduleLogger *zap.Logger
}

func newContextMenuHandler(ctrl api.ContextMenuCtrl, moduleLogger *zap.Logger) *contextMenuWrapper {
	return &contextMenuWrapper{
		ctrl:         ctrl,
		moduleLogger: moduleLogger,
	}
}

func (h *contextMenuWrapper) filterChanged(msg *pb.FilterData) error {
	h.moduleLogger.Error("Unexpected grpc message for context menu handler", zap.String("Message", "FilterChanged"))
	return errors.New("Unexpected grpc message")
}

func (h *contextMenuWrapper) rootListRowActivated(msg *pb.RootListRowGlobalID) error {
	h.moduleLogger.Error("Unexpected grpc message for context menu handler", zap.String("Message", "RootListRowActivated"))
	return errors.New("Unexpected grpc message")
}

func (h *contextMenuWrapper) rootListMenuActivated(msg *pb.RootListRowGlobalID) error {
	h.moduleLogger.Error("Unexpected grpc message for context menu handler", zap.String("Message", "RootListMenuActivated"))
	return errors.New("Unexpected grpc message")
}

func (h *contextMenuWrapper) contextMenuRowActivated(msg *pb.ContextMenuRowID) error {
	h.ctrl.OnRowActivate(api.ContextMenuRowID(msg.RowID))
	return nil
}
