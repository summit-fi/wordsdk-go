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
[restaurant] Welcome to our { table }
[tennis]  Welcome to your { court }
    }.
    
room = { $count ->
[0] no rooms
*[one] one room
[other] {$count} rooms
}

table = { $count ->
[0] no tables
*[one] one table
[two] two tables
[few] {$count} tables
[other] {$count} tables
}

court = { $count ->
[0] no courts
*[one] tennis court
[two] two tennis courts
[few] {$count} tennis courts
[other] {$count} tennis courts
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
	testCases := []struct {
		place    string
		count    float64
		expected string
	}{
		{"hotel", 0, "Welcome to your no rooms."},
		{"hotel", 1, "Welcome to your one room."},
		{"hotel", 3, "Welcome to your 3 rooms."},
		{"hotel", 55, "Welcome to your 55 rooms."},

		{"restaurant", 0, "Welcome to our no tables."},
		{"restaurant", 1, "Welcome to our one table."},
		{"restaurant", 2, "Welcome to our 2 tables."},
		{"restaurant", 10, "Welcome to our 10 tables."},

		{"tennis", 0, "Welcome to your no courts."},
		{"tennis", 1, "Welcome to your tennis court."},
		{"tennis", 2, "Welcome to your 2 tennis courts."},
		{"tennis", 5, "Welcome to your 5 tennis courts."},
		{"tennis", 100, "Welcome to your 100 tennis courts."},
	}

	for _, tc := range testCases {
		msg, _, fatalErr := bundle.FormatMessage("spot", fluent.WithVariable("count", tc.count), fluent.WithVariable("center", tc.place))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for score %s - %v: %s", tc.place, tc.count, fatalErr)
			continue
		}
		if msg != tc.expected {
			t.Errorf("Score formatting error. Expected: %s, Got: %s", tc.expected, msg)
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
		//{1.0, "You scored 1.0 points."},
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

func TestEmailPatternsNumberFormatting(t *testing.T) {
	ftl := `emails = You have { $unreadEmails } unread emails.
emails2 = You have { NUMBER($unreadEmails) } unread emails.`

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
		unreadEmails     int
		expectedImplicit string
		expectedExplicit string
	}{
		{0, "You have 0 unread emails.", "You have 0 unread emails."},
		{1, "You have 1 unread emails.", "You have 1 unread emails."},
		{5, "You have 5 unread emails.", "You have 5 unread emails."},
		{42, "You have 42 unread emails.", "You have 42 unread emails."},
		{1000, "You have 1000 unread emails.", "You have 1,000 unread emails."},
		{1234567, "You have 1234567 unread emails.", "You have 1,234,567 unread emails."},
	}

	for _, tc := range testCases {
		// Test implicit number formatting (emails)
		msg1, _, fatalErr := bundle.FormatMessage("emails", fluent.WithVariable("unreadEmails", tc.unreadEmails))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for emails with %d: %s", tc.unreadEmails, fatalErr)
			continue
		}

		// Test explicit number formatting (emails2)
		msg2, _, fatalErr := bundle.FormatMessage("emails2", fluent.WithVariable("unreadEmails", tc.unreadEmails))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for emails2 with %d: %s", tc.unreadEmails, fatalErr)
			continue
		}

		// Check implicit formatting
		if msg1 != tc.expectedImplicit {
			t.Errorf("Implicit formatting error for %d emails. Expected: %s, Got: %s",
				tc.unreadEmails, tc.expectedImplicit, msg1)
		}

		// Check explicit formatting
		if msg2 != tc.expectedExplicit {
			t.Errorf("Explicit formatting error for %d emails. Expected: %s, Got: %s",
				tc.unreadEmails, tc.expectedExplicit, msg2)
		}

		t.Logf("Emails %d - Implicit: %s, Explicit: %s", tc.unreadEmails, msg1, msg2)
	}
}

func TestOrdinalNumberWithSelect(t *testing.T) {
	ftl := `your-rank = { NUMBER($pos, style: "ordinal") ->
   [1] You finished first!
   [one] You finished {$pos}st
   [two] You finished {$pos}nd
   [few] You finished {$pos}rd
  *[other] You finished {$pos}th
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
		position int
		expected string
	}{
		//{1, "You finished first!"},
		{2, "You finished 2nd"},
		{3, "You finished 3rd"},
		{4, "You finished 4th"},
		{11, "You finished 11th"},
		{12, "You finished 12th"},
		{13, "You finished 13th"},
		{21, "You finished 21st"},
		{22, "You finished 22nd"},
		{23, "You finished 23rd"},
		{24, "You finished 24th"},
		{31, "You finished 31st"},
		{42, "You finished 42nd"},
		{53, "You finished 53rd"},
		{100, "You finished 100th"},
		{101, "You finished 101st"},
		{102, "You finished 102nd"},
		{103, "You finished 103rd"},
		{111, "You finished 111th"},
		{1000, "You finished 1000th"},
	}

	for _, tc := range testCases {
		msg, _, fatalErr := bundle.FormatMessage("your-rank", fluent.WithVariable("pos", tc.position))
		if fatalErr != nil {
			t.Errorf("bundle.FormatMessage fatal error for position %d: %s", tc.position, fatalErr)
			continue
		}

		if msg != tc.expected {
			t.Errorf("Expected: %s, Got: %s",
				tc.expected, msg)
		}

	}
}
