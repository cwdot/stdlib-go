package wood

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestPrefix(t *testing.T) {
	Init(logrus.InfoLevel)

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

func TestPrintf(t *testing.T) {
	Init(logrus.InfoLevel)

	Printf("test %s", "f")
}
