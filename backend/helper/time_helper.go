package helper

import "time"

// NormalizeToNoonUTC takes a time of a day and converts it to 12pm of UTC for the ease of comparison
func NormalizeToNoonUTC(t time.Time) time.Time {
	// This is utterly terrible, but it seems to work
	return time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, time.UTC)
}
