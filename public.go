package wordsdk_go

import "github.com/summit-fi/wordsdk-go/fluent"

type WordSDK interface {
	Translate(key string, ctx ...*fluent.FormatContext) (string, error)
}

func Index(i int) *fluent.FormatContext {
	return fluent.WithVariable("cldr-month-narrow", i)
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
