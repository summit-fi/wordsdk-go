package unifiedTime

import (
	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"testing"
	"time"
)

func TestUnifiedTimeFunctions(t *testing.T) {
	cases := []struct {
		name        string
		fnName      string
		messageID   string
		resourceFTL string
		time        UnifiedTime
		expected    string
	}{

		{
			name:        "Test MMMD",
			fnName:      "MMMD",
			messageID:   "mmmd",
			resourceFTL: "mmmd = { MMMD($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Jun 15",
		},
		{
			name:        "Test MMMMEEEED",
			fnName:      "MMMMEEEED",
			messageID:   "mmmmeeeed",
			resourceFTL: "mmmmeeeed = { MMMMEEEED($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Saturday, June 15",
		},
		{
			name:        "Test YMMMD",
			fnName:      "YMMMD",
			messageID:   "ymmmd",
			resourceFTL: "ymmmd = { YMMMD($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Jun 15, 2024",
		},
		{
			name:        "Test YMMMMEEEED",
			fnName:      "YMMMMEEEED",
			messageID:   "ymmmmeeeed",
			resourceFTL: "ymmmmeeeed = { YMMMMEEEED($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Saturday, June 15, 2024",
		},
		{
			name:        "Test JM",
			fnName:      "JM",
			messageID:   "jm",
			resourceFTL: "jm = { JM($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "12:00 AM",
		},
		{
			name:        "Test HHMM",
			fnName:      "HHMM",
			messageID:   "hhmm",
			resourceFTL: "hhmm = { HHMM($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "00:00",
		},
		{
			name:        "Test MMMED",
			fnName:      "MMMED",
			messageID:   "mmmed",
			resourceFTL: "mmmed = { MMMED($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Sat, Jun 15",
		},
		{
			name:        "Test YMMMED",
			fnName:      "YMMMED",
			messageID:   "ymmmed",
			resourceFTL: "ymmmed = { YMMMED($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Sat, Jun 15, 2024",
		},
		{
			name:        "Test JMS",
			fnName:      "JMS",
			messageID:   "jms",
			resourceFTL: "jms = { JMS($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "12:00:00 AM",
		},
		{
			name:        "Test YMD",
			fnName:      "YMD",
			messageID:   "ymd",
			resourceFTL: "ymd = { YMD($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "6/15/2024",
		},
		{
			name:        "Test E",
			fnName:      "E",
			messageID:   "e",
			resourceFTL: "e = { E($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Sat",
		},
		{
			name:        "Test MD",
			fnName:      "MD",
			messageID:   "md",
			resourceFTL: "md = { MD($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "6/15",
		},
		{
			name:        "Test YM",
			fnName:      "YM",
			messageID:   "ym",
			resourceFTL: "ym = { YM($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "6/2024",
		},
		{
			name:        "Test EEEEE",
			fnName:      "EEEEE",
			messageID:   "eeeee",
			resourceFTL: "eeeee = { EEEEE($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "S",
		},
		{
			name:        "Test Y",
			fnName:      "Y",
			messageID:   "y",
			resourceFTL: "y = { Y($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "2024",
		},
		{
			name:        "Test LLL",
			fnName:      "LLL",
			messageID:   "lll",
			resourceFTL: "lll = { LLL($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Jun",
		},
		{
			name:        "Test YMMMM",
			fnName:      "YMMMM",
			messageID:   "ymmmm",
			resourceFTL: "ymmmm = { YMMMM($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "June 2024",
		},
		{
			name:        "Test MMM",
			fnName:      "MMM",
			messageID:   "mmm",
			resourceFTL: "mmm = { MMM($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Jun",
		},
		{
			name:        "Test MMMMD",
			fnName:      "MMMMD",
			messageID:   "mmmmd",
			resourceFTL: "mmmmd = { MMMMD($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "June 15",
		},
		{
			name:        "Test YMMMMD",
			fnName:      "YMMMMD",
			messageID:   "ymmmmd",
			resourceFTL: "ymmmmd = { YMMMMD($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "June 15, 2024",
		},
		{
			name:        "Test EEE_D",
			fnName:      "EEE_D",
			messageID:   "eee_d",
			resourceFTL: "eee_d = { EEE_D($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Sat 15",
		},
		{
			name:        "Test YMMM",
			fnName:      "YMMM",
			messageID:   "ymmm",
			resourceFTL: "ymmm = { YMMM($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Jun 2024",
		},
	}

	for _, c := range cases {
		bundle := fluent.NewBundle(cldr.LanguageEnUS)
		resource, errs := fluent.NewResource(c.resourceFTL)
		if len(errs) > 0 {
			t.Fatalf("failed to create resource: %v", errs)
		}
		bundle.AddResource(resource)

		result, _, _ := bundle.FormatMessage(c.messageID, fluent.WithVariables(map[string]any{"date": c.time.Time}))
		if result != c.expected {
			t.Fatalf("test failed: %s. expected: %s, got: %s", c.name, c.expected, result)
		}
	}
}
