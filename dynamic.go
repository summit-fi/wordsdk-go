package word

import "github.com/summit-fi/wordsdk-go/source"

type DynamicContent struct {
	*Client
}

func (d *DynamicContent) T(lang string, key string) string {
	datum, err := d.source.LoadOneDynamic(d.dynamicContentAccessKey, lang, key)
	if err != nil {
		d.logger.Errorf("Failed to get dynamic content: %v", err)
		return key
	}

	if len(datum) == 0 {
		d.logger.Debugf("translation '%s' not found", key)
		return key
	}
	return datum
}

func (d *DynamicContent) SaveTranslation(lang string, key string, value string) error {
	return d.source.SaveDynamic(d.dynamicContentAccessKey, []source.Object{{LocaleCode: lang, Key: key, Value: value}})
}
