package color

import (
	"fmt"
	"testing"
)

func TestLutMarker(t *testing.T) {
	r := NewLutMarker()
	r.Set("hello", Red)
	r.Set("there", Green)
	tests := []struct {
		name  string
		value string
		color Color
	}{
		{"red", "hello", Red},
		{"green", "there", Green},
		{"default", "billy", Normal},
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
