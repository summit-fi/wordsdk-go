// Function represents a function that builds a Value based on parameters
package fluent

import (
	"testing"

	"golang.org/x/text/language"
)

// Function represents a function that builds a Value based on parameters
func TestBundle_RegisterFunction(t *testing.T) {
	bundle := NewBundle(language.English)

	testCases := []struct {
		name string
		fn   Function
	}{
		{
			name: "NUMBER",
			fn: func(positional []Value, named map[string]Value, params ...string) Value {
				return Number(42)
			},
		},
		{
			name: "UPPERCASE",
			fn: func(positional []Value, named map[string]Value, params ...string) Value {
				return String("UPPERCASE")
			},
		},
	}

	for _, tc := range testCases {
		bundle.RegisterFunction(tc.name, tc.fn)
	}

	if len(bundle.functions) != 2 {
		t.Errorf("Expected 2 functions, got %d", len(bundle.functions))
	}

	if bundle.functions["NUMBER"] == nil {
		t.Errorf("Expected function NUMBER to be registered")
	}

	if bundle.functions["UPPERCASE"] == nil {
		t.Errorf("Expected function UPPERCASE to be registered")
	}
}
