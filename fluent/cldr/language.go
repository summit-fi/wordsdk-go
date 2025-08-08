package cldr

import "golang.org/x/text/language"

type Language string

const (
	LanguageEnCo Language = "en_CO"
	LanguageEnEu Language = "en_EU"
	LanguageEnUa Language = "en_UA"
	LanguageEnUS Language = "en_US"
	LanguageEsCo Language = "es_CO"
	LanguageUkUa Language = "uk_UA"
	LanguageRuUa Language = "ru_UA"
)

func (l Language) BCP47() language.Tag {
	switch l {
	case LanguageEnUS:
		return language.MustParse("en-US")
	case LanguageEnCo:
		return language.MustParse("en-CO")
	case LanguageEsCo:
		return language.MustParse("es-CO")
	case LanguageEnEu:
		return language.MustParse("en-EU")
	case LanguageEnUa:
		return language.MustParse("en-UA")
	case LanguageRuUa:
		return language.MustParse("ru-UA")
	case LanguageUkUa:
		return language.MustParse("uk-UA")
	default:
		return language.Und
	}
}

func (l Language) GetNumberRules() Numbers {
	switch l {
	case LanguageEnUS:
		return Numbers{}.NumberRules(LanguageEnUS)
	case LanguageUkUa:
		return Numbers{}.NumberRules(LanguageUkUa)
	case LanguageEnCo:
		return Numbers{}.NumberRules(LanguageEnCo)
	case LanguageEsCo:
		return Numbers{}.NumberRules(LanguageEsCo)
	case LanguageEnEu:
		return Numbers{}.NumberRules(LanguageEnEu)
	case LanguageEnUa:
		return Numbers{}.NumberRules(LanguageEnUa)
	case LanguageRuUa:
		return Numbers{}.NumberRules(LanguageRuUa)

	default:
		return Numbers{}
	}
}

func (l Language) OrdinalRules(num int) string {
	switch l {
	case LanguageEnUS, LanguageEnEu, LanguageEnUa, LanguageEnCo:
		if num%10 == 1 && num%100 != 11 {
			return "one"
		}
		if num%10 == 2 && num%100 != 12 {
			return "two"
		}
		if num%10 == 3 && num%100 != 13 {
			return "few"
		}
		return "other"
	case LanguageUkUa:
		if num%10 == 3 && num%100 != 13 {
			return "few"
		}
		return "other"
	case LanguageRuUa:
		return "other" // Russian does not have specific ordinal rules in this context
	case LanguageEsCo:
		return "other" // Spanish does not have specific ordinal rules in this context
	}
	return "other"
}
