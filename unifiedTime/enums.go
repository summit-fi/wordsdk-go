package unifiedTime

/*
ENUM(
year, month, day, hour, minute, second, millisecond, microsecond
)
*/
type TimeUnit string

/*
ENUM(

	year,
	month,
	day,
	hour,
	minute,
	second,
	millisecond,
	microsecond,
	weekday,
	weekOfYear,
	ordinalDate,
	millisecondsSinceEpoch,
	timeZoneOffsetMilliseconds,

)
*/
type TimeValue string

/*
ENUM(

	before,
	after,
	identical,
	sameToYear,
	sameToMonth,
	sameToDay,
	sameToHour,
	sameToMinute,
	sameToSecond,
	sameToMillisecond,
	sameToMillisecondFromEpoch,
	leapYear,
	inWeekend,

)
*/
type Relation string

/*
ENUM(

	firstMonthDay,
	firstWeekDay,
	lastMonthDay,
	lastWeekDay,
	nextDay,
	previousDay,
	nextMonthFirstDay,
	firstDayOfYear,
	lastDayOfYear,

)
*/
type Anchor string

/*
ENUM( toYear, toMonth, toDay, toHour, toMinute, toSecond, toMillisecond )
*/
type TruncationUnit string

/*
ENUM( currentMonth, previousMonth, pastQuarter, previousYear )
*/
type RangeAnchor string

// Defines which boundaries of the range are inclusive.
//
//  - [closed]    `[start, end]`
//  - [open]      `(start, end)`
//  - [openStart] `(start, end]`
//  - [openEnd]   `[start, end)`
/*
ENUM( closed, open, openStart, openEnd )
*/
type RangeInterval string

/*
ENUM(

	monday,
	tuesday,
	wednesday,
	thursday,
	friday,
	saturday,
	sunday,

)
*/
type Weekday string

func (w Weekday) ToISO() int {
	switch w {
	case "monday":
		return 1
	case "tuesday":
		return 2
	case "wednesday":
		return 3
	case "thursday":
		return 4
	case "friday":
		return 5
	case "saturday":
		return 6
	case "sunday":
		return 7
	default:
		panic("invalid weekday: " + w)
	}

}

/*
ENUM(
day, week, month, year
)
*/
type TimePeriod string
