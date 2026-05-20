package fluent_testing

import (
	"path/filepath"
	"testing"

	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"github.com/summit-fi/wordsdk-go/test/fluent-testing/scenario"
	"github.com/summit-fi/wordsdk-go/utils/dir"
)

func Test(t *testing.T) {
	var scenarios = []scenario.Scenario{
		scenario.SimpleScenario,
		scenario.BrowserScenario,
		scenario.PreferencesScenario,
		scenario.EmptyResourceOneLocaleScenario,
		scenario.EmptyResourceAllLocalesScenario,
		scenario.MissingOptionalOneLocaleScenario,
		scenario.MissingOptionalAllLocalesScenario,
		scenario.MissingRequiredOneLocaleScenario,
		scenario.MissingRequiredAllLocalesScenario,
	}

	for _, s := range scenarios {
		t.Run(s.Name, func(t *testing.T) {
			MessageFull(t, s)
		})
	}

}

func MessageFull(t *testing.T, cases scenario.Scenario) *testing.T {
	sources := "./test/fluent-testing/resources"
	bundle := scenario.LoadBundle(filepath.Join(dir.Root(), sources), cases)

	bundle.RegisterFunction("PLATFORM", func(
		_ []fluent.Value,
		_ map[string]fluent.Value,
		_ cldr.Language,
		_ ...string,
	) fluent.Value {
		return fluent.String("linux")
	})
	for _, q := range cases.Queries {
		t.Run(q.Key, func(t *testing.T) {
			msg, errs, err := bundle.FormatFullMessage(q.Key, fluent.WithVariables(q.Args))
			if err != nil {
				t.Fatal(err)
			}
			if len(errs) > 0 {
				t.Fatalf("resolver errors: %v", errs)
			}

			if q.Value != nil {
				if msg.Value == nil {
					t.Fatalf("value is nil, want %q", *q.Value)
				}
				if *msg.Value != *q.Value {
					t.Fatalf("value = %q, want %q", *msg.Value, *q.Value)
				}
			}

			if q.Attrs != nil {
				for attrID, expected := range q.Attrs {
					got, ok := msg.Attributes[attrID]
					if !ok {
						t.Fatalf("missing attribute %q", attrID)
					}
					if got != expected {
						t.Fatalf("attribute %q = %q, want %q", attrID, got, expected)
					}
				}
			}
		})
	}

	return t
}
