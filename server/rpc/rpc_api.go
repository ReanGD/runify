package rpc

import (
	"context"

	"github.com/ReanGD/runify/server/pb"
	"github.com/ReanGD/runify/server/provider"
)

type runifyServer struct {
	provider *provider.Provider

	pb.UnimplementedRunifyServer
}

func newRunifyServer(provider *provider.Provider) *runifyServer {
	return &runifyServer{
		provider: provider,
	}
}

func (s *runifyServer) GetRoot(context.Context, *pb.Empty) (*pb.Commands, error) {
	cmds := <-s.provider.GetRoot()
	return &pb.Commands{
		Data: cmds,
	}, nil
}

func (s *runifyServer) GetActions(ctx context.Context, commandID *pb.CommandID) (*pb.Actions, error) {
	return nil, nil
}

func (s *runifyServer) Execute(ctx context.Context, actionID *pb.ActionID) (*pb.Result, error) {
	return nil, nil
}
