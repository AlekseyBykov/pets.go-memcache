package utils

import "time"

const TimeFormatMinutesSeconds = "04:05"

func FormatDuration(d time.Duration, format string) string {
	return time.Unix(0, 0).UTC().Add(d).Format(format)
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func GetExpirationTime(duration time.Duration) time.Time {
	return GetCurrentTime().Add(duration)
}
