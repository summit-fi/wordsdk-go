package unifiedTime

type UnifiedWhenRuleValidation struct {
	Errors []UnifiedWhenRuleError
}

func (v UnifiedWhenRuleValidation) IsValid() bool { return len(v.Errors) == 0 }

func (r UnifiedWhenRule) Validation() (v UnifiedWhenRuleValidation) {
	if r.Start != nil && r.End != nil && r.Start.CompareTo(*r.End) >= 0 {
		v.Errors = append(v.Errors, StartNotBeforeEnd)
	}
	if r.Repeat != nil {
		if r.Start == nil && r.Repeat.Count != nil {
			v.Errors = append(v.Errors, CountWithoutStart)
		}
		v.Errors = append(v.Errors, r.Repeat.Validation().Errors...)
	}
	return
}

func allValid[T interface{ IsValid() bool }](values []T) bool {
	for _, value := range values {
		if !value.IsValid() {
			return false
		}
	}
	return true
}
func (r UnifiedWhenRepeat) Validation() (v UnifiedWhenRuleValidation) {
	if r.Count != nil && *r.Count <= 0 {
		v.Errors = append(v.Errors, CountNotPositive)
	}
	if !allValid(r.Hourly) || !allValid(r.Weekly) || !allValid(r.Monthly) || !allValid(r.Yearly) {
		v.Errors = append(v.Errors, InvalidValue)
	}
	patterns := 0
	for _, n := range []int{len(r.Weekly), len(r.Monthly), len(r.Yearly)} {
		if n > 0 {
			patterns++
		}
	}
	if patterns > 1 {
		v.Errors = append(v.Errors, MultiplePatterns)
	}
	if r.Every != nil {
		v.Errors = append(v.Errors, EveryNotSupported)
	}
	return
}
