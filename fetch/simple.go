package fetch

import (
	"encoding/json"
	"io"
	"net/http"
)

// TextSimple - fetches response from given target url and returns response body as string, will return error
// if any occurs or response status is not 2xx
func TextSimple(target string) (string, error) {
	res, err := http.Get(target)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", ErrNon200Status
	}

	return string(body), nil
}

// ResponseSimple - fetches response from target and returns it.
func ResponseSimple(target string) (*http.Response, error) {
	res, err := http.Get(target)

	return res, err
}

// JSONSimple - fetches response from target, attempts to unmarshall it as JSON and returns it as map[string]interface{}
func JSONSimple(target string) (map[string]interface{}, error) {
	res, err := http.Get(target)
	if err != nil {
		return map[string]interface{}{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}
	resJSON := map[string]interface{}{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if res.StatusCode != 200 {
		return map[string]interface{}{}, ErrNon200Status
	}
	return resJSON, nil
}
