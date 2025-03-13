package word

import (
	"github.com/summit-fi/wordsdk-go/source"
)

type SDK interface {
	T(lang string, key string) string
	SaveTranslations(data []source.Object) error
	SaveTranslation(lang string, key string, value string) error
	SetLogger(logger Logger)
	Flush() error
}
