package fluent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

const (
	numberStyle         = "style"    // Named parameter for number style
	numberStyleCurrency = "currency" // Currency style
	numberStylePercent  = "percent"  // Percent style
	numberStyleDecimal  = "decimal"  // Decimal style
	numberStyleOrdinal  = "ordinal"  // Ordinal style

	numberCurrencyDisplay       = "currencyDisplay" // Named parameter for currency display
	numberCurrencyDisplaySymbol = "symbol"          // Display currency as symbol
	numberCurrencyDisplayCode   = "code"            // Display currency as code

	numberPattern = "pattern" // Named parameter for custom pattern

	numberParameterMinimumFractionDigits = "minimumFractionDigits" // Named parameter for minimum fraction digits
	numberParameterMaximumFractionDigits = "maximumFractionDigits" // Named parameter for maximum fraction digits

)

func LoadNumberRules(lang cldr.Language) (cldr.Numbers, error) {
	root := func() string {
		current, _ := os.Getwd()
		for {
			goModPath := filepath.Join(current, "go.mod")
			if _, err := os.Stat(goModPath); err == nil {
				return current // Found go.mod, so this is the root
			}

			// Move up one directory
			parent := filepath.Dir(current)
			if parent == current { // Reached root of filesystem
				return current
			}
			current = parent
		}
	}

	path := filepath.Join(root(), "test", "fixtures", "custom", "custom_data", fmt.Sprintf("%s.ftl", lang))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cldr.Numbers{}, fmt.Errorf("file for language '%s' does not exist at path: %s", lang, path)
	}
	f, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return cldr.Numbers{}, fmt.Errorf("failed to open file for language '%s': %v", lang, err)
	}

	defer f.Close()

	// Read the file content
	b, err := io.ReadAll(f)
	if err != nil {
		return cldr.Numbers{}, fmt.Errorf("failed to read file for language '%s': %v", lang, err)
	}

	bundle := NewBundle(cldr.Language(lang))
	if bundle == nil {
		return cldr.Numbers{}, fmt.Errorf("bundle for language '%s' is nil", lang)
	}
	resource, errs := NewResource(string(b))
	if errs != nil {
		return cldr.Numbers{}, fmt.Errorf("failed to create resource for language '%s': %v", lang, errs)
	}
	if err := bundle.AddResource(resource); err != nil {
		return cldr.Numbers{}, fmt.Errorf("failed to add resource for language '%s': %v", lang, err)
	}
	messages := bundle.RetrieveMessages()
	// Extract numbers from the bundle
	numbers := cldr.Numbers{}

	numbers.DecimalSep = messages["cldr-decimal-sep"]
	numbers.GroupSep = messages["cldr-group-sep"]
	numbers.Percent = messages["cldr-percent"]
	numbers.ZeroDigit = messages["cldr-zero-digit"]
	numbers.PlusSign = messages["cldr-plus-sign"]
	numbers.MinusSign = messages["cldr-minus-sign"]
	numbers.ExpSymbol = messages["cldr-exp-symbol"]
	numbers.Permill = messages["cldr-permill"]
	numbers.Infinity = messages["cldr-infinity"]
	numbers.NaN = messages["cldr-nan"]
	numbers.DecimalPattern = messages["cldr-decimal-pattern"]
	numbers.ScientificPattern = messages["cldr-scientific-pattern"]
	numbers.PercentPattern = messages["cldr-percent-pattern"]
	numbers.CurrencyPattern = messages["cldr-currency-pattern"]
	numbers.CurrencyCode = messages["cldr-def-currency-code"]
	numbers.SelectedCurrencyCode = numbers.CurrencyCode
	numbers.DisplaySymbol = true
	numbers.Currency = make(map[string]*cldr.Currency)

	return numbers, nil
}
