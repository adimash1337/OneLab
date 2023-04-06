package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func Logger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
	return logger
}
