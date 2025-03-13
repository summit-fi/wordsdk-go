package word

import (
	"fmt"
	"reflect"

	"github.com/gotnospirit/messageformat"
	"github.com/summit-fi/wordsdk-go/source"
)

func (c *Client) T(lang string, key string) string {
	c.localizersLock.RLock()
	defer c.localizersLock.RUnlock()

	datum := c.cache.Get(lang + key)

	if len(datum) == 0 {
		c.logger.Debugf("translation '%s' not found", key)
		return key
	}
	return datum
}

func (c *Client) TA(lang string, key string, args any) string {
	// temporary
	datum := c.cache.Get(lang + key)

	if len(datum) == 0 {
		c.logger.Debugf("translation '%s' not found", key)
		return key
	}

	parser, err := messageformat.NewWithCulture(lang[:2])
	if err != nil {
		c.logger.Errorf("Failed to create ICU parser: %v", err)
		return ""
	}

	icu, err := parser.Parse(datum)
	if err != nil {
		c.logger.Errorf("Failed to translate %s: %v", key, err)
		return key
	}

	var template map[string]interface{}

	switch temp := args.(type) {
	case string:
		return c.T(lang, key)

	case map[string]interface{}:
		template = temp
	default:
		c.logger.Errorf("Unsupported type of args -> key: %s, sent type: %s", key, reflect.TypeOf(datum).String())
	}

	result, err := icu.FormatMap(template)
	if err != nil {
		c.logger.Errorf("Failed to translate %s: %v", key, err)
		return key
	}

	return result
}

func (c *Client) SaveTranslations(data []source.Object) error {
	err := c.saveObjects(data)
	if err != nil {
		return err
	}

	// Update local cache
	err = c.UpdateBundle(data)
	if err != nil {
		return err
	}

	for _, d := range data {
		c.logger.Debugf("SaveTranslations(%s, %s, %s)", d.LocaleCode, d.Key, d.Value)
	}
	return nil
}

func (c *Client) SaveTranslation(lang string, key string, value string) error {
	data := []source.Object{
		{
			LocaleCode: lang,
			Key:        key,
			Value:      value,
		},
	}
	err := c.saveObjects(data)
	if err != nil {
		return err
	}

	// Update local cache
	err = c.UpdateBundle(data)
	if err != nil {
		return err
	}

	c.logger.Debugf("SaveTranslation(%s, %s, %s)", lang, key, value)
	return nil
}

func (c *Client) saveObjects(data []source.Object) error {
	if c.saveStrategy == SaveStrategyImmediate {
		err := c.source.Save(data)
		if err != nil {
			return err
		}
	} else if c.saveStrategy == SaveStrategyOnDemand {
		c.updateSaveBundleWithData(data)
	} else {
		return fmt.Errorf("unknown save strategy: %v", c.saveStrategy)
	}
	return nil
}

func (c *Client) updateSaveBundleWithData(data []source.Object) {
	c.saveBundleLock.Lock()
	defer c.saveBundleLock.Unlock()
	var updatedSaveBundle []source.Object
	updatedSaveBundle = append(updatedSaveBundle, c.saveBundle...)
	for _, d := range data {
		found := false
		for i, s := range updatedSaveBundle {
			if s.LocaleCode == d.LocaleCode && s.Key == d.Key {
				updatedSaveBundle[i] = d
				found = true
				break
			}
		}
		if !found {
			updatedSaveBundle = append(updatedSaveBundle, d)
		}
	}
	c.saveBundle = updatedSaveBundle
}

// Flush saves all pending translations to the database.
// This method is useful when SaveStrategy is set to SaveStrategyOnDemand.
func (c *Client) Flush() error {
	c.saveBundleLock.Lock()
	defer c.saveBundleLock.Unlock()

	if len(c.saveBundle) == 0 {
		return nil
	}

	err := c.source.Save(c.saveBundle)
	if err != nil {
		return err
	}

	c.saveBundle = nil
	return nil
}
