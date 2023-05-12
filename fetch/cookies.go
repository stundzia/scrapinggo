package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// CookiesExample - shows how to handle cookies via CookieJar on http.Client
func CookiesExample() {
	cookieJar, _ := cookiejar.New(nil)
	c := http.Client{
		Jar: cookieJar,
	}

	res, _ := c.Get(TargetGoogleSearch)
	fmt.Println("Google status: ", res.StatusCode)
	googleURL, _ := url.Parse(TargetGoogleSearch)
	googleCookies := c.Jar.Cookies(googleURL)

	targetUrl, _ := url.Parse(TargetHTTPBinCookies)
	ourCookies := []*http.Cookie{
		{
			Name:  "testCookie",
			Value: "nomnom",
		},
	}

	c.Jar.SetCookies(targetUrl, ourCookies)
	c.Jar.SetCookies(targetUrl, googleCookies)
	res, err := c.Get(TargetHTTPBinCookies)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("Status: ", res.StatusCode)
	b, _ := io.ReadAll(res.Body)
	fmt.Println("Cookies: ", string(b))
}
