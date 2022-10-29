package app

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (a *Application) initLogger() error {
	loggerCfg := zap.NewProductionConfig()

	loggerCfg.Encoding = a.config.Log.Format
	loggerCfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	if a.config.Log.Debug {
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := loggerCfg.Build()
	if err != nil {
		return err
	}

	a.logger = logger
	return nil
}
