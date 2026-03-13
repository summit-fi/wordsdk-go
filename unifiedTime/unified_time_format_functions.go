package unifiedTime

import (
	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

type UnifiedTimeFormatFunctions struct {
	time UnifiedTime
}

const (
	PatternMMMMEEEEd  = "MMMMEEEEd"
	PatternyMMMMEEEEd = "yMMMMEEEEd"
	PatternMMMd       = "MMMd"
	PatternyMMMd      = "yMMMd"
	Patternjm         = "jm"
	PatternHHmm       = "HHmm"
	PatternMMMEd      = "MMMEd"
	PatternyMMMEd     = "yMMMEd"
	Patternjms        = "jms"
	PatternyMd        = "yMd"
	PatternE          = "E"
	PatternMd         = "Md"
	PatternyM         = "yM"
	PatternEEEEE      = "EEEEE"
	Patterny          = "y"
	PatternLLL        = "LLL"
	PatternyMMMM      = "yMMMM"
	PatternMMM        = "MMM"
	PatternMMMMd      = "MMMMd"
	PatternyMMMMd     = "yMMMMd"
	PatternEEE_d      = "EEE_d"
	PatternyMMM       = "yMMM"
)

func (u UnifiedTimeFormatFunctions) UT_DATETIME(positional []fluent.Value, named map[string]fluent.Value, language cldr.Language, params ...string) fluent.Value {
	if len(named) == 0 || named["pattern"] == nil {
		return &fluent.StringValue{Value: "error: missing pattern"}
	}
	pattern := named["pattern"].String()
	date := named["date"].String()
	t, err := UnifiedTime{}.Parse(date, nil)
	if err != nil {
		return &fluent.StringValue{Value: "error: " + err.Error()}
	}
	f, err := fluent.CLDRDateTimeFormatter(language.BCP47(), pattern)
	if err != nil {
		return &fluent.StringValue{Value: "error: " + err.Error()}
	}
	return &fluent.StringValue{Value: f.Format(t.Time)}

}

func (u UnifiedTimeFormatFunctions) UT_MMMMEEEEd(t UnifiedTime) fluent.Function {
	return func(positional []fluent.Value, named map[string]fluent.Value, language cldr.Language, params ...string) fluent.Value {
		f, err := fluent.CLDRDateTimeFormatter(language.BCP47(), PatternMMMMEEEEd)
		if err != nil {
			return &fluent.StringValue{Value: "error: " + err.Error()}
		}
		return &fluent.StringValue{Value: f.Format(t.Time)}
	}
}

func (u UnifiedTimeFormatFunctions) UT_yMMMMEEEEd(t UnifiedTime) fluent.Function {
	return func(positional []fluent.Value, named map[string]fluent.Value, language cldr.Language, params ...string) fluent.Value {
		f, err := fluent.CLDRDateTimeFormatter(language.BCP47(), PatternyMMMMEEEEd)
		if err != nil {
			return &fluent.StringValue{Value: "error: " + err.Error()}
		}
		return &fluent.StringValue{Value: f.Format(t.Time)}
	}
}

func (u UnifiedTimeFormatFunctions) UT_MMMd(t UnifiedTime) fluent.Function {
	return func(positional []fluent.Value, named map[string]fluent.Value, language cldr.Language, params ...string) fluent.Value {
		f, err := fluent.CLDRDateTimeFormatter(language.BCP47(), PatternMMMd)
		if err != nil {
			return &fluent.StringValue{Value: "error: " + err.Error()}
		}
		return &fluent.StringValue{Value: f.Format(t.Time)}
	}
}

func (u UnifiedTimeFormatFunctions) UT_yMMMd(t UnifiedTime) fluent.Function {
	return func(positional []fluent.Value, named map[string]fluent.Value, language cldr.Language, params ...string) fluent.Value {
		f, err := fluent.CLDRDateTimeFormatter(language.BCP47(), PatternyMMMd)
		if err != nil {
			return &fluent.StringValue{Value: "error: " + err.Error()}
		}
		return &fluent.StringValue{Value: f.Format(t.Time)}
	}
}

func (u UnifiedTimeFormatFunctions) UT_jm(t UnifiedTime) fluent.Function {
	return func(positional []fluent.Value, named map[string]fluent.Value, language cldr.Language, params ...string) fluent.Value {
		f, err := fluent.CLDRDateTimeFormatter(language.BCP47(), Patternjm)
		if err != nil {
			return &fluent.StringValue{Value: "error: " + err.Error()}
		}
		return &fluent.StringValue{Value: f.Format(t.Time)}
	}
}

func (u UnifiedTimeFormatFunctions) UT_HHmm(t UnifiedTime) fluent.Function {
	return func(positional []fluent.Value, named map[string]fluent.Value, language cldr.Language, params ...string) fluent.Value {
		f, err := fluent.CLDRDateTimeFormatter(language.BCP47(), PatternHHmm)
		if err != nil {
			return &fluent.StringValue{Value: "error: " + err.Error()}
		}
		return &fluent.StringValue{Value: f.Format(t.Time)}
	}
}
