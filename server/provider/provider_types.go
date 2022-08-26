package provider

import "github.com/ReanGD/runify/server/pb"

const (
	desktopEntryID = uint64(1) << 32
	providerIDMask = uint64(0xFFFF) << 32
	commandIDMask  = (uint64(1) << 32) - 1
)

type getRootCmd struct {
	result chan<- []*pb.Command
}

type getActionsCmd struct {
	commandID uint64
	result    chan<- []*pb.Action
}

type executeCmd struct {
	commandID uint64
	actionID  uint32
	result    chan<- *pb.Result
}
