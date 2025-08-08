package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"golang.org/x/text/language"

	"github.com/summit-fi/wordsdk-go/fluent"
)

func TestCLDRDateTimeFormatter_Format(t *testing.T) {
	type Test struct {
		name     string
		time     time.Time
		pattern  string
		expected string
		locale   language.Tag
	}

	var tests []Test

	for code, model := range ReadScenarios(t) {
		tag := language.MustParse(strings.ReplaceAll(code, "_", "-"))

		for _, jsonTest := range model.Tests {
			if jsonTest.Type != "DateFormatter::Skeleton" {
				continue
			}

			for _, scenario := range jsonTest.Scenarios {
				tInt, err := strconv.ParseInt(scenario.Value, 10, 64)
				if err != nil {
					t.Fatalf("invalid timestamp in %s: %v", code, err)
				}

				name := fmt.Sprintf("%s_%s", code, scenario.Pattern)
				tests = append(tests, Test{
					name:     name,
					pattern:  scenario.Pattern,
					expected: scenario.Expected,
					locale:   tag,
					time:     time.Unix(tInt/1000, 0).In(time.UTC),
				})
			}
		}
	}

	for _, tt := range tests {
		tt := tt // capture loop variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := fluent.CLDRDateTimeFormatter(tt.locale, tt.pattern)
			result := f.Format(tt.time)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
