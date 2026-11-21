package unifiedTime

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeekdayISO(t *testing.T) {
	cases := []struct {
		weekday Weekday
		iso     int
	}{
		{WeekdayMonday, 1},
		{WeekdayTuesday, 2},
		{WeekdayWednesday, 3},
		{WeekdayThursday, 4},
		{WeekdayFriday, 5},
		{WeekdaySaturday, 6},
		{WeekdaySunday, 7},
	}

	t.Run("ToISO", func(t *testing.T) {
		for _, tc := range cases {
			t.Run(string(tc.weekday), func(t *testing.T) {
				assert.Equal(t, tc.iso, tc.weekday.ToISO())
			})
		}

		t.Run("invalid", func(t *testing.T) {
			assert.Panics(t, func() { Weekday("invalid").ToISO() })
		})
	})

	t.Run("FromISO", func(t *testing.T) {
		for _, tc := range cases {
			t.Run(strconv.Itoa(tc.iso), func(t *testing.T) {
				var weekday Weekday
				assert.Equal(t, tc.weekday, weekday.FromISO(tc.iso))
			})
		}

		for _, iso := range []int{0, 8} {
			t.Run(strconv.Itoa(iso), func(t *testing.T) {
				var weekday Weekday
				assert.Panics(t, func() { weekday.FromISO(iso) })
			})
		}
	})
}
