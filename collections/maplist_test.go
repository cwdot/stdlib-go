package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapList(t *testing.T) {
	ml := NewMapList[string]()

	ml.Add("k", "a")
	ml.Add("k", "b")
	ml.Add("k", "c")
	ml.Add("x", "x")
	require.Equal(t, 2, ml.Size())

	k, ok := ml.GetList("k")
	require.True(t, ok)
	require.Equal(t, []string{"a", "b", "c"}, k)
	require.Equal(t, 3, ml.Count("k"))

	k, ok = ml.GetList("x")
	require.True(t, ok)
	require.Equal(t, []string{"x"}, k)
	require.Equal(t, 1, ml.Count("x"))

	c := ml.Copy()
	require.Equal(t, 2, len(c))
	assert.Equal(t, []string{"a", "b", "c"}, c["k"])
	assert.Equal(t, []string{"x"}, c["x"])

	require.Equal(t, 2, ml.Size())
	require.Equal(t, 4, ml.GrandSize())
}
