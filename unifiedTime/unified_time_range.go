package unifiedTime

// UnifiedTimeRange represents a range between two UnifiedTime values.
type UnifiedTimeRange struct {
	Start UnifiedTime
	End   UnifiedTime
}

func NewUnifiedTimeRange(start, end UnifiedTime) UnifiedTimeRange {
	return UnifiedTimeRange{Start: start, End: end}
}

func (r UnifiedTimeRange) String() string {
	return r.Start.String() + " - " + r.End.String()
}

// Duration is calculated as the difference between the end and start times, in milliseconds.
func (r UnifiedTimeRange) Duration() int64 {
	return int64(r.End.Diff(r.Start))
}

// IsTimeInRange checks if a given UnifiedTime falls within the range, based on the specified interval type.
// The interval type determines whether the endpoints of the range are included or excluded in the check:
//
//   - Closed: [start,end] → start <= other <= end (inclusive)
//   - Open: (start,end) → start < other < end (exclusive)
//   - OpenStart: (start,end] → start < other <= end (exclusive start, inclusive end)
//   - OpenEnd: [start,end) → start <= other < end (inclusive start, exclusive end)
func (r UnifiedTimeRange) IsTimeInRange(other UnifiedTime, interval RangeInterval) bool {
	switch interval {
	case RangeIntervalClosed:
		return !other.Is(RelationBefore, &r.Start) && !other.Is(RelationAfter, &r.End)
	case RangeIntervalOpen:
		return other.Is(RelationAfter, &r.Start) && other.Is(RelationBefore, &r.End)
	case RangeIntervalOpenStart:
		return other.Is(RelationAfter, &r.Start) && !other.Is(RelationAfter, &r.End)
	case RangeIntervalOpenEnd:
		return !other.Is(RelationBefore, &r.Start) && other.Is(RelationBefore, &r.End)
	default:
		return false
	}
}

// IsCollidingWith checks if two time ranges overlap.
//
//   - closed: [start1,end1] and [start2,end2] collide if start1 <= end2 and start2 <= end1
//   - open, open start, open end: (start1,end1) and (start2,end2) collide if start1 < end2 and start2 < end1
func (r UnifiedTimeRange) IsCollidingWith(other UnifiedTimeRange, interval RangeInterval) bool {
	switch interval {
	case RangeIntervalClosed:
		return !r.Start.Is(RelationAfter, &other.End) && !other.Start.Is(RelationAfter, &r.End)
	case RangeIntervalOpen, RangeIntervalOpenStart, RangeIntervalOpenEnd:
		return r.Start.Is(RelationBefore, &other.End) && other.Start.Is(RelationBefore, &r.End)
	default:
		return false
	}
}

// GetRange returns a UnifiedTimeRange for the given anchor.
func GetRange(anchor RangeAnchor, now *UnifiedTime) UnifiedTimeRange {
	if now == nil {
		init := Now(nil)
		now = &init
	}
	switch anchor {
	case RangeAnchorCurrentMonth:
		start := now.Truncate(TruncationUnitToMonth)
		end := start.Add(TimeUnitMonth, 1)
		return NewUnifiedTimeRange(start, end)
	case RangeAnchorPreviousMonth:
		end := now.Truncate(TruncationUnitToMonth)
		start := end.Sub(TimeUnitMonth, 1)
		return NewUnifiedTimeRange(start, end)
	case RangeAnchorPastQuarter:
		quarter := ((now.Value(TimeValueMonth) - 1) / 3) - 1
		startMonth := quarter*3 + 1
		start := NewUnifiedTime(now.Value(TimeValueYear), startMonth, 1, 0, 0, 0, 0, 0, now.Time.Location())
		end := start.Add(TimeUnitMonth, 3)

		return NewUnifiedTimeRange(start, end)
	case RangeAnchorPreviousYear:
		end := now.Truncate(TruncationUnitToYear)
		start := end.Sub(TimeUnitYear, 1)
		return NewUnifiedTimeRange(start, end)
	default:
		return NewUnifiedTimeRange(*now, *now)
	}
}
