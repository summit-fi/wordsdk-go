package test

import (
	"fmt"
	"testing"

	"github.com/summit-fi/wordsdk-go/fluent"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestCLDR(t *testing.T) {
	ftl := `cldr-month-narrow = { $index ->
*[0] J
[1] F
[2] M
[3] A
[4] M
[5] J
[6] J
[7] A
[8] S
[9] O
[10] N
[11] D
}`
	resource, err := fluent.NewResource(ftl)
	if err != nil {
		t.Fatalf("Failed to create resource: %v", err)
	}
	bundle := fluent.NewBundle(language.English)

	errs := bundle.AddResource(resource)
	if errs != nil {
		for _, errt := range errs {
			if errt != nil {
				t.Errorf("bundle.AddResource error: %s", errt)
			}
		}
	}

	for i := 0; i < 12; i++ {
		msg, _, fatalErr := bundle.FormatMessage("cldr-month-narrow", fluent.WithVariable("index", i))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error: %s", fatalErr)
		}
		expected := "JFMAMJJASOND"[i : i+1]
		if msg != expected {
			t.Errorf("Expected '%s', got '%s'", expected, msg)
		}
	}

}

func TestCLDRCurrency(t *testing.T) {
	ftl := `cldr-currency = { $amount ->
*[0] { $symbol -> $amount }
[1] { $symbol -> $amount }
[2] { $symbol -> $amount }
}`
	resource, err := fluent.NewResource(ftl)
	if err != nil {
		t.Fatalf("Failed to create resource: %v", err)
	}
	bundle := fluent.NewBundle(language.English)
	errs := bundle.AddResource(resource)
	if errs != nil {
		for _, errt := range errs {
			if errt != nil {
				t.Errorf("bundle.AddResource error: %s", errt)
			}
		}
	}
	amount := 1234567.89
	symbol := "$"
	msg, _, fatalErr := bundle.FormatMessage("cldr-currency", fluent.WithVariable("amount", amount), fluent.WithVariable("symbol", symbol))
	if fatalErr != nil {
		t.Errorf("bundle.FormatMessage fatal error: %s", fatalErr)
	}
	expected := fmt.Sprintf("%s %.2f", symbol, amount)
	if msg != expected {
		t.Errorf("Expected '%s', got '%s'", expected, msg)
	}
}

func formatCurrency(amount float64, symbol string) string {
	// Format the number: #,##0.00 and add the currency symbol
	p := message.NewPrinter(language.English)
	return fmt.Sprintf("%s %s", p.Sprintf("%0.2f", amount), symbol)
}

func TestCurrency(t *testing.T) {
	amount := 1234567.89

	// Use language-specific printer for currency formatting
	p := message.NewPrinter(language.English)

	// Format: grouped thousands, 2 decimals, currency symbol
	p.Printf("%s\n", formatCurrency(amount, "$"))

}
