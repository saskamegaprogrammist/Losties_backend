package utils

import (
	"github.com/google/logger"
	"os"
)

const logPath  = "log.log"
const verbose = true
var loggerLosties *logger.Logger
var loggerFile *os.File

func LoggerSetup() {
	var err error
	loggerFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	loggerLosties = logger.Init("LoggerExample", verbose, true, loggerFile)
}

func WriteError (fatal bool, errString string, err error) {
	if fatal {
		loggerLosties.Fatalf("%s: %v", errString, err)
		return
	}
	loggerLosties.Errorf("%s: %v", errString, err)
}

func Close() {
	loggerFile.Close()
	loggerLosties.Close()
}