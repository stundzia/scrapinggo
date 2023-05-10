package fetch

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

// TlsExample - shows examples how to configure custom TLS for TLS fingerprint spoofing
func TlsExample() {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				CipherSuites: []uint16{
					tls.TLS_RSA_WITH_AES_128_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				},
				MaxVersion:         tls.VersionTLS13,
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true,
			},
		},
	}
	res, err := c.Get(TargetGoogleSearch)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("Status: ", res.StatusCode, " Err: ", err)
}
