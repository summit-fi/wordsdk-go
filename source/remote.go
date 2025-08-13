package source

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Remote struct {
	ApiBaseUrl string
	AccessKey  string
	httpClient *http.Client
}

func NewRemote(apiBaseUrl, accessKey string) *Remote {
	return &Remote{
		ApiBaseUrl: apiBaseUrl,
		AccessKey:  accessKey,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

type response struct {
	Key         string `json:"key"`
	HasComments bool   `json:"hasComments"`
	Value       []struct {
		Value       string `json:"value"`
		LocaleCode  string `json:"locale"`
		Status      string `json:"status"`
		HasComments bool   `json:"hasComments"`
	} `json:"values"`
}

func (c *Remote) LoadAllStatic(checksumIn string) (result []Object, checksumOut string, err error) {

	url := fmt.Sprintf("%s/static/values", c.ApiBaseUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessKey)
	if checksumIn != "" {
		req.Header.Set("If-None-Match", checksumIn)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return nil, checksumIn, nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, "", errors.New("Error from server returned: " + string(b))
	}

	var data []response
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, "", err
	}

	for _, d := range data {
		for _, v := range d.Value {
			result = append(result, Object{
				LocaleCode: v.LocaleCode,
				Key:        d.Key,
				Value:      v.Value,
			})
		}
	}

	checksumOut = resp.Header.Get("ETag")
	return result, checksumOut, nil
}

func (c *Remote) LoadAllDynamic(dynamicKey string, checksumIn string) (result []Object, checkSumOut string, err error) {
	url := fmt.Sprintf("%s/dynamic/values", c.ApiBaseUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessKey)

	req.Header.Set("X-Dynamic-Key", dynamicKey)
	if checksumIn != "" {
		req.Header.Set("If-None-Match", checksumIn)
	}
	return nil, "", nil
}

type singleKeyResponse struct {
}

func (c *Remote) LoadOneDynamic(dynamicKey, lang, key string) (string, error) {
	url := fmt.Sprintf("%s/dynamic/value?lang=%s&key=%s", c.ApiBaseUrl, lang, key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return key, err
	}

	req.Header.Set("X-Dynamic-Key", dynamicKey)
	req.Header.Set("Authorization", "Bearer "+c.AccessKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {

		return key, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return key, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return key, err

	}
	return string(b), err
}

func (c *Remote) SaveDynamic(dynamicKey string, data []Object) error {
	var r = struct {
		Values []Object `json:"values"`
	}{

		Values: data,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/dynamic/values", c.ApiBaseUrl)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("X-Dynamic-Key", dynamicKey)
	req.Header.Set("Authorization", "Bearer "+c.AccessKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New("Error from server returned: " + string(b))
	}

	return nil
}

func (c *Remote) Save(data []Object) error {

	var r = struct {
		Values []Object `json:"values"`
	}{
		Values: data,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", c.ApiBaseUrl+"/values", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New("Error from server returned: " + string(b))
	}

	return nil
}
