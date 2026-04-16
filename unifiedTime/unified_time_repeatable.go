package unifiedTime

import (
	"fmt"
	"strings"
)

// UnifiedTimeRepeatable represents a repeatable time range (e.g., for schedules).
type UnifiedTimeRepeatable struct {
	TimeStart  string
	TimeEnd    string
	TimeFormat string
	TimeZone   string
}

func NewUnifiedTimeRepeatable(timeStart, timeEnd, timeFormat, timeZone string) UnifiedTimeRepeatable {
	return UnifiedTimeRepeatable{
		TimeStart:  timeStart,
		TimeEnd:    timeEnd,
		TimeFormat: timeFormat,
		TimeZone:   timeZone,
	}
}

func (utr UnifiedTimeRepeatable) HasRepeatedHour() bool {
	return strings.Contains(utr.TimeStart, "*") || strings.Contains(utr.TimeEnd, "*")
}

// IsVisibleOnDay returns true if the time range should be displayed on the given date.
func (utr UnifiedTimeRepeatable) IsVisibleOnDay(date UnifiedTime, transitions UnifiedTimeTransitions) bool {
	if !utr.HasRepeatedHour() {
		return true
	}
	if transitions.DST == nil {
		return false
	}
	// Show only on DST day
	return date.Equal(*transitions.DST)
}

// UpdateTimeStart returns a copy with updated timeStart.
func (utr UnifiedTimeRepeatable) UpdateTimeStart(newTimeStart string) UnifiedTimeRepeatable {
	utr.TimeStart = newTimeStart
	return utr
}

// UpdateTimeEnd returns a copy with updated timeEnd.
func (utr UnifiedTimeRepeatable) UpdateTimeEnd(newTimeEnd string) UnifiedTimeRepeatable {
	utr.TimeEnd = newTimeEnd
	return utr
}

// GetTimeRangeForDate returns both [TimeStart] and [TimeEnd] as [UnifiedTime] instances for a given [date].
//
// Both times are computed together so that [timeEnd] is always >= [timeStart].
// This correctly handles:
// - DST transitions (via the `*` suffix - wall-clock correction: gap auto-resolved on spring-forward, second occurrence on fall-back)
// - Overnight schedules (timeEnd on the next calendar day)
func (utr UnifiedTimeRepeatable) GetTimeRangeForDate(date UnifiedTime, transitions UnifiedTimeTransitions) (UnifiedTime, UnifiedTime) {
	// Parse timeStart and timeEnd as "HH:mm" or "HH:mm*" relative to date
	start, err := parseTimeStringForDate(utr.TimeStart, date, transitions)
	if err != nil {
		// TODO: handle error
	}
	end, err := parseTimeStringForDate(utr.TimeEnd, date, transitions)

	// If end is before start, move end to next day
	if end.Is(RelationBefore, &start) {
		nextDay := date.Add(TimeUnitDay, 1)
		end, err = parseTimeStringForDate(utr.TimeEnd, nextDay, transitions)
	}
	return start, end
}

// parseTimeStringForDate parses a time string ("HH:mm" or "HH:mm*") for a given date and DST transitions.
func parseTimeStringForDate(timeString string, date UnifiedTime, transitions UnifiedTimeTransitions) (UnifiedTime, error) {
	hasAsterisk := strings.Contains(timeString, "*")
	clean := timeString
	if hasAsterisk {
		clean = strings.ReplaceAll(timeString, "*", "")
	}
	var hour, minute int
	_, err := fmt.Sscanf(clean, "%d:%d", &hour, &minute)
	if err != nil {
		return UnifiedTime{}, fmt.Errorf("invalid time format: %s", timeString)
	}
	midnight := date.Truncate(TruncationUnitToDay)
	t := midnight.Add(TimeUnitHour, hour).Add(TimeUnitMinute, minute)

	ut := NewUnifiedTime(
		date.Value(TimeValueYear),
		date.Value(TimeValueMonth),
		date.Value(TimeValueDay),
		0,
		0,
		0,
		0,
		0,
		date.Time.Location())

	if hasAsterisk {
		// Get the time one hour before the target time
		prevHour := ut.Add(TimeUnitHour, -1)
		// Calculate the offset difference for the transition
		offsetDiff := ut.Value(TimeValueTimeZoneOffsetMilliseconds) - prevHour.Value(TimeValueTimeZoneOffsetMilliseconds)
		// If the offset difference is positive, it means we are in a spring-forward transition (gap), so we need to add the offset difference to get the correct time.
		t = t.Sub(TimeUnitMillisecond, abs(offsetDiff))
	}

	return t, nil
}
