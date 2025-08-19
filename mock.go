package word

import (
	"github.com/summit-fi/wordsdk-go/source"
)

func NewWordSDKMock(sources map[string]string) (SDK, error) {
	db := source.NewFtl()

	err := db.AddLocaleFiles(sources)
	if err != nil {
		return nil, err
	}

	c := Config{
		Source:       db,
		SaveStrategy: SaveStrategyOnDemand,
	}
	return NewClient(&c)
}
