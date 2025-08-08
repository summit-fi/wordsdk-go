package word

import (
	"fmt"

	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"github.com/summit-fi/wordsdk-go/source"
)

func (c *Client) T(lang string, key string) string {

	bundle := c.cache.Get(cldr.Language(lang))

	message, errs, err := bundle.FormatMessage(key)
	if err != nil {
		c.logger.Debugf("Failed to format message for key '%s': %v, stack:%v", key, err, errs)
		return ""
	}

	return message
}

func (c *Client) TA(lang string, key string, args any) string {
	// temporary
	bundle := c.cache.Get(cldr.Language(lang))

	var FormatContext []*fluent.FormatContext

	switch args.(type) {
	case map[string]interface{}:
		for k, v := range args.(map[string]interface{}) {
			FormatContext = append(FormatContext, fluent.WithVariable(k, v))
		}
	}

	message, errs, err := bundle.FormatMessage(key, FormatContext...)
	if err != nil {
		c.logger.Debugf("Failed to format message for key '%s': %v, stack:%v", key, err, errs)
		return ""
	}

	return message
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
	return fmt.Errorf("SaveTranslation is deprecated, DynamicContent.SaveTranslation should be used instead")
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
	for _, item := range data {
		bundle := c.cache.Get(cldr.Language(item.LocaleCode))
		if !bundle.HasMessage(item.Key) {
			resource, errs := fluent.NewResource(fmt.Sprintf("%s = %s", item.Key, item.Value))
			if errs != nil {
				c.logger.Errorf("Failed to create resource for language %s: %v", item.LocaleCode, errs)
				continue
			}
			if err := bundle.AddResource(resource); err != nil {
				c.logger.Errorf("Failed to add resource for language %s: %v", item.LocaleCode, err)
				continue
			}
		} else {
			c.logger.Debugf("Key '%s' already exists in the bundle for language '%s'", item.Key, item.LocaleCode)
		}
	}

}

// Flush saves all pending translations to the database.
// This method is useful when SaveStrategy is set to SaveStrategyOnDemand.
func (c *Client) Flush() error {

	//if len(c.saveBundle) == 0 {
	//	return nil
	//}

	//err := c.source.Save(c.saveBundle)
	//if err != nil {
	//	return err
	//}
	//
	//c.saveBundle = nil
	return nil
}
