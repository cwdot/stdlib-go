package color

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPen(t *testing.T) {
	r := NewPen(Red, Green)

	t.Run("TernaryTrue", func(t *testing.T) {
		got := r.Ternary(true, "hello", "there")
		wanted := fmt.Sprintf("%s%s%s", Red, "hello", Reset)
		assert.Equal(t, wanted, got)
	})
	t.Run("TernaryFalse", func(t *testing.T) {
		got := r.Ternary(false, "hello", "there")
		wanted := fmt.Sprintf("%s%s%s", Green, "there", Reset)
		assert.Equal(t, wanted, got)
	})

	t.Run("MarkTrue", func(t *testing.T) {
		got := r.Mark(true, "value")
		wanted := fmt.Sprintf("%s%s%s", Red, "value", Reset)
		assert.Equal(t, wanted, got)
	})
	t.Run("MarkFalse", func(t *testing.T) {
		got := r.Mark(false, "value")
		wanted := fmt.Sprintf("%s%s%s", Green, "value", Reset)
		assert.Equal(t, wanted, got)
	})
}
