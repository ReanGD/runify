package rpc

import (
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/pb"
)

type protoClient struct {
	id      uint32
	outCh   chan<- *pb.SrvMessage
	storage *formStorage
}

func newProtoClient(id uint32, outCh chan<- *pb.SrvMessage, storage *formStorage) *protoClient {
	return &protoClient{
		id:      id,
		outCh:   outCh,
		storage: storage,
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
	formID := c.storage.addRootList(ctrl)
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
	if !c.storage.isExists(formID) {
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
	if !c.storage.isExists(formID) {
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
	if !c.storage.isExists(formID) {
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
	formID := c.storage.addContextMenu(ctrl)
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

func (c *protoClient) UserMessage(text string) {
	c.outCh <- &pb.SrvMessage{
		FormID: 0,
		Payload: &pb.SrvMessage_UserMessage{
			UserMessage: &pb.UserMessage{
				Message:     text,
				MessageType: pb.MessageType_TYPE_ERROR,
			},
		},
	}
}

func (c *protoClient) CloseForm(formID api.FormID) {
	if !c.storage.remove(formID) {
		return
	}

	c.outCh <- &pb.SrvMessage{
		FormID: uint32(formID),
		Payload: &pb.SrvMessage_CloseForm{
			CloseForm: &pb.CloseForm{},
		},
	}
}

func (c *protoClient) HideUI(msg error) {
	c.storage.removeAll()

	var pbMsg *pb.UserMessage
	if msg != nil {
		pbMsg = &pb.UserMessage{
			MessageType: pb.MessageType_TYPE_ERROR,
			Message:     msg.Error(),
		}
	}

	c.outCh <- &pb.SrvMessage{
		FormID: 0,
		Payload: &pb.SrvMessage_HideUI{
			HideUI: &pb.HideUI{
				Message: pbMsg,
			},
		},
	}
}

func (c *protoClient) CloseUI() {
	c.storage.removeAll()

	c.outCh <- &pb.SrvMessage{
		FormID: 0,
		Payload: &pb.SrvMessage_CloseUI{
			CloseUI: &pb.CloseUI{},
		},
	}
}
