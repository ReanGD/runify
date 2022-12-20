package rpc

import (
	"errors"
	"io"
	"reflect"
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/api"
	"github.com/ReanGD/runify/server/pb"
)

type runifyServer struct {
	rpc          *Rpc
	cfg          *config.Configuration
	uiLogger     *zap.Logger
	moduleLogger *zap.Logger
	lastClientID uint32

	pb.UnimplementedRunifyServer
}

func newRunifyServer(
	rpc *Rpc,
	cfg *config.Configuration,
	uiLogger *zap.Logger,
	moduleLogger *zap.Logger,
) *runifyServer {
	return &runifyServer{
		cfg:          cfg,
		rpc:          rpc,
		uiLogger:     uiLogger,
		moduleLogger: moduleLogger,
		lastClientID: 0,
	}
}

func (s *runifyServer) logGrpcError(messageType string, err error) {
	s.moduleLogger.Warn("Failed process grpc message", zap.String("MessageType", messageType), zap.Error(err))
}

func (s *runifyServer) writeUILog(msg *pb.WriteLog) error {
	var level zapcore.Level
	switch msg.Level {
	case pb.LogLevel_DEBUG:
		level = zapcore.DebugLevel
	case pb.LogLevel_INFO:
		level = zapcore.InfoLevel
	case pb.LogLevel_WARNING:
		level = zapcore.WarnLevel
	case pb.LogLevel_ERROR:
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	if ce := s.uiLogger.Check(level, msg.Message); ce != nil {
		ce.Write()
	}

	return nil
}

func (s *runifyServer) Connect(stream pb.Runify_ConnectServer) error {
	processor := newStreamProcessor(stream.Context())
	forms := newFormStorage(s.moduleLogger)
	outCh := make(chan *pb.SrvMessage, s.cfg.Rpc.SendMsgChLen)
	id := atomic.AddUint32(&s.lastClientID, 1)
	pClient := newProtoClient(id, outCh, forms)
	s.rpc.uiClientConnected(pClient)

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
				s.moduleLogger.Warn("Failed receive grpc message", zap.Error(err))
				return err
			}

			switch m := req.Payload.(type) {
			case *pb.UIMessage_WriteLog:
				if err = s.writeUILog(m.WriteLog); err != nil {
					s.logGrpcError("WriteLog", err)
					return err
				}
			case *pb.UIMessage_FilterChanged:
				if err = forms.filterChanged(api.FormID(req.FormID), m.FilterChanged); err != nil {
					s.logGrpcError("FilterChanged", err)
					return err
				}
			case *pb.UIMessage_RootListRowActivated:
				if err = forms.rootListRowActivated(api.FormID(req.FormID), m.RootListRowActivated); err != nil {
					s.logGrpcError("RootListRowActivated", err)
					return err
				}
			case *pb.UIMessage_RootListMenuActivated:
				if err = forms.rootListMenuActivated(api.FormID(req.FormID), m.RootListMenuActivated); err != nil {
					s.logGrpcError("RootListMenuActivated", err)
					return err
				}
			case *pb.UIMessage_ContextMenuRowActivated:
				if err = forms.contextMenuRowActivated(api.FormID(req.FormID), m.ContextMenuRowActivated); err != nil {
					s.logGrpcError("ContextMenuRowActivated", err)
					return err
				}
			case *pb.UIMessage_FormClosed:
				if err = forms.formClosed(api.FormID(req.FormID)); err != nil {
					s.logGrpcError("FormClosed", err)
					return err
				}
			default:
				err = errors.New("recv unknown message type")
				s.logGrpcError(reflect.TypeOf(m).String(), err)
				return err
			}
		}
	})

	processor.runSend(func(doneCh <-chan struct{}) error {
		for {
			select {
			case msg := <-outCh:
				if err := stream.Send(msg); err != nil {
					s.moduleLogger.Info("Failed send grpc message", zap.Error(err))
					return err
				}
			case <-doneCh:
				return nil
			}
		}
	})

	result := processor.wait()
	s.rpc.uiClientDisconnected(pClient.id)
	return result
}
