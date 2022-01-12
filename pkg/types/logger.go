package types

import (
	"log"
	"os"
)

type BasicLogger struct {
	*log.Logger
}

func InitBasicLogger(logPath string) *BasicLogger {
	logger := &BasicLogger{log.Default()}
	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}
	logger.SetOutput(f)
	return logger
}

func (l *BasicLogger) Debug(format string, v ...interface{}) {
	l.SetPrefix("[DEBUG] ")
	l.Printf(format, v)
}

func (l *BasicLogger) Error(format string, v ...interface{}) {
	l.SetPrefix("[ERROR] ")
	l.Printf(format, v)
}

func (l *BasicLogger) Info(format string, v ...interface{}) {
	l.SetPrefix("[INFO] ")
	l.Printf(format, v)
}
