package word

import (
	"reflect"

	"github.com/gotnospirit/messageformat"
	"github.com/summit-fi/wordsdk-go/source"
)

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
	d.logger.Debugf("Translated %s: %s", key, datum)
	return datum
}

func (d *DynamicContent) TA(lang, key string, args any) string {
	datum, err := d.source.LoadOneDynamic(d.dynamicContentAccessKey, lang, key)
	if err != nil {
		d.logger.Errorf("Failed to get dynamic content: %v", err)
		return key
	}

	if len(datum) == 0 {
		d.logger.Debugf("translation '%s' not found", key)
		return key
	}

	parser, err := messageformat.NewWithCulture(lang[:2])
	if err != nil {
		d.logger.Errorf("Failed to create ICU parser: %v", err)
		return key
	}

	icu, err := parser.Parse(datum)
	if err != nil {
		d.logger.Errorf("Failed to translate %s: %v", key, err)
		return key
	}

	var template map[string]interface{}

	switch temp := args.(type) {
	case string:
		return d.T(lang, key)

	case map[string]interface{}:
		template = temp
	default:
		d.logger.Errorf("Unsupported type of args -> key: %s, sent type: %s", key, reflect.TypeOf(datum).String())
		return key
	}

	result, err := icu.FormatMap(template)
	if err != nil {
		d.logger.Errorf("Failed to translate %s: %v", key, err)
		return key
	}

	d.logger.Debugf("Translated %s: %s", key, result)

	return result
}

func (d *DynamicContent) SaveTranslation(lang string, key string, value string) error {
	return d.source.SaveDynamic(d.dynamicContentAccessKey, []source.Object{{LocaleCode: lang, Key: key, Value: value}})
}

func (d *DynamicContent) SaveTranslations(data []source.Object) error {
	err := d.source.SaveDynamic(d.dynamicContentAccessKey, data)
	if err != nil {
		return err
	}
	d.logger.Debugf("Saved %d translations", len(data))
	return nil
}
