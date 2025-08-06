package fluent

import (
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

func DateTimeFunc(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	return &NumberValue{Value: positional[0].(*NumberValue).Value}
}
