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
		Prefix(p)
		Info("Add: ", expectedLevel, "  ", fmt.Sprintf("buff=`%s`", buff))
		require.Equal(t, expectedLevel, len(buff))
	}
	pop := func() {
		expectedLevel--
		Reset()
		Info("Pop: ", expectedLevel, "  ", fmt.Sprintf("buff=`%s`", buff))
		require.Equal(t, expectedLevel, len(buff))
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
		Prefix(p)
		Infof("Add: %d  %s", expectedLevel, fmt.Sprintf("buff=`%s`", buff))
		require.Equal(t, expectedLevel, len(buff))
	}
	pop := func() {
		expectedLevel--
		Reset()
		Infof("Pop: %d  %s", expectedLevel, fmt.Sprintf("buff=`%s`", buff))
		require.Equal(t, expectedLevel, len(buff))
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
	Prefix("Reset1")
	Printf("test %s", "1")
	Println("test", "2")
	Increment()
	Increment()
	Println("test", "f")
	Increment()
	Println("test", "increment")
	Decrement()
	Println("test", "decremented")
	Reset()

	Prefix("Reset2")
	Printf("test %s", "f")
	Println("reset", "f")
}

func TestPrint(t *testing.T) {
	Init(InfoLevel)

	Printf("test %s", "f")
	Println("test", "f")
}
