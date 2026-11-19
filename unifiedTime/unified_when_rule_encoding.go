package unifiedTime

import (
	"encoding/json"
	"fmt"
	"io"
)

// MarshalJSON encodes the rule using the WhenRule field names.
func (r UnifiedWhenRule) MarshalJSON() ([]byte, error) {
	type fields UnifiedWhenRule
	return json.Marshal(fields(r))
}

// UnmarshalJSON replaces the rule after all fields have been decoded.
func (r *UnifiedWhenRule) UnmarshalJSON(data []byte) error {
	type fields UnifiedWhenRule
	var value fields
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = UnifiedWhenRule(value)
	return nil
}

func (r UnifiedWhenRule) MarshalGQL(w io.Writer) { writeGQL(w, r) }
func (r *UnifiedWhenRule) UnmarshalGQL(value any) error {
	switch value.(type) {
	case string:
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return r.UnmarshalJSON(data)
	case []byte:
		return r.UnmarshalJSON(value.([]byte))
	case json.RawMessage:
		return r.UnmarshalJSON(value.(json.RawMessage))
	default:
		return fmt.Errorf("expected type string, []byte, or json.RawMessage for UnifiedWhenRule, got %T", value)
	}
}

func writeGQL(w io.Writer, value any) {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	if _, err = w.Write(data); err != nil {
		panic(err)
	}
}

// rangeToken is a non-null string inside a nullable list of range scalars.
type rangeToken string

func (t *rangeToken) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		return fmt.Errorf("range element must be a string")
	}
	*t = rangeToken(*value)
	return nil
}

type repeatFields struct {
	Count    *int         `json:"count"`
	Hourly   []rangeToken `json:"hourly"`
	Weekly   []rangeToken `json:"weekly"`
	Monthly  []rangeToken `json:"monthly"`
	Yearly   []rangeToken `json:"yearly"`
	Every    *string      `json:"every"`
	TimeZone *string      `json:"timeZone"`
}

func convertRanges[A, B any](values []A, convert func(A) B) []B {
	if values == nil {
		return nil
	}
	out := make([]B, len(values))
	for i, v := range values {
		out[i] = convert(v)
	}
	return out
}

func (r UnifiedWhenRepeat) MarshalJSON() ([]byte, error) {
	return json.Marshal(repeatFields{Count: r.Count, Every: r.Every, TimeZone: r.TimeZone,
		Hourly:  convertRanges(r.Hourly, func(r UnifiedTimeRepeatable) rangeToken { return rangeToken(r.ToRangeString()) }),
		Weekly:  convertRanges(r.Weekly, func(r UnifiedWeek) rangeToken { return rangeToken(r.String()) }),
		Monthly: convertRanges(r.Monthly, func(r UnifiedMonth) rangeToken { return rangeToken(r.String()) }),
		Yearly:  convertRanges(r.Yearly, func(r UnifiedYear) rangeToken { return rangeToken(r.String()) }),
	})
}
func (r *UnifiedWhenRepeat) UnmarshalJSON(data []byte) error {
	var fields repeatFields
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	*r = UnifiedWhenRepeat{Count: fields.Count, Every: fields.Every, TimeZone: fields.TimeZone,
		Hourly:  convertRanges(fields.Hourly, func(v rangeToken) UnifiedTimeRepeatable { return ParseUnifiedTimeRepeatable(string(v)) }),
		Weekly:  convertRanges(fields.Weekly, func(v rangeToken) UnifiedWeek { return ParseUnifiedWeek(string(v)) }),
		Monthly: convertRanges(fields.Monthly, func(v rangeToken) UnifiedMonth { return ParseUnifiedMonth(string(v)) }),
		Yearly:  convertRanges(fields.Yearly, func(v rangeToken) UnifiedYear { return ParseUnifiedYear(string(v)) }),
	}
	return nil
}
func (r UnifiedTimeRepeatable) MarshalGQL(w io.Writer) { writeGQL(w, r.ToRangeString()) }
func (r *UnifiedTimeRepeatable) UnmarshalGQL(value any) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("expected string for UnifiedTimeRepeatable, got %T", value)
	}
	*r = ParseUnifiedTimeRepeatable(s)
	return nil
}
func (r UnifiedWeek) MarshalGQL(w io.Writer) { writeGQL(w, r.String()) }
func (r *UnifiedWeek) UnmarshalGQL(value any) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("expected string for UnifiedWeek, got %T", value)
	}
	*r = ParseUnifiedWeek(s)
	return nil
}
func (r UnifiedMonth) MarshalGQL(w io.Writer) { writeGQL(w, r.String()) }
func (r *UnifiedMonth) UnmarshalGQL(value any) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("expected string for UnifiedMonth, got %T", value)
	}
	*r = ParseUnifiedMonth(s)
	return nil
}
func (r UnifiedYear) MarshalGQL(w io.Writer) { writeGQL(w, r.String()) }
func (r *UnifiedYear) UnmarshalGQL(value any) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("expected string for UnifiedYear, got %T", value)
	}
	*r = ParseUnifiedYear(s)
	return nil
}
