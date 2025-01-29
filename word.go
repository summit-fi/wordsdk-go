package wordsdk_go

import (
	"fmt"
	"sync"

	"github.com/summit-fi/wordsdk-go/fluent"
	"golang.org/x/text/language"
)

type Config struct {
}
type localize struct {
	bundles map[language.Tag]*fluent.Bundle
	sync.RWMutex
}
type Word struct {
	localize localize
}

func (w *Word) Translate(key string, ctx ...*fluent.FormatContext) (string, error) {
	w.localize.RLock()
	defer w.localize.RUnlock()
	bundle := w.localize.bundles[language.English]
	if bundle == nil {
		return key, fmt.Errorf("bundle not found")
	}

	msg, _, err := bundle.FormatMessage(key, ctx...)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func InitBundles(source map[language.Tag]string) map[language.Tag]*fluent.Bundle {
	bundles := make(map[language.Tag]*fluent.Bundle)
	for lang, src := range source {
		resource, err := fluent.NewResource(src)
		if err != nil {
			panic(err)
		}
		bundle := fluent.NewBundle(lang)
		errs := bundle.AddResource(resource)
		if errs != nil {
			for _, errt := range errs {
				if errt != nil {
					panic(err)
				}
			}
		}
		bundles[lang] = bundle

	}
	return bundles
}

func New(config Config) WordSDK {
	return &Word{
		localize: localize{
			bundles: InitBundles(
				map[language.Tag]string{
					language.English: mock,
				}),
		}}
}

var mock = `spot =
{ $center ->
*[hotel] Welcome to your { room }
[restaurant] Welcome to your { table }
[tennis]  Welcome to your  { court }
   }.

room = { $count ->
*[one] one room
[other] {$count} rooms
}

table = { $count ->
*[one] one table
[other] {$count} tables
}

court = { $count ->
*[one] one court
[other] {$count} courts
}

cldr-month-narrow = { $index ->
    *[0] J
     [1] F
     [2] M
     [3] A
     [4] M
     [5] J
     [6] J
     [7] A
     [8] S
     [9] O
    [10] N
    [11] D
}
`
