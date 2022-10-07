package rpc

import (
	"context"

	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/pb"
	"github.com/ReanGD/runify/server/system/module"
)

type runifyServer struct {
	provider         module.Provider
	showUIMultiplier *showUIMultiplier
	moduleLogger     *zap.Logger

	pb.UnimplementedRunifyServer
}

func newRunifyServer(provider module.Provider, showUIMultiplier *showUIMultiplier, moduleLogger *zap.Logger) *runifyServer {
	return &runifyServer{
		provider:         provider,
		showUIMultiplier: showUIMultiplier,
		moduleLogger:     moduleLogger,
	}
}

func (s *runifyServer) WaitShowWindow(empty *pb.Empty, stream pb.Runify_WaitShowWindowServer) error {
	id, ch := s.showUIMultiplier.subscribe()
	defer s.showUIMultiplier.unsubscribe(id)

	isLive := true
	for isLive {
		select {
		case msg := <-ch:
			if err := stream.Send(msg); err != nil {
				s.moduleLogger.Warn("Failed send ShowWindow to client", zap.Error(err))
				return err
			}
		case <-stream.Context().Done():
			isLive = false
		}
	}

	return nil
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
