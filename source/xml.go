package source

import (
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type LocalXml struct {
	path string
}

func NewLocalXml(path string) (*LocalXml, error) {
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

	return &LocalXml{path: path}, nil
}

func (x *LocalXml) LoadAll(checksumIn string) ([]Object, string, error) {
	file, err := os.Open(x.path)
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
	err = xml.Unmarshal(b, &objects)
	if err != nil {
		return nil, "", err
	}

	return objects, checksum, nil
}

func (x *LocalXml) Save(data []Object) error {

	var objects []Object

	file, err := os.Open(x.path)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(b, &objects)
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
	b, err = xml.Marshal(objects)
	if err != nil {
		return err
	}
	return os.WriteFile(x.path, b, 0755)
}
