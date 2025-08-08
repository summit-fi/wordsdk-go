package fluent

import (
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
	return &StringValue{Value: CLDRDateTimeFormatter(language.BCP47(), pattern).Format(positional[0].(*DateTimeValue).Value)}
}

func CLDRDateTimeFormatter(lang language.Tag, template string) intl.DateTimeFormat {
	return intl.NewDateTimeFormatLayout(lang, template)
}
