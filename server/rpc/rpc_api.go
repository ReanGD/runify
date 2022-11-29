package rpc

import (
	"context"
	"errors"
	"io"
	"reflect"

	"go.uber.org/zap"

	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/pb"
)

type runifyServer struct {
	provider         api.Provider
	showUIMultiplier *showUIMultiplier
	handler          *uiHandler
	moduleLogger     *zap.Logger

	pb.UnimplementedRunifyServer
}

func newRunifyServer(
	provider api.Provider, showUIMultiplier *showUIMultiplier, handler *uiHandler, moduleLogger *zap.Logger) *runifyServer {

	return &runifyServer{
		provider:         provider,
		showUIMultiplier: showUIMultiplier,
		handler:          handler,
		moduleLogger:     moduleLogger,
	}
}

func (s *runifyServer) ServiceChannel(stream pb.Runify_ServiceChannelServer) error {
	processor := newStreamProcessor(stream.Context())
	log := s.moduleLogger.With(zap.String("Method", "ServiceChannel"))

	processor.runRecv(func(doneCh <-chan struct{}) error {
		for {
			select {
			case <-doneCh:
				return nil
			default:
			}

			req, err := stream.Recv()
			if err == io.EOF {
				return nil
			}

			if err != nil {
				log.Warn("Failed receive grpc message", zap.Error(err))
				return err
			}

			switch m := req.Payload.(type) {
			case *pb.ServiceMsgUI_WriteLog:
				if err = s.handler.writeLog(m.WriteLog); err != nil {
					log.Warn("Failed process grpc message", zap.String("MessageType", "WriteLog"), zap.Error(err))
					return err
				}
			default:
				err = errors.New("recv unknown message type")
				log.Warn("Failed process grpc message", zap.String("MessageType", reflect.TypeOf(m).String()), zap.Error(err))
				return err
			}
		}
	})

	id, sendMsgCh := s.showUIMultiplier.subscribe()
	defer s.showUIMultiplier.unsubscribe(id)

	processor.runSend(func(doneCh <-chan struct{}) error {
		for {
			select {
			case <-sendMsgCh:
				msg := &pb.ServiceMsgSrv{
					Payload: &pb.ServiceMsgSrv_SetFormState{
						SetFormState: &pb.SetFormState{
							State: pb.FormStateType_SHOW,
						},
					},
				}
				if err := stream.Send(msg); err != nil {
					log.Warn("Failed send grpc message", zap.String("MessageType", "SetFormState"), zap.Error(err))
					return err
				}
			case <-doneCh:
				return nil
			}
		}
	})

	return processor.wait()
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
