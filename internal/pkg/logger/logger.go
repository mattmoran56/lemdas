package logger

import (
	"os"

	"go.uber.org/zap"
)

func Init() {
	var logger *zap.Logger

	if os.Getenv("DEV") == "true" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	zap.ReplaceGlobals(logger)
}
