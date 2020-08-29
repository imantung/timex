package timex_test

import (
	"fmt"
	"time"

	"github.com/imantung/timex"
)

func ExampleParseDuration() {
	fmt.Printf("%d\n", timex.ParseDuration("1h"))

	// Output: 3600000000000
}

func ExampleParse() {
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	fmt.Println(timex.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)"))

	// Output: 2013-02-03 19:54:00 +0000 PST
}

func ExampleParseInLocation() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	// This will look for the name CEST in the Europe/Berlin time zone.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	fmt.Println(timex.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc))

	// Output: 2012-07-09 05:02:00 +0200 CEST
}
