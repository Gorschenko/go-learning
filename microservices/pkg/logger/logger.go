package logger

import (
	"log/slog"
	"os"
	"pkg/configs"
)

type LoggerServiceDependencies struct {
	Config *configs.Config
}

func SetLogger(dependencies *LoggerServiceDependencies) {
	logLevel := slog.LevelInfo
	addSource := false

	if dependencies.Config.Software.Logger.Level == "debug" {
		logLevel = slog.LevelDebug
		addSource = true
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: addSource, // Добавить информацию о файле и строке
	}))

	slog.SetDefault(logger)

	slog.Info(
		"Logger is set up",
		"Level",
		logLevel,
	)
}
