package wood

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

const expectedF = "%s\x1b[%dm%s\x1b[0m%s test %s"

func TestSinglePrefix(t *testing.T) {
	Init()

	var f string
	var args []any

	Push("a")
	f, args = TLogF("test %s", "x")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"", 33, "a", "                                 ", "x"}, args)

	Push("b")
	f, args = TLogF("test %s", "y")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{" ", 33, "b", "                                ", "y"}, args)

	Push("c")
	f, args = TLogF("test %s", "z")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"  ", 33, "c", "                               ", "z"}, args)

	Pop()
	f, args = TLogF("test %s", "decremented")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{" ", 33, "b", "                                ", "decremented"}, args)

	Pop()

	f, args = TLogF("test %s", "decremented")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"", 33, "a", "                                 ", "decremented"}, args)
}

func TestMultiplePrefixes(t *testing.T) {
	Init()

	var f string
	var args []any

	f, args = TLogF("test %s", "x")
	require.Equal(t, "test %s", f)
	require.Equal(t, []any{"x"}, args)

	Push("a")
	f, args = TLogF("test %s", "x")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"", 33, "a", "                                 ", "x"}, args)

	Push("a", "b")
	f, args = TLogF("test %s", "y")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"  ", 33, "b", "                               ", "y"}, args)

	Push("a", "b", "c")
	f, args = TLogF("test %s", "z")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"     ", 33, "c", "                            ", "z"}, args)

	Pop()
	f, args = TLogF("test %s", "decremented")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"    ", 33, "b", "                             ", "decremented"}, args)

	Pop()

	f, args = TLogF("test %s", "decremented")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"   ", 33, "a", "                              ", "decremented"}, args)
}

func TestLogPrefixPeriods(t *testing.T) {
	Init()

	var f string
	var args []any

	Push("a.b")
	f, args = TLogF("test %s", "y")
	require.Equal(t, expectedF, f)
	require.Equal(t, []any{"", 33, "a.b", "                               ", "y"}, args)
}

func Test_decorate(t *testing.T) {
	Init()

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

func WithHarness(th *TestHarness) func(*Opts) {
	return func(opts *Opts) {
		opts.output = th.SB
		opts.formatter = &logrus.JSONFormatter{}
	}
}
