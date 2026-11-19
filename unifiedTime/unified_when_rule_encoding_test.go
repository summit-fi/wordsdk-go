package unifiedTime

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestWhenRuleEncoding(t *testing.T) {
	input := `{"start":"2026-01-01T00:00:00Z","end":null,"repeat":{"count":1,"hourly":["09:00-10:00","bad"],"weekly":[],"monthly":["end"],"yearly":null,"every":null,"timeZone":"useCenter"}}`
	var r UnifiedWhenRule
	if err := json.Unmarshal([]byte(input), &r); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	r.MarshalGQL(&output)
	if output.String() != input {
		t.Fatal(output.String())
	}
	if err := r.UnmarshalGQL(map[string]any{"repeat": map[string]any{"weekly": []any{nil}}}); err == nil {
		t.Fatal("accepted null range")
	}
	if r.Start == nil || r.Repeat.Monthly[0].String() != "end" {
		t.Fatal("failed decode modified rule")
	}
}
