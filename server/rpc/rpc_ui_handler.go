package rpc

import (
	"github.com/ReanGD/runify/server/pb"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type uiHandler struct {
	uiLogger     *zap.Logger
	moduleLogger *zap.Logger
}

func newUIHandler(uiLogger *zap.Logger, moduleLogger *zap.Logger) *uiHandler {
	return &uiHandler{
		uiLogger:     uiLogger,
		moduleLogger: moduleLogger,
	}
}

func (h *uiHandler) writeLog(msg *pb.WriteLog) error {
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

	if ce := h.uiLogger.Check(level, msg.Message); ce != nil {
		ce.Write()
	}

	return nil
}
