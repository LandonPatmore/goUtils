package networking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Handles response and errors of HTTP requests.
func handleResponse(response *http.Response, err error) [] byte {
	if err != nil {
		fmt.Println("The request could not be sent due to an error.")
		return nil
	} else {
		defer response.Body.Close()

		if response.StatusCode >= 200 && response.StatusCode <= 299 {

			body, err := ioutil.ReadAll(response.Body)

			if err != nil {
				fmt.Println("The message body could not be read due to an error.")
				return nil
			}

			return body
		} else {
			fmt.Println("The response code was: " + strconv.Itoa(response.StatusCode) + ".  The response message is: " + response.Status)
		}
	}

	return nil
}

// GET request.
func Get(url string) [] byte {
	return handleResponse(http.Get(url))
}

// POST request.
func Post(url string, bodyData string) [] byte {
	return handleResponse(http.Post(url, "text/plain", strings.NewReader(bodyData)))
}

// POST request.
func PostJson(url string, jsonData interface{}) [] byte {

	marshaledJson, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println("JSON data could not be marshaled correctly.  JSON: " + jsonData.(string))
		return nil
	}

	return handleResponse(http.Post(url, "text/json", bytes.NewBuffer(marshaledJson)))
}
