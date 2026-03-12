package ternary

func If[T any](value bool, then, els T) T {
	if value {
		return then
	}
	return els
}
