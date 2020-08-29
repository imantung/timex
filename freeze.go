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

// FreezeTime freeze with specific time
func FreezeTime(t time.Time) {
	Freeze(func() time.Time {
		return t
	})
}

// FreezeRFC3339 freeze with specific RFC3339 time string
func FreezeRFC3339(s string) {
	Freeze(func() time.Time {
		t, _ := time.Parse(time.RFC3339, s)
		return t
	})
}

// Unfreeze restore `timex.Now()`
func Unfreeze() {
	Now = time.Now
}
