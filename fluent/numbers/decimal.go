package numbers

import "github.com/summit-fi/wordsdk-go/fluent/cldr"

type DecimalFormatter struct {
	Pattern string
	Base    cldr.Numbers
}

func (f DecimalFormatter) Format(num float64, opt ...Option) string {
	var pattern string

	if len(f.Pattern) == 0 {
		pattern = f.Base.DecimalPattern
	} else {
		pattern = f.Pattern
	}

	pf := PatternFormatter{
		Pattern: pattern,
		Base:    f.Base,
	}

	return pf.Format(num, opt...)
}
