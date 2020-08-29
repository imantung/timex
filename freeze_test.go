package timex_test

import (
	"fmt"
	"time"

	"github.com/imantung/timex"
)

func ExampleFreeze() {
	timex.Freeze(func() time.Time {
		return time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	})
	defer timex.Unfreeze()

	fmt.Println(timex.Now().Format(time.RFC3339))

	// Output: 2009-11-10T23:00:00Z
}

func ExampleFreezeByTime() {
	timex.FreezeByTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	defer timex.Unfreeze()

	fmt.Println(timex.Now().Format(time.RFC3339))

	// Output: 2009-11-10T23:00:00Z
}

func ExampleFreezeByRFC3339() {
	timex.FreezeByRFC3339("2020-08-29T14:21:00Z")
	defer timex.Unfreeze()

	fmt.Println(timex.Now().Format(time.RFC3339))

	// Output: 2020-08-29T14:21:00Z
}
