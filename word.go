package wordsdk_go

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/summit-fi/wordsdk-go/fluent"
	"golang.org/x/text/language"
)

type Config struct {
	FTLFilePath []string
}

type localize struct {
	bundles map[language.Tag]*fluent.Bundle
	sync.RWMutex
}

type Word struct {
	ctx      context.Context
	localize localize
}

func New(config Config) (WordSDK, error) {
	files, err := LoadFtl(config.FTLFilePath)
	if err != nil {
		return nil, err
	}
	return &Word{
		localize: localize{
			bundles: InitBundles(files),
		}}, nil
}

func InitBundles(source map[language.Tag]string) map[language.Tag]*fluent.Bundle {

	bundles := make(map[language.Tag]*fluent.Bundle)
	for lang, src := range source {
		resource, err := fluent.NewResource(src)
		if err != nil {
			panic(err)
		}
		bundle := fluent.NewBundle(lang)
		errs := bundle.AddResource(resource)
		if errs != nil {
			for _, errt := range errs {
				if errt != nil {
					panic(err)
				}
			}
		}
		bundles[lang] = bundle

	}
	return bundles
}

func LoadFtl(filepath []string) (map[language.Tag]string, error) {
	if len(filepath) == 0 {
		return map[language.Tag]string{}, fmt.Errorf("no file path provided")
	}
	source := make(map[language.Tag]string)
	for _, path := range filepath {
		file, err := os.Open(path)
		if err != nil {
			return map[language.Tag]string{}, err
		}
		defer file.Close()

		b, err := io.ReadAll(file)
		if err != nil {
			return map[language.Tag]string{}, err
		}
		source[language.English] = string(b)
	}
	return source, nil
}

func (w *Word) WithContext(ctx context.Context) *Word {
	w.ctx = ctx
	return w
}

func (w *Word) Translate(lang language.Tag, key string, ctx ...*fluent.FormatContext) (string, error) {
	w.localize.RLock()
	defer w.localize.RUnlock()

	bundle := w.localize.bundles[lang]
	if bundle == nil {
		return "", fmt.Errorf("bundle not found")
	}

	msg, _, err := bundle.FormatMessage(key, ctx...)
	if err != nil {
		return "", err
	}

	return msg, nil
}
