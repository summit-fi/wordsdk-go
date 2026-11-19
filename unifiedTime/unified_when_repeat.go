package unifiedTime

import (
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type UnifiedWhenRepeat struct {
	Count           *int
	Hourly          []UnifiedTimeRepeatable
	Weekly          []UnifiedWeek
	Monthly         []UnifiedMonth
	Yearly          []UnifiedYear
	Every, TimeZone *string
}

var hourlyToken = regexp.MustCompile(`^(([01][0-9]|2[0-3]):[0-5][0-9]|24:00)\*?$`)

func (r UnifiedTimeRepeatable) IsValid() bool {
	return hourlyToken.MatchString(r.TimeStart) && hourlyToken.MatchString(r.TimeEnd)
}
func ParseUnifiedTimeRepeatable(value string) UnifiedTimeRepeatable {
	start, end, found := strings.Cut(value, "-")
	if !found {
		return UnifiedTimeRepeatable{TimeStart: start, TimeFormat: "hhmm", TimeZone: "useCenter", invalidRange: &value}
	}
	return NewUnifiedTimeRepeatable(start, end, "hhmm", "useCenter")
}

// ToRangeString returns the hourly scalar in its wire format.
func (r UnifiedTimeRepeatable) ToRangeString() string {
	if r.invalidRange != nil {
		return *r.invalidRange
	}
	return r.TimeStart + "-" + r.TimeEnd
}

type UnifiedWeek struct {
	Start, End Weekday
	invalid    *string
}

func ParseUnifiedWeek(value string) UnifiedWeek {
	first, last, pair := strings.Cut(value, "-")
	if !pair {
		last = first
	}
	var w UnifiedWeek
	for _, day := range AllWeekdayValues() {
		if string(day)[:3] == first {
			w.Start = day
		}
		if string(day)[:3] == last {
			w.End = day
		}
	}
	if !w.IsValid() {
		return UnifiedWeek{WeekdayMonday, WeekdayMonday, &value}
	}
	return w
}
func (w UnifiedWeek) IsValid() bool {
	days := AllWeekdayValues()
	return w.invalid == nil && slices.Contains(days, w.Start) && slices.Contains(days, w.End) && w.Start.ToISO() <= w.End.ToISO()
}
func (w UnifiedWeek) Days() []Weekday {
	if !w.IsValid() {
		return nil
	}
	return AllWeekdayValues()[w.Start.ToISO()-1 : w.End.ToISO()]
}
func (w UnifiedWeek) String() string {
	if w.invalid != nil {
		return *w.invalid
	}
	if !w.IsValid() {
		return ""
	}
	return rangeString(string(w.Start)[:3], string(w.End)[:3])
}

type calendarRange struct {
	Start, End *int
	invalid    *string
}
type UnifiedMonth calendarRange
type UnifiedYear calendarRange

func parseCalendarRange(value string, max int) calendarRange {
	first, last, pair := strings.Cut(value, "-")
	if !pair {
		last = first
	}
	parse := func(token string) (*int, bool) {
		if token == "end" {
			return nil, true
		}
		if strings.Trim(token, "0123456789") != "" {
			return nil, false
		}
		n, err := strconv.Atoi(token)
		return &n, err == nil && n >= 1 && n <= max
	}
	a, ok := parse(first)
	b, valid := parse(last)
	if !ok || !valid || pair && a == nil || a != nil && b != nil && *a > *b {
		return calendarRange{invalid: &value}
	}
	return calendarRange{Start: a, End: b}
}
func rangeString(a, b string) string {
	if a == b {
		return a
	}
	return a + "-" + b
}
func (r calendarRange) String() string {
	if r.invalid != nil {
		return *r.invalid
	}
	token := func(n *int) string {
		if n == nil {
			return "end"
		}
		return strconv.Itoa(*n)
	}
	return rangeString(token(r.Start), token(r.End))
}
func (r calendarRange) contains(day, last int) bool {
	a, b := last, last
	if r.Start != nil {
		a = *r.Start
	}
	if r.End != nil {
		b = *r.End
	}
	return day >= a && day <= b
}
func ParseUnifiedMonth(value string) UnifiedMonth  { return UnifiedMonth(parseCalendarRange(value, 31)) }
func (r UnifiedMonth) IsValid() bool               { return r.invalid == nil }
func (r UnifiedMonth) Contains(day, last int) bool { return calendarRange(r).contains(day, last) }
func (r UnifiedMonth) String() string              { return calendarRange(r).String() }
func ParseUnifiedYear(value string) UnifiedYear    { return UnifiedYear(parseCalendarRange(value, 366)) }
func (r UnifiedYear) IsValid() bool                { return r.invalid == nil }
func (r UnifiedYear) Contains(day, last int) bool  { return calendarRange(r).contains(day, last) }
func (r UnifiedYear) String() string               { return calendarRange(r).String() }

func (r UnifiedWhenRepeat) CopyWith(overrides UnifiedWhenRepeat) UnifiedWhenRepeat {
	if overrides.Count != nil {
		r.Count = overrides.Count
	}
	if overrides.Hourly != nil {
		r.Hourly = overrides.Hourly
	}
	if overrides.Weekly != nil {
		r.Weekly = overrides.Weekly
	}
	if overrides.Monthly != nil {
		r.Monthly = overrides.Monthly
	}
	if overrides.Yearly != nil {
		r.Yearly = overrides.Yearly
	}
	if overrides.Every != nil {
		r.Every = overrides.Every
	}
	if overrides.TimeZone != nil {
		r.TimeZone = overrides.TimeZone
	}
	return r
}
func (r UnifiedWhenRepeat) Equal(other UnifiedWhenRepeat) bool { return reflect.DeepEqual(r, other) }
