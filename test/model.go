package test

const fixturesPath = "/test/fixtures/custom"

type MainCustomTestModel struct {
	Locale string            `json:"locale"`
	Tests  []CustomTestModel `json:"tests"`
}

type CustomTestModel struct {
	Scenarios []TestScenarios `json:"scenarios"` // Changed from TestScenarios to []TestScenarios
	Type      string          `json:"type"`
}

type TestScenarios struct {
	Expected string `json:"expected"`
	Pattern  string `json:"pattern,omitempty"`
	Value    string `json:"value"`
	Timezone string `json:"timezone,omitempty"`
}
