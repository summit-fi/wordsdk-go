package unifiedTime

import "time"

// UnifiedTimeTransitions stores DST and EST transition times and duration.
type UnifiedTimeTransitions struct {
	DST                *UnifiedTime
	EST                *UnifiedTime
	TransitionDuration time.Duration
}

func NewUnifiedTimeTransitions(dst, est *UnifiedTime, duration time.Duration) UnifiedTimeTransitions {
	return UnifiedTimeTransitions{
		DST:                dst,
		EST:                est,
		TransitionDuration: duration,
	}
}

func (utt UnifiedTimeTransitions) WithoutTransitions() bool {
	return utt.DST == nil && utt.EST == nil
}
