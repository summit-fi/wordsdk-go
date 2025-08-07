package numbers

import "github.com/summit-fi/wordsdk-go/fluent/cldr"

type OrdinalFormatter struct {
	Pattern  string
	Language cldr.Language
}

func (f OrdinalFormatter) Format(num float64, opt ...Option) string {
	return f.Language.OrdinalRules(int(num))
}
