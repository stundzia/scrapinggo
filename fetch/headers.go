package fetch

import (
	"fmt"
	"io"
	"net/http"
)

// TextWithCustomHeaders - makes a GET request to provided target, reads the response body and returns it as string
func TextWithCustomHeaders(target string, headers http.Header) (string, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return "", err
	}
	req.Header = headers
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return "", ErrNon2xxStatus
	}
	return string(body), nil
}

// ResponseWithCustomHeaders - makes a GET request to provided target with custom headers and returns the response
func ResponseWithCustomHeaders(target string, headers http.Header) (*http.Response, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	return c.Do(req)
}

// HeadersExample - shows example of how to use custom headers when fetching performing requests via http.Client
func HeadersExample() {
	headers := http.Header{
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"Accept-Encoding":           {"gzip,deflate,sdch"},
		"Accept-Language":           {"en-US,en;q=0.8"},
		"Connection":                {"keep-alive"},
		"Cache-Control":             {"max-age=0"},
		"Upgrade-Insecure-Requests": {"1"},
		"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36"},
	}
	res, err := ResponseWithCustomHeaders(TargetHTTPBinHeaders, headers)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("Resp Code: ", res.StatusCode)
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("Headers: ", string(b))
}
