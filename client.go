package word

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/summit-fi/wordsdk-go/source"

	"golang.org/x/text/language"
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
		Source:         source.NewRemote("https://dev.wordapi.thesumm.it/api/v1", apiKey),
		UpdateInterval: 10 * time.Second,
		MaxCacheSizeMB: 256,
	}
}

type Client struct {
	httpClient     *http.Client
	source         source.Source
	logger         Logger
	checksum       string
	updateInterval time.Duration
	logLevel       int
	maxCacheSizeMB int
	localizers     map[string]*i18n.Localizer // localCode -> localizer
	bundles        map[string]*i18n.Bundle    // localCode -> bundle
	localizersLock sync.RWMutex
	saveStrategy   SaveStrategy
	saveBundle     []source.Object
	saveBundleLock sync.Mutex
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
		localizers:     make(map[string]*i18n.Localizer),
		bundles:        make(map[string]*i18n.Bundle),
		saveStrategy:   config.SaveStrategy,
	}

	data, checksum, err := config.Source.LoadAll("")
	if err != nil {
		return nil, fmt.Errorf("failed to load translations: %v", err)
	}

	c.checksum = checksum
	err = c.UpdateBundle(data)
	if err != nil {
		return nil, err
	}

	c.runSyncTranslationsJob()
	return &c, nil
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
	data, checksum, err := c.source.LoadAll(c.checksum)
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
	sizeMB := float64(c.GetBundleSize()) / 1024 / 1024
	c.logger.Infof("Translations synced, count: %d, size %f MB", len(data), sizeMB)
	//c.logger.Infof("Translations synced, count: %d", len(data))

	if sizeMB > float64(c.maxCacheSizeMB) {
		c.logger.Errorf("Warning! Cache size exceeded! %f MB", sizeMB)
	}
}

func (c *Client) UpdateBundle(data []source.Object) error {
	c.localizersLock.Lock()
	defer c.localizersLock.Unlock()

	var localizeTags = make(map[string]language.Tag)

	var err error
	for _, datum := range data {

		var tag language.Tag
		tag, ok := localizeTags[datum.LocaleCode]
		if !ok {
			tag, err = language.Parse(datum.LocaleCode)
			if err != nil {
				return err
			}
			localizeTags[datum.LocaleCode] = tag
		}

		localizer, ok := c.localizers[datum.LocaleCode]
		if !ok {
			c.bundles[datum.LocaleCode] = i18n.NewBundle(tag)
			localizer = i18n.NewLocalizer(c.bundles[datum.LocaleCode])
			c.localizers[datum.LocaleCode] = localizer
		}
		bundle := c.bundles[datum.LocaleCode]

		switch datum.Value.(type) {
		case string:
			err = bundle.AddMessages(tag, &i18n.Message{
				ID:    datum.Key,
				Other: datum.Value.(string),
			})
			if err != nil {
				return err
			}
		case map[string]interface{}:
			value := datum.Value.(map[string]interface{})
			err = bundle.AddMessages(tag, &i18n.Message{
				ID:    datum.Key,
				Zero:  fmt.Sprint(value["zero"]),
				One:   fmt.Sprint(value["one"]),
				Two:   fmt.Sprint(value["two"]),
				Few:   fmt.Sprint(value["few"]),
				Many:  fmt.Sprint(value["many"]),
				Other: fmt.Sprint(value["other"]),
			})
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported value type %T", datum.Value)
		}
	}

	return err
}

func (c *Client) GetBundleSize() int {
	c.localizersLock.RLock()
	defer c.localizersLock.RUnlock()

	size := 0
	for _, bundle := range c.bundles {
		size += int(reflect.TypeOf(bundle).Size())
	}
	return size
}
