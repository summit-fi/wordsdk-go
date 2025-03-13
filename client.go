package word

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"time"

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
	cache          Map[string, string]
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
		saveStrategy:   config.SaveStrategy,
	}

	data, checksum, err := config.Source.LoadAll("")
	if err != nil {
		return nil, fmt.Errorf("failed to load translations: %v", err)
	}

	c.cache = NewMap[string, string]()

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

	sizeMB := float64(c.GetCacheSize()) / 1024 / 1024
	c.logger.Infof("Translations synced, count: %d, size %f MB", len(data), sizeMB)
	//c.logger.Infof("Translations synced, count: %d", len(data))

	if sizeMB > float64(c.maxCacheSizeMB) {
		c.logger.Errorf("Warning! Cache size exceeded! %f MB", sizeMB)
	}
}

func (c *Client) UpdateBundle(data []source.Object) error {
	c.localizersLock.Lock()
	defer c.localizersLock.Unlock()
	for _, d := range data {
		c.cache.Set(d.LocaleCode+d.Key, d.Value)
	}
	return nil
}

func (c *Client) GetCacheSize() int {
	c.localizersLock.RLock()
	defer c.localizersLock.RUnlock()

	size := 0
	for _, bundle := range c.cache.GetAll() {
		size += int(reflect.TypeOf(bundle).Size())
	}
	return size
}
