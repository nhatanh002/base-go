package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogType = string
type LogLevel = string

const (
	AccessLog LogType = "access_log"
	AppLog    LogType = "app_log"
	ErrorLog  LogType = "error_log"

	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelError LogLevel = "error"
)

var log *zap.Logger

// Init initialize log settings of zap
// TODO: LogConfig
func Init() {
	var err error

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         "json",
		DisableCaller:    true,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:       "level",
			TimeKey:        "timestamp",
			MessageKey:     "message",
			CallerKey:      "caller",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	log, err = config.Build()
	if err != nil {
		panic(any(err))
	}
}

// Output write log contents according to given log level
func Output(level LogLevel, msg string, fields ...any) {
	if log == nil {
		Init()
	}
	switch level {
	case LevelDebug:
		Debug(msg, fields...)
	case LevelInfo:
		Info(msg, fields...)
	case LevelError:
		Error(msg, fields...)
	}
}

func Debug(msg string, args ...any) {
	defer log.Sync()
	log.Debug(fmt.Sprintf(msg, args...))
}

func Info(msg string, args ...any) {
	defer log.Sync()
	log.Info(fmt.Sprintf(msg, args...))
}

func Error(msg string, args ...any) {
	defer log.Sync()
	log.Error(fmt.Sprintf(msg, args...))
}
