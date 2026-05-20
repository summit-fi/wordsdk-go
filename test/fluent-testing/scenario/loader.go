package scenario

import (
	"log"
	"os"
	"path/filepath"

	"github.com/summit-fi/wordsdk-go/fluent"
	"github.com/summit-fi/wordsdk-go/fluent/cldr"
)

func LoadBundle(root string, s Scenario) *fluent.Bundle {
	fluentBundle := fluent.NewBundle(cldr.Language(s.Locale))

	err := s.LoadResources(root)
	if err != nil {
		log.Fatalf("build ress: %v", err)
		return nil
	}

	for _, res := range s.Resources {

		path := filepath.Join(root, res)

		b, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("read %s: %v", res, err)
		}
		content := string(b)

		if len(content) == 0 {
			log.Fatalf("resource not found: %s", res)
		}

		res, errs := fluent.NewResource(content)
		if len(errs) > 0 {
			log.Fatalf("parse %s: %v", res, errs)
		}
		if errs := fluentBundle.AddResource(res); len(errs) > 0 {
			log.Fatalf("add %s: %v", res, errs)
		}
	}

	return fluentBundle
}
