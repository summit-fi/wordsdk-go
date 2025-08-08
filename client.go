package word

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"github.com/summit-fi/wordsdk-go/source"
)

type Config struct {
	Source         source.Source
	UpdateInterval time.Duration
	MaxCacheSizeMB int
	SaveStrategy   SaveStrategy
}

type SaveStrategy int

const (
	SaveStrategyImmediate SaveStrategy = iota
	SaveStrategyOnDemand
)

var (
	ErrNoConfig = errors.New("no config provided")
)

func GetDefaultConfig(apiKey string) *Config {
	return &Config{
		Source: source.NewRemote(
			"https://dev.wordapi.thesumm.it/api/v1",
			apiKey,
		),
		UpdateInterval: 10 * time.Second,
		MaxCacheSizeMB: 256,
	}
}

type Client struct {
	httpClient              *http.Client
	source                  source.Source
	dynamicContentAccessKey string
	logger                  Logger
	checksum                string
	updateInterval          time.Duration
	logLevel                int
	maxCacheSizeMB          int
	cache                   fluent.Map[cldr.Language, *fluent.Bundle]
	saveStrategy            SaveStrategy
}

func NewClient(config *Config) (SDK, error) {
	if config == nil {
		return nil, ErrNoConfig
	}

	var c = Client{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		source: config.Source,
		logger: &DefaultLogger{
			LogLevel: LogLevelError,
		},
		updateInterval: config.UpdateInterval,
		maxCacheSizeMB: config.MaxCacheSizeMB,
		saveStrategy:   config.SaveStrategy,
	}

	data, checksum, err := config.Source.LoadAllStatic("")
	if err != nil {
		return nil, fmt.Errorf("failed to load translations: %v", err)
	}

	c.cache = fluent.NewMap[cldr.Language, *fluent.Bundle]()

	c.checksum = checksum
	err = c.UpdateBundle(data)
	if err != nil {
		return nil, err
	}

	c.runSyncTranslationsJob()
	return &c, nil
}

func (c *Client) EnableDynamicContent(key string) *DynamicContent {
	cli := c
	cli.dynamicContentAccessKey = key
	return &DynamicContent{
		Client: cli,
	}
}
func (c *Client) Dynamic() *DynamicContent {
	return &DynamicContent{
		Client: c,
	}
}

func (c *Client) runSyncTranslationsJob() {
	c.syncTranslations()
	go func() {
		for {
			time.Sleep(c.updateInterval)
			c.syncTranslations()
		}
	}()
}

func (c *Client) syncTranslations() {
	data, checksum, err := c.source.LoadAllStatic(c.checksum)
	if err != nil {
		c.logger.Errorf("Failed to sync translations: %v", err)
		return
	}

	if checksum == c.checksum {
		c.logger.Infof("Translations are up to date")
		return
	}

	c.checksum = checksum

	err = c.UpdateBundle(data)
	if err != nil {
		c.logger.Errorf("Failed to update bundle: %v", err)
		return
	}

	sizeMB := float64(c.GetCacheSize()) / 1024 / 1024
	c.logger.Infof("Translations synced, count: %d, size %f MB", len(data), sizeMB)
	//c.logger.Infof("Translations synced, count: %d", len(data))

	if sizeMB > float64(c.maxCacheSizeMB) {
		c.logger.Errorf("Warning! Cache size exceeded! %f MB", sizeMB)
	}
}

func (c *Client) UpdateBundle(data []source.Object) error {

	var mapData = make(map[string]*strings.Builder)

	for _, item := range data {

		if _, ok := mapData[item.LocaleCode]; !ok {
			mapData[item.LocaleCode] = &strings.Builder{}
		}
		mapData[item.LocaleCode].WriteString(item.Key)
		mapData[item.LocaleCode].WriteString("=")
		mapData[item.LocaleCode].WriteString(item.Value)
		mapData[item.LocaleCode].WriteString("\n")
	}

	for lang, sb := range mapData {
		if len(sb.String()) == 0 {
			continue
		}
		bundle := fluent.NewBundle(cldr.Language(lang))

		resource, errs := fluent.NewResource(sb.String())
		if errs != nil {
			return fmt.Errorf("failed to create resource for language %s: %v", lang, errs)
		}
		if err := bundle.AddResource(resource); err != nil {
			return fmt.Errorf("failed to add resource for language %s: %v", lang, err)
		}
		c.cache.Set(cldr.Language(lang), bundle)
	}

	return nil
}

func (c *Client) GetCacheSize() int {

	size := 0
	for _, bundle := range c.cache.GetAll() {
		size += int(reflect.TypeOf(bundle).Size())
	}
	return size
}
