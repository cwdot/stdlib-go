package wood

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefixFormatted(t *testing.T) {
	Init(InfoLevel)

	var expectedLevel int
	add := func(p string) {
		expectedLevel++
		Push(p)
	}
	pop := func() {
		if expectedLevel > 0 {
			expectedLevel--
		}
		Pop()
	}
	check := func(verb string, canonical string, cd string, dws string) {
		testLevel := expectedLevel - 1
		if testLevel < 0 {
			testLevel = 0
		}
		Infof("%s: %d cid=`%s` cdisplay=`%s` dws=`%s`", verb, expectedLevel, currentCanonical, currentDisplay, displayWhitespace)
		Info(verb, expectedLevel, currentCanonical, currentDisplay, displayWhitespace)
		require.Equal(t, testLevel, len(displayWhitespace))
		require.Equal(t, canonical, currentCanonical)
		require.Equal(t, cd, currentDisplay)
		require.Equal(t, dws, displayWhitespace)
	}

	add("a")
	check("Add", "a", "a", "")

	add("b")
	check("Add", "a.b", "b", " ")

	add("c")
	check("Add", "a.b.c", "c", "  ")

	pop()
	check("Pop", "a.b", "b", " ")

	pop()
	check("Pop", "a", "a", "")

	pop()
	check("Pop", "", "", "")

	pop()
	check("Pop", "", "", "")
}

func TestIndent(t *testing.T) {
	Init(InfoLevel)

	Increment()
	Printf("test %s", "f")
	Increment()
	Println("test", "f")

	// should reset the text indention to 0
	Push("Reset1")
	Printf("test %s", "1")
	Println("test", "2")
	Increment()
	Increment()
	Println("test", "f")
	Increment()
	Println("test", "increment")
	Decrement()
	Println("test", "decremented")
	Pop()

	Push("Reset2")
	Printf("test %s", "f")
	Println("reset", "f")
}

func TestPrint(t *testing.T) {
	Init(InfoLevel)

	Printf("test %s", "f")
	Println("test", "f")
}
