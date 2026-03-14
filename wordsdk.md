
### WordSDK Go

This package provides simple access to the Word translation service API, allowing you to fetch translations with [Fluent](https://projectfluent.org/) localization system for your application. 

## Install:

```bash
go get github.com/summit-fi/wordsdk-go
```

## Basic setup API usage:

```go

package main

import (
	"fmt"
	"log"

	word "github.com/summit-fi/wordsdk-go"
)

func main() {
	cfg := word.GetDefaultConfig("<API_KEY>")
	cfg.SaveStrategy = word.SaveStrategyImmediate

	sdk, err := word.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sdk.T("en_US", "hello_world"))
}
```

# Client

The client is the entry point of the SDK. It loads static translations, keeps an in-memory bundle cache, and runs periodic sync.

## API surface

Main constructor and config are implemented in:
- `client.go` (`Config`, `GetDefaultConfig`, `NewClient`)
- `public.go` (`SDK` interface)

### `Config`

```go
type Config struct {
    Source         source.Source
    UpdateInterval time.Duration
    MaxCacheSizeMB int
    SaveStrategy   SaveStrategy
}
```

### Create a client (remote source)

```go
cfg := word.GetDefaultConfig(apiKey) // uses remote source + defaults
cfg.SaveStrategy = word.SaveStrategyImmediate

sdk, err := word.NewClient(cfg)
if err != nil {
    // handle error
}
```

### Create a client (local FTL source)

```go
src := source.NewFtl()
_ = src.AddLocaleFiles(map[string]string{
    "en_US": "./locales/en_US.ftl",
    "es_CO": "./locales/es_CO.ftl",
})

sdk, err := word.NewClient(&word.Config{
    Source:         src,
    UpdateInterval: 30 * time.Second,
    MaxCacheSizeMB: 128,
    SaveStrategy:   word.SaveStrategyOnDemand,
})
if err != nil {
    // handle error
}
```

## Save strategies

```
const (
    SaveStrategyImmediate SaveStrategy = iota
    SaveStrategyOnDemand
)
```
### Save strategy behavior
#### SaveStrategyImmediate
- Saves to Source.SaveDynamic(...) immediately.
- Updates local bundle cache.

#### SaveStrategyOnDemand
- Updates local bundle cache first.
- Intended to persist via Flush().
### Flush()
```go
err := sdk.Flush()
```
Flush forces saving pending translations to the source.
Only relevant for SaveStrategyOnDemand.

### Logger
Set custom logger through SetLogger (logger.go):

```go
type MyLogger struct{}
func (l *MyLogger) Errorf(f string, a ...interface{}) {}
func (l *MyLogger) Infof(f string, a ...interface{})  {}
func (l *MyLogger) Debugf(f string, a ...interface{}) {}

sdk.SetLogger(&MyLogger{})
```

# Static translations

Static translations are read from the in-memory bundle loaded/synced by `Client`.

Implementation:
- `static.go` (`T`, `TA`)

## Simple lookup

```go
text := sdk.T("en_US", "checkout_total")
```
If locale bundle or key is missing, the SDK returns the key as fallback.

## Lookup with arguments

```go
text := sdk.TA("en_US", "user_name", map[string]interface{
    "name": "Olivia",
})
```
TA expects map[string]any. If another type is passed, the key is returned.

# Dynamic translations

Dynamic translations are managed via `DynamicContent`.

Implementation:
- `client.go` (`EnableDynamicContent`, `Dynamic`)
- `dynamic.go` (`DynamicContent` methods)

## Get dynamic handler

```go
dynamicKey := word.XKeyGen("tenant", "project", "namespace")
dyn := sdk.EnableDynamicContent(dynamicKey)
```
Use EnableDynamicContent(...) when your source requires a dynamic access key (remote source).

## Read dynamic values

With variables:

```go
value := dyn.TA("en_US", "promo_with_name", map[string]any{
    "name": "Alice",
})
```

## Save dynamic values

Single value:

```go
err := dyn.SaveTranslation("en_US", "promo_banner", "Limited offer")
```

Batch save:

```go
err := dyn.SaveTranslations([]word.Translation{
    {
        Locale: "en_US",
        Key:    "promo_banner",
        Value:  "Limited offer",
    },
    {
        Locale: "es_CO",
        Key:    "promo_banner",
        Value:  "Oferta limitada",
    },
})
```