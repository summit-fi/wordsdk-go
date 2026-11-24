package unifiedTime

import (
	"slices"
	"strconv"
	"strings"
	"time"
)

// UnifiedWhenRule describes an active window with optional bounds and repetition.
type UnifiedWhenRule struct {
	Start  *UnifiedTime       `json:"start"`
	End    *UnifiedTime       `json:"end"`
	Repeat *UnifiedWhenRepeat `json:"repeat"`
}

func (r UnifiedWhenRepeat) matches(day UnifiedTime) bool {
	weekday := isoWeekday(day.Time)
	for _, week := range r.Weekly {
		if week.Start.ToISO() <= weekday && weekday <= week.End.ToISO() {
			return true
		}
	}
	for _, month := range r.Monthly {
		if month.Contains(day.Time.Day(), day.Get(AnchorLastMonthDay, nil).Time.Day()) {
			return true
		}
	}
	for _, year := range r.Yearly {
		if year.Contains(day.Time.YearDay(), day.Get(AnchorLastDayOfYear, nil).Time.YearDay()) {
			return true
		}
	}
	return len(r.Weekly)+len(r.Monthly)+len(r.Yearly) == 0
}

// activeDay checks the calendar bounds, repeat pattern, and matching-day count.
func (r UnifiedWhenRule) activeDay(day UnifiedTime) bool {
	if !r.Validation().IsValid() {
		return false
	}
	day = ruleDayStart(day, 0)
	if r.Start != nil && day.CompareTo(ruleDayStart(*r.Start, 0)) < 0 {
		return false
	}
	if r.End != nil && day.CompareTo(ruleDayStart(*r.End, 0)) > 0 {
		return false
	}
	if r.Repeat == nil {
		return true
	}
	if !r.Repeat.matches(day) {
		return false
	}
	if r.Repeat.Count == nil {
		return true
	}

	// The current day is active if the earlier matching days leave a slot for it.
	remaining := *r.Repeat.Count
	for candidate := ruleDayStart(*r.Start, 0); candidate.CompareTo(day) < 0; candidate = ruleDayStart(candidate, 1) {
		if !r.Repeat.matches(candidate) {
			continue
		}
		remaining--
		if remaining == 0 {
			return false
		}
	}
	return true
}

// ruleDayStart resolves midnight in the same location, including a skipped midnight.
func ruleDayStart(day UnifiedTime, offset int) UnifiedTime {
	t := day.Time
	noon := time.Date(t.Year(), t.Month(), t.Day()+offset, 12, 0, 0, 0, t.Location())
	midnight := time.Date(noon.Year(), noon.Month(), noon.Day(), 0, 0, 0, 0, t.Location())
	if midnight.Day() != noon.Day() {
		_, next := midnight.ZoneBounds()
		if !next.IsZero() {
			midnight = next
		}
	}
	return UnifiedTime{Time: midnight}
}

// ruleHour resolves a validated storage time relative to the calendar or business day.
func ruleHour(token string, anchor UnifiedTime) UnifiedTime {
	h, _ := strconv.Atoi(token[:2])
	m, _ := strconv.Atoi(token[3:5])
	day := anchor.Time
	beforeAnchor := h*60+m < day.Hour()*60+day.Minute()
	if beforeAnchor && !anchor.Equal(ruleDayStart(anchor, 0)) {
		day = day.AddDate(0, 0, 1)
	}
	wall := time.Date(day.Year(), day.Month(), day.Day(), h, m, 0, 0, time.UTC)
	local := time.Date(day.Year(), day.Month(), day.Day(), h, m, 0, 0, day.Location())
	// Resolve wall time with the pre-transition offset. A skipped time moves forward.
	before := timeZoneOffset(local.Add(-24 * time.Hour))
	after := timeZoneOffset(local.Add(24 * time.Hour))
	first := wall.Add(-time.Duration(before) * time.Second).In(day.Location())
	second := wall.Add(-time.Duration(after) * time.Second).In(day.Location())
	if before != after && timeZoneOffset(second) == after && (timeZoneOffset(first) != before || strings.HasSuffix(token, "*")) {
		local = second
	} else {
		local = first
	}
	// In a gap neither candidate has the requested wall clock.
	if after > before && timeZoneOffset(first) == after && timeZoneOffset(second) == before {
		local = first
	}
	return UnifiedTime{Time: local}
}

// GetActiveRanges returns the day's matching windows clipped to the rule bounds.
func (r UnifiedWhenRule) GetActiveRanges(day UnifiedTime) []UnifiedTimeRange {
	if !r.activeDay(day) {
		return nil
	}

	var ranges []UnifiedTimeRange
	appendClipped := func(window UnifiedTimeRange) {
		if r.Start != nil && window.Start.CompareTo(*r.Start) < 0 {
			window.Start = *r.Start
		}
		if r.End != nil && window.End.CompareTo(*r.End) > 0 {
			window.End = *r.End
		}
		if window.Start.CompareTo(window.End) < 0 {
			ranges = append(ranges, window)
		}
	}

	tomorrow := day.Add(TimeUnitDay, 1)
	if r.Repeat == nil || len(r.Repeat.Hourly) == 0 {
		if day.Equal(ruleDayStart(day, 0)) {
			tomorrow = ruleDayStart(day, 1)
		}
		appendClipped(NewUnifiedTimeRange(day, tomorrow))
		return ranges
	}

	hasClockRollback := timeZoneOffset(day.Time) > timeZoneOffset(tomorrow.Time)
	for _, hourly := range r.Repeat.Hourly {
		if hourly.HasRepeatedHour() && !hasClockRollback {
			continue
		}
		appendClipped(resolveHourlyWindow(hourly, day))
	}
	return ranges
}

func resolveHourlyWindow(hourly UnifiedTimeRepeatable, day UnifiedTime) UnifiedTimeRange {
	start := ruleHour(hourly.TimeStart, day)
	end := ruleHour(hourly.TimeEnd, day)
	if end.CompareTo(start) < 0 {
		end = ruleHour(hourly.TimeEnd, ruleDayStart(day, 1))
	}

	_, transition := start.Time.ZoneBounds()
	if transition.IsZero() {
		return NewUnifiedTimeRange(start, end)
	}

	// Use the offset change to preserve wall-clock duration across the transition.
	offsetChange := time.Duration(timeZoneOffset(transition)-timeZoneOffset(transition.Add(-time.Second))) * time.Second
	boundary := transition
	if offsetChange < 0 {
		boundary = boundary.Add(offsetChange)
	}
	if start.Time.Before(boundary) && !end.Time.Before(boundary) {
		end.Time = end.Time.Add(offsetChange)
	}
	return NewUnifiedTimeRange(start, end)
}

// IsActiveAt checks whether an instant falls within an active window for dayStart.
func (r UnifiedWhenRule) IsActiveAt(t, dayStart UnifiedTime) bool {
	for _, w := range r.GetActiveRanges(dayStart) {
		if w.IsTimeInRange(t, RangeIntervalOpenEnd) {
			return true
		}
	}
	return false
}

// IsActiveDuring checks the requested coverage for dayStart and defaults to CoverageFull.
func (r UnifiedWhenRule) IsActiveDuring(target UnifiedTimeRange, dayStart UnifiedTime, coverage ...UnifiedWhenRuleCoverage) bool {
	if target.Start.CompareTo(target.End) >= 0 {
		return false
	}
	windows := r.GetActiveRanges(dayStart)
	if len(coverage) > 0 && coverage[0] == CoveragePartial {
		for _, w := range windows {
			if w.IsCollidingWith(target, RangeIntervalOpen) {
				return true
			}
		}
		return false
	}
	return r.covers(target, windows)
}

// covers reports whether adjacent or overlapping windows span the entire target.
func (r UnifiedWhenRule) covers(target UnifiedTimeRange, windows []UnifiedTimeRange) bool {
	sorted := slices.Clone(windows)
	slices.SortFunc(sorted, func(a, b UnifiedTimeRange) int { return a.Start.CompareTo(b.Start) })
	coveredUntil := target.Start
	for _, w := range sorted {
		if w.Start.CompareTo(coveredUntil) > 0 {
			return false
		}
		if w.End.CompareTo(coveredUntil) > 0 {
			coveredUntil = w.End
		}
		if coveredUntil.CompareTo(target.End) >= 0 {
			return true
		}
	}
	return false
}

func (r UnifiedWhenRule) Equal(other UnifiedWhenRule) bool {
	sameTime := func(a, b *UnifiedTime) bool {
		if a == nil || b == nil {
			return a == b
		}
		return a.Equal(*b)
	}
	if !sameTime(r.Start, other.Start) || !sameTime(r.End, other.End) {
		return false
	}
	if r.Repeat == nil || other.Repeat == nil {
		return r.Repeat == other.Repeat
	}
	return r.Repeat.Equal(*other.Repeat)
}
