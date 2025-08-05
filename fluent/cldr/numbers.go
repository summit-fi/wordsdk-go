package cldr

import (
	"fmt"
)

type Currency struct {
	Code string
	//Narrow string
	Short  string
	Long   string
	Symbol string
}

type Numbers struct {
	DecimalSep           string
	GroupSep             string
	Percent              string
	ZeroDigit            string
	PlusSign             string
	MinusSign            string
	ExpSymbol            string
	Permill              string
	Infinity             string
	NaN                  string
	DecimalPattern       string
	ScientificPattern    string
	PercentPattern       string
	CurrencyPattern      string
	CurrencyCode         string
	SelectedCurrencyCode string
	Currency             map[string]*Currency
	DisplaySymbol        bool // else use code
}

func (n Numbers) NumberRules(l Language) Numbers {
	switch l {
	case LanguageEnUa:
		return Numbers{
			DecimalSep:           ",",
			GroupSep:             " ",
			Percent:              "%",
			ZeroDigit:            "0",
			PlusSign:             "+",
			MinusSign:            "-",
			ExpSymbol:            "E",
			Permill:              "‰",
			Infinity:             "∞",
			NaN:                  "NaN",
			DecimalPattern:       "#,##0.###",
			ScientificPattern:    "#E0",
			PercentPattern:       "#,##0%",
			CurrencyPattern:      "#,##0.00 ¤",
			CurrencyCode:         "UAH",
			SelectedCurrencyCode: "USD",
			DisplaySymbol:        true,
			Currency: map[string]*Currency{
				"USD": {
					Code:   "USD",
					Short:  "$",
					Long:   "US Dollar",
					Symbol: "$",
				},
				"UAH": {
					Code:   "UAH",
					Short:  "uah",
					Long:   "Ukrainian Hryvnia",
					Symbol: "₴",
				},
				"EUR": {
					Code:   "EUR",
					Short:  "EUR",
					Long:   "euro",
					Symbol: "€",
				},
			},
		}
	case LanguageEnEu:
		return Numbers{
			DecimalSep:           ".",
			GroupSep:             ",",
			Percent:              "%",
			ZeroDigit:            "0",
			PlusSign:             "+",
			MinusSign:            "-",
			ExpSymbol:            "E",
			Permill:              "‰",
			Infinity:             "∞",
			NaN:                  "NaN",
			DecimalPattern:       "#,##0.###",
			ScientificPattern:    "#E0",
			PercentPattern:       "#,##0%",
			CurrencyPattern:      "¤#,##0.00",
			CurrencyCode:         "EUR",
			SelectedCurrencyCode: "EUR",
			DisplaySymbol:        true,
			Currency: map[string]*Currency{
				"USD": {
					Code:   "USD",
					Short:  "$",
					Long:   "US Dollar",
					Symbol: "$",
				},
				"UAH": {
					Code:   "UAH",
					Short:  "uah",
					Long:   "Ukrainian Hryvnia",
					Symbol: "₴",
				},
				"EUR": {
					Code:   "EUR",
					Short:  "EUR",
					Long:   "euro",
					Symbol: "€",
				},
			},
		}
	case LanguageEnCo:
		return Numbers{
			DecimalSep:           ",",
			GroupSep:             ".",
			Percent:              "%",
			ZeroDigit:            "0",
			PlusSign:             "+",
			MinusSign:            "-",
			ExpSymbol:            "E",
			Permill:              "‰",
			Infinity:             "∞",
			NaN:                  "NaN",
			DecimalPattern:       "#,##0.###",
			ScientificPattern:    "#E0",
			PercentPattern:       "#,##0%",
			CurrencyPattern:      "¤ #,##0",
			CurrencyCode:         "COP",
			SelectedCurrencyCode: "COP",
			DisplaySymbol:        true,
			Currency: map[string]*Currency{
				"COP": {
					Code:   "COP",
					Short:  "$",
					Long:   "colombian Peso",
					Symbol: "$",
				},
				"USD": {
					Code:   "USD",
					Short:  "USD",
					Long:   "dólar estadounidense",
					Symbol: "$",
				},
			},
		}
	case LanguageUkUa:
		return Numbers{
			DecimalSep:           ",",
			GroupSep:             " ",
			Percent:              "%",
			ZeroDigit:            "0",
			PlusSign:             "+",
			MinusSign:            "-",
			ExpSymbol:            "E",
			Permill:              "‰",
			Infinity:             "∞",
			NaN:                  "NaN",
			DecimalPattern:       "###,##0.###",
			ScientificPattern:    "#E0",
			PercentPattern:       "#,##0%",
			CurrencyPattern:      "#,##0.00 ¤",
			CurrencyCode:         "UAH",
			SelectedCurrencyCode: "UAH",
			DisplaySymbol:        true,
			Currency: map[string]*Currency{
				"UAH": {
					Code:   "UAH",
					Short:  "грн",
					Long:   "українська гривня",
					Symbol: "₴",
				},
				"USD": {
					Code:   "USD",
					Short:  "дол. США",
					Long:   "долар США",
					Symbol: "₴",
				},
				"COP": {
					Code:   "COP",
					Short:  "COP",
					Long:   "колумбійський песо",
					Symbol: "$",
				},
				"EUR": {
					Code:   "EUR",
					Short:  "EUR",
					Long:   "євро",
					Symbol: "€",
				},
			},
		}
	case LanguageEsCo:
		return Numbers{
			DecimalSep:           ",",
			GroupSep:             ".",
			Percent:              "%",
			ZeroDigit:            "0",
			PlusSign:             "+",
			MinusSign:            "-",
			ExpSymbol:            "E",
			Permill:              "‰",
			Infinity:             "∞",
			NaN:                  "NaN",
			DecimalPattern:       "#,##0.###",
			ScientificPattern:    "#E0",
			PercentPattern:       "#,##0%",
			CurrencyPattern:      "¤ #,##0",
			CurrencyCode:         "COP",
			SelectedCurrencyCode: "COP",
			DisplaySymbol:        true,
			Currency: map[string]*Currency{
				"COP": {
					Code:   "COP",
					Short:  "$",
					Long:   "colombian Peso",
					Symbol: "$",
				},
				"USD": {
					Code:   "USD",
					Short:  "USD",
					Long:   "dólar estadounidense",
					Symbol: "$",
				},
			},
		}
	case LanguageRuUa:
		return Numbers{
			DecimalSep:           ",",
			GroupSep:             " ",
			Percent:              "%",
			ZeroDigit:            "0",
			PlusSign:             "+",
			MinusSign:            "-",
			ExpSymbol:            "E",
			Permill:              "‰",
			Infinity:             "∞",
			NaN:                  "NaN",
			DecimalPattern:       "###,##0.###",
			ScientificPattern:    "#E0",
			PercentPattern:       "#,##0%",
			CurrencyPattern:      "#,##0.00 ¤",
			CurrencyCode:         "UAH",
			SelectedCurrencyCode: "UAH",
			DisplaySymbol:        true,
			Currency: map[string]*Currency{
				"UAH": {
					Code:   "UAH",
					Short:  "грн",
					Long:   "українська гривня",
					Symbol: "₴",
				},
				"USD": {
					Code:   "USD",
					Short:  "дол. США",
					Long:   "долар США",
					Symbol: "₴",
				},
				"COP": {
					Code:   "COP",
					Short:  "COP",
					Long:   "колумбійський песо",
					Symbol: "$",
				},
				"EUR": {
					Code:   "EUR",
					Short:  "EUR",
					Long:   "євро",
					Symbol: "€",
				},
			},
		}
	case LanguageEnUS:
		return Numbers{
			DecimalSep:           ".",
			GroupSep:             ",",
			Percent:              "%",
			ZeroDigit:            "0",
			PlusSign:             "+",
			MinusSign:            "-",
			ExpSymbol:            "E",
			Permill:              "‰",
			Infinity:             "∞",
			NaN:                  "NaN",
			DecimalPattern:       "#,##0.###",
			ScientificPattern:    "#E0",
			PercentPattern:       "#,##0%",
			CurrencyPattern:      "¤#,##0.00",
			CurrencyCode:         "USD",
			SelectedCurrencyCode: "USD",
			DisplaySymbol:        true,
			Currency: map[string]*Currency{
				"USD": {
					Code:   "USD",
					Short:  "$",
					Long:   "US Dollar",
					Symbol: "$",
				},
				"UAH": {
					Code:   "UAH",
					Short:  "uah",
					Long:   "Ukrainian Hryvnia",
					Symbol: "₴",
				},
				"COP": {
					Code:   "COP",
					Short:  "COP",
					Long:   "colombian Peso",
					Symbol: "$",
				},
			},
		}
	default:
		return Numbers{}
	}

}

func (n *Numbers) ModifyCurrencySymbol(symbol string) error {
	if n.Currency == nil {
		return fmt.Errorf("currency map is nil")
	}
	if _, exists := n.Currency[n.SelectedCurrencyCode]; !exists {
		return fmt.Errorf("currency code '%s' does not exist", n.SelectedCurrencyCode)
	}
	n.Currency[n.SelectedCurrencyCode].Symbol = symbol
	return nil
}

func (n *Numbers) ModifyCurrencyCode(code string) error {
	if n.Currency == nil {
		return fmt.Errorf("currency map is nil")
	}
	if _, exists := n.Currency[n.SelectedCurrencyCode]; !exists {
		return fmt.Errorf("currency code '%s' does not exist", n.SelectedCurrencyCode)
	}
	n.Currency[n.SelectedCurrencyCode].Code = code
	return nil
}

func (n *Numbers) EnsureCurrencyExists() error {
	if n.Currency == nil {
		return fmt.Errorf("currency map is nil")
	}
	if _, exists := n.Currency[n.SelectedCurrencyCode]; !exists {
		n.Currency[n.SelectedCurrencyCode] = &Currency{
			Code:  n.SelectedCurrencyCode,
			Short: n.SelectedCurrencyCode,
		}
	}
	return nil
}
