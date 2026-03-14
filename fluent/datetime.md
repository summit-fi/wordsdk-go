
## DateTime in Fluent messages

DateTime formatting is built into the Fluent resolver (`fluent/bundle.go` + `fluent/datetime_func.go`).

### Main function: `DATETIME`

Use in FTL:

```ftl
order-date = Ordered on { DATETIME($date, pattern: "yMMMMd") }
clock = Time: { DATETIME($date, pattern: "Hm") }
weekday-time = { DATETIME($date, pattern: "Ehms") }
```

Then pass `time.Time` from Go:

```go
msg := sdk.TA("es_CO", "order-date", map[string]any{
	"date": time.Date(2025, time.March, 8, 14, 7, 9, 0, time.UTC),
})
```

`time.Time` is converted to Fluent `DateTimeValue` by `resolveValue(...)` in `fluent/bundle.go`.

### Locale behavior

Date formatting locale is taken from bundle language (for example `es_CO`, `en_US`) via `cldr.Language.BCP47()` in `fluent/cldr/language.go`.

### Pattern/skeleton:

- `yMMMMd` -> full date
- `Hms` -> time with seconds (24h)
- `yMd` -> short date
- `Hm` -> short time
- `yMMMM` -> year + month
- `Ehms` -> weekday + time

Example expected output in `es_CO`:
- `"8 de marzo de 2025"` for `yMMMMd`
- `"14:07:09"` for `Hms`
- `"lun, 9:45:05 p.m."` for `Ehms`

### Error behavior

If pattern is invalid or missing:
- missing pattern -> `"error: missing pattern"`
- unsupported symbol -> `"error: unsupported skeleton symbol ..."`
- invalid date variable type -> `"error: invalid datetime value"`

### Named shortcut datetime functions

Besides `DATETIME`, these function names are pre-registered in `fluent/bundle.go`:
- `MMMMEEEED`, `YMMMMEEEED`
- `MMMD`, `YMMMD`, `YMMMED`
- `JM`, `JMS`, `HHMM`
- `YMD`, `MD`, `YM`, `Y`, `E`, `EEEEE`
- `LLL`, `MMM`, `YMMMM`, `MMMMD`, `YMMMMD`, `YMMM`
- `EEE_D` (maps to `EEEd` pattern)

These can be used directly in FTL if needed, but `DATETIME($date, pattern: "...")` is usually clearer and more flexible.


## End-to-end example with DateTime

FTL (`en_US.ftl`):

```ftl
welcome = Hello, { $name }!
report-date = Report date: { DATETIME($date, pattern: "yMMMMd") }
report-time = Report time: { DATETIME($date, pattern: "Hms") }
next-week = Next week: { YMMMMD($date) }
submission-deadline = Submission deadline: { YMD($date) }
```

Go:

```go
msg1 := sdk.TA("en_US", "welcome", map[string]any{
	"name": "Olivia",
})

msg2 := sdk.TA("en_US", "report-date", map[string]any{
	"date": time.Now(),
})

msg3 := sdk.TA("en_US", "report-time", map[string]any{
	"date": time.Now(),
})
msg4 := sdk.TA("en_US", "next-week", map[string]any{
    "date": time.Now().AddDate(0, 0, 7),
})
msg5 := sdk.TA("en_US", "submission-deadline", map[string]any{
    "date": time.Now().AddDate(0, 0, 3),
})
```