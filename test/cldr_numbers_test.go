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

func ReadScenarios(t *testing.T) map[string]MainCustomTestModel {
	readDir, err := os.ReadDir(filepath.Join(word.Root(), fixturesPath))
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	cldrTestMap := make(map[string]MainCustomTestModel)

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

		if _, exists := cldrTestMap[code]; !exists {
			cldrTestMap[code] = model
		}
	}
	return cldrTestMap
}

func TestCurrencyCLDRTest(t *testing.T) {

	for key, tests := range ReadScenarios(t) {
		for _, test := range tests.Tests {
			if test.Type != "NumberFormatter::currency" {
				continue
			}

			bundle := fluent.NewBundle(cldr.Language(key))

			errors := executeNumberFormatterCurrency(t, bundle, test.Scenarios)
			for _, err := range errors {
				t.Error(err)
			}

		}
	}
	t.Logf("Tested %d scenarios for currency formatting in CLDR", len(ReadScenarios(t)))
}

func executeNumberFormatterCurrency(t *testing.T, bundle *fluent.Bundle, testScenarios []TestScenarios) []error {
	var errors []error

	for i, scenario := range testScenarios {
		t.Run(fmt.Sprintf("Scenario-%d", i), func(t *testing.T) {
			name := fmt.Sprintf("%s-%d", "scenario", i)

			// Add currency formatting message to bundle
			currencyMessage := fmt.Sprintf(`%s = { NUMBER($amount, style: "currency", currencyDisplay: "symbol") }`, name)
			resource, errs := fluent.NewResource(currencyMessage)
			if errs != nil {
				errors = append(errors, fmt.Errorf("failed to create currency resource: %v", errs))
				return
			}

			bundle.AddResource(resource)

			amount, err := strconv.ParseFloat(scenario.Value, 64)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to parse amount %s: %v", scenario.Value, err))
				return
			}
			msg, _, fatalErr := bundle.FormatMessage(name, fluent.WithVariable("amount", amount))
			if fatalErr != nil {
				errors = append(errors, fmt.Errorf("bundle.FormatMessage fatal error: %s", fatalErr))
				return
			}

			if msg != scenario.Expected {
				t.Errorf("%s - %+v : expected %s, got %s",
					bundle.PrimaryLocale(), scenario, scenario.Expected, msg)
			}
		})
	}

	return errors
}

func TestDecimalCLDR(t *testing.T) {

	for key, tests := range ReadScenarios(t) {
		for _, test := range tests.Tests {

			if test.Type != "NumberFormatter::decimal" {
				continue
			}

			bundle := fluent.NewBundle(cldr.Language(key))

			errors := executeNumberFormatterDecimal(t, bundle, test.Scenarios)
			for _, err := range errors {
				t.Error(err)
			}

		}
	}
}

func executeNumberFormatterDecimal(t *testing.T, bundle *fluent.Bundle, testScenarios []TestScenarios) []error {
	var errors []error

	for i, scenario := range testScenarios {
		t.Run(fmt.Sprintf("Scenario-%s", scenario.Value), func(t *testing.T) {
			name := fmt.Sprintf("%s-%d", "scenario", i)

			// Add decimal formatting message to bundle
			decimalMessage := fmt.Sprintf(`%s = { NUMBER($amount, style: "decimal") }`, name)
			resource, errs := fluent.NewResource(decimalMessage)
			if errs != nil {
				errors = append(errors, fmt.Errorf("failed to create decimal resource: %v", errs))
				return
			}

			bundle.AddResource(resource)

			amount, err := strconv.ParseFloat(scenario.Value, 64)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to parse amount %s: %v", scenario.Value, err))
				return
			}
			msg, _, fatalErr := bundle.FormatMessage(name, fluent.WithVariable("amount", amount))
			if fatalErr != nil {
				errors = append(errors, fmt.Errorf("bundle.FormatMessage fatal error: %s", fatalErr))
				return
			}

			if msg != scenario.Expected {
				t.Errorf("%s - %+v : expected %s, got %s",
					bundle.PrimaryLocale(), scenario, scenario.Expected, msg)
			}
		})
	}

	return errors
}
