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

func GetLanguageTypeByCode(code string) Language {
	if len(code) == 0 {
		return LanguageEnUS // Default to English US if empty
	}

	switch code {
	case "en_CO":
		return LanguageEnCo
	case "es_CO":
		return LanguageEsCo
	case "en_EU":
		return LanguageEnEu
	case "en_UA":
		return LanguageEnUa
	case "en_US":
		return LanguageEnUS
	case "uk_UA":
		return LanguageUkUa
	case "ru_UA":
		return LanguageRuUa
	default:
		return LanguageEnUS // Default to English US if unknown
	}
}

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
