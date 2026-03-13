package fluent

import (
	"fmt"
	"time"

	"github.com/boltegg/intl"
	"golang.org/x/text/language"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

type DateTimeValue struct {
	Value time.Time
}

func (value *DateTimeValue) String() string {
	return value.Value.Format(time.RFC3339)
}

func DateTimeFunc(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	return &StringValue{}
}

func CLDRDateTimeFunc(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	pattern := named["pattern"].String()
	if pattern == "" {
		return &StringValue{Value: fmt.Sprintf("error: missing pattern")}
	}

	f, err := CLDRDateTimeFormatter(language.BCP47(), pattern)
	if err != nil {
		return &StringValue{Value: fmt.Sprintf("error: %v", err)}
	}

	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func CLDRDateTimeFormatter(lang language.Tag, template string) (intl.DateTimeFormat, error) {
	return intl.NewDateTimeFormatLayout(lang, template)
}
