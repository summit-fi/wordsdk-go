package unifiedTime

// UnifiedWhenRuleError describes why a UnifiedWhenRule or UnifiedWhenRepeat is invalid.
type UnifiedWhenRuleError string

const (
	// StartNotBeforeEnd indicates that Start is equal to or later than End.
	StartNotBeforeEnd UnifiedWhenRuleError = "startNotBeforeEnd"

	// CountWithoutStart indicates that Repeat.Count is set without a Start to count from.
	CountWithoutStart UnifiedWhenRuleError = "countWithoutStart"

	// CountNotPositive indicates that Repeat.Count is zero or negative, so the rule can never be active.
	CountNotPositive UnifiedWhenRuleError = "countNotPositive"

	// MultiplePatterns indicates that more than one of Weekly, Monthly, and Yearly is non-empty.
	MultiplePatterns UnifiedWhenRuleError = "multiplePatterns"

	// EveryNotSupported indicates that Every is set. Rules do not yet support this interval.
	EveryNotSupported UnifiedWhenRuleError = "everyNotSupported"

	// InvalidValue indicates that Hourly, Weekly, Monthly, or Yearly contains a value that could not be parsed.
	InvalidValue UnifiedWhenRuleError = "invalidValue"
)
