package wood

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	// label Text shown
	label string
	// buff cache of spaces in front of the label; computed from stack
	buff string

	// stack of labels
	stack []string
	// indent custom indention for a particular prefix; resets when the label changes
	indent int

	// prefixLevel Sets the logging level for specific prefixes
	prefixLevel map[string]Level

	// std is the name of the standard logger in stdlib `log`
	std = logrus.New()
)

const yellow = 33

const pad = 35

func decorate(args ...interface{}) []any {
	c := make([]any, 0, len(args)+1)

	f, a := createLabel()
	if f != "" {
		c = append(c, fmt.Sprintf(f, a...))
	}
	c = append(c, fmt.Sprint(args...))
	return c
}

func decorateF(level Level, args []interface{}, fn func(format string, args []any)) {
	if ignored(level) {
		return
	}

	formats := make([]string, 0, len(args)+1)
	arguments := make([]any, 0, len(args)+1)

	f, a := createLabel()
	if f != "" {
		// add extra space to align with decorate. log adds space between the []any elements
		formats = append(formats, f+" ")
		arguments = append(arguments, a...)
	}

	formats = append(formats, args[0].(string))
	arguments = append(arguments, args[1:]...)

	fn(strings.Join(formats, ""), arguments)
}

func Init(level Level) {
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
	std.SetLevel(logrus.Level(level))
	stack = make([]string, 0, 5)
	prefixLevel = make(map[string]Level)
}
