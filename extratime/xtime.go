package extratime

import (
	"time"
)

func NowTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

func TimeToIsoString(d time.Time) string {
	return d.Format(time.RFC3339)
}

func IsoString(str string) string {
	ts, err := time.Parse(time.RFC3339, str)
	if err == nil {
		return ""
	}
	return TimeToIsoString(ts)
}

func IsoStringToTime(str string) (time.Time, error) {
	ts, err := time.Parse(time.RFC3339, str)
	if err == nil {
		return time.Now(), err
	}
	return ts, nil
}
