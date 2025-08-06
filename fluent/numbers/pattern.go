package numbers

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

type PatternFormatter struct {
	Pattern string
	Base    cldr.Numbers
}

type PatternInfo struct {
	Pattern          string
	HasPercent       bool
	HasCurrency      bool
	CurrencyPosition string // "prefix" or "suffix"
	DecimalPrecision int
	HasDecimal       bool
	MinIntegerDigits int
	MaxIntegerDigits int
	GroupingSize     int
	HasGrouping      bool
	HasNBSP          bool   // Non-breaking space
	SpacePosition    string // "prefix" or "suffix"
}

func (f PatternFormatter) Format(n float64, opts ...Option) string {
	// Merge options
	options := mergeOptions(opts)

	// Analyze the pattern once
	info := AnalyzePattern(f.Pattern)

	isNeg := n < 0
	if isNeg {
		n = -n
	}

	s := f.Base

	// Apply percent conversion
	if info.HasPercent {
		n *= 100
	}

	// Format number with options
	intPart, fracPart := formatWithOptions(n, info, options)

	// Apply grouping if needed
	if info.HasGrouping {
		groupSep := s.GroupSep
		intPart = insertGrouping(intPart, groupSep)
	}

	// Handle maximum fraction digits
	if options.MaximumFractionDigits != nil && len(fracPart) > *options.MaximumFractionDigits {
		fracPart = fracPart[:*options.MaximumFractionDigits]
	}
	if options.MaximumFractionDigits == nil && info.HasDecimal && len(fracPart) > info.DecimalPrecision {
		fracPart = fracPart[:info.DecimalPrecision]
	}

	if options.MinimumFractionDigits == nil && info.HasDecimal && len(fracPart) > 0 {
		// If there are decimal digits, we need to ensure we have at least one digit
		if len(fracPart) == 0 {
			fracPart = "0"
		}
	}

	// Build result
	res := intPart
	if (info.HasDecimal && len(fracPart) > 0) || options.MinimumFractionDigits != nil {
		if len(fracPart) > 0 {
			res += s.DecimalSep + fracPart
		} else if options.MinimumFractionDigits != nil {
			res += s.DecimalSep + strings.Repeat("0", *options.MinimumFractionDigits)
		}
	}

	// Add currency
	if info.HasCurrency {
		display := s.Currency[s.SelectedCurrencyCode].Code
		if s.DisplaySymbol {
			display = s.Currency[s.SelectedCurrencyCode].Symbol
		}

		if info.CurrencyPosition == "prefix" {
			if info.HasNBSP {
				display = display + " "
			}

			res = display + res
		} else {
			res = res + " " + display
		}
	}

	// Add percent symbol
	if info.HasPercent {
		res += s.Percent
	}

	// Add minus sign for negative numbers
	if isNeg {
		res = s.MinusSign + res
	}

	return res
}

func mergeOptions(opts []Option) Option {
	var result Option
	for _, opt := range opts {
		if opt.MinimumFractionDigits != nil {
			result.MinimumFractionDigits = opt.MinimumFractionDigits
		}
		if opt.MaximumFractionDigits != nil {
			result.MaximumFractionDigits = opt.MaximumFractionDigits
		}
		if opt.MinimumIntegerDigits > 0 {
			result.MinimumIntegerDigits = opt.MinimumIntegerDigits
		}
	}
	return result
}

func AnalyzePattern(pattern string) PatternInfo {

	info := PatternInfo{
		Pattern:     pattern,
		HasPercent:  strings.Contains(pattern, "%"),
		HasCurrency: strings.Contains(pattern, "¤"),
	}

	if info.HasCurrency {
		info.CurrencyPosition = "suffix"
		if strings.HasPrefix(pattern, "¤") {
			info.CurrencyPosition = "prefix"
		}
	}

	// Find decimal position and calculate precision
	decimalPos := strings.Index(pattern, ".")
	info.HasDecimal = decimalPos >= 0

	if info.HasDecimal {
		// Count decimal places (0 and # after decimal point)
		for _, ch := range pattern[decimalPos+1:] {
			if ch == '0' || ch == '#' {
				info.DecimalPrecision++
			}
		}
	}

	// Analyze integer part
	integerPart := pattern
	if info.HasDecimal {
		integerPart = pattern[:decimalPos]
	}

	// Remove currency and percent symbols for analysis
	integerPart = strings.ReplaceAll(integerPart, "¤", "")
	integerPart = strings.ReplaceAll(integerPart, "%", "")

	info.HasGrouping = strings.Contains(integerPart, ",")

	if strings.Contains(pattern, " ") {
		info.HasNBSP = true
		info.SpacePosition = "suffix"
		if strings.LastIndex(pattern, " ") == len(pattern)-1 {
			info.SpacePosition = "prefix"
		}
	}

	if strings.Contains(pattern, " ") {
		info.HasNBSP = true
		info.SpacePosition = "suffix"
		if strings.LastIndex(pattern, " ") == len(pattern)-1 {
			info.SpacePosition = "prefix"
		}
	}

	// Count digits in integer part
	digitCount := 0
	zeroCount := 0

	for _, ch := range integerPart {
		if ch == '0' {
			digitCount++
			zeroCount++
		} else if ch == '#' {
			digitCount++
		}
	}

	info.MinIntegerDigits = zeroCount
	info.MaxIntegerDigits = digitCount

	// Determine grouping size (typically 3 for most locales)
	if info.HasGrouping {
		info.GroupingSize = 3 // Standard grouping size
	}

	return info
}

func formatWithOptions(n float64, info PatternInfo, opts Option) (string, string) {
	// Use significant digits if specified

	// Determine decimal precision
	precision := info.DecimalPrecision
	if opts.MaximumFractionDigits != nil {
		precision = *opts.MaximumFractionDigits
	} else if opts.MinimumFractionDigits != nil && precision < *opts.MinimumFractionDigits {
		precision = *opts.MinimumFractionDigits
	}

	// Format number
	factor := math.Pow(10, float64(precision))
	rounded := math.Round(n*factor) / factor

	sfmt := "%." + strconv.Itoa(precision) + "f"
	numStr := fmt.Sprintf(sfmt, rounded)

	// Split into integer and fractional parts
	intPart, fracPart := splitDecimal(numStr)

	// Handle minimum integer digits
	minIntDigits := info.MinIntegerDigits
	if opts.MinimumIntegerDigits > 0 {
		minIntDigits = opts.MinimumIntegerDigits
	}

	// Pad integer part with zeros if needed
	if len(intPart) < minIntDigits {
		intPart = strings.Repeat("0", minIntDigits-len(intPart)) + intPart
	}

	if info.HasDecimal && !strings.Contains(fractionPattern(info), "0") {
		fracPart = strings.TrimRight(fracPart, "0")
		if fracPart == "" {
			// Remove decimal point if no fraction left
			fracPart = ""
		}
	}

	return intPart, fracPart
}

func fractionPattern(info PatternInfo) string {
	decimalPos := strings.Index(info.Pattern, ".")
	if decimalPos >= 0 {
		return info.Pattern[decimalPos+1:]
	}
	return ""
}

func splitDecimal(s string) (string, string) {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

func insertGrouping(s, sep string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var b strings.Builder
	pre := n % 3
	if pre > 0 {
		b.WriteString(s[:pre])
		if n > pre {
			b.WriteString(sep)
		}
	}
	for i := pre; i < n; i += 3 {
		b.WriteString(s[i : i+3])
		if i+3 < n {
			b.WriteString(sep)
		}
	}
	return b.String()
}
