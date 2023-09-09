package wood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixLevel(t *testing.T) {
	run := func(labels ...string) {
		Push("a")
		Info(A)
		Push("b")
		Info(B)
		Push("c")
		Info(C)
		Push("d")
		Info(D)
	}

	t.Run("NoRestrictions", func(t *testing.T) {
		th := setup(InfoLevel, DebugLevel)
		run()
		entries := th.Pop()
		assert.Len(t, entries, 4)
		th.Assert(t, entries, 0, InfoLevel, "a", A)
		th.Assert(t, entries, 1, InfoLevel, "b", B)
		th.Assert(t, entries, 2, InfoLevel, "c", C)
		th.Assert(t, entries, 3, InfoLevel, "d", D)
	})
	t.Run("RestrictA", func(t *testing.T) {
		th := setup(InfoLevel, DebugLevel, "a")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 0)
	})
	t.Run("RestrictB", func(t *testing.T) {
		th := setup(InfoLevel, DebugLevel, "a.b")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 1)
		th.Assert(t, entries, 0, InfoLevel, "a", A)
	})
	t.Run("RestrictC", func(t *testing.T) {
		th := setup(InfoLevel, DebugLevel, "a.b.c")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 2)
		th.Assert(t, entries, 0, InfoLevel, "a", A)
		th.Assert(t, entries, 1, InfoLevel, "b", B)
	})
	t.Run("RestrictD", func(t *testing.T) {
		th := setup(InfoLevel, DebugLevel, "a.b.c.d")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 3)
		th.Assert(t, entries, 0, InfoLevel, "a", A)
		th.Assert(t, entries, 1, InfoLevel, "b", B)
		th.Assert(t, entries, 2, InfoLevel, "c", C)
	})
	t.Run("RestrictA_BoostC", func(t *testing.T) {
		th := setup(InfoLevel, DebugLevel, "a")
		PrefixLevel(InfoLevel, "a.b.c")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 2)
		th.Assert(t, entries, 0, InfoLevel, "c", C)
		th.Assert(t, entries, 1, InfoLevel, "d", D)
	})
	t.Run("RestrictA_BoostD", func(t *testing.T) {
		th := setup(InfoLevel, DebugLevel, "a")
		PrefixLevel(InfoLevel, "a.b.c.d")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 1)
		th.Assert(t, entries, 0, InfoLevel, "d", D)
	})
}

func TestComponentLevel(t *testing.T) {
	run := func(lvl Level, labels ...string) {
		Push("a")
		Log(lvl, A)
		Push("b")
		Log(lvl, B)
		Push("c")
		Log(lvl, C)
		Push("d")
		Log(lvl, D)
	}

	t.Run("WithLowerLevel", func(t *testing.T) {
		th := setup(InfoLevel, TraceLevel, "a")
		run(TraceLevel)
		entries := th.Pop()
		assert.Len(t, entries, 0)
	})

	t.Run("WithSameLevel", func(t *testing.T) {
		t.Run("RestrictA", func(t *testing.T) {
			th := setup(InfoLevel, InfoLevel, "a")
			run(InfoLevel)
			entries := th.Pop()
			assert.Len(t, entries, 4)
			th.Assert(t, entries, 0, InfoLevel, "a", A)
			th.Assert(t, entries, 1, InfoLevel, "b", B)
			th.Assert(t, entries, 2, InfoLevel, "c", C)
			th.Assert(t, entries, 3, InfoLevel, "d", D)
		})
		t.Run("RestrictD", func(t *testing.T) {
			th := setup(InfoLevel, InfoLevel, "d")
			run(InfoLevel)
			entries := th.Pop()
			assert.Len(t, entries, 4)
			th.Assert(t, entries, 0, InfoLevel, "a", A)
			th.Assert(t, entries, 1, InfoLevel, "b", B)
			th.Assert(t, entries, 2, InfoLevel, "c", C)
			th.Assert(t, entries, 3, InfoLevel, "d", D)
		})
	})

	t.Run("WithUpperLevel", func(t *testing.T) {
		t.Run("RestrictA", func(t *testing.T) {
			th := setup(InfoLevel, ErrorLevel, "a")
			run(InfoLevel)
			entries := th.Pop()
			assert.Len(t, entries, 0)
		})
		t.Run("RestrictD", func(t *testing.T) {
			th := setup(InfoLevel, ErrorLevel, "d")

			Push("a")
			Log(InfoLevel, A)
			Push("b")
			Log(InfoLevel, B)
			Push("c")
			Log(InfoLevel, C)
			Push("d")
			Log(ErrorLevel, D)

			entries := th.Pop()
			assert.Len(t, entries, 4)
			th.Assert(t, entries, 0, InfoLevel, "a", A)
			th.Assert(t, entries, 1, InfoLevel, "b", B)
			th.Assert(t, entries, 2, InfoLevel, "c", C)
			th.Assert(t, entries, 3, ErrorLevel, "d", D)
		})
	})
}
