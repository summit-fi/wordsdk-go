package unifiedTime

import (
	"slices"
	"strconv"
	"strings"
	"time"
)

type UnifiedWhenRule struct {
	Start  *UnifiedTime       `json:"start"`
	End    *UnifiedTime       `json:"end"`
	Repeat *UnifiedWhenRepeat `json:"repeat"`
}

func (r UnifiedWhenRepeat) matches(day UnifiedTime) bool {
	for _, w := range r.Weekly {
		n := isoWeekday(day.Time)
		if w.Start.ToISO() <= n && n <= w.End.ToISO() {
			return true
		}
	}
	for _, m := range r.Monthly {
		if m.Contains(day.Time.Day(), day.Get(AnchorLastMonthDay, nil).Time.Day()) {
			return true
		}
	}
	for _, y := range r.Yearly {
		if y.Contains(day.Time.YearDay(), day.Get(AnchorLastDayOfYear, nil).Time.YearDay()) {
			return true
		}
	}
	return len(r.Weekly)+len(r.Monthly)+len(r.Yearly) == 0
}
func (r UnifiedWhenRule) activeDay(day UnifiedTime) bool {
	if !r.Validation().IsValid() {
		return false
	}
	day = ruleDayStart(day, 0)
	if r.Start != nil && day.CompareTo(ruleDayStart(*r.Start, 0)) < 0 || r.End != nil && day.CompareTo(ruleDayStart(*r.End, 0)) > 0 {
		return false
	}
	if r.Repeat == nil {
		return true
	}
	if !r.Repeat.matches(day) {
		return false
	}
	if r.Repeat.Count != nil {
		remaining := *r.Repeat.Count
		for d := ruleDayStart(*r.Start, 0); d.CompareTo(day) < 0; d = ruleDayStart(d, 1) {
			if r.Repeat.matches(d) {
				remaining--
				if remaining == 0 {
					return false
				}
			}
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
	if !anchor.Equal(ruleDayStart(anchor, 0)) && (h < day.Hour() || h == day.Hour() && m < day.Minute()) {
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
func (r UnifiedWhenRule) GetActiveRanges(day UnifiedTime) []UnifiedTimeRange {
	var ranges []UnifiedTimeRange
	if !r.activeDay(day) {
		return ranges
	}
	add := func(a, b UnifiedTime) {
		if r.Start != nil && a.CompareTo(*r.Start) < 0 {
			a = *r.Start
		}
		if r.End != nil && b.CompareTo(*r.End) > 0 {
			b = *r.End
		}
		if a.CompareTo(b) < 0 {
			ranges = append(ranges, NewUnifiedTimeRange(a, b))
		}
	}
	if r.Repeat == nil || len(r.Repeat.Hourly) == 0 {
		end := day.Add(TimeUnitDay, 1)
		if day.Equal(ruleDayStart(day, 0)) {
			end = ruleDayStart(day, 1)
		}
		add(day, end)
		return ranges
	}
	for _, hour := range r.Repeat.Hourly {
		if hour.HasRepeatedHour() && timeZoneOffset(day.Time) <= timeZoneOffset(day.Add(TimeUnitDay, 1).Time) {
			continue
		}
		a, b := ruleHour(hour.TimeStart, day), ruleHour(hour.TimeEnd, day)
		if b.CompareTo(a) < 0 {
			b = ruleHour(hour.TimeEnd, ruleDayStart(day, 1))
		}
		// Dart maintains the wall-clock duration when a window crosses an offset change.
		_, transition := a.Time.ZoneBounds()
		if !transition.IsZero() {
			delta := timeZoneOffset(transition) - timeZoneOffset(transition.Add(-time.Second))
			boundary := transition
			if delta < 0 {
				boundary = boundary.Add(time.Duration(delta) * time.Second)
			}
			if a.Time.Before(boundary) && !b.Time.Before(boundary) {
				b.Time = b.Time.Add(time.Duration(delta) * time.Second)
			}
		}
		add(a, b)
	}
	return ranges
}
func (r UnifiedWhenRule) IsActiveAt(t, dayStart UnifiedTime) bool {
	for _, w := range r.GetActiveRanges(dayStart) {
		if w.IsTimeInRange(t, RangeIntervalOpenEnd) {
			return true
		}
	}
	return false
}
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

func (r UnifiedWhenRule) CopyWith(start, end *UnifiedTime, repeat *UnifiedWhenRepeat) UnifiedWhenRule {
	if start != nil {
		r.Start = start
	}
	if end != nil {
		r.End = end
	}
	if repeat != nil {
		r.Repeat = repeat
	}
	return r
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
