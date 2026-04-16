package cldr

import "testing"

func TestLanguageByCode(t *testing.T) {
	tests := []struct {
		code     string
		expected Language
	}{
		{"en_CO", LanguageEnCo},
		{"es_CO", LanguageEsCo},
		{"en_EU", LanguageEnEu},
		{"en_UA", LanguageEnUa},
		{"en_US", LanguageEnUS},
		{"uk_UA", LanguageUkUa},
		{"ru_UA", LanguageRuUa},
		{"", LanguageEnUS}, // Default case for empty string
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			if got := Language(tt.code); got != tt.expected {
				t.Errorf("Language(%q) = %v, want %v", tt.code, got, tt.expected)
			}
		})
	}
}

func TestLanguage_BCP47(t *testing.T) {
	tests := []struct {
		language Language
		expected string
	}{
		{LanguageEnUS, "en-US"},
		{LanguageEnCo, "en-CO"},
		{LanguageEsCo, "es-CO"},
		{LanguageEnEu, "en-EU"},
		{LanguageEnUa, "en-UA"},
		{LanguageRuUa, "ru-UA"},
		{LanguageUkUa, "uk-UA"},
	}

	for _, tt := range tests {
		t.Run(string(tt.language), func(t *testing.T) {
			if got := tt.language.BCP47().String(); got != tt.expected {
				t.Errorf("BCP47() = %v, want %v", got, tt.expected)
			}
		})
	}
}
