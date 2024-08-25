package wood

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/porfirion/trie"
	"github.com/sirupsen/logrus"

	"github.com/cwdot/stdlib-go/color"
)

var (
	currentCanonical string
	// currentDisplay Text shown
	currentDisplay string
	// displayWhitespace cache of spaces in front of the currentDisplay; computed from stack
	displayWhitespace string

	// stack of labels
	stack []*logStack
	// indent custom indention for a particular prefix; resets when the currentDisplay changes
	indent int

	components map[string]Level

	// prefixLevel Sets the logging level for specific prefixes
	prefixes trie.Trie[Level]

	// std is the name of the standard logger in stdlib `log`
	std = logrus.New()
)

func WithLevel(level Level) func(*Opts) {
	return func(opts *Opts) {
		opts.level = level
	}
}

func WithNoColors() func(*Opts) {
	return func(opts *Opts) {
		opts.noColors = true
	}
}

func WithVerbosity(verbosity int) func(*Opts) {
	return func(opts *Opts) {
		switch verbosity {
		case 3:
			opts.level = TraceLevel
		case 2:
			opts.level = TraceLevel
		case 1:
			opts.level = DebugLevel
		default:
			opts.level = InfoLevel
		}
	}
}

type Opts struct {
	level     Level
	output    io.Writer
	formatter logrus.Formatter
	noColors  bool
}

func Init(opts ...func(opts *Opts)) {
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

	settings := &Opts{level: InfoLevel}
	for _, opt := range opts {
		opt(settings)
	}
	std.SetLevel(logrus.Level(settings.level))
	if settings.output != nil {
		std.SetOutput(io.MultiWriter(os.Stderr, settings.output))
	}
	if settings.formatter != nil {
		std.SetFormatter(settings.formatter)
	}
	if settings.noColors {
		std.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
		})
		color.Disable()
	}

	stack = make([]*logStack, 0, 5)
	currentDisplay = ""
	currentCanonical = ""
	displayWhitespace = ""

	prefixes = trie.Trie[Level]{}
	components = make(map[string]Level)
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

func DumpState() {
	fmt.Println("Stack:", stack)
	fmt.Println("components:", components)
	fmt.Println("Prefixes:", prefixes)
}
