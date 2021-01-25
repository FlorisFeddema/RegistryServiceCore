package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
var logger *zap.Logger

func SetupLogger()  {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ = config.Build()
}

func Logger() *zap.Logger {
	return logger
}

