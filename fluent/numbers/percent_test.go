package numbers

import (
	"testing"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

func TestPercentFormatter_Format(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		locale   cldr.Numbers
		num      float64
		expected string
	}{
		{"en-US", "#,##0%", cldr.LanguageEnUS.GetNumbers(), 0.127, "13%"},
		{"uk-UA", "#,##0%", cldr.LanguageUkUa.GetNumbers(), 1, "100%"},
		{"es-CO", "#,##0%", cldr.LanguageEsCo.GetNumbers(), 0.12, "12%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := PercentFormatter{
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
