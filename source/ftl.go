package source

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Ftl struct {
	files []file
	sync.RWMutex
}

func (f *Ftl) SaveDynamic(accessKey string, data []Object) error {
	//TODO implement me
	panic("implement me")
}

func (f *Ftl) LoadAllDynamic(key string, checksumIn string) (result []Object, checksumOut string, err error) {
	//TODO implement me
	panic("implement me")
}
func (f *Ftl) LoadOneDynamic(accessKey, lang, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

type file struct {
	localeCode string
	path       string
}

func NewFtl() *Ftl {
	return &Ftl{}
}

func (f *Ftl) AddLocaleFile(localeCode string, path string) error {
	f.RLock()
	defer f.RUnlock()

	// Check if file exists
	fi, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	fi.Close()

	// Add to files
	f.files = append(f.files, file{
		localeCode: localeCode,
		path:       path,
	})
	return nil
}

func (f *Ftl) AddLocaleFiles(paths map[string]string) error {
	for locale, path := range paths {
		err := f.AddLocaleFile(locale, path)
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

func FtlParse(data []byte) []Object {
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var foundedKey string
	var foundedValue json.RawMessage

	// Use a map to store key-value pairs instead of the undefined ptype.Rows
	entries := make(map[string]json.RawMessage)

	for scanner.Scan() {
		line := scanner.Text()
		lineWithoutSpace := strings.TrimSpace(line)

		// Skip comments
		if strings.HasPrefix(lineWithoutSpace, "#") {
			continue
		}

		if strings.Contains(lineWithoutSpace, "=") && !strings.HasPrefix(lineWithoutSpace, ".") {
			split := strings.Split(lineWithoutSpace, "=")

			if len(split) > 0 {
				// When founded new key, founded value should be reset
				foundedValue = nil
				key := strings.TrimSpace(split[0])
				foundedKey = key

				for idx, parts := range split {
					if idx == 0 {
						continue
					}
					foundedValue = json.RawMessage(strings.TrimPrefix(parts, " ") + "\n")
				}

				if foundedKey != "" {
					entries[foundedKey] = foundedValue
				}
			}
		}

		if foundedKey != "" && (!strings.Contains(lineWithoutSpace, "=") || strings.HasPrefix(lineWithoutSpace, ".")) {
			// Skip lines without any characters
			if len(lineWithoutSpace) < 1 {
				continue
			}

			foundedValue = append(foundedValue, []byte(line+"\n")...)
			entries[foundedKey] = foundedValue
		}
	}

	// Convert the map to []Object
	var result []Object
	for key, value := range entries {
		result = append(result, Object{
			Key:   key,
			Value: string(value),
			// Note: LocaleCode isn't set here since it's not available in the parsing context
			// You might need to pass it as a parameter if needed
		})
	}

	return result
}

func (f *Ftl) LoadAllStatic(checksumIn string) ([]Object, string, error) {
	f.RLock()
	defer f.RUnlock()

	// Read all files
	var files []openedFile
	for _, file := range f.files {
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
		objects = FtlParse(file.bytes)
	}

	return objects, checksum, nil
}

func (f *Ftl) Save(data []Object) error {
	f.Lock()
	defer f.Unlock()

	var dataMap = make(map[string]map[string]interface{}) // localeCode -> key -> value
	for _, datum := range data {
		if _, ok := dataMap[datum.LocaleCode]; !ok {
			dataMap[datum.LocaleCode] = make(map[string]interface{})
		}
		dataMap[datum.LocaleCode][datum.Key] = datum.Value
	}

	for _, fi := range f.files {
		if _, ok := dataMap[fi.localeCode]; !ok {
			// File with this locale is not loaded/attached
			continue
		}

		file, err := os.Open(fi.path)
		if err != nil {
			return err
		}
		defer file.Close()

		b, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		builder, err := strings.Builder{}, nil

		for key, value := range dataMap[fi.localeCode] {
			builder.Write(b)
			builder.WriteString("\n")
			builder.WriteString(key)
			builder.WriteString(" = ")
			if strValue, ok := value.(string); ok {
				builder.WriteString(strValue)
			}
			builder.WriteString("\n")
		}

		os.WriteFile(fi.path, []byte(builder.String()), 0755)
	}

	return nil
}
