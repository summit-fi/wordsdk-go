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
