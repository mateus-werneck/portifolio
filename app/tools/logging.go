package tools

import (
	"log"
	"log/slog"
	"os"
)

var GlobalLogger *slog.Logger

func init() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Failed to initiate log file: %v", err)
	}

	handler := slog.NewJSONHandler(logFile, nil)
	GlobalLogger = slog.New(handler)
}
