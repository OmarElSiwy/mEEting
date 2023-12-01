package WaterlooAPI

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://api.example.com/v3"
	apiKey  = "D0ED10AD93A34867BC9BFA4211A32B81"
)

/*
* Abstraction of all GET requests
 */
func AbstractedGET(Address string) ([]byte, error) {
	// Creates a http request
	req, err := http.NewRequest("GET", baseURL+Address, nil)
	if err != nil {
		return nil, err
	}
	// Adds the API Key to http request
	req.Header.Add("x-api-key", apiKey)
	client := &http.Client{} // sets up http client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body) // return json in string format
}

/*
* Abstraction of all POST requests
 */
func AbstractedPOST(requestBody []byte, err error, Address string) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(Address, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
