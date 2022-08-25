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
	data := <-s.provider.GetRoot()
	return &pb.Commands{
		Data: data,
	}, nil
}

func (s *runifyServer) GetActions(ctx context.Context, selectedCommand *pb.SelectedCommand) (*pb.Actions, error) {
	data := <-s.provider.GetActions(selectedCommand.CommandID)
	return &pb.Actions{
		Data: data,
	}, nil
}

func (s *runifyServer) Execute(ctx context.Context, selectedAction *pb.SelectedAction) (*pb.Result, error) {
	return nil, nil
}
