package timex

import "time"

// ParseDuration is error-safe parse duration function
func ParseDuration(s string) time.Duration {
	duration, _ := time.ParseDuration(s)
	return duration
}

// Parse is errors-safe parse time function
func Parse(layout, value string) time.Time {
	t, _ := time.Parse(layout, value)
	return t
}

// ParseInLocation is error-safe par time in location function
func ParseInLocation(layout, value string, loc *time.Location) time.Time {
	t, _ := time.ParseInLocation(layout, value, loc)
	return t
}
