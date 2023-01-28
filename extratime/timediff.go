package extratime

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02T15:04:05Z"

type TimeDiffOpts struct {
	Compact       bool
	EpochRounding bool
}

// TimeDiff computes the difference
// Generally: a < b
func TimeDiff(a time.Time, b time.Time, fn ...func(*TimeDiffOpts)) string {
	opts := &TimeDiffOpts{
		Compact:       false,
		EpochRounding: false,
	}
	for _, f := range fn {
		f(opts)
	}

	hours := b.Sub(a).Hours()

	var years, days float64

	years = math.Floor(hours / 365 / 24)
	hours = hours - (years * 365 * 24)

	if hours >= 24 {
		days = math.Floor(hours / 24)
		hours = hours - (days * 24)
	}

	hours, mf := math.Modf(hours)
	mins := mf * 60

	mins, sf := math.Modf(mins)
	secs := sf * 60

	parts := make([]string, 0, 5)
	if years != 0 {
		parts = append(parts, fmt.Sprintf("%.0fy", years))
	}
	if days != 0 {
		parts = append(parts, fmt.Sprintf("%.0fd", days))
	}
	if hours != 0 {
		parts = append(parts, fmt.Sprintf("%.0fh", hours))
	}
	if mins != 0 {
		parts = append(parts, fmt.Sprintf("%.0fm", mins))
	}
	if secs != 0 {
		parts = append(parts, fmt.Sprintf("%.0fs", secs))
	}

	// trim to largest epochs
	if opts.EpochRounding && len(parts) > 3 {
		parts = parts[0:2]
	}
	if opts.Compact {
		return strings.Join(parts, "")
	}
	return strings.Join(parts, " ")
}
