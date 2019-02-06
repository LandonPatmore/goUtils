package networking

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Handles response and errors of HTTP requests.
func handleResponse(response *http.Response, err error) ([] byte, error) {
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if response.StatusCode >= 200 && response.StatusCode <= 299 {

			body, err := ioutil.ReadAll(response.Body)

			if err != nil {
				return nil, err
			}

			return body, nil
		} else {
			return nil, errors.New("The response code is: " + strconv.Itoa(response.StatusCode) + ".  The response message is: " + response.Status)
		}
	}
}

// GET request.
func Get(url string) ([] byte, error) {
	return handleResponse(createGetClient(url))
}

// POST request.
func Post(url string, bodyData string) ([] byte, error) {
	return handleResponse(createPostClient(url, "text/plain", []byte(bodyData)))
}

// POST request.
func PostJson(url string, jsonData interface{}) ([] byte, error) {

	marshaledJson, err := json.Marshal(jsonData)

	if err != nil {
		return nil, err
	}

	return handleResponse(createPostClient(url, "application/json", marshaledJson))
}

func createPostClient(url string, contentType string, data [] byte) (*http.Response, error) {
	httpClient := http.Client{Timeout: 10 * time.Second}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", contentType)

	return httpClient.Do(request)
}

func createGetClient(url string) (*http.Response, error) {
	httpClient := http.Client{Timeout: 10 * time.Second}
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	return httpClient.Do(request)
}
