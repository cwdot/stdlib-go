package wood

import (
	"fmt"
	"strings"

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

	f, a := createLabel()
	if f != "" {
		c = append(c, fmt.Sprintf(f, a...))
	}
	c = append(c, fmt.Sprint(args...))
	return c
}

func decorateF(level logrus.Level, args []interface{}, fn func(format string, args []any)) {
	if ignored(level) {
		return
	}

	formats := make([]string, 0, len(args)+1)
	arguments := make([]any, 0, len(args)+1)

	f, a := createLabel()
	if f != "" {
		formats = append(formats, f)
		arguments = append(arguments, a...)
	}

	formats = append(formats, args[0].(string))
	arguments = append(arguments, args[1:]...)

	fn(strings.Join(formats, ""), arguments)
}

func createLabel() (string, []any) {
	if label == "" {
		return "", []any{}
	}

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

	format := "%s\x1b[%dm%s\x1b[0m%s"
	arguments := []any{buff, yellow, trimmed, whitespace}
	return format, arguments
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
