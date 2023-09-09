package wood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogf(t *testing.T) {
	th := setup(InfoLevel, InfoLevel)

	Logf(InfoLevel, "%s", A)
	Push("b")
	Logf(InfoLevel, "%s", B)

	entries := th.Pop()
	assert.Len(t, entries, 2)
	th.Assert(t, entries, 0, InfoLevel, "", A)
	th.Assert(t, entries, 1, InfoLevel, "b", B)
}
