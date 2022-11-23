package logger

import (
	"fmt"
	"os"

	"github.com/ReanGD/runify/server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	level *zap.AtomicLevel
	root  *zap.Logger
}

func newEncoder(format config.LoggerFormat) zapcore.Encoder {
	if format == config.JSONFormat {
		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func newCore(cfg *config.LoggerCfg, level zapcore.LevelEnabler, kraangID zap.Field) (zapcore.Core, error) {
	encoder := newEncoder(cfg.Format)

	var sync zapcore.WriteSyncer

	if cfg.EnableRotate && cfg.Output != "stderr" && cfg.Output != "stdout" {
		sync = zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Output,
			MaxSize:    cfg.MaxSizeMb,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAgeDays,
			LocalTime:  cfg.LocalTimeInBackupFilename,
			Compress:   cfg.Compress,
		})
	} else {
		if cfg.EnableRotate {
			fmt.Fprintf(os.Stdout, "Can't apply rotate logs for '%s' output", cfg.Output)
		}

		var err error
		sync, _, err = zap.Open(cfg.Output)
		if err != nil {
			return nil, err
		}
	}

	return zapcore.NewCore(encoder, sync, level).With([]zap.Field{kraangID}), nil
}

func newOptions(cfg *config.LoggerCfg) ([]zap.Option, error) {
	opts := []zap.Option{zap.AddStacktrace(cfg.LevelStacktrace)}

	if cfg.AddCallerInfo {
		opts = append(opts, zap.AddCaller())
	}

	loggerErrSink, _, err := zap.Open("stderr")
	if err != nil {
		return nil, err
	}
	opts = append(opts, zap.ErrorOutput(loggerErrSink))

	return opts, nil
}

func New(cfg *config.Config, appID string) (*Logger, error) {
	logCfg := cfg.Get().Logger
	level := zap.NewAtomicLevelAt(logCfg.Level)
	obj := &Logger{
		level: &level,
		root:  nil,
	}

	core, err := newCore(logCfg, obj.level, zap.String("AppID", appID))
	if err != nil {
		return nil, err
	}

	opts, err := newOptions(logCfg)
	if err != nil {
		return nil, err
	}
	obj.root = zap.New(core, opts...)

	return obj, nil
}

func (l *Logger) SetLevel(lvl zapcore.Level) {
	l.level.SetLevel(lvl)
}

func (l *Logger) GetLevel() zapcore.Level {
	return l.level.Level()
}

func (l *Logger) GetRoot() *zap.Logger {
	return l.root
}
