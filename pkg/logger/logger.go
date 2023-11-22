package logger

import "go.uber.org/zap"

func NewSugaredLogger(production bool, module string) *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	if production {
		logger, _ = zap.NewProduction()
	}
	suggarLogger := logger.Sugar()
	suggarLogger.With("module", module)
	return logger.Sugar()
}
