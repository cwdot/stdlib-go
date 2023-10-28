package color

import (
	"fmt"
	"testing"
)

func TestRainbowMarker(t *testing.T) {
	r := NewRainbowMarker()
	tests := []struct {
		name  string
		value string
		color Color
	}{
		{"red", "hello", Red},
		{"green", "hello", Green},
		{"yellow", "hello", Yellow},
		{"blue", "hello", Blue},
		{"magenta", "hello", Magenta},
		{"cyan", "hello", Cyan},
		{"red-loop", "hello", Red},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := r.Mark(tt.value)
			want := fmt.Sprintf("%s%s%s", tt.color, tt.value, Reset)
			if got != want {
				t.Errorf("Mark() = %v, want %v", got, want)
			}
		})
	}
}
