package word

import (
	"fmt"

	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
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

	return datum
}

func (d *DynamicContent) saveObjects(data []source.Object) error {
	if d.saveStrategy == SaveStrategyImmediate {
		err := d.source.SaveDynamic(d.dynamicContentAccessKey, data)
		if err != nil {
			return err
		}

		// Update local cache
		err = d.UpdateBundle(data)
		if err != nil {
			return err
		}
	} else if d.saveStrategy == SaveStrategyOnDemand {
		d.updateSaveBundleWithData(data)
	} else {
		return fmt.Errorf("unknown save strategy: %v", d.saveStrategy)
	}
	return nil
}

func (d *DynamicContent) SaveTranslation(lang string, key string, value string) error {
	data := []source.Object{{LocaleCode: lang, Key: key, Value: value}}
	if err := d.saveObjects([]source.Object{{LocaleCode: lang, Key: key, Value: value}}); err != nil {
		return err
	}
	for _, datum := range data {
		d.logger.Debugf("SaveTranslations(%s, %s, %s)", datum.LocaleCode, datum.Key, datum.Value)
	}
	return nil
}

func (d *DynamicContent) SaveTranslations(data []source.Object) error {
	err := d.saveObjects(data)
	if err != nil {
		d.logger.Errorf("Failed to save translations: %v", err)
		return err
	}

	for _, datum := range data {
		d.logger.Debugf("SaveTranslations(%s, %s, %s)", datum.LocaleCode, datum.Key, datum.Value)
	}

	return nil
}

func (d *DynamicContent) updateSaveBundleWithData(data []source.Object) {
	for _, item := range data {
		bundle := d.cache.Get(cldr.Language(item.LocaleCode))
		if !bundle.HasMessage(item.Key) {
			d.logger.Debugf("Adding key '%s' for language '%s'", item.Key, item.LocaleCode)
			resource, errs := fluent.NewResource(fmt.Sprintf("%s = %s", item.Key, item.Value))
			if errs != nil {
				d.logger.Errorf("Failed to create resource for language %s: %v", item.LocaleCode, errs)
				continue
			}
			if err := bundle.AddResource(resource); err != nil {
				d.logger.Errorf("Failed to add resource for language %s: %v", item.LocaleCode, err)
				continue
			}
		} else {
			d.logger.Debugf("Key '%s' already exists in the bundle for language '%s'", item.Key, item.LocaleCode)
		}
	}

}
