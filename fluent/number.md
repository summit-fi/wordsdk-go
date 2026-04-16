## Numbers & Currencies in Fluent messages

Number formatting is built into the Fluent resolver (`fluent/bundle.go` + `fluent/number_func.go`).

The `NUMBER` function handles all numeric styles: **decimal**, **currency**, **percent**, and **ordinal**.
The locale is taken from the bundle language via `cldr.Language.GetNumberRules()` in `fluent/cldr/numbers.go`.

---

### Main function: `NUMBER`

```ftl
price        = Price: { NUMBER($amount, style: "currency") }
discount     = Discount: { NUMBER($rate, style: "percent") }
score        = Score: { NUMBER($value, minimumFractionDigits: 2) }
rank         = { NUMBER($pos, style: "ordinal") ->
   [one]   { $pos }st
   [two]   { $pos }nd
   [few]   { $pos }rd
  *[other] { $pos }th
}
```

Pass a numeric value from Go:

```go
msg := sdk.TA("en_US", "price", map[string]any{
    "amount": 1234.56,
})
```

Any Go numeric type (`int`, `float32`, `float64`, `uint`, etc.) is accepted.
`resolveValue(...)` in `fluent/bundle.go` converts them to `NumberValue` internally.

---

### Styles

#### `decimal` (default)

Formats a plain number using locale grouping and decimal separators.
No `style` parameter needed — decimal is the default when `style` is omitted.

```ftl
emails = You have { NUMBER($count) } unread emails.
score  = Score: { NUMBER($value, minimumFractionDigits: 1) }
```

| Input | `en_US` | `en_EU` | `uk_UA` | `es_CO` |
|-------|---------|---------|---------|---------|
| `1000` | `1,000` | `1,000` | `1 000` | `1.000` |
| `1234567` | `1,234,567` | `1,234,567` | `1 234 567` | `1.234.567` |
| `1234.5` + `minimumFractionDigits: 2` | `1,234.50` | `1,234.50` | `1 234,50` | `1.234,50` |

---

#### `currency`

Formats a number as a monetary value using the locale's default currency and pattern.

```ftl
price-simple  = { NUMBER($amount, style: "currency") }
price-code    = { NUMBER($amount, style: "currency", currencyDisplay: "code") }
price-symbol  = { NUMBER($amount, style: "currency", currency: "UAH", currencySymbol: "₴", currencyDisplay: "symbol") }
price-pattern = { NUMBER($amount, style: "currency", currency: "UAH", currencySymbol: "₴", currencyDisplay: "symbol", pattern: "#,##0.00 ¤") }
```

Go:

```go
msg := sdk.TA("en_US", "price-simple", map[string]any{
    "amount": 1234.56,
})
// → "Price: $1,234.56"

msg = sdk.TA("uk_UA", "price-simple", map[string]any{
    "amount": 1234.56,
})
// → "1 234,56 ₴"
```

##### Named parameters for `currency`

| Parameter | Type | Description |
|---|---|---|
| `style` | `"currency"` | Enables currency formatting |
| `currency` | string | Override currency code (e.g. `"UAH"`, `"USD"`, `"COP"`) |
| `currencySymbol` | string | Override the displayed symbol (e.g. `"₴"`, `"$"`) |
| `currencyDisplay` | `"symbol"` / `"code"` | Display symbol (default) or ISO code |
| `pattern` | string | Custom CLDR pattern (e.g. `"#,##0.00 ¤"`) |
| `minimumFractionDigits` | int | Min decimal places |
| `maximumFractionDigits` | int | Max decimal places |

##### Default currency and pattern per locale

| Locale | Default currency | Pattern | Separator (dec / group) |
|--------|-----------------|---------|------------------------|
| `en_US` | USD `$` | `¤#,##0.00` (prefix) | `.` / `,` |
| `en_EU` | EUR `€` | `¤#,##0.00` (prefix) | `.` / `,` |
| `en_UA` | UAH `₴` | `#,##0.00 ¤` (suffix) | `,` / ` ` |
| `uk_UA` | UAH `₴` | `#,##0.00 ¤` (suffix) | `,` / ` ` |
| `ru_UA` | UAH `₴` | `#,##0.00 ¤` (suffix) | `,` / ` ` |
| `en_CO` | COP `$` | `¤ #,##0` (prefix) | `,` / `.` |
| `es_CO` | COP `$` | `¤ #,##0` (prefix) | `,` / `.` |

##### Supported currencies per locale

| Locale | USD | EUR | UAH | COP |
|--------|-----|-----|-----|-----|
| `en_US` | `$` | — | `₴` | `$` |
| `en_EU` | `$` | `€` | `₴` | — |
| `en_UA` | `$` | `€` | `₴` | — |
| `uk_UA` | `₴`* | `€` | `₴` | `$` |
| `ru_UA` | `₴`* | `€` | `₴` | `$` |
| `en_CO` | `$` | — | — | `$` |
| `es_CO` | `$` | — | — | `$` |

> \* `uk_UA`/`ru_UA` have USD symbol mapped to `₴` in CLDR data — override with `currencySymbol` if needed.

---

#### `percent`

Pass a decimal ratio (e.g. `0.15` for 15%). The formatter multiplies by 100 internally.

```ftl
tax-rate    = Tax: { NUMBER($rate, style: "percent") }
completion  = { NUMBER($done, style: "percent") } complete
```

Go:

```go
msg := sdk.TA("en_US", "tax-rate", map[string]any{
    "rate": 0.15,
})
// → "Tax: 15%"
```

| Input | `en_US` | `en_UA` / `uk_UA` / `ru_UA` | `es_CO` / `en_CO` |
|-------|---------|------------------------------|-------------------|
| `0.12` | `12%` | `12%` | `12%` |
| `1.0` | `100%` | `100%` | `100%` |
| `-0.12` | `-12%` | `-12%` | `-12%` |
| `-1234567.89` | `-123,456,789%` | `-123 456 789%` | `-123.456.789%` |

Custom pattern in FTL:

```ftl
discount = { NUMBER($rate, style: "percent", pattern: "#,##0.00%") }
```

---

#### `ordinal`

Returns the CLDR ordinal category (`"one"`, `"two"`, `"few"`, `"other"`) for use in select expressions.

```ftl
your-rank = { NUMBER($pos, style: "ordinal") ->
   [one]   { $pos }st
   [two]   { $pos }nd
   [few]   { $pos }rd
  *[other] { $pos }th
}
```

Go:

```go
msg := sdk.TA("en_US", "your-rank", map[string]any{
    "pos": 3,
})
// → "3rd"
```

| Input | `en_US` / `en_EU` / `en_UA` / `en_CO` | `uk_UA` | `ru_UA` / `es_CO` |
|-------|---------------------------------------|---------|-------------------|
| 1 | `one` | `other` | `other` |
| 2 | `two` | `other` | `other` |
| 3 | `few` | `few` | `other` |
| 11 | `other` | `other` | `other` |
| 21 | `one` | `other` | `other` |
| 23 | `few` | `few` | `other` |

---

### Fraction digit parameters

These work with `decimal` and `currency` styles.

| Parameter | FTL | Description |
|---|---|---|
| `minimumFractionDigits` | `NUMBER($v, minimumFractionDigits: 2)` | Pads trailing zeros up to N |
| `maximumFractionDigits` | `NUMBER($v, maximumFractionDigits: 1)` | Rounds/truncates to N decimal places |

Examples (`en_US`):

| Input | `minimumFractionDigits: 2` | `maximumFractionDigits: 1` |
|-------|---------------------------|---------------------------|
| `1234` | `1,234.00` | `1,234` |
| `1234.5` | `1,234.50` | `1,234.5` |
| `1234.5678` | `1,234.5678` | `1,234.6` |

---

### Custom number patterns

Both `currency` and `percent` (and `decimal`) accept a `pattern` override.
Pattern syntax follows CLDR conventions (`¤` = currency placeholder):

| Pattern | Meaning |
|---------|---------|
| `¤#,##0.00` | Currency symbol as prefix, 2 decimal places |
| `#,##0.00 ¤` | Currency symbol as suffix, 2 decimal places, space separator |
| `¤ #,##0` | Currency symbol prefix with space, no decimals |
| `#,##0%` | Percent with grouping |
| `#,##0.00%` | Percent with 2 decimal places |

FTL example:

```ftl
price-custom = { NUMBER($amount, style: "currency", currency: "UAH", currencySymbol: "₴", currencyDisplay: "symbol", pattern: "#,##0.00 ¤") }
```

---

### Error behavior

| Situation | Returned value |
|-----------|---------------|
| `minimumFractionDigits` < 0 | `"func NUMBER: minimum fraction digits cannot be negative -> ..."` |
| Invalid number passed | `"func NUMBER: invalid number cloneFormat -> ..."` |
| Unknown currency code | `"func NUMBER: invalid currency code -> ..."` |
| Invalid currency symbol | `"func NUMBER: invalid currency symbol -> ..."` |

---

### End-to-end example

FTL (`en_US.ftl`):

```ftl
invoice-total   = Total: { NUMBER($amount, style: "currency") }
invoice-tax     = Tax ({ NUMBER($rate, style: "percent") }): { NUMBER($taxAmount, style: "currency") }
invoice-items   = { NUMBER($count) } item(s)
invoice-rank    = Invoice #{ NUMBER($pos, style: "ordinal") ->
   [one]   { $pos }st
   [two]   { $pos }nd
   [few]   { $pos }rd
  *[other] { $pos }th
}
```

Go:

```go
total, _ := sdk.TA("en_US", "invoice-total", map[string]any{
    "amount": 1500.00,
})
// → "Total: $1,500.00"

tax, _ := sdk.TA("en_US", "invoice-tax", map[string]any{
    "rate":      0.21,
    "taxAmount": 315.00,
})
// → "Tax (21%): $315.00"

items, _ := sdk.TA("en_US", "invoice-items", map[string]any{
    "count": 3,
})
// → "3 item(s)"

rank, _ := sdk.TA("en_US", "invoice-rank", map[string]any{
    "pos": 23,
})
// → "Invoice #23rd"
```
