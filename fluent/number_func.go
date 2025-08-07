package fluent

import (
	"fmt"
	"strconv"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"github.com/summit-fi/wordsdk-go/fluent/numbers"
)

// NumberValue wraps a number (float32 at the moment) in order to comply with the Value API
type NumberValue struct {
	Value float32
}

// String formats a NumberValue into a string
func (value *NumberValue) String() string {
	// TODO: Simulate JavaScript number formatting
	return strconv.FormatFloat(float64(value.Value), 'f', -1, 32)
}

// Number returns a new NumberValue with the given value; used for variables
func NumberLiteral(val float32) *NumberValue {
	return &NumberValue{
		Value: val,
	}
}

func NumberFunc(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	if len(positional) < 1 {
		return &StringValue{Value: positional[0].String()}
	}

	options := numbers.Option{}

	if num, hasMinimumFractionDigits := named[numberParameterMinimumFractionDigits]; hasMinimumFractionDigits {
		minFractionDigits, err := strconv.Atoi(num.String())
		if err != nil {
			return &NoValue{value: fmt.Sprintf("func NUMBER: invalid minimum fraction digits -> %s", num.String())}
		}
		if minFractionDigits < 0 {
			return &NoValue{value: fmt.Sprintf("func NUMBER: minimum fraction digits cannot be negative -> %d", minFractionDigits)}
		}
		options.MinimumFractionDigits = &minFractionDigits
	}

	if num, hasMaximumFractionDigits := named[numberParameterMaximumFractionDigits]; hasMaximumFractionDigits {
		maxFractionDigits, err := strconv.Atoi(num.String())
		if err != nil {
			return &NoValue{value: fmt.Sprintf("func NUMBER: invalid maximum fraction digits -> %s", num.String())}
		}
		if maxFractionDigits < 0 {
			return &NoValue{value: fmt.Sprintf("func NUMBER: maximum fraction digits cannot be negative -> %d", maxFractionDigits)}
		}
		options.MaximumFractionDigits = &maxFractionDigits
	}

	// Default if parameter style is not provided, we use decimal mode
	if _, hasStyle := named[numberStyle]; !hasStyle {
		named[numberStyle] = &StringValue{Value: numberStyleDecimal}
	}

	switch named[numberStyle].String() {
	case numberStyleCurrency:
		// clone needs to be cloned because it is mutable
		cloneFormat := language.GetNumberRules()

		if currency, hasCurrency := named[numberStyleCurrency]; hasCurrency {
			cloneFormat.SelectedCurrencyCode = currency.String()
		}

		cloneFormat.SelectedCurrencyCode = cloneFormat.CurrencyCode
		if err := cloneFormat.EnsureCurrencyExists(); err != nil {
			return &NoValue{value: fmt.Sprintf("func NUMBER: invalid currency code -> %s", cloneFormat.SelectedCurrencyCode)}
		}

		if symbol, hasCurrencySymbol := named[numberCurrencyDisplaySymbol]; hasCurrencySymbol {
			err := cloneFormat.ModifyCurrencySymbol(symbol.String())
			if err != nil {
				return &NoValue{value: fmt.Sprintf("func NUMBER: invalid currency symbol -> %s", symbol.String())}
			}
		}

		if code, hasCurrencyCode := named[numberCurrencyDisplayCode]; hasCurrencyCode {
			err := cloneFormat.ModifyCurrencyCode(code.String())
			if err != nil {
				return &NoValue{value: fmt.Sprintf("func NUMBER: invalid currency code -> %s", code.String())}
			}
		}

		num, err := strconv.ParseFloat(positional[0].String(), 64)
		if err != nil {
			return &NoValue{value: fmt.Sprintf("func NUMBER: invalid number cloneFormat -> %s", positional[0].String())}
		}

		currencyFormatter := numbers.CurrencyFormatter{
			Base: cloneFormat,
		}

		if pattern, hasPattern := named[numberPattern]; hasPattern {
			currencyFormatter.Pattern = pattern.String()
		}

		result := currencyFormatter.Format(num, options)

		return &StringValue{result}

	case numberStylePercent:
		num, _ := strconv.ParseFloat(positional[0].String(), 64)

		percentFormatter := numbers.PercentFormatter{
			Base: language.GetNumberRules(),
		}

		if pattern, hasPattern := named[numberPattern]; hasPattern {
			percentFormatter.Pattern = pattern.String()
		}

		result := percentFormatter.Format(num, options)

		return &StringValue{result}

	case numberStyleDecimal:
		num, _ := strconv.ParseFloat(positional[0].String(), 64)

		decimalFormatter := numbers.DecimalFormatter{
			Base: language.GetNumberRules(),
		}
		if pattern, hasPattern := named[numberPattern]; hasPattern {
			decimalFormatter.Pattern = pattern.String()
		}
		result := decimalFormatter.Format(num, options)
		return &StringValue{result}

	case numberStyleOrdinal:
		num, err := strconv.ParseFloat(positional[0].String(), 64)
		if err != nil {
			return &NoValue{value: fmt.Sprintf("func NUMBER: invalid number cloneFormat -> %s", positional[0].String())}
		}
		ordinalFormatter := numbers.OrdinalFormatter{
			Language: language,
		}
		return &StringValue{
			Value: ordinalFormatter.Format(num),
		}
	}

	numStr := positional[0].String()
	num, err := strconv.ParseFloat(numStr, 32)
	if err != nil {
		return &NoValue{value: fmt.Sprintf("func NUMBER: invalid number cloneFormat -> %s", numStr)}
	}

	return &NumberValue{Value: float32(num)}
}
