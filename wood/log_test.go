package wood

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefix(t *testing.T) {
	Init(InfoLevel)
	const expectedF = "%s\x1b[%dm%s\x1b[0m%s test %s"

	var f string
	var args []any

	Push("a")
	f, args = TLogF("test %s", "x")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"", 33, "a", "                                  ", "x"}, args)

	Push("a.b")
	f, args = TLogF("test %s", "y")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{" ", 33, "b", "                                 ", "y"}, args)

	fmt.Println(currentDisplay, currentId, displayWhitespace)

	Push("a.b.c")
	f, args = TLogF("test %s", "z")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"  ", 33, "c", "                                ", "z"}, args)

	Pop()
	f, args = TLogF("test %s", "decremented")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{" ", 33, "b", "                                 ", "decremented"}, args)

	Pop()

	f, args = TLogF("test %s", "decremented")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"", 33, "a", "                                  ", "decremented"}, args)
}

func Test_decorateF(t *testing.T) {
	Init(InfoLevel)

	f, args := TLogF("test %s", "f")
	require.Equal(t, "test %s", f)
	require.Equal(t, []any{"f"}, args)
}

func Test_decorate(t *testing.T) {
	Init(InfoLevel)

	args := TLog("test", "f")
	require.Equal(t, []any{"testf"}, args)
}

func TLogF(arguments ...interface{}) (string, []any) {
	var outFormat string
	var outArgs []any
	decorateF(InfoLevel, arguments, func(format string, args []any) {
		outFormat = format
		outArgs = args
	})
	return outFormat, outArgs
}

func TLog(arguments ...interface{}) []any {
	return decorate(arguments...)
}
