package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func StringViaProxy(target, proxyURLString string) (string, error) {
	proxyUrl, err := url.Parse(proxyURLString)
	c := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	res, err := c.Get(target)
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

func ResponseViaProxy(target, proxyURLString string) (*http.Response, error) {
	proxyUrl, err := url.Parse(proxyURLString)
	c := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	res, err := c.Get(target)

	return res, err
}

// ProxyExample - shows example of how to make requests via proxy
func ProxyExample() {
	// Proxy string, i.e. http://username:password@pr.oxylabs.io:7777
	res, err := ResponseViaProxy(TargetIPInfoJSON, ProxyOxylabsHTTP)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("Status: ", res.StatusCode)
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("IP Info: ", string(b))
}
