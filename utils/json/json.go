package json

import "encoding/json"

func ParseBytesWithError[T interface{}](data []byte) (T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	return t, err
}
