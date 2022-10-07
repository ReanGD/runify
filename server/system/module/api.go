package module

import (
	"github.com/ReanGD/runify/server/pb"
	"github.com/ReanGD/runify/server/system/mime"
)

type Provider interface {
	GetRoot() <-chan []*pb.CardItem
	GetActions(cardID uint64) <-chan *pb.Actions
	Execute(cardID uint64, actionID uint32) <-chan *pb.Result
}

type Rpc interface {
	ShowUI()
}

type X11 interface {
	WriteToClipboard(isPrimary bool, data *mime.Data)
}
