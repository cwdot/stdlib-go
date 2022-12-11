package wood

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

var (
	label       string
	buff        string
	stack       []string
	prefixLevel map[string]logrus.Level

	// std is the name of the standard logger in stdlib `log`
	std = logrus.New()
)

const yellow = 33

const pad = 35

func decorate(args ...interface{}) []any {
	c := make([]any, 0, len(args)+1)
	if label != "" {
		var trimmed string
		var whitespace string
		m := len(label) + len(buff)
		if m < pad {
			trimmed = label
			m = pad - m
			if m > 0 {
				whitespace = strings.Repeat(" ", m)
			}
		} else {
			trimmed = label[0 : pad-len(buff)]
			whitespace = ">"
		}

		text := fmt.Sprintf("%s\x1b[%dm%s\x1b[0m%s", buff, yellow, trimmed, whitespace)
		c = append(c, text)
	}
	c = append(c, fmt.Sprint(args...))
	return c
}

func Prefix(newPrefix string) {
	label = newPrefix
	stack = append(stack, label)
	refreshLabel()
}

func Reset() {
	l := len(stack)
	if len(stack) == 0 {
		label = ""
		refreshLabel()
		return
	}
	stack = slices.Delete(stack, l-1, l)
	l = len(stack)
	if l == 0 {
		label = ""
	} else {
		label = stack[l-1]
	}
	refreshLabel()
}

func refreshLabel() {
	buff = strings.Repeat(" ", len(stack))
}

func Init(level logrus.Level) {
	std.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		//DisableColors: !colors,
		//ForceQuote:                false,
		//DisableQuote:              false,
		//EnvironmentOverrideColors: false,
		DisableTimestamp: true,
		FullTimestamp:    false,
		//TimestampFormat:           "",
		//DisableSorting:            false,
		//SortingFunc:               nil,
		//DisableLevelTruncation:    false,
		//PadLevelText:              false,
		//QuoteEmptyFields:          false,
		//FieldMap:                  nil,
	})
	std.SetLevel(level)
	stack = make([]string, 0, 5)
	prefixLevel = make(map[string]logrus.Level)
}

func ignored(action logrus.Level) bool {
	if label == "" {
		return false
	}
	val, ok := prefixLevel[label]
	if !ok {
		return false
	}
	return val < action
}

func PrefixLevel(label string, level logrus.Level) {
	prefixLevel[label] = level
}

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

//// Tracef logs a message at level Trace on the standard logger.
//func Tracef(format string, args ...interface{}) {
//	std.Tracef(format, args...)
//}
//
//// Debugf logs a message at level PrintDebug on the standard logger.
//func Debugf(format string, args ...interface{}) {
//	std.Debugf(format, args...)
//}
//
//// Printf logs a message at level Info on the standard logger.
//func Printf(format string, args ...interface{}) {
//	std.Printf(format, args...)
//}
//
//// Infof logs a message at level Info on the standard logger.
//func Infof(format string, args ...interface{}) {
//	std.Infof(format, args...)
//}
//
//// Warnf logs a message at level Warn on the standard logger.
//func Warnf(format string, args ...interface{}) {
//	std.Warnf(format, args...)
//}
//
//// Warningf logs a message at level Warn on the standard logger.
//func Warningf(format string, args ...interface{}) {
//	std.Warningf(format, args...)
//}
//
//// Errorf logs a message at level Error on the standard logger.
//func Errorf(format string, args ...interface{}) {
//	std.Errorf(format, args...)
//}
//
//// Panicf logs a message at level Panic on the standard logger.
//func Panicf(format string, args ...interface{}) {
//	std.Panicf(format, args...)
//}
//
//// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
//func Fatalf(format string, args ...interface{}) {
//	std.Fatalf(format, args...)
//}

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
