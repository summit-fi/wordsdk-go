package test

import (
	"testing"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"github.com/summit-fi/wordsdk-go/fluent/numbers"
)

func TestCurrencyMaximumFractionDigits(t *testing.T) {
	rules := cldr.LanguageEnUS.GetNumberRules()
	formatter := numbers.CurrencyFormatter{
		Pattern: rules.CurrencyPattern,
		Base:    rules,
	}

	tests := []struct {
		value    float64
		maxFrac  int
		expected string
	}{
		{1234.5678, 2, "$1,234.57"},
		{1234.5, 0, "$1,235"},
		{1234.5678, 3, "$1,234.568"},
		{1234.5678, 1, "$1,234.6"},
	}

	for _, tt := range tests {
		got := formatter.Format(tt.value, numbers.MaximumFractionDigits(tt.maxFrac))
		if got != tt.expected {
			t.Errorf("value=%v, maxFrac=%d: got %q, want %q", tt.value, tt.maxFrac, got, tt.expected)
		}
	}
}

func TestCurrencyMinimumFractionDigits(t *testing.T) {
	rules := cldr.LanguageEnUS.GetNumberRules()
	formatter := numbers.CurrencyFormatter{
		Pattern: rules.CurrencyPattern,
		Base:    rules,
	}

	tests := []struct {
		value    float64
		minFrac  int
		expected string
	}{
		{0.0, 1, "$0.0"},
		{1234, 2, "$1,234.00"},
		{1234.5, 2, "$1,234.50"},
		{1234.5678, 3, "$1,234.568"},
		{1234.5, 0, "$1,235"},
		{1234.5678, 1, "$1,234.6"},
	}

	for _, tt := range tests {
		got := formatter.Format(tt.value, numbers.MinimumFractionDigits(tt.minFrac))
		if got != tt.expected {
			t.Errorf("value=%v, minFrac=%d: got %q, want %q", tt.value, tt.minFrac, got, tt.expected)
		}
	}
}
