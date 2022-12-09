package rpc

import (
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

type protoClient struct {
	outCh        chan<- *pb.FormDataMsgSrv
	forms        map[uint32]formHandler
	moduleLogger *zap.Logger
}

func newProtoClient(outCh chan<- *pb.FormDataMsgSrv, moduleLogger *zap.Logger) *protoClient {
	return &protoClient{
		outCh: outCh,
		forms: make(map[uint32]formHandler),
	}
}

func (c *protoClient) getHandler(formID uint32, msgName string) (formHandler, bool) {
	handler, ok := c.forms[formID]
	if !ok {
		c.moduleLogger.Debug("Grpc message for unknown form",
			zap.Uint32("FormID", formID),
			zap.String("Message", msgName),
		)
		return nil, false
	}

	return handler, true
}

func (c *protoClient) filterChanged(formID uint32, msg *pb.FilterData) error {
	if handler, ok := c.getHandler(formID, "FilterChanged"); ok {
		return handler.filterChanged(msg)
	}

	return nil
}

func (c *protoClient) rootListRowActivated(formID uint32, msg *pb.RootListRowGlobalID) error {
	if handler, ok := c.getHandler(formID, "RootListRowActivated"); ok {
		return handler.rootListRowActivated(msg)
	}

	return nil
}

func (c *protoClient) rootListMenuActivated(formID uint32, msg *pb.RootListRowGlobalID) error {
	if handler, ok := c.getHandler(formID, "RootListMenuActivated"); ok {
		return handler.rootListMenuActivated(msg)
	}

	return nil
}

func (c *protoClient) contextMenuRowActivated(formID uint32, msg *pb.ContextMenuRowID) error {
	if handler, ok := c.getHandler(formID, "ContextMenuRowActivated"); ok {
		return handler.contextMenuRowActivated(msg)
	}

	return nil
}

func (c *protoClient) formClosed(formID uint32) {
	delete(c.forms, formID)
}

func (c *protoClient) rootListRowsToProtobuf(rows []*api.RootListRow) []*pb.RootListRow {
	pbRows := make([]*pb.RootListRow, len(rows))
	for i, row := range rows {
		pbRows[i] = row.ToProtobuf()
	}

	return pbRows
}

func (c *protoClient) rootListRowGlobalIDsToProtobuf(rows []*api.RootListRowGlobalID) []*pb.RootListRowGlobalID {
	pbRows := make([]*pb.RootListRowGlobalID, len(rows))
	for i, row := range rows {
		pbRows[i] = row.ToProtobuf()
	}

	return pbRows
}

func (c *protoClient) contextMenuRowsToProtobuf(rows []*api.ContextMenuRow) []*pb.ContextMenuRow {
	pbRows := make([]*pb.ContextMenuRow, len(rows))
	for i, row := range rows {
		pbRows[i] = row.ToProtobuf()
	}

	return pbRows
}

func (c *protoClient) RootListOpen(formID uint32, ctrl api.RootListCtrl, rows []*api.RootListRow) {
	c.forms[formID] = newRootListHandler(ctrl, c.moduleLogger.With(zap.Uint32("FormID", formID)))
	c.outCh <- &pb.FormDataMsgSrv{
		FormID: formID,
		Payload: &pb.FormDataMsgSrv_RootListOpen{
			RootListOpen: &pb.RootListOpen{
				Rows: c.rootListRowsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) RootListAddRows(formID uint32, rows ...*api.RootListRow) {
	c.outCh <- &pb.FormDataMsgSrv{
		FormID: formID,
		Payload: &pb.FormDataMsgSrv_RootListAddRows{
			RootListAddRows: &pb.RootListAddRows{
				Rows: c.rootListRowsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) RootListChangeRows(formID uint32, rows ...*api.RootListRow) {
	c.outCh <- &pb.FormDataMsgSrv{
		FormID: formID,
		Payload: &pb.FormDataMsgSrv_RootListChangeRows{
			RootListChangeRows: &pb.RootListChangeRows{
				Rows: c.rootListRowsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) RootListRemoveRows(formID uint32, rows ...*api.RootListRowGlobalID) {
	c.outCh <- &pb.FormDataMsgSrv{
		FormID: formID,
		Payload: &pb.FormDataMsgSrv_RootListRemoveRows{
			RootListRemoveRows: &pb.RootListRemoveRows{
				Rows: c.rootListRowGlobalIDsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) ContextMenuOpen(formID uint32, ctrl api.ContextMenuCtrl, rows ...*api.ContextMenuRow) {
	c.forms[formID] = newContextMenuHandler(ctrl, c.moduleLogger.With(zap.Uint32("FormID", formID)))
	c.outCh <- &pb.FormDataMsgSrv{
		FormID: formID,
		Payload: &pb.FormDataMsgSrv_ContextMenuOpen{
			ContextMenuOpen: &pb.ContextMenuOpen{
				Rows: c.contextMenuRowsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) CloseAll(msg error) {
	c.forms = make(map[uint32]formHandler)
	var pbMsg *pb.UserMessage
	if msg != nil {
		pbMsg = &pb.UserMessage{
			MessageType: pb.MessageType_TYPE_ERROR,
			Message:     msg.Error(),
		}
	}

	c.outCh <- &pb.FormDataMsgSrv{
		FormID: 0,
		Payload: &pb.FormDataMsgSrv_FormAction{
			FormAction: &pb.FormAction{
				ActionType: pb.FormActionType_CLOSE_ALL,
				Message:    pbMsg,
			},
		},
	}
}

func (c *protoClient) CloseOne(formID uint32, msg error) {
	delete(c.forms, formID)
	var pbMsg *pb.UserMessage
	if msg != nil {
		pbMsg = &pb.UserMessage{
			MessageType: pb.MessageType_TYPE_ERROR,
			Message:     msg.Error(),
		}
	}

	c.outCh <- &pb.FormDataMsgSrv{
		FormID: formID,
		Payload: &pb.FormDataMsgSrv_FormAction{
			FormAction: &pb.FormAction{
				ActionType: pb.FormActionType_CLOSE_ONE,
				Message:    pbMsg,
			},
		},
	}
}

func (c *protoClient) ShowMessage(msg error) {
	c.outCh <- &pb.FormDataMsgSrv{
		FormID: 0,
		Payload: &pb.FormDataMsgSrv_FormAction{
			FormAction: &pb.FormAction{
				ActionType: pb.FormActionType_SHOW_MESSAGE,
				Message: &pb.UserMessage{
					MessageType: pb.MessageType_TYPE_ERROR,
					Message:     msg.Error(),
				},
			},
		},
	}
}
