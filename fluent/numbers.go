package fluent

import (
	"fmt"
	"strconv"
)

// NumberValue wraps a number (float32 at the moment) in order to comply with the Value API
type NumberValue struct {
	Value float32
}

// String formats a NumberValue into a string
func (value *NumberValue) String() string {
	// TODO: Simulate JavaScript number formatting
	return strconv.FormatFloat(float64(value.Value), 'f', -1, 32)
}

// Number returns a new NumberValue with the given value; used for variables
func NumberLiteral(val float32) *NumberValue {
	return &NumberValue{
		Value: val,
	}
}

func NumberFunc(positional []Value, named map[string]Value, params ...string) Value {
	if len(positional) < 1 {
		return &NoValue{value: "func NUMBER: missing number argument"}
	}
	fmt.Println("positional:", positional)
	fmt.Println("named:", named)
	fmt.Println("params:", params)
	numStr := positional[0].String()
	num, err := strconv.ParseFloat(numStr, 32)
	if err != nil {
		return &NoValue{value: fmt.Sprintf("func NUMBER: invalid number format -> %s", numStr)}
	}

	return &NumberValue{Value: float32(num)}
}
