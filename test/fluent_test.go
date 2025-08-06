package test

import (
	"fmt"
	"testing"

	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

func TestHotelSpot(t *testing.T) {
	ftl := `
spot =
{ $center ->
*[hotel] Welcome to your { room }
[restaurant] Welcome to your { table }
[tennis]  Welcome to your  { court }
    }.
    
room = { $count ->
[0] no rooms
*[one] one room
[other] {$count} rooms
}

table = { $count ->
[0] no tables
*[one] one table
[other] {$count} tables
}

court = { $count ->
[0] no courts
*[one] one court
[other] {$count} courts
}`
	resource, err := fluent.NewResource(string(ftl))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	bundle := fluent.NewBundle(cldr.LanguageEnUS)

	errs := bundle.AddResource(resource)
	if errs != nil {
		for _, errt := range errs {
			if errt != nil {
				t.Errorf("bundle.AddResource error: %s", err)
			}
		}
	}

	count := []int{0, 1, 3, 55}

	result := make([]string, len(count))
	for i, c := range count {
		msg, _, fatalErr := bundle.FormatMessage("spot", fluent.WithVariable("count", c))

		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error: %s", fatalErr)
			panic(fatalErr)
		}

		result[i] = msg
	}
	expected := []string{
		"Welcome to your no rooms.",
		"Welcome to your one room.",
		"Welcome to your 3 rooms.",
		"Welcome to your 55 rooms.",
	}

	for i, r := range result {
		fmt.Println(i, r)
		if r != expected[i] {
			t.Errorf("bundle.FormatMessage error: %s", r)
		}
	}
}

func TestImplicitlyCallNumber(t *testing.T) {
	ftl := `emails = Number of unread emails { $unreadEmails }`
	resource, err := fluent.NewResource(string(ftl))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	bundle := fluent.NewBundle(cldr.LanguageEnUS)
	bundle.AddResource(resource)
	bundle.FormatMessage("emails", fluent.WithVariable("unreadEmails", 5))
	msg, _, fatalErr := bundle.FormatMessage("emails", fluent.WithVariable("unreadEmails", 5))
	if fatalErr != nil {
		t.Errorf("bundle.FormatMessage fatal error: %s", fatalErr)
		panic(fatalErr)
	}
	expected := "Number of unread emails 5"
	if msg != expected {
		t.Errorf("bundle.FormatMessage error: %s", msg)
	}
}

func TestNumberWithSelectExpression(t *testing.T) {
	ftl := `your-score =  
      { NUMBER($score, minimumFractionDigits: 1) ->
          [0.0]   You scored zero points. What happened?
          *[other] You scored { NUMBER($score, minimumFractionDigits: 1) } points.
      }`
	resource, err := fluent.NewResource(string(ftl))
	if err != nil {
		t.Error(err)
		return
	}
	bundle := fluent.NewBundle(cldr.LanguageEnUS)
	errs := bundle.AddResource(resource)
	if errs != nil {
		for _, errt := range errs {
			if errt != nil {
				t.Errorf("bundle.AddResource error: %s", errt)
			}
		}
		return
	}

	testCases := []struct {
		score    float64
		expected string
	}{
		{0.0, "You scored zero points. What happened?"},
		{5.0, "You scored 5.0 points."},
		{10.5, "You scored 10.5 points."},
		{100.0, "You scored 100.0 points."},
	}

	for _, tc := range testCases {
		msg, _, fatalErr := bundle.FormatMessage("your-score", fluent.WithVariable("score", tc.score))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for score %v: %s", tc.score, fatalErr)
			continue
		}
		if msg != tc.expected {
			t.Errorf("Score formatting error. Expected: %s, Got: %s", tc.expected, msg)
		}
	}
}

func TestNumberFormattingVariations(t *testing.T) {
	ftl := `price-info-code =
      price with code - { NUMBER($amount, style: "currency", currencyDisplay: "code") }
price-info-simple =
    price simple - { NUMBER($amount, style: "currency") }

price-info-symbol =
      price with symbol - { NUMBER($amount, style: "currency", currency: "COP", currencySymbol: "$", minimumFractionDigits: 0, currencyDisplay: "symbol") }

percent-info = percent : {NUMBER($amount, style: "percent")}

price-info-custom-pattern =
      price with custom pattern - { NUMBER($amount, style: "currency", currency: "UAH", currencySymbol: "₴", currencyDisplay: "symbol", pattern: "#,##0.00 ¤") }
price-info-code-pattern =
        price with code pattern - { NUMBER($amount, style: "currency", currency: "LKSDFLSDJF", currencyDisplay: "code",pattern: "#,##0.00 ¤") }

price-with-min-fraction = 
        price with minimumFractionDigits 

percent-info-pattern =
        percent info pattern - {NUMBER($amount,style:"percent")}`

	resource, err := fluent.NewResource(string(ftl))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	bundle := fluent.NewBundle(cldr.LanguageEnUS)

	errs := bundle.AddResource(resource)
	if errs != nil {
		for _, errt := range errs {
			if errt != nil {
				t.Errorf("bundle.AddResource error: %s", errt)
			}
		}
	}

	// Test currency formatting variations
	t.Run("CurrencyFormatting", func(t *testing.T) {
		amount := 1234.56

		// Test price-info-code
		msg, _, fatalErr := bundle.FormatMessage("price-info-code", fluent.WithVariable("amount", amount))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for price-info-code: %s", fatalErr)
		} else {
			fmt.Printf("price-info-code: %s\n", msg)
		}

		// Test price-info-simple
		msg, _, fatalErr = bundle.FormatMessage("price-info-simple", fluent.WithVariable("amount", amount))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for price-info-simple: %s", fatalErr)
		} else {
			fmt.Printf("price-info-simple: %s\n", msg)
		}

		// Test price-info-symbol
		msg, _, fatalErr = bundle.FormatMessage("price-info-symbol", fluent.WithVariable("amount", amount))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for price-info-symbol: %s", fatalErr)
		} else {
			fmt.Printf("price-info-symbol: %s\n", msg)
		}

		// Test price-info-custom-pattern
		msg, _, fatalErr = bundle.FormatMessage("price-info-custom-pattern", fluent.WithVariable("amount", amount))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for price-info-custom-pattern: %s", fatalErr)
		} else {
			fmt.Printf("price-info-custom-pattern: %s\n", msg)
		}

		// Test price-info-code-pattern
		msg, _, fatalErr = bundle.FormatMessage("price-info-code-pattern", fluent.WithVariable("amount", amount))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for price-info-code-pattern: %s", fatalErr)
		} else {
			fmt.Printf("price-info-code-pattern: %s\n", msg)
		}
	})

	// Test percent formatting
	t.Run("PercentFormatting", func(t *testing.T) {
		testCases := []struct {
			amount float64
		}{
			{0.15}, // 15%
			{0.5},  // 50%
			{1.0},  // 100%
			{2.5},  // 250%
		}

		for _, tc := range testCases {
			// Test percent-info
			msg, _, fatalErr := bundle.FormatMessage("percent-info", fluent.WithVariable("amount", tc.amount))
			if fatalErr != nil {
				t.Errorf("bundle.FormatMessage fatal error for percent-info with amount %v: %s", tc.amount, fatalErr)
			} else {
				fmt.Printf("percent-info with %v: %s\n", tc.amount, msg)
			}

			// Test percent-info-pattern
			msg, _, fatalErr = bundle.FormatMessage("percent-info-pattern", fluent.WithVariable("amount", tc.amount))
			if fatalErr != nil {
				t.Errorf("bundle.FormatMessage fatal error for percent-info-pattern with amount %v: %s", tc.amount, fatalErr)
			} else {
				fmt.Printf("percent-info-pattern with %v: %s\n", tc.amount, msg)
			}
		}
	})
}
