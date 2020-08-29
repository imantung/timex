package timex

import "time"

// Now returns the current local time.
var Now DateFn = time.Now

type (
	// DateFn date function
	DateFn func() time.Time
)

// Freeze to override `timex.Now()` function
func Freeze(fn DateFn) {
	Now = fn
}

// FreezeByTime freeze by specific time
func FreezeByTime(t time.Time) {
	Freeze(func() time.Time {
		return t
	})
}

// FreezeByRFC3339 freeze by specific RFC3339 formatted string
func FreezeByRFC3339(s string) {
	Freeze(func() time.Time {
		t, _ := time.Parse(time.RFC3339, s)
		return t
	})
}

// Unfreeze restore `timex.Now()`
func Unfreeze() {
	Now = time.Now
}
