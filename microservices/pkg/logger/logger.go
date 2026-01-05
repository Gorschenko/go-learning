package logger

import (
	"context"
	"log/slog"
	"os"
	"pkg/configs"
	"pkg/static"
)

func SetupLogger(config *configs.Config) {
	addSource := config.Software.Logger.AddSource || false
	logLevel := slog.LevelInfo

	if config.Software.Logger.Level == "debug" {
		logLevel = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: addSource, // Добавить информацию о файле и строке
	}))

	slog.SetDefault(logger)
}

func GetLogger(ctx context.Context) *slog.Logger {
	defaultLogger := slog.Default()
	correlationID, ok := ctx.Value(static.ContextCorrelationID).(string)

	if ok {
		return defaultLogger.With("CorrelationID", correlationID)
	}

	return slog.Default()
}
