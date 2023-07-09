package wood

import (
	"strings"

	"golang.org/x/exp/slices"
)

func Push(name string) func() {
	ls := splitName(name)
	//currentDisplay = ls.Display
	stack = append(stack, ls)
	indent = 0
	computeLabel()

	return Pop
}

func Pop() {
	indent = 0
	l := len(stack)
	if len(stack) == 0 {
		//currentDisplay = ""
		computeLabel()
		return
	}

	stack = slices.Delete(stack, l-1, l)

	computeLabel()
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
	if currentDisplay == "" {
		return "", []any{}
	}

	var trimmed string
	var whitespace string
	m := len(currentDisplay) + len(displayWhitespace)
	if m < pad {
		trimmed = currentDisplay
		m = pad - m + indent
		if m > 0 {
			whitespace = strings.Repeat(" ", m)
		}
	} else {
		trimmed = currentDisplay[0 : pad-len(displayWhitespace)]
		whitespace = ">"
	}

	format := "%s\x1b[%dm%s\x1b[0m%s"
	arguments := []any{displayWhitespace, yellow, trimmed, whitespace}
	return format, arguments
}

func computeLabel() {
	l := len(stack)
	if l < 1 {
		currentDisplay = ""
		currentId = ""
		displayWhitespace = ""
	} else {
		currentDisplay = stack[l-1].Display
		currentId = stack[l-1].Id
		displayWhitespace = strings.Repeat(" ", len(stack)-1)
	}
}
