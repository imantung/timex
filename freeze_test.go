package timex_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/imantung/timex"
)

func ExampleNow() {
	type Model struct {
		Name      string
		CreatedAt time.Time
	}

	newModel := func(name string) *Model {
		return &Model{
			Name:      name,
			CreatedAt: timex.Now(),
		}
	}

	// freeze time
	timex.FreezeRFC3339("2020-08-29T14:21:00Z")
	defer timex.Unfreeze()

	b, _ := json.Marshal(newModel("some-name"))
	fmt.Println(string(b))

	// Output: {"Name":"some-name","CreatedAt":"2020-08-29T14:21:00Z"}
}

func TestFreezeTime(t *testing.T) {
	timex.FreezeTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	defer timex.Unfreeze()

	want := "2009-11-10T23:00:00Z"
	got := timex.Now().Format(time.RFC3339)

	if want != got {
		t.Fatalf("%s != %s", want, got)
	}
}

func TestFreezeRFC3339(t *testing.T) {
	timex.FreezeRFC3339("2020-08-29T14:21:00Z")
	defer timex.Unfreeze()

	want := "2020-08-29T14:21:00Z"
	got := timex.Now().Format(time.RFC3339)

	if want != got {
		t.Fatalf("%s != %s", want, got)
	}
}
