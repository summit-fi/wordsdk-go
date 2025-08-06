package fluent

import (
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

// Function represents a function that builds a Value based on parameters
type Function func(positional []Value, named map[string]Value, primaryLanguage cldr.Language, params ...string) Value

// A Value is the result of a resolving operation performed by the Resolver.
// It represents either a string, a number or a date time.
type Value interface {
	String() string
}

// StringValue wraps a string in order to comply with the Value API
type StringValue struct {
	Value string
}

// String returns the wrapped value of a StringValue
func (value *StringValue) String() string {
	return value.Value
}

// String returns a new StringValue with the given value; used for variables
func String(val string) *StringValue {
	return &StringValue{
		Value: val,
	}
}

// NoValue is used whenever no "real" value could be built
type NoValue struct {
	value string
}

// String returns the NoValue's string representation
func (value *NoValue) String() string {
	return "{" + value.value + "}"
}
