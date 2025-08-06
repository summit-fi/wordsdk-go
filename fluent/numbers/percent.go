package numbers

import "github.com/summit-fi/wordsdk-go/fluent/cldr"

// PercentFormatter formats percent values
type PercentFormatter struct {
	Pattern string
	Base    cldr.Numbers
}

func (f PercentFormatter) Format(num float64, opt ...Option) string {
	var pattern string

	if len(f.Pattern) == 0 {
		pattern = f.Base.PercentPattern
	} else {
		pattern = f.Pattern
	}

	pf := PatternFormatter{
		Pattern: pattern,
		Base:    f.Base,
	}
	return pf.Format(num, opt...)
}
