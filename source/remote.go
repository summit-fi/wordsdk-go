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
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Values []struct {
		Id          int         `json:"id"`
		Value       interface{} `json:"value"`
		Locale      string      `json:"locale"`
		Status      string      `json:"status"`
		HasComments bool        `json:"hasComments"`
	} `json:"values"`
}

func (c *Remote) LoadAll(checksumIn string) (result []Object, checksumOut string, err error) {

	url := fmt.Sprintf("%s/values", c.ApiBaseUrl)
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
		for _, v := range d.Values {
			result = append(result, Object{
				LocaleCode: v.Locale,
				Key:        d.Name,
				Value:      v.Value,
			})
		}
	}

	checksumOut = resp.Header.Get("ETag")
	return result, checksumOut, nil
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
