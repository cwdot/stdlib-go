package wood

import (
	"fmt"
	"strings"

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

	// level for any member of the stack
	componentLevel map[string]Level

	// prefixLevel Sets the logging level for specific prefixes
	prefixes trie.Trie[Level]

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
	stack = make([]*logStack, 0, 5)
	componentLevel = make(map[string]Level)
	currentDisplay = ""
	currentCanonical = ""
	displayWhitespace = ""
}
