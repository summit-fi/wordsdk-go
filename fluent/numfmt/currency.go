package numfmt

import "github.com/summit-fi/wordsdk-go/fluent/cldr"

// CurrencyFormatter formats currency numbers
type CurrencyFormatter struct {
	Pattern string
	Base    cldr.Numbers
}

func (f CurrencyFormatter) Format(num float64, opt ...Option) string {
	var pattern string

	if len(f.Pattern) == 0 {
		pattern = f.Base.CurrencyPattern
	} else {
		pattern = f.Pattern
	}

	pf := PatternFormatter{
		Pattern: pattern,
		Base:    f.Base,
	}

	return pf.Format(num, opt...)
}
