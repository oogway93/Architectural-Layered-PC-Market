package logger

import (
	"log/slog"
	"os"

	"github.com/oogway93/golangArchitecture/internal/adapter/config"
	multi "github.com/samber/slog-multi"
	"gopkg.in/natefinch/lumberjack.v2"
)

// logger is the default logger used by the application
var logger *slog.Logger

// Set sets the logger configuration based on the environment
func Set(config *config.App) {
	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)
	// wd, _ := os.Getwd()
	// slog.Info(wd)
	// path := fmt.Sprintf("%s%s", wd, "./logs/application.log")
	if config.Env == "production" {
		logRotate := &lumberjack.Logger{
			Filename:   "logs/application.log",
			MaxSize:    200, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
			Compress:   true,
		}

		logger = slog.New(
			multi.Fanout(
				slog.NewJSONHandler(logRotate, nil),
				slog.NewTextHandler(os.Stderr, nil),
			),
		)
	}
	
	slog.SetDefault(logger)
}