package unifiedTime

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEncodingUnifiedTime(t *testing.T) {

	inputs := []string{
		`2024-07-25T12:00:00Z`,
		`2024-12-31T23:59:59Z`,
		`2024-01-01T00:00:00Z`,
		`2024-01-01T00:00:00.59Z`,
		`2024-01-01T00:00:00.000000001Z`,
		`0001-01-01T17:00:00Z`,
	}

	var builder strings.Builder
	for _, input := range inputs {
		// decode
		var t1 UnifiedTime
		err := t1.UnmarshalGQL(input)
		assert.Nil(t, err, "Failed to unmarshal input %q: %v", input, err)

		// encode
		t1.MarshalGQL(&builder)
		output := builder.String()
		output = output[1 : len(output)-1]
		assert.Equal(t, input, output, "reencoded value doesn't match")

		// reset builder
		builder.Reset()
	}
}
