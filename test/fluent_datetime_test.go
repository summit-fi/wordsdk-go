package test

import (
	"testing"
	"time"

	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

func TestCLDRDateTimeFormatter(t *testing.T) {
	ftl := `yMMMMd = full: { CLDRDATETIME($date, pattern: "yMMMMd") }
Hms = duration: { CLDRDATETIME($date, pattern: "Hms") }
yMd = short: { CLDRDATETIME($date, pattern: "yMd") }
Hm = time short: { CLDRDATETIME($date, pattern: "Hm") }
yMMMM = year month: { CLDRDATETIME($date, pattern: "yMMMM") }
Ehms = weekday and time: { CLDRDATETIME($date, pattern: "Ehms") }
htrbfgsdnr = error-date: { CLDRDATETIME($date, pattern: "htrbfgsdnr") }
`

	tests := []struct {
		name     string
		vars     map[string]any
		expected string
	}{
		{
			name: "yMMMMd",
			vars: map[string]any{
				"date": time.Date(2025, time.March, 8, 0, 0, 0, 0, time.UTC),
			},
			expected: "full: 8 de marzo de 2025",
		},
		{
			name: "Hms",
			vars: map[string]any{
				"date": time.Date(2025, time.July, 15, 14, 7, 9, 0, time.UTC),
			},
			expected: "duration: 14:07:09",
		},
		{
			name: "yMd",
			vars: map[string]any{
				"date": time.Date(1999, time.December, 31, 23, 59, 59, 0, time.UTC),
			},
			expected: "short: 31/12/1999",
		},
		{
			name: "Hm",
			vars: map[string]any{
				"date": time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: "time short: 0:00",
		},
		{
			name: "yMMMM",
			vars: map[string]any{
				"date": time.Date(2023, time.June, 5, 12, 0, 0, 0, time.UTC),
			},
			expected: "year month: junio de 2023",
		},
		{
			name: "Ehms",
			vars: map[string]any{
				"date": time.Date(2025, time.February, 17, 21, 45, 5, 0, time.UTC),
			},
			expected: "weekday and time: lun, 9:45:05 p.m.",
		},
		{
			name: "htrbfgsdnr",
			vars: map[string]any{
				"date": time.Date(2025, time.February, 17, 21, 45, 5, 0, time.UTC),
			},
			expected: "error-date: error: unsupported skeleton symbol \"t\"",
		},
	}

	resource, err := fluent.NewResource(ftl)
	if err != nil {
		t.Fatalf("NewResource error: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bundle := fluent.NewBundle(cldr.LanguageEsCo)
			bundle.AddResource(resource)

			msg, _, fatalErr := bundle.FormatMessage(tt.name, fluent.WithVariables(tt.vars))
			if fatalErr != nil {
				t.Fatalf("FormatMessage fatal error: %v", fatalErr)
			}
			if msg != tt.expected {
				t.Errorf("unexpected result:\ngot:      %q\nexpected: %q", msg, tt.expected)
			}
		})
	}
}
