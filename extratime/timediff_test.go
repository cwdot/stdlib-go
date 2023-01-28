package extratime

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeDiffFormatting(t *testing.T) {
	y1 := time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	y2 := time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
	y3 := time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	m1 := time.Date(2022, 2, 1, 1, 1, 1, 1, time.UTC)
	m2 := time.Date(2022, 3, 1, 1, 1, 1, 1, time.UTC)
	md := time.Date(2022, 3, 5, 1, 1, 1, 1, time.UTC)
	mdh := time.Date(2022, 3, 5, 10, 1, 1, 1, time.UTC)

	tests := []struct {
		name string
		a    time.Time
		b    time.Time
		want string
	}{
		{"Years (Leap Year)", y1, y2, "1y 1d"},
		{"Years", y2, y3, "1y"},
		{"Years,Month", y1, m1, "2y 32d"},
		{"Month diff", m1, m2, "28d"},
		{"Month,Day diff", m2, md, "4d"},
		{"Hour diff", md, mdh, "9h"},
		{"Month,Day,Hour", m1, mdh, "32d 9h"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeDiff(tt.a, tt.b); got != tt.want {
				t.Errorf("TimeDiff() = %v, want %v", got, tt.want)
			}

			compactVersion := strings.ReplaceAll(tt.want, " ", "")
			if got := TimeDiff(tt.a, tt.b, func(opts *TimeDiffOpts) {
				opts.Compact = true
			}); got != compactVersion {
				t.Errorf("TimeDiff() = %v, want %v", got, compactVersion)
			}
		})
	}
}

func TestTimeDiffEpoch(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		a := time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
		b := time.Date(2022, 3, 5, 10, 22, 4, 1, time.UTC)

		tdOff := TimeDiff(a, b)
		require.Equal(t, "1y 63d 9h 21m 3s", tdOff)

		tdOn := TimeDiff(a, b, func(opts *TimeDiffOpts) {
			opts.EpochRounding = true
		})
		require.Equal(t, "1y 63d", tdOn)
	})
	t.Run("FewerTokens", func(t *testing.T) {
		a := time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
		b := time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC)

		tdOn := TimeDiff(a, b, func(opts *TimeDiffOpts) {
			opts.EpochRounding = true
		})
		require.Equal(t, "2y", tdOn)

	})
}
