// The logging package provides functions for logging
package logging

import (
	"os"
	"path/filepath"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
)

// The name of the log file to write
const FileName = "app.log"

// Init initializes the Logger for use
func Init(debugLogging bool) {
	logLevels := slog.DangerLevels
	minLevel := slog.WarnLevel
	if debugLogging {
		logLevels = append(logLevels, slog.InfoLevel, slog.NoticeLevel, slog.DebugLevel)
		minLevel = slog.DebugLevel
	}

	consoleFormatter := slog.NewTextFormatter()
	consoleFormatter.EnableColor = true
	consoleFormatter.SetTemplate("[{{level}}] ({{caller}}) {{message}} {{data}} {{extra}}\n")

	consoleHandler := handler.NewConsoleHandler(logLevels)
	consoleHandler.SetFormatter(consoleFormatter)

	fileFormatter := slog.NewTextFormatter()
	fileFormatter.EnableColor = false
	fileFormatter.TimeFormat = "2006-01-02T15:04:05.000"
	fileFormatter.SetTemplate("{{datetime}} ({{caller}}) [{{level}}] {{message}} {{data}} {{extra}}\n")
	err := os.MkdirAll(folderstructure.OutputFolderName, 0755)
	if err != nil {
		panic(err)
	}
	fileStream, err := os.OpenFile(filepath.Join(folderstructure.OutputFolderName, FileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileHandler := handler.NewSimple(fileStream, minLevel)
	fileHandler.SetFormatter(fileFormatter)

	slog.Std().ResetHandlers()
	slog.AddHandlers(fileHandler, consoleHandler)
}
