# timex

Time extended package

## Freeze Time 

Replace `time.Now` with `timex.Now`

```go
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
```