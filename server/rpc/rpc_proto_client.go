package rpc

import (
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/pb"
)

type protoClient struct {
	outCh chan<- *pb.SrvMessage
	forms *formStorage
}

func newProtoClient(outCh chan<- *pb.SrvMessage, forms *formStorage) *protoClient {
	return &protoClient{
		outCh: outCh,
		forms: forms,
	}
}

func (c *protoClient) rootListRowsToProtobuf(rows []*api.RootListRow) []*pb.RootListRow {
	pbRows := make([]*pb.RootListRow, len(rows))
	for i, row := range rows {
		pbRows[i] = row.ToProtobuf()
	}

	return pbRows
}

func (c *protoClient) rootListRowGlobalIDsToProtobuf(rows []api.RootListRowGlobalID) []*pb.RootListRowGlobalID {
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

func (c *protoClient) AddRootList(ctrl api.RootListCtrl) {
	formID := c.forms.addRootList(ctrl)
	rows := ctrl.OnOpen(formID, c)
	c.outCh <- &pb.SrvMessage{
		FormID: uint32(formID),
		Payload: &pb.SrvMessage_RootListOpen{
			RootListOpen: &pb.RootListOpen{
				Rows: c.rootListRowsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) RootListAddRows(formID api.FormID, rows ...*api.RootListRow) {
	if !c.forms.isExists(formID) {
		return
	}

	c.outCh <- &pb.SrvMessage{
		FormID: uint32(formID),
		Payload: &pb.SrvMessage_RootListAddRows{
			RootListAddRows: &pb.RootListAddRows{
				Rows: c.rootListRowsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) RootListChangeRows(formID api.FormID, rows ...*api.RootListRow) {
	if !c.forms.isExists(formID) {
		return
	}

	c.outCh <- &pb.SrvMessage{
		FormID: uint32(formID),
		Payload: &pb.SrvMessage_RootListChangeRows{
			RootListChangeRows: &pb.RootListChangeRows{
				Rows: c.rootListRowsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) RootListRemoveRows(formID api.FormID, rows ...api.RootListRowGlobalID) {
	if !c.forms.isExists(formID) {
		return
	}

	c.outCh <- &pb.SrvMessage{
		FormID: uint32(formID),
		Payload: &pb.SrvMessage_RootListRemoveRows{
			RootListRemoveRows: &pb.RootListRemoveRows{
				Rows: c.rootListRowGlobalIDsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) AddContextMenu(ctrl api.ContextMenuCtrl) {
	formID := c.forms.addContextMenu(ctrl)
	rows := ctrl.OnOpen(formID, c)
	c.outCh <- &pb.SrvMessage{
		FormID: uint32(formID),
		Payload: &pb.SrvMessage_ContextMenuOpen{
			ContextMenuOpen: &pb.ContextMenuOpen{
				Rows: c.contextMenuRowsToProtobuf(rows),
			},
		},
	}
}

func (c *protoClient) CloseAll(msg error) {
	c.forms.removeAll()

	var pbMsg *pb.UserMessage
	if msg != nil {
		pbMsg = &pb.UserMessage{
			MessageType: pb.MessageType_TYPE_ERROR,
			Message:     msg.Error(),
		}
	}

	c.outCh <- &pb.SrvMessage{
		FormID: 0,
		Payload: &pb.SrvMessage_FormAction{
			FormAction: &pb.FormAction{
				ActionType: pb.FormActionType_CLOSE_ALL,
				Message:    pbMsg,
			},
		},
	}
}

func (c *protoClient) CloseOne(formID api.FormID, msg error) {
	if !c.forms.remove(formID) {
		return
	}

	var pbMsg *pb.UserMessage
	if msg != nil {
		pbMsg = &pb.UserMessage{
			MessageType: pb.MessageType_TYPE_ERROR,
			Message:     msg.Error(),
		}
	}

	c.outCh <- &pb.SrvMessage{
		FormID: uint32(formID),
		Payload: &pb.SrvMessage_FormAction{
			FormAction: &pb.FormAction{
				ActionType: pb.FormActionType_CLOSE_ONE,
				Message:    pbMsg,
			},
		},
	}
}

func (c *protoClient) ShowMessage(msg error) {
	c.outCh <- &pb.SrvMessage{
		FormID: 0,
		Payload: &pb.SrvMessage_FormAction{
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
