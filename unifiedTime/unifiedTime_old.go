package unifiedTime

//import (
//	"time"
//
//	word "github.com/summit-fi/wordsdk-go"
//	"github.com/summit-fi/wordsdk-go/fluent"
//	"github.com/summit-fi/wordsdk-go/fluent/cldr"
//)
//
///*
//	TRANSLATION FUNCTIONS
//*/
//
//func (u UnifiedTime) FormatDateTime(word word.SDK, timezone *time.Location, language string) string {
//	return word.TA(language, "date-format-datetime", map[string]interface{}{
//		"date": u.Time.In(timezone),
//	})
//}
//
//func (u UnifiedTime) FormatDate(word word.SDK, timezone *time.Location, language string) string {
//	return word.TA(language, "date-format-date", map[string]interface{}{
//		"date": u.Time.In(timezone),
//	})
//}
//
//func (u UnifiedTime) FormatTime(word word.SDK, timezone *time.Location, language string) string {
//	return word.TA(language, "date-format-time", map[string]interface{}{
//		"date": u.Time.In(timezone),
//	})
//}
//
//func (u UnifiedTime) FormatSkeleton(skeleton string, timezone *time.Location, language string) (string, error) {
//	lang := cldr.Language(language)
//	f, err := fluent.CLDRDateTimeFormatter(lang.BCP47(), skeleton)
//	return f.Format(u.Time.In(timezone)), err
//}
