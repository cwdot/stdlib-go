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
func SetLevel(level logrus.Level) {
	std.SetLevel(level)
}

// GetLevel returns the standard logger level.
func GetLevel() logrus.Level {
	return std.GetLevel()
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level param
func IsLevelEnabled(level logrus.Level) bool {
	return std.IsLevelEnabled(level)
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
func Trace(args ...interface{}) {
	if ignored(logrus.TraceLevel) {
		return
	}
	std.Trace(decorate(args...)...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if ignored(logrus.DebugLevel) {
		return
	}
	std.Debug(decorate(args...)...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	if ignored(logrus.InfoLevel) {
		return
	}
	std.Print(decorate(args...)...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if ignored(logrus.InfoLevel) {
		return
	}
	std.Info(decorate(args...)...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if ignored(logrus.WarnLevel) {
		return
	}
	std.Warn(decorate(args...)...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	if ignored(logrus.WarnLevel) {
		return
	}
	std.Warning(decorate(args...)...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	std.Error(decorate(args...)...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	std.Panic(decorate(args...)...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	std.Fatal(decorate(args...)...)
}

// Tracef logs a message with formatting at level Trace on the standard logger.
func Tracef(args ...interface{}) {
	decorateF(logrus.TraceLevel, args, func(format string, args []any) {
		std.Tracef(format, args...)
	})
}

// Debugf logs a message with formatting at level Debug on the standard logger.
func Debugf(args ...interface{}) {
	decorateF(logrus.InfoLevel, args, func(format string, args []any) {
		std.Debugf(format, args...)
	})
}

// Printf logs a message with formatting at level Info on the standard logger.
func Printf(args ...interface{}) {
	decorateF(logrus.InfoLevel, args, func(format string, args []any) {
		std.Printf(format, args...)
	})
}

// Infof logs a message with formatting at level Info on the standard logger.
func Infof(args ...interface{}) {
	decorateF(logrus.InfoLevel, args, func(format string, args []any) {
		std.Infof(format, args...)
	})
}

// Warnf logs a message at level Info on the standard logger.
func Warnf(args ...interface{}) {
	decorateF(logrus.WarnLevel, args, func(format string, args []any) {
		std.Warnf(format, args...)
	})
}

// Errorf logs a message with formatting at level Error on the standard logger.
func Errorf(args ...interface{}) {
	decorateF(logrus.ErrorLevel, args, func(format string, args []any) {
		std.Errorf(format, args...)
	})
}

// Panicf logs a message with formatting at level Panic on the standard logger.
func Panicf(args ...interface{}) {
	decorateF(logrus.PanicLevel, args, func(format string, args []any) {
		std.Panicf(format, args...)
	})
}

// Fatalf logs a message with formatting at level Fatal on the standard logger.
func Fatalf(args ...interface{}) {
	decorateF(logrus.FatalLevel, args, func(format string, args []any) {
		std.Fatalf(format, args...)
	})
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	if ignored(logrus.TraceLevel) {
		return
	}
	std.Traceln(decorate(args...)...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	if ignored(logrus.DebugLevel) {
		return
	}
	std.Debugln(decorate(args...)...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	if ignored(logrus.InfoLevel) {
		return
	}
	std.Println(decorate(args...)...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	if ignored(logrus.InfoLevel) {
		return
	}
	std.Infoln(decorate(args...)...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	if ignored(logrus.WarnLevel) {
		return
	}
	std.Warnln(decorate(args...)...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	if ignored(logrus.WarnLevel) {
		return
	}
	std.Warningln(decorate(args...)...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	std.Errorln(decorate(args...)...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	std.Panicln(decorate(args...)...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	std.Fatalln(decorate(args...)...)
}
