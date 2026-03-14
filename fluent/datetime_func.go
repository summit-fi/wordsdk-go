package fluent

import (
	"fmt"
	"time"

	"github.com/boltegg/intl"
	"golang.org/x/text/language"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

/*
US:

	        y  2018
	       yM  6/2018
	      yMM  06/2018
	     yMMM  Jun 2018
	    yMMMM  June 2018
	   GyMMMM  June 2018 AD
	     yMdE  Wed, 6/27/2018
	   yMMdEE  Wed, 06/27/2018
	 yMMMdEEE  Wed, Jun 27, 2018
	yMMMdEEEE  Wednesday, Jun 27, 2018
	       Bh  4 at night
	      Bhm  4:23 at night
	     EBhm  Wed, 4:23 at night
	      Hms  04:23:00
	     Hmsv  04:23:00 ET
	  Hmsvvvv  04:23:00 Eastern Time
	     Hmsz  04:23:00 EDT
	  Hmszzzz  04:23:00 Eastern Daylight Time
	 yMMMMEdB  Wed, June 27, 2018 at 4 at night

https://www.unicode.org/reports/tr35/tr35-dates.html#Date_Field_Symbol_Table
*/
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
	PatternEEEd       = "EEEd"
	PatternyMMM       = "yMMM"
)

type DateTimeValue struct {
	Value time.Time
}

func (value *DateTimeValue) String() string {
	return value.Value.Format(time.RFC3339)
}

func DateTimeFunc(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	return &StringValue{}
}

func CLDRDateTimeFunc(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	pattern := named["pattern"].String()
	if pattern == "" {
		return &StringValue{Value: fmt.Sprintf("error: missing pattern")}
	}

	f, err := CLDRDateTimeFormatter(language.BCP47(), pattern)
	if err != nil {
		return &StringValue{Value: fmt.Sprintf("error: %v", err)}
	}

	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func DATETIME(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	if len(named) == 0 || named["pattern"] == nil {
		return &StringValue{Value: "error: missing pattern"}
	}

	pattern := named["pattern"].String()

	f, err := CLDRDateTimeFormatter(language.BCP47(), pattern)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func CLDRDateTimeFormatter(lang language.Tag, template string) (intl.DateTimeFormat, error) {
	return intl.NewDateTimeFormatLayout(lang, template)
}

func MMMMEEEED(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternMMMMEEEEd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}

}

func YMMMMEEEED(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternyMMMMEEEEd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func MMMd(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternMMMd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func YMMMd(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternyMMMd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func JM(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), Patternjm)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func HHMM(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternHHmm)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func MMMED(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternMMMEd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func YMMMED(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternyMMMEd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func JMS(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), Patternjms)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func YMD(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternyMd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func E(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternE)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func Md(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternMd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func YM(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternyM)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func EEEEE(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternEEEEE)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func Y(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), Patterny)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func LLL(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternLLL)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func YMMMM(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternyMMMM)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func MMM(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternMMM)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func MMMMD(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternMMMMd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func YMMMMD(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternyMMMMd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func EEE_D(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternEEEd)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}

	return &StringValue{Value: f.Format(dt.Value)}
}

func YMMM(positional []Value, named map[string]Value, language cldr.Language, params ...string) Value {
	f, err := CLDRDateTimeFormatter(language.BCP47(), PatternyMMM)
	if err != nil {
		return &StringValue{Value: "error: " + err.Error()}
	}
	dt, ok := positional[0].(*DateTimeValue)
	if !ok {
		return &StringValue{Value: "error: invalid datetime value"}
	}
	return &StringValue{Value: f.Format(dt.Value)}
}
