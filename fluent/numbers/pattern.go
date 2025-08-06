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

func mergeOptions(opts []Option) Option {
	var result Option
	for _, opt := range opts {
		if opt.MinimumFractionDigits > 0 {
			result.MinimumFractionDigits = opt.MinimumFractionDigits
		}
		if opt.MaximumFractionDigits > 0 {
			result.MaximumFractionDigits = opt.MaximumFractionDigits
		}
		if opt.MinimumIntegerDigits > 0 {
			result.MinimumIntegerDigits = opt.MinimumIntegerDigits
		}
	}
	return result
}

func formatWithOptions(n float64, info PatternInfo, opts Option) (string, string) {
	// Use significant digits if specified

	// Determine decimal precision
	precision := info.DecimalPrecision
	if opts.MaximumFractionDigits > 0 {
		precision = opts.MaximumFractionDigits
	} else if opts.MinimumFractionDigits > 0 && precision < opts.MinimumFractionDigits {
		precision = opts.MinimumFractionDigits
	}

	// Format number
	factor := math.Pow(10, float64(precision))
	rounded := math.Round(n*factor) / factor

	sfmt := "%." + strconv.Itoa(precision) + "f"
	numStr := fmt.Sprintf(sfmt, rounded)

	// Split into integer and fractional parts
	ip, fp := splitDecimal(numStr)

	// Handle minimum integer digits
	minIntDigits := info.MinIntegerDigits
	if opts.MinimumIntegerDigits > 0 {
		minIntDigits = opts.MinimumIntegerDigits
	}

	// Pad integer part with zeros if needed
	if len(ip) < minIntDigits {
		ip = strings.Repeat("0", minIntDigits-len(ip)) + ip
	}

	// Ensure minimum fraction digits
	if opts.MinimumFractionDigits > 0 && len(fp) < opts.MinimumFractionDigits {
		fp = fp + strings.Repeat("0", opts.MinimumFractionDigits-len(fp))
	}

	return ip, fp
}

func AnalyzePattern(pattern string) PatternInfo {
	info := PatternInfo{}

	// Check for percent symbol
	info.HasPercent = strings.Contains(pattern, "%")

	// Check for currency symbol and position
	info.HasCurrency = strings.Contains(pattern, "¤")
	if info.HasCurrency {
		if strings.HasPrefix(pattern, "¤") {
			info.CurrencyPosition = "prefix"
		} else {
			info.CurrencyPosition = "suffix"
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

	// Count digits in integer part
	digitCount := 0
	zeroCount := 0
	hasComma := strings.Contains(integerPart, ",")

	info.HasGrouping = hasComma

	if strings.Index(pattern, " ") >= 0 { // Non-breaking space
		if strings.LastIndex(pattern, " ") == len(pattern)-1 {
			info.SpacePosition = "prefix"
		} else {
			info.SpacePosition = "suffix"
		}
		info.HasNBSP = true

	}

	if strings.Index(pattern, " ") >= 0 { // Non-breaking space
		if strings.LastIndex(pattern, " ") == len(pattern)-1 {
			info.SpacePosition = "prefix"
		} else {
			info.SpacePosition = "suffix"
		}
		info.HasNBSP = true
	}

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

func (f PatternFormatter) Format(n float64, opts ...Option) string {
	// Merge options
	opt := mergeOptions(opts)

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
	ip, fp := formatWithOptions(n, info, opt)

	// Apply grouping if needed
	if info.HasGrouping {
		groupSep := s.GroupSep
		ip = insertGrouping(ip, groupSep)
	}

	// Handle maximum fraction digits
	if opt.MaximumFractionDigits > 0 && len(fp) > opt.MaximumFractionDigits {
		fp = fp[:opt.MaximumFractionDigits]
	}

	// Build result
	res := ip
	if (info.HasDecimal && len(fp) > 0) || opt.MinimumFractionDigits > 0 {
		if len(fp) > 0 {
			res += s.DecimalSep + fp
		} else if opt.MinimumFractionDigits > 0 {
			res += s.DecimalSep + strings.Repeat("0", opt.MinimumFractionDigits)
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
