package source

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type LocalYaml struct {
	path string
}

func NewLocalYaml(path string) (*LocalYaml, error) {
	// Check if file exists
	file, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, create it
			file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				return nil, fmt.Errorf("failed to create file: %w", err)
			}

			fmt.Println("DB file created: " + path)
		}
	}

	file.Close()

	return &LocalYaml{path: path}, nil
}

func (y *LocalYaml) LoadAll(checksumIn string) ([]Object, string, error) {
	file, err := os.Open(y.path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, "", err
	}

	// Compare checksum
	h := sha256.New()
	h.Write(b)
	checksum := fmt.Sprintf("%x", h.Sum(nil))
	if checksum == checksumIn {
		return nil, checksum, nil
	}

	// Format result
	var objects []Object
	err = yaml.Unmarshal(b, &objects)
	if err != nil {
		return nil, "", err
	}

	return objects, checksumIn, nil
}

func (y *LocalYaml) Save(data []Object) error {

	var objects []Object

	file, err := os.Open(y.path)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &objects)
	if err != nil {
		return err
	}

	for i, object := range objects {
		for j, datum := range data {
			if object.LocaleCode == datum.LocaleCode && object.Key == datum.Key {
				objects[i] = datum

				// Remove from data
				data = append(data[:j], data[j+1:]...)
				break
			}
		}
	}

	// Append new objects
	objects = append(objects, data...)

	// Write all
	b, err = yaml.Marshal(objects)
	if err != nil {
		return err
	}
	return os.WriteFile(y.path, b, 0755)
}
