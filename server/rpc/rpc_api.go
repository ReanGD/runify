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

func (s *runifyServer) GetRoot(context.Context, *pb.Empty) (*pb.Form, error) {
	data := <-s.provider.GetRoot()
	return &pb.Form{
		Cards: data,
	}, nil
}

func (s *runifyServer) GetActions(ctx context.Context, selectedCard *pb.SelectedCard) (*pb.Actions, error) {
	data := <-s.provider.GetActions(selectedCard.CardID)
	return data, nil
}

func (s *runifyServer) ExecuteDefault(ctx context.Context, selectedCard *pb.SelectedCard) (*pb.Result, error) {
	data := <-s.provider.Execute(selectedCard.CardID, 1)
	return data, nil
}

func (s *runifyServer) Execute(ctx context.Context, selectedAction *pb.SelectedAction) (*pb.Result, error) {
	data := <-s.provider.Execute(selectedAction.CardID, selectedAction.ActionID)
	return data, nil
}
