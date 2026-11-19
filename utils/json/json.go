package json

import "encoding/json"

func ParseBytesWithError[T any](data []byte) (T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	return t, err
}

func Bytes[T any](v T) []byte {
	data, _ := json.Marshal(v)
	return data
}
