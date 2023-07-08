package wood

import (
	"strings"

	"golang.org/x/exp/slices"
)

func Prefix(newPrefix string) {
	label = newPrefix
	stack = append(stack, label)
	indent = 0
	refreshLabel()
}

func Reset() {
	indent = 0
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

func Increment() {
	indent += 2
}

func Decrement() {
	indent -= 2
	if indent < 0 {
		indent = 0
	}
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
		m = pad - m + indent
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

func refreshLabel() {
	if len(stack) < 1 {
		buff = ""
	} else {
		buff = strings.Repeat(" ", len(stack)-1)
	}
}

func ignored(action Level) bool {
	if label == "" {
		return false
	}
	val, ok := prefixLevel[label]
	if !ok {
		return false
	}
	return val < action
}

func PrefixLevel(label string, level Level) {
	prefixLevel[label] = level
}
