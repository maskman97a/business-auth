package time_utils

import (
	"time"
)

func ToString(time time.Time, pattern string) string {
	return time.Format(pattern)
}
func ToDate(timeStr string, pattern string) time.Time {
	t, _ := time.Parse(pattern, timeStr)
	return t
}
