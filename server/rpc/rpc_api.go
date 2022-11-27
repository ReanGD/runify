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
	doneCh := stream.Context().Done()
	recvFinishCh := make(chan struct{})
	sendFinishCh := make(chan struct{})
	errorCh := make(chan error, 2)
	zapMethod := zap.String("Method", "ServiceChannel")

	go func() {
		defer close(recvFinishCh)
		for {
			select {
			case <-doneCh:
				return
			case <-sendFinishCh:
				return
			default:
			}

			req, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				errorCh <- err
				s.moduleLogger.Warn("Failed receive grpc message", zapMethod, zap.Error(err))
				return
			}

			switch m := req.Payload.(type) {
			case *pb.ServiceMsgUI_WriteLog:
				if err = s.handler.writeLog(m.WriteLog); err != nil {
					s.moduleLogger.Warn("Failed process grpc message",
						zapMethod, zap.String("MessageType", "WriteLog"), zap.Error(err))
					errorCh <- err
					return
				}
			default:
				err = errors.New("recv unknown message type")
				s.moduleLogger.Warn("Failed process grpc message",
					zapMethod, zap.String("MessageType", reflect.TypeOf(m).String()), zap.Error(err))
				errorCh <- err
				return
			}
		}
	}()

	id, sendMsgCh := s.showUIMultiplier.subscribe()
	defer s.showUIMultiplier.unsubscribe(id)

loop:
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
				s.moduleLogger.Warn("Failed send grpc message",
					zapMethod, zap.String("MessageType", "SetFormState"), zap.Error(err))
				return err
			}
		case <-doneCh:
			break loop
		case <-recvFinishCh:
			break loop
		}
	}
	close(sendFinishCh)

	<-recvFinishCh
	select {
	case err := <-errorCh:
		return err
	default:
		return nil
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
	data := <-s.provider.Execute(selectedCard.CardID, 0)
	return data, nil
}

func (s *runifyServer) Execute(ctx context.Context, selectedAction *pb.SelectedAction) (*pb.Result, error) {
	data := <-s.provider.Execute(selectedAction.CardID, selectedAction.ActionID)
	return data, nil
}
