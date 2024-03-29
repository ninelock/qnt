package config

import "net/url"

var (
	ProxyURL *url.URL
)

func init() {
	ProxyURL, _ = url.Parse("http://127.0.0.1:9910")
}
