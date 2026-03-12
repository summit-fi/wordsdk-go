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
			name:        "Test UT_MMMD",
			fnName:      "UT_MMMD",
			messageID:   "mmmd",
			resourceFTL: "mmmd = { UT_MMMD($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "Jun 15",
		},
		{
			name:        "Test UT_YMMMD",
			fnName:      "UT_YMMMD",
			messageID:   "ymmmd",
			resourceFTL: "ymmmd = { UT_YMMMD($date) }",
			time:        UnifiedTime{Time: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			expected:    "2024 Jun15",
		},
	}

	for _, c := range cases {
		bundle := fluent.NewBundle(cldr.LanguageEnUS)
		resource, errs := fluent.NewResource(c.resourceFTL)
		if len(errs) > 0 {
			t.Fatalf("failed to create resource: %v", errs[0])
		}
		bundle.AddResource(resource)

		f := UnifiedTimeFormatFunctions{time: c.time}

		switch c.fnName {
		case "UT_MMMD":
			bundle.RegisterFunction("UT_MMMD", f.MMMd(c.time))
		case "UT_YMMMD":
			bundle.RegisterFunction("UT_YMMMD", f.UT_yMMMd(c.time))
		}

		result, _, _ := bundle.FormatMessage(c.messageID, fluent.WithVariables(map[string]any{"date": c.time}))
		if result != c.expected {
			t.Fatalf("test failed: %s. expected: %s, got: %s", c.name, c.expected, result)
		}
	}
}
