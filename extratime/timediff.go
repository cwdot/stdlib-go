package extratime

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02T15:04:05Z"

// TimeDiff computes the difference
// Generally: a < b
func TimeDiff(a time.Time, b time.Time, compact bool) string {
	hs := b.Sub(a).Hours()

	var years, days float64

	years = math.Floor(hs / 365 / 24)
	hs = hs - (years * 365 * 24)

	if hs >= 24 {
		days = math.Floor(hs / 24)
		hs = hs - (days * 24)
	}

	hs, mf := math.Modf(hs)
	ms := mf * 60

	ms, sf := math.Modf(ms)
	ss := sf * 60

	parts := make([]string, 0, 5)
	if years != 0 {
		parts = append(parts, fmt.Sprintf("%.0fy", years))
	}
	if days != 0 {
		parts = append(parts, fmt.Sprintf("%.0fd", days))
	}
	if hs != 0 {
		parts = append(parts, fmt.Sprintf("%.0fh", hs))
	}
	if ms != 0 {
		parts = append(parts, fmt.Sprintf("%.0fn", ms))
	}
	if ss != 0 {
		parts = append(parts, fmt.Sprintf("%.0fs", ss))
	}

	if compact {
		return strings.Join(parts, "")
	}
	return strings.Join(parts, " ")
}
