package tools

import (
	"log"
	"log/slog"
	"os"
)

var GlobalLogger *slog.Logger
var GinLogger *slog.Logger

func init() {
	logFile, err := os.OpenFile("gin.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Failed to initiate log file: %v", err)
	}

	handler := slog.NewJSONHandler(logFile, nil)
	GinLogger = slog.New(handler)

	internalFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Failed to initiate log file: %v", err)
	}

	textHandler := slog.NewTextHandler(internalFile, nil)
	GlobalLogger = slog.New(textHandler)
}
