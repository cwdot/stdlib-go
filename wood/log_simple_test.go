package wood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	th := setup(InfoLevel, InfoLevel)

	Log(InfoLevel, A)
	Push("b")
	Log(InfoLevel, B)

	entries := th.Pop()
	assert.Len(t, entries, 2)
	th.Assert(t, entries, 0, InfoLevel, "", A)
	th.Assert(t, entries, 1, InfoLevel, "b", B)
}
