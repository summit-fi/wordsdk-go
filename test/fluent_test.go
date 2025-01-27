package test

import (
	"fmt"
	"testing"

	"github.com/summit-fi/wordsdk-go/fluent"
	"golang.org/x/text/language"
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
	bundle := fluent.NewBundle(language.English)

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

func TestDateTime(t *testing.T) {
	ftl := `
datetime = { DATETIME($date, pattern:"") }
`
	resource, err := fluent.NewResource(string(ftl))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	bundle := fluent.NewBundle(language.English)
	bundle.AddResource(resource)

	date := "2021-09-01T12:00:00Z"
	msg, _, fatalErr := bundle.FormatMessage("datetime", fluent.WithVariable("date", date))
	if fatalErr != nil {
		t.Errorf("bundle.FormatMessage fatal error: %s", fatalErr)
		panic(fatalErr)
	}

	expected := "2021-09-01"
	if msg != expected {
		t.Errorf("bundle.FormatMessage error: %s", msg)
	}
}
