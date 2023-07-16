package wood

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComponentLevel(t *testing.T) {
	Init(InfoLevel)
	PrefixLevel("a.b.c.d", DebugLevel)

	currentCanonical = "a.b.c.d"
	require.True(t, ignored(InfoLevel))

	currentCanonical = "d"
	ComponentLevel("d", InfoLevel)
	require.False(t, ignored(InfoLevel))
}

func Test_ignored(t *testing.T) {
	PrefixLevel("a", ErrorLevel)
	PrefixLevel("a.b", WarnLevel)
	PrefixLevel("a.b.c", InfoLevel)
	PrefixLevel("a.b.c.d", DebugLevel)

	tests := []struct {
		name   string
		id     string
		action Level
		want   bool
	}{
		// show everything
		{"DebugLevel1", "a", DebugLevel, false},
		{"DebugLevel2", "a.b", DebugLevel, false},
		{"DebugLevel3", "a.b.c", DebugLevel, false},
		{"DebugLevel4", "a.b.c.d", DebugLevel, false},

		// a.b.c.d is below requested
		{"InfoLevel1", "a", InfoLevel, false},
		{"InfoLevel2", "a.b", InfoLevel, false},
		{"InfoLevel3", "a.b.c", InfoLevel, false},
		{"InfoLevel4", "a.b.c.d", InfoLevel, true},

		// a.b.c/d are below requested
		{"WarnLevel1", "a", WarnLevel, false},
		{"WarnLevel2", "a.b", WarnLevel, false},
		{"WarnLevel3", "a.b.c", WarnLevel, true},
		{"WarnLevel4", "a.b.c.d", WarnLevel, true},

		// a.b/c/d are below requested
		{"ErrorLevel1", "a", ErrorLevel, false},
		{"ErrorLevel2", "a.b", ErrorLevel, true},
		{"ErrorLevel3", "a.b.c", ErrorLevel, true},
		{"ErrorLevel4", "a.b.c.d", ErrorLevel, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentCanonical = tt.id
			if got := ignored(tt.action); got != tt.want {
				t.Errorf("ignored() = %v, want %v", got, tt.want)
			}
			currentCanonical = ""
		})
	}
}
