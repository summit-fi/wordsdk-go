package source

import (
	"bufio"
	"bytes"
	"crypto/sha256"
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

const ftlBlockIndent = "    "

// FormatFTLEntry formats one key/value pair as a valid FTL entry.
// Multi-line values are emitted as block patterns so line breaks are part of
// the value instead of being interpreted as the end of the entry.
func FormatFTLEntry(key, value string) string {
	value = normalizeFTLNewlines(value)
	if !strings.Contains(value, "\n") {
		return fmt.Sprintf("%s = %s\n", key, value)
	}

	var builder strings.Builder
	builder.WriteString(key)
	builder.WriteString(" =\n")
	for _, line := range strings.Split(value, "\n") {
		builder.WriteString(ftlBlockIndent)
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	return builder.String()
}

func normalizeFTLNewlines(value string) string {
	value = strings.ReplaceAll(value, "\r\n", "\n")
	return strings.ReplaceAll(value, "\r", "\n")
}

func (f *Ftl) SaveDynamic(accessKey string, data []Object) error {
	f.Lock()
	defer f.Unlock()

	var dataMap = make(map[string]map[string]string) // localeCode -> key -> value
	for _, datum := range data {
		if _, ok := dataMap[datum.LocaleCode]; !ok {
			dataMap[datum.LocaleCode] = make(map[string]string)
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

		var builder strings.Builder
		builder.Write(b)

		for key, value := range dataMap[fi.localeCode] {
			builder.WriteString("\n")
			builder.WriteString(FormatFTLEntry(key, value))
		}

		if err := os.WriteFile(fi.path, []byte(builder.String()), 0755); err != nil {
			return err
		}
	}

	return nil
}

func (f *Ftl) LoadAllDynamic(key string, checksumIn string) (result []Object, checksumOut string, err error) {
	return f.LoadAllStatic(checksumIn)
}
func (f *Ftl) LoadOneDynamic(accessKey, lang, key string) (string, error) {
	f.RLock()
	defer f.RUnlock()

	for _, file := range f.files {
		if file.localeCode == lang {
			f, err := os.Open(file.path)
			if err != nil {
				return key, err
			}
			defer f.Close()

			b, err := io.ReadAll(f)
			if err != nil {
				return key, err
			}

			objects := FtlParse(lang, b)
			for _, obj := range objects {
				if obj.Key == key {
					return strings.Trim(obj.Value, "\n"), nil
				}
			}
		}
	}

	return key, fmt.Errorf("Key: %s, nof found", key) // Not found
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

func FtlParse(locale string, data []byte) []Object {
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var foundedKey string
	var foundedValue strings.Builder
	var foundedValueBlock bool

	entries := make(map[string]string)
	flush := func() {
		if foundedKey != "" {
			entries[foundedKey] = foundedValue.String()
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineWithoutSpace := strings.TrimSpace(line)

		// Skip top-level comments.
		if !startsWithFTLIndent(line) && strings.HasPrefix(lineWithoutSpace, "#") {
			continue
		}

		if key, value, ok := parseFTLEntryLine(line); ok {
			flush()
			foundedKey = key
			foundedValue.Reset()
			foundedValueBlock = value == ""
			if value != "" {
				foundedValue.WriteString(value)
				foundedValue.WriteString("\n")
			}
			continue
		}

		if foundedKey != "" {
			// Skip lines without any characters
			if len(lineWithoutSpace) < 1 && !foundedValueBlock {
				continue
			}

			if foundedValueBlock {
				line = strings.TrimPrefix(line, ftlBlockIndent)
			}
			foundedValue.WriteString(line)
			foundedValue.WriteString("\n")
		}
	}
	flush()

	// Convert the map to []Object
	var result []Object
	for key, value := range entries {
		result = append(result, Object{
			LocaleCode: locale,
			Key:        key,
			Value:      value,
		})
	}

	return result
}

func parseFTLEntryLine(line string) (key, value string, ok bool) {
	if startsWithFTLIndent(line) {
		return "", "", false
	}

	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ".") {
		return "", "", false
	}

	idx := strings.Index(line, "=")
	if idx < 0 {
		return "", "", false
	}

	key = strings.TrimSpace(line[:idx])
	if key == "" {
		return "", "", false
	}

	value = strings.TrimPrefix(line[idx+1:], " ")
	return key, value, true
}

func startsWithFTLIndent(line string) bool {
	return strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t")
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
		objects = append(objects, FtlParse(file.localeCode, file.bytes)...)
	}

	return objects, checksum, nil
}
