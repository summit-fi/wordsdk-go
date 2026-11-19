package unifiedTime

// UnifiedWhenRuleCoverage specifies how much of a time range must be active.
type UnifiedWhenRuleCoverage string

const (
	// CoverageFull requires continuous activity across the entire range.
	CoverageFull UnifiedWhenRuleCoverage = "full"

	// CoveragePartial requires an overlap of positive duration with an active window.
	CoveragePartial UnifiedWhenRuleCoverage = "partial"
)
