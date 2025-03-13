package source

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

type LocalJson struct {
	files []file
	sync.RWMutex
}

type file struct {
	localeCode string
	path       string
}

func NewJson() *LocalJson {
	return &LocalJson{}
}

func (j *LocalJson) AddLocaleFile(localeCode string, path string) error {
	j.RLock()
	defer j.RUnlock()

	// Check if file exists
	f, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	f.Close()

	// Add to files
	j.files = append(j.files, file{
		localeCode: localeCode,
		path:       path,
	})
	return nil
}

func (j *LocalJson) AddLocaleFiles(paths map[string]string) error {
	for locale, path := range paths {
		err := j.AddLocaleFile(locale, path)
		if err != nil {
			return err
		}
	}
	return nil
}

type openedFile struct {
	localeCode string
	bytes      []byte
}

func (j *LocalJson) LoadAll(checksumIn string) ([]Object, string, error) {
	j.RLock()
	defer j.RUnlock()

	// Read all files
	var files []openedFile
	for _, file := range j.files {
		f, err := os.Open(file.path)
		if err != nil {
			return nil, "", fmt.Errorf("failed to open file %s: %v", file.path, err)
		}
		defer f.Close()

		b, err := io.ReadAll(f)
		if err != nil {
			return nil, "", fmt.Errorf("failed to read file %s: %v", file.path, err)
		}

		files = append(files, openedFile{
			localeCode: file.localeCode,
			bytes:      b,
		})
	}

	// Compare checksum
	b := []byte{}
	for _, file := range files {
		b = append(b, file.bytes...)
	}
	h := sha256.New()
	h.Write(b)
	checksum := fmt.Sprintf("%x", h.Sum(nil))
	if checksum == checksumIn {
		return nil, checksum, nil
	}

	// Format result
	var objects []Object
	for _, file := range files {
		var data map[string]string
		err := json.Unmarshal(file.bytes, &data)
		if err != nil {
			return nil, "", fmt.Errorf("failed to unmarshal file %s: %v", file.localeCode, err)
		}

		for key, value := range data {
			objects = append(objects, Object{
				LocaleCode: file.localeCode,
				Key:        key,
				Value:      value,
			})
		}
	}

	return objects, checksum, nil
}

func (j *LocalJson) Save(data []Object) error {
	j.Lock()
	defer j.Unlock()

	var dataMap = make(map[string]map[string]interface{}) // localeCode -> key -> value
	for _, datum := range data {
		if _, ok := dataMap[datum.LocaleCode]; !ok {
			dataMap[datum.LocaleCode] = make(map[string]interface{})
		}
		dataMap[datum.LocaleCode][datum.Key] = datum.Value
	}

	for _, f := range j.files {
		if _, ok := dataMap[f.localeCode]; !ok {
			// File with this locale is not loaded/attached
			continue
		}

		file, err := os.Open(f.path)
		if err != nil {
			return err
		}
		defer file.Close()

		b, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		var objects map[string]interface{}
		err = json.Unmarshal(b, &objects)
		if err != nil {
			return err
		}

		for key, value := range dataMap[f.localeCode] {
			objects[key] = value
		}

		// Write all
		b, err = json.MarshalIndent(objects, "", "  ")
		if err != nil {

			return err
		}

		os.WriteFile(f.path, b, 0755)
	}

	return nil
}
