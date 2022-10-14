package helpers

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.SetReportCaller(true)
	standardLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
	})
	standardLogger.SetOutput(os.Stdout)

	// standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	Error = Event{1, "Error ▶ %s"}
	Debug = Event{2, "Info ▶ %s"}
)

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(Error.message, argumentName)
}

// MissingArg is a standard error message
func (l *StandardLogger) InfoArg(argumentName string) {
	l.Infof(Debug.message, argumentName)
}