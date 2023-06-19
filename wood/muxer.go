package wood

import (
	"github.com/sirupsen/logrus"
)

// Log logs a message the provided text using the provided log level
func Log(level Level, args ...interface{}) {
	switch level {
	case logrus.TraceLevel:
		Trace(args...)
	case logrus.DebugLevel:
		Debug(args...)
	case logrus.InfoLevel:
		Info(args...)
	case logrus.WarnLevel:
		Warn(args...)
	case logrus.ErrorLevel:
		Error(args...)
	case logrus.FatalLevel:
		Fatal(args...)
	case logrus.PanicLevel:
		Panic(args...)
	}
}

// Logln logs a message with newline the provided log level
func Logln(level Level, args ...interface{}) {
	switch level {
	case logrus.TraceLevel:
		Traceln(args...)
	case logrus.DebugLevel:
		Debugln(args...)
	case logrus.InfoLevel:
		Infoln(args...)
	case logrus.WarnLevel:
		Warnln(args...)
	case logrus.ErrorLevel:
		Errorln(args...)
	case logrus.FatalLevel:
		Fatalln(args...)
	case logrus.PanicLevel:
		Panicln(args...)
	}
}

// Logf logs a message with formatting using the provided log level
func Logf(level Level, args ...interface{}) {
	decorateF(level, args, func(format string, args []any) {
		std.Infof(format, args...)
	})
}
