package wordsdk_go

import (
	"github.com/summit-fi/wordsdk-go/fluent"
	"golang.org/x/text/language"
)

type WordSDK interface {
	Translate(lang language.Tag, key string, ctx ...*fluent.FormatContext) (string, error)
}

func Index(i int) *fluent.FormatContext {
	return fluent.WithVariable("index", i)
}

func Count(i int) *fluent.FormatContext {
	return fluent.WithVariable("count", i)
}

func Custom(input map[string]interface{}) *fluent.FormatContext {
	return fluent.WithVariables(input)
}

func CustomFunction(name string, function fluent.Function) *fluent.FormatContext {
	return fluent.WithFunction(name, function)
}

func Argument(name string, value interface{}) *fluent.FormatContext {
	return fluent.WithVariable(name, value)
}

func Arguments(args map[string]interface{}) *fluent.FormatContext {
	return fluent.WithVariables(args)
}
