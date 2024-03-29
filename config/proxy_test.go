package config

import (
	"testing"
)

func TestProxy(t *testing.T) {
	t.Log(ProxyURL.String())
}
