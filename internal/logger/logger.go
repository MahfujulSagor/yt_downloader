package logger

import (
	"github/MahfujulSagor/video_downloader/internal/config"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func Init(cfg *config.Config) {
	//? Determina log file path
	logFilePath := cfg.LoggingConfig.File
	if logFilePath == "" {
		logFilePath = "logs/app.log"
	}

	//? Ensure log directory exists
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Failed to create log directory:", err)
	}

	logFile, err := os.OpenFile(filepath.Join("logs", "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file", err)
	}

	//? In development, log to both console and file
	var writer io.Writer
	if cfg.Env == "development" {
		writer = io.MultiWriter(os.Stdout, logFile)
	} else {
		writer = logFile
	}

	//? Define loggers with standard flags
	flags := log.Ldate | log.Ltime | log.Lshortfile
	Info = log.New(writer, "INFO: ", flags)
	Error = log.New(writer, "ERROR: ", flags)
}
