# timex

Time extended package

## Freeze Time 

Replace `time.Now` with `timex.Now`

Freeze by function:
```go
timex.Freeze(func() time.Time {
    return time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
})
defer timex.Unfreeze()

fmt.Println(timex.Now().Format(time.RFC3339))

// Output: 2009-11-10T23:00:00Z
```

Freeze by specific time:
```go
timex.FreezeByTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
defer timex.Unfreeze()

fmt.Println(timex.Now().Format(time.RFC3339))

// Output: 2009-11-10T23:00:00Z
```

Freeze by specific rfc3339 formatted string:
```go
timex.FreezeByRFC3339("2020-08-29T14:21:00Z")
defer timex.Unfreeze()

fmt.Println(timex.Now().Format(time.RFC3339))

// Output: 2020-08-29T14:21:00Z
```

## Error-Safe Parse Functions

Standard time parse function without error

Use `timex.ParseDuration` to parses a duration string:
```go
fmt.Printf("%d\n", timex.ParseDuration("1h"))

// Output: 3600000000000
```

Use `timex.Parse` to parses a formatted string and returns the time value it represents:
```go
const longForm = "Jan 2, 2006 at 3:04pm (MST)"
fmt.Println(timex.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)"))

// Output: 2013-02-03 19:54:00 +0000 PST
```

Use `timex.ParseInLocation` to parsed a formatted string using location:
```go
loc, _ := time.LoadLocation("Europe/Berlin")

// This will look for the name CEST in the Europe/Berlin time zone.
const longForm = "Jan 2, 2006 at 3:04pm (MST)"
fmt.Println(timex.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc))

// Output: 2012-07-09 05:02:00 +0200 CEST
```