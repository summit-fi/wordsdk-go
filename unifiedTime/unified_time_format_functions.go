package unifiedTime

import (
	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

type UnifiedTimeFormatFunctions struct {
	time UnifiedTime
}

const (
	PatternMMMMEEEEd  = "MMMMEEEEd "
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
	pattern := named["pattern"].String()
	f, err := fluent.CLDRDateTimeFormatter(language.BCP47(), pattern)
	if err != nil {
		return &fluent.StringValue{Value: "error: " + err.Error()}
	}
	return &fluent.StringValue{Value: f.Format(u.time.Time)}

}

func (u UnifiedTimeFormatFunctions) MMMd(t UnifiedTime) fluent.Function {
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

//func (u UnifiedTimeFormatFunctions) UT_jm(t UnifiedTime) string    { return t.Format("15:04") }
