package wood

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) {
	std.SetFormatter(formatter)
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include bool) {
	std.SetReportCaller(include)
}

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
	std.SetLevel(logrus.Level(level))
}

// GetLevel returns the standard logger level.
func GetLevel() Level {
	return Level(std.GetLevel())
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level param
func IsLevelEnabled(level Level) bool {
	return std.IsLevelEnabled(logrus.Level(level))
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook logrus.Hook) {
	std.AddHook(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *logrus.Entry {
	return std.WithField(logrus.ErrorKey, err)
}

// WithContext creates an entry from the standard logger and adds a context to it.
func WithContext(ctx context.Context) *logrus.Entry {
	return std.WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return std.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields logrus.Fields) *logrus.Entry {
	return std.WithFields(fields)
}

// WithTime creates an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *logrus.Entry {
	return std.WithTime(t)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(arguments ...interface{}) {
	if ignored(TraceLevel) {
		return
	}
	std.Trace(decorate(arguments...)...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(arguments ...interface{}) {
	if ignored(DebugLevel) {
		return
	}
	std.Debug(decorate(arguments...)...)
}

// Print logs a message at level Info on the standard logger.
func Print(arguments ...interface{}) {
	if ignored(InfoLevel) {
		return
	}
	std.Print(decorate(arguments...)...)
}

// Info logs a message at level Info on the standard logger.
func Info(arguments ...interface{}) {
	if ignored(InfoLevel) {
		return
	}
	std.Info(decorate(arguments...)...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(arguments ...interface{}) {
	if ignored(WarnLevel) {
		return
	}
	std.Warn(decorate(arguments...)...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(arguments ...interface{}) {
	if ignored(WarnLevel) {
		return
	}
	std.Warning(decorate(arguments...)...)
}

// Error logs a message at level Error on the standard logger.
func Error(arguments ...interface{}) {
	std.Error(decorate(arguments...)...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(arguments ...interface{}) {
	std.Panic(decorate(arguments...)...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(arguments ...interface{}) {
	std.Fatal(decorate(arguments...)...)
}

// Tracef logs a message with formatting at level Trace on the standard logger.
func Tracef(arguments ...interface{}) {
	decorateF(TraceLevel, arguments, func(format string, args []any) {
		std.Tracef(format, args...)
	})
}

// Debugf logs a message with formatting at level Debug on the standard logger.
func Debugf(arguments ...interface{}) {
	decorateF(InfoLevel, arguments, func(format string, args []any) {
		std.Debugf(format, args...)
	})
}

// Printf logs a message with formatting at level Info on the standard logger.
func Printf(arguments ...interface{}) {
	decorateF(InfoLevel, arguments, func(format string, args []any) {
		std.Printf(format, args...)
	})
}

// Infof logs a message with formatting at level Info on the standard logger.
func Infof(arguments ...interface{}) {
	decorateF(InfoLevel, arguments, func(format string, args []any) {
		std.Infof(format, args...)
	})
}

// Warnf logs a message at level Info on the standard logger.
func Warnf(arguments ...interface{}) {
	decorateF(WarnLevel, arguments, func(format string, args []any) {
		std.Warnf(format, args...)
	})
}

// Errorf logs a message with formatting at level Error on the standard logger.
func Errorf(arguments ...interface{}) {
	decorateF(ErrorLevel, arguments, func(format string, args []any) {
		std.Errorf(format, args...)
	})
}

// Panicf logs a message with formatting at level Panic on the standard logger.
func Panicf(arguments ...interface{}) {
	decorateF(PanicLevel, arguments, func(format string, args []any) {
		std.Panicf(format, args...)
	})
}

// Fatalf logs a message with formatting at level Fatal on the standard logger.
func Fatalf(arguments ...interface{}) {
	decorateF(FatalLevel, arguments, func(format string, args []any) {
		std.Fatalf(format, args...)
	})
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(arguments ...interface{}) {
	if ignored(TraceLevel) {
		return
	}
	std.Traceln(decorate(arguments...)...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(arguments ...interface{}) {
	if ignored(DebugLevel) {
		return
	}
	std.Debugln(decorate(arguments...)...)
}

// Println logs a message at level Info on the standard logger.
func Println(arguments ...interface{}) {
	if ignored(InfoLevel) {
		return
	}
	std.Println(decorate(arguments...)...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(arguments ...interface{}) {
	if ignored(InfoLevel) {
		return
	}
	std.Infoln(decorate(arguments...)...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(arguments ...interface{}) {
	if ignored(WarnLevel) {
		return
	}
	std.Warnln(decorate(arguments...)...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(arguments ...interface{}) {
	if ignored(WarnLevel) {
		return
	}
	std.Warningln(decorate(arguments...)...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(arguments ...interface{}) {
	std.Errorln(decorate(arguments...)...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(arguments ...interface{}) {
	std.Panicln(decorate(arguments...)...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(arguments ...interface{}) {
	std.Fatalln(decorate(arguments...)...)
}
