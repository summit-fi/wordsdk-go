package test

type UTimeTest struct {
	Suite string      `json:"suite"`
	Cases []UTimeCase `json:"cases"`
}
type UTimeCase struct {
	Id          string             `json:"id"`
	Description string             `json:"description"`
	Input       uTimeTestInput     `json:"input"`
	Op          uTimeTestOperation `json:"op"`
	Expected    uTimeTestExpected  `json:"expected"`
}

type uTimeTestInput struct {
	UTC string `json:"utc"`
	TZ  string `json:"tz"`
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
	Input2 *uTimeTestInput `json:"input2,omitempty"`

	// Third operand for isCollidesWith
	Input3 *uTimeTestInput `json:"input3,omitempty"`

	// Upper bound for iterator
	Till *uTimeTestInput `json:"till,omitempty"`

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
}

type uTimeTestExpected struct {
	Utc             string   `json:"utc"`
	Int             int      `json:"int"`
	DurationMS      int64    `json:"duration_ms"`
	Bool            *bool    `json:"bool"`
	String          string   `json:"string"`
	List            []string `json:"list"`
	Throws          *bool    `json:"throws"`
	Start           string   `json:"start"`
	End             string   `json:"end"`
	TransitionDurMs int64    `json:"transition_duration_ms"`
}
