package unifiedTime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	word "github.com/summit-fi/wordsdk-go"
)

func sdk() (word.SDK, error) {
	wsdk, err := word.NewWordSDKMock(
		map[string]string{
			"en_US": word.Root() + "/test/fixtures/custom/custom_data/en_EU.ftl",
			"uk_UA": word.Root() + "/test/fixtures/custom/custom_data/uk_UA.ftl",
			"en_CO": word.Root() + "/test/fixtures/custom/custom_data/en_CO.ftl",
		},
	)
	if err != nil {
		return nil, err
	}
	return wsdk, nil
}
func TestUTime_FormatDateTime(t *testing.T) {
	sdk, err := sdk()
	if err != nil {
		t.Fatalf("Failed to create SDK: %v", err)
		return
	}
	tcases := []struct {
		name     string
		utime    UTime
		language string
		timezone *time.Location
		expected string
	}{
		{
			name:     "UTC to New York",
			utime:    UTime(time.Date(2024, time.July, 4, 15, 0, 0, 0, time.UTC)),
			language: "en_US",
			timezone: time.FixedZone("America/New_York", -4*60*60),
			expected: "July 4, 2024, 11:00 AM",
		},
		{
			name:     "UTC to Kyiv",
			utime:    UTime(time.Date(2024, time.December, 25, 12, 0, 0, 0, time.UTC)),
			language: "uk_UA",
			timezone: time.FixedZone("Europe/Kyiv", 2*60*60),
			expected: "25 грудня 2024\u202fр., 14:00",
		},
		{
			name:     "UTC to Bogotá",
			utime:    UTime(time.Date(2024, time.November, 1, 18, 30, 0, 0, time.UTC)),
			language: "en_CO",
			timezone: time.FixedZone("America/Bogota", -5*60*60),
			expected: "November 1, 2024, 1:30 PM",
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.utime.FormatDateTime(sdk, tc.timezone, tc.language)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestUTime_FormatSkeleton(t *testing.T) {
	tcases := []struct {
		name     string
		utime    UTime
		language string
		timezone *time.Location
		skeleton string
		expected string
	}{
		{
			name:     "Full date in New York",
			utime:    UTime(time.Date(2024, time.July, 4, 15, 0, 0, 0, time.UTC)),
			language: "en_US",
			timezone: time.FixedZone("America/New_York", -4*60*60),
			skeleton: "yMMMMdjm",
			expected: "July 4, 2024 at 11:00 AM",
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.utime.FormatSkeleton(tc.skeleton, tc.timezone, tc.language)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
