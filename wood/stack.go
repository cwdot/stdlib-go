package wood

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type logStack struct {
	// complete ID
	CanonicalID string
	// First N-1
	FrontID string
	// Last 1
	LastID string
	// Text being shown
	Display string
}

func (ls *logStack) String() string {
	return fmt.Sprintf("{CanonicalID=`%s` LastID=`%s` Display=`%s`}", ls.CanonicalID, ls.LastID, ls.Display)
}

func Push(name string) func() {
	for _, part := range strings.Split(name, ".") {
		can := computeCanonical(name)
		ls := &logStack{
			CanonicalID: can,
			LastID:      part,
			Display:     part,
		}
		stack = append(stack, ls)
	}

	//currentDisplay = ls.Display
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
		currentCanonical = ""
		displayWhitespace = ""
	} else {
		currentDisplay = stack[l-1].Display
		currentCanonical = stack[l-1].CanonicalID
		displayWhitespace = strings.Repeat(" ", len(stack)-1)
	}
}

func computeCanonical(extra string) string {
	if len(stack) == 0 && extra == "" {
		return ""
	}

	var b strings.Builder
	b.Grow(100)
	if len(stack) > 0 {
		b.WriteString(stack[0].LastID)
		for _, s := range stack[1:] {
			b.WriteString(".")
			b.WriteString(s.LastID)
		}
		if extra != "" {
			b.WriteString(".")
			b.WriteString(extra)
		}
	} else if extra != "" {
		b.WriteString(extra)
	}
	return b.String()
}
