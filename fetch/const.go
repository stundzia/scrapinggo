package fetch

import (
	"errors"
	"net/http"
)

const (
	OxylabsUsername           = "username"
	OxylabsPassword           = "password"
	ProxyOxylabsHTTP          = "http://" + OxylabsUsername + ":" + OxylabsPassword + "@pr.oxylabs.io:7777"
	ProxyOxylabsHTTPS         = "https://" + OxylabsUsername + ":" + OxylabsPassword + "@pr.oxylabs.io:7777"
	ProxyOxylabsSOCKS5        = "socks5://" + OxylabsUsername + ":" + OxylabsPassword + "@pr.oxylabs.io:7777"
	ProxyOxylabsSOCKS5H       = "socks5h://" + OxylabsUsername + ":" + OxylabsPassword + "@pr.oxylabs.io:7777"
	TargetGoogleSearch        = "https://www.google.com/search?q=asdf"
	TargetHTTPBinHeaders      = "https://httpbin.org/headers"
	TargetHTTPBinCookies      = "http://httpbin.org/cookies"
	TargetIPInfoJSON          = "https://ipinfo.io/json"
	TargetIPInfoIP            = "https://ipinfo.io/ip"
	TargetStundziaLTStatus200 = "http://xn--stundia-hxb.lt/status_spoof?status=200"
	TargetStundziaLTStatus204 = "http://xn--stundia-hxb.lt/status_spoof?status=204"
	TargetStundziaLTStatus400 = "http://xn--stundia-hxb.lt/status_spoof?status=400"
	TargetStundziaLTStatus404 = "http://xn--stundia-hxb.lt/status_spoof?status=404"
	TargetStundziaLTStatus500 = "http://xn--stundia-hxb.lt/status_spoof?status=500"
)

var ErrNon2xxStatus = errors.New("received non-2xx status code")
var ErrNon200Status = errors.New("received non-200 status code")

var HeadersDesktop = http.Header{
	"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
	"Accept-Encoding":           {"gzip,deflate,sdch"},
	"Accept-Language":           {"en-US,en;q=0.8"},
	"Connection":                {"keep-alive"},
	"Cache-Control":             {"max-age=0"},
	"Upgrade-Insecure-Requests": {"1"},
	"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36"},
}

var HeadersDesktopMap = map[string]interface{}{
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	"Accept-Encoding":           "gzip,deflate,sdch",
	"Accept-Language":           "en-US,en;q=0.8",
	"Connection":                "keep-alive",
	"Cache-Control":             "max-age=0",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36",
}
