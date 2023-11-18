package wood

import (
	"fmt"
	"strings"
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

	switch t := args[0].(type) {
	case string:
		formats = append(formats, t)
	default:
		formats = append(formats, fmt.Sprintf("%v", t))
	}
	arguments = append(arguments, args[1:]...)

	fn(strings.Join(formats, ""), arguments)
}
