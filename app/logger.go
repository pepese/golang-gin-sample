package app

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const loggerKey = int(1)

var (
	desugaredLogger     *zap.Logger
	desugaredLoggerOnce sync.Once
	logger              *zap.SugaredLogger
	loggerOnce          sync.Once
)

func getDesugaredLogger() *zap.Logger {
	desugaredLoggerOnce.Do(func() {
		config := zap.Config{
			Level:    zap.NewAtomicLevelAt(zapcore.InfoLevel),
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "message",
				LevelKey:       "level",
				TimeKey:        "timestamp",
				NameKey:        "name",
				CallerKey:      "caller",
				StacktraceKey:  "stacktrace",
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
		zapLogger, err := config.Build()
		if err != nil {
			panic(err)
		}
		desugaredLogger = zapLogger.With(
			zap.String("app", GetConfig().AppName),
			zap.String("version", GetConfig().AppVersion),
		)
	})
	return desugaredLogger
}

func GetLogger() *zap.SugaredLogger {
	tmpLogger := getDesugaredLogger()
	loggerOnce.Do(func() {
		logger = tmpLogger.Sugar()
	})
	return logger
}

func GetLoggerFromContext(c context.Context) *zap.SugaredLogger {
	if c == nil {
		return GetLogger()
	}
	if zapLogger, ok := c.Value(loggerKey).(*zap.SugaredLogger); ok {
		return zapLogger
	}
	return GetLogger()

}

func SetLoggerToContext(c context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(c, loggerKey, logger)
}

func GetLoggerWithKeyValue(key string, value string) *zap.SugaredLogger {
	tmpLogger := getDesugaredLogger()
	tmpLogger = tmpLogger.With(
		zap.String(key, value),
	)
	return tmpLogger.Sugar()
}
