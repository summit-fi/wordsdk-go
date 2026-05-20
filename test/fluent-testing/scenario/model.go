package scenario

import (
	"fmt"
	"os"
	"strings"
)

type FileSource struct {
	Name       string
	PathScheme string
}

type Query struct {
	Key    string
	Args   map[string]any
	Value  *string
	Attrs  map[string]any
	Result string
}

type Scenario struct {
	Name        string
	FileSources []FileSource
	Locale      string
	FromFile    *string
	Resources   []string
	Queries     []Query
}

func GetResourcesPath(base string) ([]string, error) {
	var paths []string
	entries, err := os.ReadDir(base)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subBase := base + "/" + entry.Name()
			subPaths, err := GetResourcesPath(subBase)
			if err != nil {
				return nil, err
			}
			for _, subPath := range subPaths {
				path := entry.Name() + "/" + subPath
				paths = append(paths, path)
			}
		} else {
			paths = append(paths, entry.Name())
		}
	}

	if len(paths) == 0 {
		return nil, fmt.Errorf("no files found in path: %s", base)

	}

	return paths, nil
}

func (s *Scenario) LoadResources(root string) error {

	for _, src := range s.FileSources {
		base := strings.ReplaceAll(src.PathScheme, "{locale}", s.Locale)

		paths, err := GetResourcesPath(root + "/" + base)
		if err != nil {
			return err
		}

		for _, path := range paths {
			resID := base + path
			s.Resources = append(s.Resources, resID)
		}

	}
	return nil
}
