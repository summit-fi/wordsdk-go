package unifiedTime

import (
	"errors"
	"github.com/summit-fi/wordsdk-go/utils/ternary"

	"io"
	"strconv"
	"strings"
	"time"
)

const (
	daysInWeek = 7
)

// UnifiedTime wraps time.Time and provides additional utilities.
type UnifiedTime struct {
	Time time.Time
}

func Now(loc *time.Location) UnifiedTime {
	if loc == nil {
		loc = time.Local
	}
	return UnifiedTime{Time: time.Now().In(loc)}
}

func NewUnifiedTime(year, month, day, hour, minute, second, millisecond, microsecond int, loc *time.Location) UnifiedTime {
	return createTimeWithDSTCorrection(year, month, day, hour, minute, second, millisecond, microsecond, loc)
}

func (u UnifiedTime) MarshalGQL(w io.Writer) {
	io.WriteString(w, u.Time.In(time.UTC).Format(`"2006-01-02T15:04:05.999999999Z07:00"`))
}

func (u *UnifiedTime) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("UnifiedTime must be a string")
	}

	if len(str) < 10 {
		return errors.New("the string is too small for UnifiedTime field")
	}
	//TODO: check with Oleh
	//if str[0] != '"' || str[len(str)-1] != '"' {
	//	return errors.New("invalid format for TimeUTC field, string must be enclosed in double quotes")
	//}
	//

	date, err := time.Parse("2006-01-02T15:04:05.000Z", str)

	if err != nil {
		return err
	}
	*u = UnifiedTime{Time: date}
	return nil
}

func (u UnifiedTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + u.Time.In(time.UTC).Format("2006-01-02T15:04:05.000Z") + `"`), nil
}

func (u *UnifiedTime) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return errors.New("invalid size of TimeUTC field")
	}

	date, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", string(data[1:len(data)-1]), time.UTC)
	if err != nil {
		return err
	}
	*u = UnifiedTime{Time: date}

	return nil
}

func (u UnifiedTime) String() string {
	return u.Time.Format("2006-01-02T15:04:05.000Z")
}

// ToISO8601UTCString returns the ISO 8601 string representation of this UnifiedTime in UTC timezone,
// including milliseconds and timezone offset (e.g., "2024-07-25T12:00:00.000Z").
func (u UnifiedTime) ToISO8601UTCString() string {
	t := u.UTC().Time
	return t.Format("2006-01-02T15:04:05.000Z07:00")
}

func (u UnifiedTime) UTC() UnifiedTime {
	t := u.Time.In(time.UTC)
	return UnifiedTime{Time: t}
}

func (u UnifiedTime) In(loc *time.Location) UnifiedTime {
	if loc == nil {
		loc = time.Local
	}
	return UnifiedTime{Time: u.Time.In(loc)}
}

func (u UnifiedTime) Add(unit TimeUnit, amount int) UnifiedTime {
	switch unit {
	case TimeUnitYear:
		return u.addCalendarMonths(amount * 12)
	case TimeUnitMonth:
		return u.addCalendarMonths(amount)
	case TimeUnitDay:
		if u.Time.IsZero() {
			return u
		}
		newVal := u.Time.AddDate(0, 0, amount)

		u.Time = newVal
		return u

	case TimeUnitHour:
		return UnifiedTime{Time: u.Time.Add(time.Duration(amount) * time.Hour)}
	case TimeUnitMinute:
		return UnifiedTime{Time: u.Time.Add(time.Duration(amount) * time.Minute)}
	case TimeUnitSecond:
		return UnifiedTime{Time: u.Time.Add(time.Duration(amount) * time.Second)}
	case TimeUnitMillisecond:
		return UnifiedTime{Time: u.Time.Add(time.Duration(amount) * time.Millisecond)}
	case TimeUnitMicrosecond:
		return UnifiedTime{Time: u.Time.Add(time.Duration(amount) * time.Microsecond)}
	default:
		return u
	}
}

func (u UnifiedTime) addCalendarMonths(months int) UnifiedTime {
	// Calculate the target month and year
	totalMonths := (u.Value(TimeValueYear))*12 + (int(u.Value(TimeValueMonth)) - 1) + months
	year := totalMonths / 12
	month := (totalMonths % 12) + 1

	// Get the last day of the target month
	lastDay := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, u.Time.Location()).Day()
	// Adjust the day if the current day exceeds the last day of the target month
	day := u.Time.Day()
	if day > lastDay {
		day = lastDay
	}

	return createTimeWithDSTCorrection(
		year,
		month,
		day,
		u.Value(TimeValueHour),
		u.Value(TimeValueMinute),
		u.Value(TimeValueSecond),
		u.Value(TimeValueMillisecond),
		u.Value(TimeValueMicrosecond),
		u.Time.Location(),
	)
}

func (u UnifiedTime) addCalendarDays(days int) UnifiedTime {
	return createTimeWithDSTCorrection(
		u.Value(TimeValueYear),
		u.Value(TimeValueMonth),
		u.Value(TimeValueDay)+days,
		u.Value(TimeValueHour),
		u.Value(TimeValueMinute),
		u.Value(TimeValueSecond),
		u.Value(TimeValueMillisecond),
		u.Value(TimeValueMicrosecond),
		u.Time.Location(),
	).Add(TimeUnitMillisecond, u.Value(TimeValueTimeZoneOffsetMilliseconds))
}

// Sub subtracts the specified amount of the given TimeUnit from the UnifiedTime and returns UnifiedTime.
func (u UnifiedTime) Sub(unit TimeUnit, amount int) UnifiedTime {
	return u.Add(unit, -amount)
}

func (u UnifiedTime) Truncate(unit TruncationUnit) UnifiedTime {
	t := u.Time
	switch unit {
	case TruncationUnitToYear:
		t = time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	case TruncationUnitToMonth:
		t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	case TruncationUnitToDay:
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		// For sub-day units, subtract the fractional duration from the absolute epoch value
		// rather than reconstructing from wall-clock components. This preserves the UTC offset
		// of the original moment, correctly handling ambiguous DST times (e.g. asterisk times
		// during fall-back: 01:30-0500 → 01:00-0500, not 01:00-0400).
	case TruncationUnitToHour:
		t = t.Add(
			-time.Duration(t.Minute())*time.Minute -
				time.Duration(t.Second())*time.Second -
				time.Duration(t.Nanosecond()),
		)
	case TruncationUnitToMinute:
		t = t.Add(
			-time.Duration(t.Second())*time.Second -
				time.Duration(t.Nanosecond()),
		)
	case TruncationUnitToSecond:
		t = t.Add(
			-time.Duration(t.Nanosecond()),
		)
	}
	return UnifiedTime{Time: t}
}

// Value return ordinal value of the specified TimeValue field (e.g., year, month, day, etc.).
//   - [TimeValueWeekOfYear], it returns the ISO week number (1..53).
//   - [TimeValueOrdinalDate], it returns the day number in year (1..365/366).
func (u UnifiedTime) Value(val TimeValue) int {
	t := u.Time
	switch val {
	case TimeValueYear:
		return t.Year()
	case TimeValueMonth:
		return int(t.Month())
	case TimeValueDay:
		return t.Day()
	case TimeValueHour:
		return t.Hour()
	case TimeValueMinute:
		return t.Minute()
	case TimeValueSecond:
		return t.Second()
	case TimeValueMillisecond:
		return t.Nanosecond() / 1e6
	case TimeValueMicrosecond:
		return (t.Nanosecond() % 1e6) / 1e3
	case TimeValueWeekday:
		return isoWeekday(t)
	case TimeValueWeekOfYear:
		return u.weekOfYear()
	case TimeValueMillisecondsSinceEpoch:
		return int(t.UnixMilli())
	case TimeValueOrdinalDate:
		return u.ordinalDate()
	case TimeValueTimeZoneOffsetMilliseconds:
		_, offset := t.Zone()
		return offset * 1000
	default:
		return 0
	}
}

// Get returns a UnifiedTime corresponding to the specified Anchor relative to the current UnifiedTime.
//
//	Default weekStart is Monday if weekStart is nil.
//
// Example:
//   - [AnchorFirstWeekDay] and WeekStart as nil, it returns the Monday of the current week.
//   - [AnchorLastWeekDay] and WeekStart as Monday, it returns the Sunday of the current week.
//   - [AnchorFirstMonthDay], it returns the first day of the current month.
//     ... and so on for other Anchor values.
func (u UnifiedTime) Get(target Anchor, weekStart *Weekday) UnifiedTime {
	if weekStart == nil {
		defaultWeekStart := WeekdayMonday
		weekStart = &defaultWeekStart
	}
	switch target {
	case AnchorFirstWeekDay:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			int(u.Value(TimeValueMonth)),
			u.Value(TimeValueDay)-(u.Value(TimeValueWeekday)-weekStart.ToISO()+daysInWeek)%daysInWeek,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())
	case AnchorLastWeekDay:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			u.Value(TimeValueMonth),
			u.Value(TimeValueDay)+(daysInWeek-1-(u.Value(TimeValueWeekday)-weekStart.ToISO()))%daysInWeek,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())
	case AnchorFirstMonthDay:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			u.Value(TimeValueMonth),
			1,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())
	case AnchorLastMonthDay:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			u.Value(TimeValueMonth)+1,
			0,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())
	case AnchorNextMonthFirstDay:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			u.Value(TimeValueMonth)+1,
			1,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())
	case AnchorNextDay:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			u.Value(TimeValueMonth),
			u.Value(TimeValueDay)+1,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())
	case AnchorPreviousDay:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			u.Value(TimeValueMonth),
			u.Value(TimeValueDay)-1,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())
	case AnchorFirstDayOfYear:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			1,
			1,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())
	case AnchorLastDayOfYear:
		return createTimeWithDSTCorrection(
			u.Value(TimeValueYear),
			12,
			31,
			0,
			0,
			0,
			0,
			0,
			u.Time.Location())

	default:
		return u
	}
}

func (u UnifiedTime) Is(match Relation, other *UnifiedTime) bool {
	switch match {
	case RelationBefore:
		return other != nil && u.Time.Before(other.Time)
	case RelationAfter:
		return other != nil && u.Time.After(other.Time)
	case RelationIdentical:
		return other != nil && u.Time.Equal(other.Time)
	case RelationSameToMillisecondFromEpoch:
		return other != nil && u.Time.UnixMilli() == other.Time.UnixMilli()
	case RelationSameToYear:
		return other != nil && u.Truncate(TruncationUnitToYear).Equal(other.Truncate(TruncationUnitToYear))
	case RelationSameToMonth:
		return other != nil && u.Truncate(TruncationUnitToMonth).Equal(other.Truncate(TruncationUnitToMonth))
	case RelationSameToDay:
		return other != nil && u.Truncate(TruncationUnitToDay).Equal(other.Truncate(TruncationUnitToDay))
	case RelationSameToHour:
		return other != nil && u.Truncate(TruncationUnitToHour).Equal(other.Truncate(TruncationUnitToHour))
	case RelationSameToMinute:
		return other != nil && u.Truncate(TruncationUnitToMinute).Equal(other.Truncate(TruncationUnitToMinute))
	case RelationSameToSecond:
		return other != nil && u.Truncate(TruncationUnitToSecond).Equal(other.Truncate(TruncationUnitToSecond))
	case RelationSameToMillisecond:
		return other != nil && u.Truncate(TruncationUnitToMillisecond).Equal(other.Truncate(TruncationUnitToMillisecond))
	case RelationLeapYear:
		return isLeapYear(u.Time.Year())
	case RelationInWeekend:
		return u.Time.Weekday() == time.Saturday || u.Time.Weekday() == time.Sunday
	default:
		return false
	}
}

// Diff returns the duration between this UnifiedTime and another UnifiedTime.
func (u UnifiedTime) Diff(other UnifiedTime) time.Duration {
	return u.Time.Sub(other.Time)
}

// CompareTo compares this UnifiedTime with another UnifiedTime:
//   - returns -1 if this is before the other,
//   - returns 1 if this is after the other,
//   - returns 0 if they are equal.
func (u UnifiedTime) CompareTo(other UnifiedTime) int {
	if u.Is(RelationBefore, &other) {
		return -1
	}
	if u.Is(RelationAfter, &other) {
		return 1
	}
	return 0
}

// Parse attempts to parse a date string in various common formats and returns a UnifiedTime.
func (u UnifiedTime) Parse(dateString string, loc *time.Location) (UnifiedTime, error) {
	if loc == nil {
		loc = time.Local
	}

	layous := []string{
		"2006-01-02T15:04:05.000Z",
		time.RFC3339,
	}

	for _, layou := range layous {
		var parsed time.Time
		var err error
		if strings.Contains(layou, "Z07") {
			parsed, err = time.Parse(layou, dateString)
		} else {
			parsed, err = time.Parse(layou, dateString)
			if err == nil {
				parsed = parsed.In(loc)
			}
		}
		if err == nil {
			return UnifiedTime{Time: parsed.In(loc)}, nil
		}
	}

	return UnifiedTime{}, errors.New("unable to parse date string")
}

func (u UnifiedTime) fromUnixMilli(ms int64, loc *time.Location) UnifiedTime {
	if loc == nil {
		loc = time.Local
	}
	t := time.UnixMilli(ms).In(loc)
	return UnifiedTime{Time: t}
}

func (u UnifiedTime) IsAsteriskTime() bool {
	isTransition := isAsteriskTime(u.Time)
	return isTransition
}

func (u UnifiedTime) DurationFromDayStart() time.Duration {
	dayStart := u.Truncate(TruncationUnitToDay)
	return u.Diff(dayStart)
}

func (u UnifiedTime) Equal(other UnifiedTime) bool {
	return u.Time.Equal(other.Time)
}

// Formats this [UnifiedTime] as a storage time string (e.g., "13:00" or "13:00*").
// The asterisk suffix is added when this time falls in an asterisk DST slot.
func (u UnifiedTime) ToStorageTimeString() string {
	hour := padLeft(u.Value(TimeValueHour), 2)
	minute := padLeft(u.Value(TimeValueMinute), 2)

	return ternary.If(u.IsAsteriskTime(), hour+":"+minute+"*", hour+":"+minute)
}

// ordinalDate returns the day number in year (1..365/366).
func (u UnifiedTime) ordinalDate() int {
	offsets := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	month := int(u.Time.Month())
	day := u.Time.Day()
	leapAdjustment := 0
	if isLeapYear(u.Time.Year()) && month > 2 {
		leapAdjustment = 1
	}

	return offsets[month-1] + day + leapAdjustment
}

// weekOfYear returns ISO 8601 week number in range [1..53].
func (u UnifiedTime) weekOfYear() int {
	year := u.Time.Year()
	woy := (u.ordinalDate() - isoWeekday(u.Time) + 10) / 7

	if woy == 0 {
		dec28PrevYear := NewUnifiedTime(year-1, 12, 28, 0, 0, 0, 0, 0, u.Time.Location())
		return dec28PrevYear.weekOfYear()
	}

	if woy == 53 {
		jan1 := NewUnifiedTime(year, 1, 1, 0, 0, 0, 0, 0, u.Time.Location())
		dec31 := NewUnifiedTime(year, 12, 31, 0, 0, 0, 0, 0, u.Time.Location())
		if isoWeekday(jan1.Time) != 4 && isoWeekday(dec31.Time) != 4 {
			return 1
		}
	}

	return woy
}

// Parses a storage time string (e.g. `"03:30"` or `"03:30*"`) relative to [day].
// If it's dst day and the time with the asterisk falls in the repeated hour,
// the returned time will be adjusted to the second instance of that hour.
// Otherwise, the time will be parsed as-is, ignoring the asterisk if present.
func (u UnifiedTime) ParseTimeString(timeStr string, day UnifiedTime) (UnifiedTime, error) {

	var result UnifiedTime

	hasAsterisk := strings.HasSuffix(timeStr, "*")
	if hasAsterisk {
		timeStr = strings.TrimSuffix(timeStr, "*")
	}

	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return UnifiedTime{}, err
	}

	dayHour := day.Value(TimeValueHour)
	dayStartsRightAfterSkippedHour := day.IsTransitionTime() && !day.IsAsteriskTime()
	isSkippedHour := dayStartsRightAfterSkippedHour && parsedTime.Hour() == dayHour
	isOverNight := (parsedTime.Hour() < dayHour && !isSkippedHour) ||
		(parsedTime.Hour() == dayHour && parsedTime.Minute() < day.Value(TimeValueMinute))

	if isOverNight {
		result = day.Add(TimeUnitDay, 1).Truncate(TruncationUnitToDay).Add(TimeUnitHour, parsedTime.Hour()).Add(TimeUnitMinute, parsedTime.Minute())
	} else {
		result = day.Truncate(TruncationUnitToDay).Add(TimeUnitHour, parsedTime.Hour()).Add(TimeUnitMinute, parsedTime.Minute())
	}

	transitions := day.GetTimeTransitions(TimePeriodDay, nil)
	m := transitions.TransitionDuration.Minutes()
	if hasAsterisk && m != 0 {
		// Check if the day of `result` has a transition at all
		springForward := transitions.DST != nil && result.Is(RelationSameToDay, transitions.DST)
		fallBack := transitions.EST != nil && result.Is(RelationSameToDay, transitions.EST)

		if fallBack {
			// Fall-back (overlap): asterisk means the second occurrence → push forward by transition offset
			if result.Is(RelationBefore, transitions.EST) || result.Is(RelationSameToDay, transitions.EST) {
				return result.Add(TimeUnitMinute, int(m)), nil
			}
		} else if springForward {
			// Spring-forward (gap): asterisk means wall-clock time that was skipped → pull back
			if !result.Is(RelationBefore, transitions.DST) {
				return result.Add(TimeUnitMinute, -int(m)), nil
			}
		}

	}
	// For occident hemisphere in autumn, we push forward by transition offset.
	occident, _ := getHemisphere(result.Time)
	if occident {
		return result.Add(TimeUnitMinute, int(m)), nil
	}

	return result, nil
}

// isoWeekday returns the ISO 8601 weekday number (1..7) for the given time, where Monday is 1 and Sunday is 7.
func isoWeekday(t time.Time) int {
	w := int(t.Weekday())
	if w == 0 {
		return 7
	}
	return w
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// createTimeWithDSTCorrection creates a UnifiedTime with the specified components and corrects for DST transitions if necessary.
func createTimeWithDSTCorrection(year, month, day, hour, minute, second, millisecond, microsecond int, loc *time.Location) UnifiedTime {
	t := time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1e6+microsecond*1e3, loc)
	isTransition := isAsteriskTime(t)
	if isTransition {
		prevHour := t.Add(-time.Hour)
		offsetDiff := timeZoneOffset(t) - timeZoneOffset(prevHour) // Calculate the offset difference for the transition
		t = t.Add(-time.Duration(abs(offsetDiff)))                 // Adjust the time by the offset difference
	}
	return UnifiedTime{Time: t}
}

func timeZoneOffset(date time.Time) int {
	_, offset := date.Zone()
	return offset
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
 * Returns true if the given date contains an extra hour due to DST ending.
 */
func isAsteriskTime(date time.Time) bool {
	hourBefore := date.Add(-time.Hour)

	offsetNow := timeZoneOffset(date)
	offsetBefore := timeZoneOffset(hourBefore)

	if offsetNow >= offsetBefore {
		return false
	}

	offsetDiff := time.Duration(offsetBefore-offsetNow) * time.Second
	timeBefore := date.Add(-offsetDiff)

	return timeZoneOffset(date) < timeZoneOffset(timeBefore)
}

func getHemisphere(date time.Time) (occident, orient bool) {

	gap := func() int {
		start, end := date.ZoneBounds()
		_, startOffset := start.Zone()
		_, endOffset := end.Zone()

		gap := endOffset - startOffset

		if startOffset < 0 && endOffset < 0 {
			gap = startOffset - endOffset
		}
		return gap
	}()

	if gap == 0 {
		return false, false
	}

	_, baseOffset := date.Zone()
	hour := date.Add(-time.Duration(gap) * time.Second)
	_, beforeOffset := hour.Zone()

	after := date.Add(time.Duration(gap) * time.Second)
	_, afterOffset := after.Zone()

	if abs(baseOffset) < abs(beforeOffset) {
		return false, true
	}
	if abs(beforeOffset) > abs(afterOffset) {
		return true, false
	}

	return false, false
}

func (u UnifiedTime) IsTransitionTime() bool {
	previousHour := u.Add(TimeUnitHour, -1)

	offNow := u.Value(TimeValueTimeZoneOffsetMilliseconds)
	offPrev := previousHour.Value(TimeValueTimeZoneOffsetMilliseconds)

	if offNow == offPrev {
		return false
	}

	offsetDiff := offNow - offPrev
	if offsetDiff < 0 {
		offsetDiff = -offsetDiff
	}

	offBefore := u.Sub(TimeUnitMillisecond, int(offsetDiff)).Value(TimeValueTimeZoneOffsetMilliseconds)

	return offNow != offBefore
}

func (u UnifiedTime) GetTimeTransitions(period TimePeriod, loc *time.Location) UnifiedTimeTransitions {
	end := UnifiedTime{}
	switch period {
	case TimePeriodDay:
		end = u.Add(TimeUnitDay, 1)
	case TimePeriodWeek:
		end = u.Add(TimeUnitDay, 7)
	case TimePeriodMonth:
		end = u.Get(AnchorLastMonthDay, nil).
			Add(TimeUnitHour, u.Value(TimeValueHour)).
			Add(TimeUnitMinute, u.Value(TimeValueMinute))
	case TimePeriodYear:
		end = u.Get(AnchorLastDayOfYear, nil).
			Add(TimeUnitHour, u.Value(TimeValueHour)).
			Add(TimeUnitMinute, u.Value(TimeValueMinute))

	}
	return u.GetTransitionTimeForRange(NewUnifiedTimeRange(u, end), loc)
}

func (u UnifiedTime) GetTransitionTimeForRange(timeRange UnifiedTimeRange, loc *time.Location) UnifiedTimeTransitions {
	if loc == nil {
		loc = u.Time.Location()
	}
	var edt, est UnifiedTime

	firstStart, firstEnd := timeRange.Start.Time.ZoneBounds()

	firstDateStart := NewUnifiedTime(
		firstStart.Year(),
		int(firstStart.Month()),
		firstStart.Day(),
		firstStart.Hour(),
		firstStart.Minute(),
		firstStart.Second(),
		firstStart.Nanosecond()/1e6,
		(firstStart.Nanosecond()%1e6)/1e3,
		loc)
	firstDateEnd := NewUnifiedTime(
		firstEnd.Year(),
		int(firstEnd.Month()),
		firstEnd.Day(),
		firstEnd.Hour(),
		firstEnd.Minute(),
		firstEnd.Second(),
		firstEnd.Nanosecond()/1e6,
		(firstEnd.Nanosecond()%1e6)/1e3,
		loc)

	if timeZoneOffset(firstStart) > timeZoneOffset(firstEnd) {
		// firstStart has higher offset → firstStart is DST side, firstEnd is EST side
		edt = firstDateStart
		est = firstDateEnd
	} else {
		edt = firstDateEnd
		est = firstDateStart
	}

	return UnifiedTimeTransitions{
		DST:                &edt,
		EST:                &est,
		TransitionDuration: ternary.If(firstStart.IsDST() != firstEnd.IsDST(), time.Duration(abs(timeZoneOffset(firstStart)-timeZoneOffset(firstEnd)))*time.Second, 0),
	}
}

func padLeft(value int, length int) string {
	str := strconv.Itoa(value)
	if len(str) >= length {
		return str
	}
	return strings.Repeat("0", length-len(str)) + str
}
