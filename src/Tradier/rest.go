package Tradier

import (
	"bytes"
	"net/http"
)

const URL = "https://api.tradier.com/v1"

func (client *TradierClient) signRequest(method string, path string, body []byte) *http.Request {

	req, _ := http.NewRequest(method, URL+path, bytes.NewBuffer(body))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.AccessToken)

	return req
}

func (client *TradierClient) _get(path string, body []byte) (*http.Response, error) {

	preparedRequest := client.signRequest("GET", path, body)

	resp, err := client.Client.Do(preparedRequest)

	return resp, err

}
