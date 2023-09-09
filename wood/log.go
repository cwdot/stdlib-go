package wood

import (
	"io"

	"github.com/porfirion/trie"
	"github.com/sirupsen/logrus"
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

type Opts struct {
	level     Level
	output    io.Writer
	formatter logrus.Formatter
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

	base := &Opts{level: InfoLevel}
	for _, opt := range opts {
		opt(base)
	}
	std.SetLevel(logrus.Level(base.level))
	if base.output != nil {
		std.SetOutput(base.output)
	}
	if base.formatter != nil {
		std.SetFormatter(base.formatter)
	}

	stack = make([]*logStack, 0, 5)
	currentDisplay = ""
	currentCanonical = ""
	displayWhitespace = ""

	prefixes = trie.Trie[Level]{}
	components = make(map[string]Level)
}
