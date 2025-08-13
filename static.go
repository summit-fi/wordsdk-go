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
		return key
	}

	return message
}

func (c *Client) TA(lang string, key string, args any) string {
	// temporary
	bundle := c.cache.Get(cldr.Language(lang))

	switch args.(type) {
	case map[string]any:
		message, errs, err := bundle.FormatMessage(key, fluent.WithVariables(args.(map[string]any)))
		if err != nil {
			c.logger.Debugf("Failed to format message for key '%s': %v, stack:%v", key, err, errs)
			return key
		}
		return message
	default:
		c.logger.Debugf("TA function expects a map[string]any for args, got %T", args)
		return key
	}

}

func (c *Client) SaveTranslations(data []source.Object) error {
	return fmt.Errorf("use DynamicContent.SaveTranslations() instead of Client.SaveTranslations()")
}

func (c *Client) SaveTranslation(lang string, key string, value string) error {
	return fmt.Errorf("use DynamicContent.SaveTranslation() instead of Client.SaveTranslation()")
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
