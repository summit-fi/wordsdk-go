package test

import (
	"encoding/json"
	"fmt"
)

type UTimeTest struct {
	Suite string      `json:"suite"`
	Cases []UTimeCase `json:"cases"`
}
type UTimeCase struct {
	Id          string             `json:"id"`
	Description string             `json:"description"`
	Input       TestTimePoint      `json:"input"`
	Op          uTimeTestOperation `json:"op"`
	Expected    uTimeTestExpected  `json:"expected"`
}

// TestTimePoint identifies a fixture instant and its time zone.
type TestTimePoint struct {
	// UTC is an ISO 8601 string with a Z suffix, e.g. "2026-02-23T12:00:00.000Z".
	UTC string `json:"utc"`
	// TZ is an IANA time zone name, e.g. "Europe/Kyiv", "America/New_York", or "UTC".
	TZ string `json:"tz"`
}

// UnmarshalJSON requires utc and tz to be JSON strings.
func (p *TestTimePoint) UnmarshalJSON(data []byte) error {
	var fields struct {
		UTC *string `json:"utc"`
		TZ  *string `json:"tz"`
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if fields.UTC == nil {
		return fmt.Errorf("test time point requires a string utc field")
	}
	if fields.TZ == nil {
		return fmt.Errorf("test time point requires a string tz field")
	}
	*p = TestTimePoint{UTC: *fields.UTC, TZ: *fields.TZ}
	return nil
}

type uTimeTestOperation struct {
	Type string `json:"type"`

	/// TimeRelation name for isRelation, e.g. "before", "sameToDay", "leapYear"
	Match string `json:"match,omitempty"`

	// Multipurpose unit string, parsed into the appropriate enum in the runner:
	// - value     → TimeValue
	// - truncate  → TruncationUnit
	// - add / sub → TimeUnit
	// - iterator  → TimeUnit
	// - getAnchor → TimeAnchor
	Unit string `json:"unit"`

	// Step size for add / sub / iterator
	Amount int `json:"amount"`

	// Second operand for isRelation / diff / compare
	Input2 *TestTimePoint `json:"input2,omitempty"`

	// Third operand for isCollidesWith
	Input3 *TestTimePoint `json:"input3,omitempty"`

	// Upper bound for iterator
	Till *TestTimePoint `json:"till,omitempty"`

	// Anchor name for getAnchor (TimeAnchor) or range (RangeAnchor)
	Anchor string `json:"anchor,omitempty"`

	// Weekday name for getAnchor(firstWeekDay / lastWeekDay), e.g. "monday", "sunday"
	WeekStart *string `json:"week_start,omitempty"`

	// RangeInterval name for isTimeInRange, e.g. "closed", "open", "openStart", "openEnd"
	Interval *string `json:"interval,omitempty"`

	// Milliseconds since epoch for fromMillisecondsSinceEpoch
	// tz comes from UtTestCase.input.tz
	DurationMs *int64 `json:"ms,omitempty"`

	// Raw string input for parse / tryParse / fromDynamic
	// When present, the operation is called with this string instead of UtTestCase.input
	InputString string `json:"input_string,omitempty"`

	// Time string for parseTimeStringForDate, e.g. "03:30" or "03:30*"
	// The date itself is taken from UtTestCase.input
	TimeString string `json:"time_string,omitempty"`

	// TimePeriod name for getTimeTransitions, e.g. "day", "week", "month", "year"
	Period string `json:"period,omitempty"`

	Locale string `json:"locale,omitempty"`

	DatePattern string `json:"date_pattern,omitempty"`
	/// WhenRepeat in the GraphQL wire format (scalars as strings), e.g.
	/// {"weekly": ["mon-fri"], "hourly": ["09:00-17:00"], "count": 10}
	WhenRepeat map[string]any `json:"repeat,omitempty"`
	/// WhenRule in the GraphQL wire format: {"start": TimeUTC?, "end": TimeUTC?, "repeat": {...}?}.
	/// start/end are parsed in UtTestCase.input.tz.
	WhenRule map[string]any `json:"rule,omitempty"`

	Coverage *string `json:"coverage,omitempty"`

	/// Range to check for whenRuleIsActiveDuring
	Range *TestRange `json:"range,omitempty"`
}

// TestRange contains the two time points bounding a fixture range.
type TestRange struct {
	Start TestTimePoint `json:"start"`
	End   TestTimePoint `json:"end"`
}

// UnmarshalJSON requires both endpoints, as in the Dart fixture model.
func (r *TestRange) UnmarshalJSON(data []byte) error {
	var fields struct {
		Start *TestTimePoint `json:"start"`
		End   *TestTimePoint `json:"end"`
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if fields.Start == nil || fields.End == nil {
		return fmt.Errorf("test range requires start and end")
	}
	*r = TestRange{Start: *fields.Start, End: *fields.End}
	return nil
}

// ExpectedTimeRange holds UTC timestamps from an expected.ranges entry.
type ExpectedTimeRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type uTimeTestExpected struct {
	Ranges          []ExpectedTimeRange `json:"ranges"`
	Utc             string              `json:"utc"`
	Int             int                 `json:"int"`
	DurationMS      int64               `json:"duration_ms"`
	Bool            *bool               `json:"bool"`
	String          string              `json:"string"`
	List            []string            `json:"list"`
	Throws          *bool               `json:"throws"`
	Start           string              `json:"start"`
	End             string              `json:"end"`
	TransitionDurMs int64               `json:"transition_duration_ms"`
	Range           *TestRange          `json:"range,omitempty"`
}
