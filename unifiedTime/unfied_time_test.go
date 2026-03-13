package unifiedTime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"github.com/summit-fi/wordsdk-go/unifiedTime/test"
	"github.com/summit-fi/wordsdk-go/utils/dir"
	"github.com/summit-fi/wordsdk-go/utils/ternary"

	jsonu "github.com/summit-fi/wordsdk-go/utils/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type UTimeTestSuite struct {
	suite.Suite
	Suites []test.UTimeTest `json:"suites"`
}

const utimeTestsDir = "unifiedTime/test"

func (s *UTimeTestSuite) SetupSuite() {
	files, err := os.ReadDir(filepath.Join(dir.Root(), utimeTestsDir))
	if err != nil {
		s.FailNow(fmt.Sprintf("Failed to read test files: %v", err))
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		read, err := os.ReadFile(filepath.Join(dir.Root(), utimeTestsDir, file.Name()))
		if err != nil {
			s.FailNow(fmt.Sprintf("Failed to read test file %s: %v", file.Name(), err))
		}

		testSuite, err := jsonu.ParseBytesWithError[test.UTimeTest](read)
		if err != nil {
			s.FailNow(fmt.Sprintf("Failed to decode test file %s: %v", file.Name(), err))
		}
		s.Suites = append(s.Suites, testSuite)
	}

}

func TestUnifiedTimeTestSuite(t *testing.T) {
	suite.Run(t, new(UTimeTestSuite))
}

func (s *UTimeTestSuite) Test() {
	for _, ts := range s.Suites {
		ts := ts
		s.Run(fmt.Sprintf("Suite: %s", ts.Suite), func() {
			for _, tc := range ts.Cases {
				tc := tc
				s.Run(fmt.Sprintf("Case: %s", tc.Id), func() {
					s.runCase(tc)
				})
			}
		})
	}
}

func (s *UTimeTestSuite) runCase(tc test.UTimeCase) {
	switch tc.Op.Type {
	case "add":
		s.runAdd(tc)
	case "sub":
		s.runSub(tc)
	case "isRelation":
		s.runIsRelation(tc)
	case "compare":
		s.runCompare(tc)
	case "truncate":
		s.runTruncate(tc)
	case "value":
		s.runValue(tc)
	case "getAnchor":
		s.runGet(tc)
	case "diff":
		s.runDiff(tc)
	case "durationFromDayStart":
		s.runDurationFromDayStart(tc)
	case "toStorageTimeString":
		s.runToStorageTimeString(tc)
	case "isTransition":
		s.runIsTransition(tc)
	case "isAsterisk":
		s.runIsAsterisk(tc)
	case "getTimeTransitions":
		s.runGetTimeTransitions(tc)
	case "iterator":
		s.runIterator(tc)
	case "fromMillisecondsSinceEpoch":
		s.runFromMillisecondsSinceEpoch(tc)
	case "toIso8601String":
		s.runToISO8601UTCString(tc)
	case "parse":
		s.runParse(tc)
	case "fromDynamic":
		s.runParse(tc)
	case "tryParse":
		s.runParse(tc)
	case "range":
		s.runRange(tc)
	case "isCollidesWith":
		s.runIsCollidesWith(tc)
	case "isTimeInRange":
		s.runIsTimeInRange(tc)
	case "parseTimeStringForDate":
		s.runParseTime(tc)
	case "formatUT":
		s.formatUT(tc)
	default:
		s.Failf("unknown op", "unsupported op type: %s", tc.Op.Type)
	}
}

func loadLocation(s *UTimeTestSuite, tz string) *time.Location {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		s.FailNowf("load location", "error loading location %s: %v", tz, err)
	}
	return loc
}

func expectedFormat(tc test.UTimeCase, result any) string {

	input2 := tc.Op.Input2 != nil
	if input2 {
		return fmt.Sprintf("Actual: %+v\nOP: %+v\nInput1: %+v\nInput2: %+v\nExpected: %+v", result, tc.Op, tc.Input, tc.Op.Input2, tc.Expected)
	}
	return fmt.Sprintf("Actual: %+v\nOP: %+v\nInput: %+v\nExpected: %+v", result, tc.Op, tc.Input, tc.Expected)
}

func (s *UTimeTestSuite) runAdd(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	utime = utime.Add(TimeUnit(tc.Op.Unit), tc.Op.Amount)

	s.Equal(
		tc.Expected.Utc,
		utime.ToISO8601UTCString(),
		expectedFormat(tc, utime.ToISO8601UTCString()),
	)
}

func (s *UTimeTestSuite) runSub(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	utime = utime.Sub(TimeUnit(tc.Op.Unit), tc.Op.Amount)

	s.Equal(
		tc.Expected.Utc,
		utime.ToISO8601UTCString(),
		expectedFormat(tc, utime.ToISO8601UTCString()),
	)
}

func (s *UTimeTestSuite) runIsRelation(tc test.UTimeCase) {

	var (
		t1, t2     UnifiedTime
		loc1, loc2 *time.Location
	)

	loc1 = loadLocation(s, tc.Input.TZ)

	utime1, err := t1.Parse(tc.Input.UTC, loc1)
	assert.NoError(s.T(), err, "parse input1 UTC")

	var utime2 *UnifiedTime
	if tc.Op.Input2 != nil {
		loc2 = loadLocation(s, tc.Op.Input2.TZ)

		parseTime2, err := t2.Parse(tc.Op.Input2.UTC, loc2)
		assert.NoError(s.T(), err, "parse input2 UTC")
		utime2 = &parseTime2
	}

	result := utime1.Is(Relation(tc.Op.Match), utime2)

	s.Equal(*tc.Expected.Bool, result, expectedFormat(tc, result))
}

func (s *UTimeTestSuite) runCompare(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1, t2 UnifiedTime
	utime1, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input1 UTC")

	utime2, err := t2.Parse(tc.Op.Input2.UTC, loc)
	assert.NoError(s.T(), err, "parse input2 UTC")

	result := utime1.CompareTo(utime2)

	s.Equal(
		tc.Expected.Int,
		result,
		expectedFormat(tc, result),
	)
}

func (s *UTimeTestSuite) runTruncate(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	utime = utime.Truncate(TruncationUnit(tc.Op.Unit))

	s.Equal(
		tc.Expected.Utc,
		utime.ToISO8601UTCString(),
		expectedFormat(tc, utime.ToISO8601UTCString()),
	)
}

func (s *UTimeTestSuite) runValue(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	result := utime.Value(TimeValue(tc.Op.Unit))

	s.Equal(
		tc.Expected.Int,
		result,
		expectedFormat(tc, result),
	)
}

func (s *UTimeTestSuite) runGet(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	var weekStart *Weekday
	if tc.Op.WeekStart != nil {
		ws := Weekday(*tc.Op.WeekStart)
		weekStart = &ws
	}
	result := utime.Get(Anchor(tc.Op.Anchor), weekStart)

	s.Equal(
		tc.Expected.Utc,
		result.ToISO8601UTCString(),
		expectedFormat(tc, result.ToISO8601UTCString()),
	)
}

func (s *UTimeTestSuite) runDiff(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1, t2 UnifiedTime
	utime1, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input1 UTC")

	utime2, err := t2.Parse(tc.Op.Input2.UTC, loc)
	assert.NoError(s.T(), err, "parse input2 UTC")

	result := utime1.Diff(utime2)

	s.Equal(
		tc.Expected.DurationMS,
		result.Milliseconds(),
		expectedFormat(tc, result),
	)
}

func (s *UTimeTestSuite) runDurationFromDayStart(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	result := utime.DurationFromDayStart()
	s.Equal(
		tc.Expected.DurationMS,
		result.Milliseconds(),
		expectedFormat(tc, result),
	)
}

func (s *UTimeTestSuite) runToStorageTimeString(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	result := utime.ToStorageTimeString()
	s.Equal(
		tc.Expected.String,
		result,
		expectedFormat(tc, result),
	)
}

func (s *UTimeTestSuite) runGetTimeTransitions(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	result := utime.GetTimeTransitions(TimePeriod(tc.Op.Period), loc)

	s.Equal(
		tc.Expected.TransitionDurMs,
		result.TransitionDuration.Milliseconds(),
		expectedFormat(tc, result),
	)
}

func (s *UTimeTestSuite) runIsTransition(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	result := utime.IsTransitionTime()

	s.Equal(
		*tc.Expected.Bool,
		result,
		expectedFormat(tc, result),
	)
}
func (s *UTimeTestSuite) runIsAsterisk(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	result := utime.IsAsteriskTime()

	s.Equal(
		*tc.Expected.Bool,
		result,
		expectedFormat(tc, result),
	)
}

func (s *UTimeTestSuite) runIterator(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	start, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	var (
		locTill *time.Location
		till    UnifiedTime
	)
	if tc.Op.Till != nil {
		locTill = loadLocation(s, tc.Op.Till.TZ)

		till, err = till.Parse(tc.Op.Till.UTC, locTill)
		assert.NoError(s.T(), err, "parse till UTC")
	}

	utime := NewUnifiedTimeIterator(start, till, TimeUnit(tc.Op.Unit), tc.Op.Amount)

	var result []string
	for iter := range utime.IteratorSeq() {
		result = append(result, iter.ToISO8601UTCString())
	}
	if len(result) == 0 {
		result = []string{}
	}

	s.Equal(
		tc.Expected.List,
		result,
		expectedFormat(tc, result),
	)
}

func (s *UTimeTestSuite) runFromMillisecondsSinceEpoch(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime

	utime := t1.fromUnixMilli(*tc.Op.DurationMs, loc)

	s.Equal(
		tc.Expected.Utc,
		utime.ToISO8601UTCString(),
		expectedFormat(tc, utime.ToISO8601UTCString()),
	)
}

func (s *UTimeTestSuite) runToISO8601UTCString(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	s.Equal(
		tc.Expected.Utc,
		utime.ToISO8601UTCString(),
		expectedFormat(tc, utime.ToISO8601UTCString()),
	)
}

func (s *UTimeTestSuite) runParse(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Op.InputString, loc)

	if tc.Expected.Throws != nil && *tc.Expected.Throws {
		s.EqualError(err, "unable to parse date string")
		return
	}
	if tc.Expected.Bool != nil && *tc.Expected.Bool == false {
		s.EqualError(err, "unable to parse date string")
		return
	}
	assert.NoError(s.T(), err, "parse input UTC")
	s.Equal(
		tc.Expected.Utc,
		utime.ToISO8601UTCString(),
		expectedFormat(tc, utime.ToISO8601UTCString()),
	)
}

func (s *UTimeTestSuite) runRange(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	//var t1, t2 UnifiedTime
	//start, err := t1.Parse(tc.Input.UTC, loc)
	//assert.NoError(s.T(), err, "parse input1 UTC")
	//end, err := t2.Parse(tc.Op.Input2.UTC, loc)
	//assert.NoError(s.T(), err, "parse input2 UTC")

	now, err := UnifiedTime{}.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	result := GetRange(RangeAnchor(tc.Op.Anchor), &now)

	if RangeAnchor(tc.Op.Anchor) == RangeAnchorPastQuarter {
		s.Equal(
			[]string{tc.Expected.Start, tc.Expected.End},
			[]string{result.Start.String(), result.End.String()},
			expectedFormat(tc, result),
		)
		return
	}

	s.Equal(
		[]string{tc.Expected.Start, tc.Expected.End},
		[]string{result.Start.ToISO8601UTCString(), result.End.ToISO8601UTCString()},
		expectedFormat(tc, result),
	)

}

func (s *UTimeTestSuite) runIsCollidesWith(tc test.UTimeCase) {

	var (
		t1, t2, t3, t4         UnifiedTime
		loc1, loc2, loc3, loc4 *time.Location
	)

	loc1 = loadLocation(s, tc.Input.TZ)
	range1Start, err := t1.Parse(tc.Input.UTC, loc1)
	assert.NoError(s.T(), err, "parse input1 UTC")

	loc2 = loadLocation(s, tc.Op.Till.TZ)
	range1End, err := t2.Parse(tc.Op.Till.UTC, loc2)
	assert.NoError(s.T(), err, "parse input2 UTC")

	loc3 = loadLocation(s, tc.Op.Input2.TZ)
	range2Start, err := t3.Parse(tc.Op.Input2.UTC, loc3)
	assert.NoError(s.T(), err, "parse input string UTC")

	loc4 = loadLocation(s, tc.Op.Input3.TZ)
	range2End, err := t4.Parse(tc.Op.Input3.UTC, loc4)
	assert.NoError(s.T(), err, "parse input string UTC")

	range1 := NewUnifiedTimeRange(range1Start, range1End)
	range2 := NewUnifiedTimeRange(range2Start, range2End)

	interval := ternary.If(tc.Op.Interval != nil, RangeInterval(*tc.Op.Interval), RangeIntervalClosed)

	result := range1.IsCollidingWith(range2, interval)

	s.Equal(
		*tc.Expected.Bool,
		result,
		expectedFormat(tc, result),
	)

}

func (s *UTimeTestSuite) runIsTimeInRange(tc test.UTimeCase) {

	var (
		t1, t2, t3       UnifiedTime
		loc1, loc2, loc3 *time.Location
	)

	loc1 = loadLocation(s, tc.Input.TZ)
	rangeStart, err := t1.Parse(tc.Input.UTC, loc1)
	assert.NoError(s.T(), err, "parse input1 UTC")

	loc2 = loadLocation(s, tc.Op.Till.TZ)
	rangeEnd, err := t2.Parse(tc.Op.Till.UTC, loc2)
	assert.NoError(s.T(), err, "parse input2 UTC")

	loc3 = loadLocation(s, tc.Op.Input2.TZ)
	otherTime, err := t3.Parse(tc.Op.Input2.UTC, loc3)
	assert.NoError(s.T(), err, "parse input string UTC")

	timeRange := NewUnifiedTimeRange(rangeStart, rangeEnd)

	interval := ternary.If(tc.Op.Interval != nil, RangeInterval(*tc.Op.Interval), RangeIntervalClosed)

	result := timeRange.IsTimeInRange(otherTime, interval)

	s.Equal(
		*tc.Expected.Bool,
		result,
		expectedFormat(tc, result),
	)

}

func (s *UTimeTestSuite) runParseTime(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	date, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	result, err := t1.ParseTimeString(tc.Op.TimeString, date)
	assert.NoError(s.T(), err, "parse time string")
	s.Equal(
		tc.Expected.Utc,
		result.ToISO8601UTCString(),
		expectedFormat(tc, result),
	)

}

func (s *UTimeTestSuite) formatUT(tc test.UTimeCase) {
	loc := loadLocation(s, tc.Input.TZ)

	var t1 UnifiedTime
	utime, err := t1.Parse(tc.Input.UTC, loc)
	assert.NoError(s.T(), err, "parse input UTC")

	bundle := fluent.NewBundle(cldr.Language(tc.Op.Locale))
	msg := fmt.Sprintf("%s = { UT_DATETIME(%s, date) }", "ut_dt", tc.Op.DatePattern)
	resource, errs := fluent.NewResource(msg)
	if len(errs) > 0 {
		s.FailNow(fmt.Sprintf("failed to create resource: %v", errs[0]))
	}

	bundle.AddResource(resource)

	f := UnifiedTimeFormatFunctions{}
	bundle.RegisterFunction("UT_DATETIME", f.UT_DATETIME)

	result, ferrs, err := bundle.FormatMessage("ut_dt",
		fluent.WithVariable("pattern", tc.Op.DatePattern),
		fluent.WithVariable("date", utime))
	if err != nil {
		s.FailNow(fmt.Sprintf("failed to format message: %v", err))
	}

	if len(ferrs) > 0 {
		s.Fail(fmt.Sprintf("failed to format message with errors: %v", ferrs))
		return
	}

	s.Equal(
		tc.Expected.String,
		result,
		expectedFormat(tc, result),
	)
}
