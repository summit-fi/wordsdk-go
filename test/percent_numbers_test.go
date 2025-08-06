package test

import (
	"testing"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"github.com/summit-fi/wordsdk-go/fluent/numbers"
)

func TestPercentFormatter_Format(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		locale   cldr.Numbers
		num      float64
		expected string
	}{
		{"en_Co_custom_pattern", "#,##0%", cldr.LanguageEnCo.GetNumberRules(), 0, "0%"},
		{"en_EU_custom_pattern", "#,##0%", cldr.LanguageEnEu.GetNumberRules(), 0.12, "12%"},
		{"en_UA_custom_pattern", "#,##0%", cldr.LanguageEnUa.GetNumberRules(), 0.12123, "12%"},
		{"en-US_custom_pattern", "#,##0%", cldr.LanguageEnUS.GetNumberRules(), 0.127, "13%"},
		{"uk-UA_custom_pattern", "#,##0%", cldr.LanguageUkUa.GetNumberRules(), 1, "100%"},
		{"ru_UA_custom_pattern", "#,##0%", cldr.LanguageRuUa.GetNumberRules(), 1, "100%"},
		{"es-CO_custom_pattern", "#,##0%", cldr.LanguageEsCo.GetNumberRules(), 0.12, "12%"},

		{"en_Co_default_pattern", "", cldr.LanguageEnCo.GetNumberRules(), 0, "0%"},
		{"en_EU_default_pattern", "", cldr.LanguageEnEu.GetNumberRules(), 0.12, "12%"},
		{"en_UA_default_pattern", "", cldr.LanguageEnUa.GetNumberRules(), 0.12123, "12%"},
		{"en-US_default_pattern", "", cldr.LanguageEnUS.GetNumberRules(), 0.127, "13%"},
		{"uk-UA_default_pattern", "", cldr.LanguageUkUa.GetNumberRules(), 1, "100%"},
		{"ru_UA_default_pattern", "", cldr.LanguageRuUa.GetNumberRules(), 1, "100%"},
		{"es-CO_default_pattern", "", cldr.LanguageEsCo.GetNumberRules(), 1, "100%"},

		{"en_Co_negative", "", cldr.LanguageEnCo.GetNumberRules(), -0.12, "-12%"},
		{"en_EU_negative", "", cldr.LanguageEnEu.GetNumberRules(), -0.12123, "-12%"},
		{"en_UA_negative", "", cldr.LanguageEnUa.GetNumberRules(), -0.127, "-13%"},
		{"en-US_negative", "", cldr.LanguageEnUS.GetNumberRules(), -1, "-100%"},
		{"uk-UA_negative", "", cldr.LanguageUkUa.GetNumberRules(), -1, "-100%"},
		{"ru_UA_negative", "", cldr.LanguageRuUa.GetNumberRules(), -1, "-100%"},
		{"es-CO_negative", "", cldr.LanguageEsCo.GetNumberRules(), -0.12, "-12%"},

		{"es-CO_negative_large", "", cldr.LanguageEsCo.GetNumberRules(), -1234567.89, "-123.456.789%"},
		{"en-US_negative_large", "", cldr.LanguageEnUS.GetNumberRules(), -1234567.89, "-123,456,789%"},
		{"en_UA_negative_large", "", cldr.LanguageEnUa.GetNumberRules(), -1234567.89, "-123 456 789%"},
		{"en_EU_negative_large", "", cldr.LanguageEnEu.GetNumberRules(), -1234567.89, "-123,456,789%"},
		{"uk_UA_negative_large", "", cldr.LanguageUkUa.GetNumberRules(), -1234567.89, "-123 456 789%"},
		{"ru_UA_negative_large", "", cldr.LanguageRuUa.GetNumberRules(), -1234567.89, "-123 456 789%"},
		{"en_Co_negative_large", "", cldr.LanguageEnCo.GetNumberRules(), -1234567.89, "-123.456.789%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := numbers.PercentFormatter{
				Pattern: tt.pattern,
				Base:    tt.locale,
			}
			result := f.Format(tt.num)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
