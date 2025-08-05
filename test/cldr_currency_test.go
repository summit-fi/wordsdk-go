package test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	word "github.com/summit-fi/wordsdk-go"
	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

func TestCurrencyCLDRTest(t *testing.T) {

	readDir, err := os.ReadDir(filepath.Join(word.Root(), fixturesPath))
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for _, entry := range readDir {
		if entry.IsDir() {
			continue
		}

		file, err := os.OpenFile(filepath.Join(word.Root(), fixturesPath, entry.Name()), os.O_RDONLY, 0644)
		if err != nil {
			t.Fatalf("Failed to open file %s: %v", entry.Name(), err)
		}
		defer file.Close()

		var model MainCustomTestModel

		if err := json.NewDecoder(file).Decode(&model); err != nil {
			t.Fatalf("Failed to decode JSON from file %s: %v", entry.Name(), err)
		}

		code := entry.Name()[:len(entry.Name())-5] // Remove .json extension

		for _, test := range model.Tests {
			if test.Type != "NumberFormatter::currency" {
				continue
			}

			bundle := fluent.NewBundle(cldr.GetLanguageTypeByCode(code))

			errors := executeNumberFormatterCurrency(t, bundle, test.Scenarios)
			for _, err := range errors {
				t.Error(err)
			}

		}
	}

}

func executeNumberFormatterCurrency(t *testing.T, bundle *fluent.Bundle, testScenarios []TestScenarios) []error {
	var errors []error

	for i, scenario := range testScenarios {
		name := fmt.Sprintf("%s-%d", "scenario", i)

		// Add currency formatting message to bundle
		currencyMessage := fmt.Sprintf(`%s = { NUMBER($amount, style: "currency", currencyDisplay: "symbol") }`, name)
		resource, errs := fluent.NewResource(currencyMessage)
		if errs != nil {
			errors = append(errors, fmt.Errorf("failed to create currency resource: %v", errs))
			continue
		}

		bundle.AddResource(resource)

		amount, err := strconv.ParseFloat(scenario.Value, 64)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to parse amount %s: %v", scenario.Value, err))
			continue
		}
		msg, _, fatalErr := bundle.FormatMessage(name, fluent.WithVariable("amount", amount))
		if fatalErr != nil {
			errors = append(errors, fmt.Errorf("bundle.FormatMessage fatal error: %s", fatalErr))
			continue
		}

		if msg != scenario.Expected {
			t.Errorf("%s - %+v : expected %s, got %s",
				bundle.RootLocale(), scenario, scenario.Expected, msg)
		}
	}

	return errors
}
