package wood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixLevel(t *testing.T) {
	setup := func(labels ...string) *TestHarness {
		th := New()
		Init(WithHarness(th))
		for _, label := range labels {
			PrefixLevel(DebugLevel, label)
		}
		return th
	}
	run := func(labels ...string) {
		Push("a")
		Info("test-a")
		Push("b")
		Info("test-b")
		Push("c")
		Info("test-c")
		Push("d")
		Info("test-d")
	}

	t.Run("NoRestrictions", func(t *testing.T) {
		th := setup()
		run()
		entries := th.Pop()
		assert.Len(t, entries, 4)
		th.Assert(t, entries, 0, InfoLevel, "a", "test-a")
		th.Assert(t, entries, 1, InfoLevel, "b", "test-b")
		th.Assert(t, entries, 2, InfoLevel, "c", "test-c")
		th.Assert(t, entries, 3, InfoLevel, "d", "test-d")
	})
	t.Run("RestrictA", func(t *testing.T) {
		th := setup("a")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 0)
	})
	t.Run("RestrictB", func(t *testing.T) {
		th := setup("a.b")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 1)
		th.Assert(t, entries, 0, InfoLevel, "a", "test-a")
	})
	t.Run("RestrictC", func(t *testing.T) {
		th := setup("a.b.c")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 2)
		th.Assert(t, entries, 0, InfoLevel, "a", "test-a")
		th.Assert(t, entries, 1, InfoLevel, "b", "test-b")
	})
	t.Run("RestrictD", func(t *testing.T) {
		th := setup("a.b.c.d")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 3)
		th.Assert(t, entries, 0, InfoLevel, "a", "test-a")
		th.Assert(t, entries, 1, InfoLevel, "b", "test-b")
		th.Assert(t, entries, 2, InfoLevel, "c", "test-c")
	})
	t.Run("RestrictA_BoostC", func(t *testing.T) {
		th := setup("a")
		PrefixLevel(InfoLevel, "a.b.c")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 2)
		th.Assert(t, entries, 0, InfoLevel, "c", "test-c")
		th.Assert(t, entries, 1, InfoLevel, "d", "test-d")
	})
	t.Run("RestrictA_BoostD", func(t *testing.T) {
		th := setup("a")
		PrefixLevel(InfoLevel, "a.b.c.d")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 1)
		th.Assert(t, entries, 0, InfoLevel, "d", "test-d")
	})
}

func TestComponentLevel(t *testing.T) {
	setup := func(labels ...string) *TestHarness {
		th := New()
		Init(WithHarness(th))
		for _, label := range labels {
			ComponentLevel(DebugLevel, label)
		}
		return th
	}
	run := func(labels ...string) {
		Push("a")
		Info("test-a")
		Push("b")
		Info("test-b")
		Push("c")
		Info("test-c")
		Push("d")
		Info("test-d")
	}

	t.Run("NoRestrictions", func(t *testing.T) {
		th := setup()
		run()
		entries := th.Pop()
		assert.Len(t, entries, 4)
		th.Assert(t, entries, 0, InfoLevel, "a", "test-a")
		th.Assert(t, entries, 1, InfoLevel, "b", "test-b")
		th.Assert(t, entries, 2, InfoLevel, "c", "test-c")
		th.Assert(t, entries, 3, InfoLevel, "d", "test-d")
	})
	t.Run("RestrictA", func(t *testing.T) {
		th := setup("a")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 0)
	})
	t.Run("RestrictB", func(t *testing.T) {
		th := setup("b")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 1)
		th.Assert(t, entries, 0, InfoLevel, "a", "test-a")
	})
	t.Run("RestrictC", func(t *testing.T) {
		th := setup("c")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 2)
		th.Assert(t, entries, 0, InfoLevel, "a", "test-a")
		th.Assert(t, entries, 1, InfoLevel, "b", "test-b")
	})
	t.Run("RestrictD", func(t *testing.T) {
		th := setup("d")
		run()
		entries := th.Pop()
		assert.Len(t, entries, 3)
		th.Assert(t, entries, 0, InfoLevel, "a", "test-a")
		th.Assert(t, entries, 1, InfoLevel, "b", "test-b")
		th.Assert(t, entries, 2, InfoLevel, "c", "test-c")
	})
}
