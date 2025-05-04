package config

import (
	"io"
	"net/http"
)

// RestClientConfig 模拟 Spring 的 RestClientConfig
type RestClientConfig struct {
	client *http.Client
}

// NewRestClientConfig 创建一个新的 RestClientConfig 实例
func NewRestClientConfig() *RestClientConfig {
	return &RestClientConfig{
		client: &http.Client{},
	}
}

// Get 发起 GET 请求
func (c *RestClientConfig) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

// Post 发起 POST 请求
func (c *RestClientConfig) Post(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}
