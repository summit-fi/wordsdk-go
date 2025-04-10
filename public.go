package word

import (
	"github.com/summit-fi/wordsdk-go/source"
)

type SDK interface {
	T(lang string, key string) string
	TA(lang string, key string, args any) string
	EnableDynamicContent(key string) *DynamicContent
	Dynamic() *DynamicContent
	SaveTranslations(data []source.Object) error
	SaveTranslation(lang string, key string, value string) error
	SetLogger(logger Logger)
	Flush() error
}
