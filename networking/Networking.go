package networking

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
	return handleResponse(http.Get(url))
}

// POST request.
func Post(url string, bodyData string) ([] byte, error) {
	return handleResponse(http.Post(url, "text/plain", strings.NewReader(bodyData)))
}

// POST request.
func PostJson(url string, jsonData interface{}) ([] byte, error) {

	marshaledJson, err := json.Marshal(jsonData)

	if err != nil {
		return nil, err
	}

	return handleResponse(http.Post(url, "text/json", bytes.NewBuffer(marshaledJson)))
}
