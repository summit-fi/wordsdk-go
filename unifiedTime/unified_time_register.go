package unifiedTime

import "github.com/summit-fi/wordsdk-go/fluent"

type UnifiedTimeFunctionsRegistrar struct {
	Bundle *fluent.Bundle
}

func (u UnifiedTimeFunctionsRegistrar) Register(bundle *fluent.Bundle, f UnifiedTimeFormatFunctions) {
	bundle.RegisterFunction("UT_DATETIME", f.UT_DATETIME)
	//bundle.RegisterFunction("UT_MMMd", f.UT_MMMd)
	//bundle.RegisterFunction("UT_yMMMd", f.UT_yMMMd)
	//bundle.RegisterFunction("UT_jm", f.UT_jm)
}
