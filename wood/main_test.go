package wood

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefixClassic(t *testing.T) {
	Init(InfoLevel)

	var expectedLevel int
	add := func(p string) {
		expectedLevel++
		Push(p)
		Info("Add: ", expectedLevel, "  ", fmt.Sprintf("buff=`%s`", displayWhitespace))
		require.Equal(t, expectedLevel, len(displayWhitespace))
	}
	pop := func() {
		expectedLevel--
		Pop()
		Info("Pop: ", expectedLevel, "  ", fmt.Sprintf("buff=`%s`", displayWhitespace))
		require.Equal(t, expectedLevel, len(displayWhitespace))
	}

	add("a")
	add("b")
	add("c")
	pop()

	add("d")
	add("e")
	pop()

	add("f")
}

func TestPrefixFormatted(t *testing.T) {
	Init(InfoLevel)

	var expectedLevel int
	add := func(p string) {
		expectedLevel++
		Push(p)
		Infof("Add: %d  %s", expectedLevel, fmt.Sprintf("buff=`%s`", displayWhitespace))
		require.Equal(t, expectedLevel, len(displayWhitespace))
	}
	pop := func() {
		expectedLevel--
		Pop()
		Infof("Pop: %d  %s", expectedLevel, fmt.Sprintf("buff=`%s`", displayWhitespace))
		require.Equal(t, expectedLevel, len(displayWhitespace))
	}

	add("a")
	add("b")
	add("c")
	pop()

	add("d")
	add("e")
	pop()

	add("f")
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
