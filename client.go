package requestfy

import "net/http"

type RequestExecuter interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	executer RequestExecuter
}

type ClientConfig func(*Client)

func NewClient(configs ...ClientConfig) *Client {
	client := &Client{}

	for i := range configs {
		configs[i](client)
	}

	return client
}

func ConfigRequestExecuter(executer RequestExecuter) ClientConfig {
	return func(c *Client) {
		c.executer = executer
	}
}
