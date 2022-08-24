package provider

import "github.com/ReanGD/runify/server/pb"

const (
	desktopEntryID = uint64(1) << 32
)

type getRootCmd struct {
	result chan<- []*pb.Command
}
