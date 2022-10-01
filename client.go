package requestfy

import "net/http"

// RequestExecuter abstracts the http client used behind the scenes to perform HTTP requests
type RequestExecuter interface {
	Do(*http.Request) (*http.Response, error)
}

// Client Allows you to perform http requests with a simple syntax
type Client struct {
	executer RequestExecuter
}

// ClientConfig is a function to configure the client
type ClientConfig func(*Client)

func NewClient(configs ...ClientConfig) *Client {
	client := &Client{}

	for i := range configs {
		configs[i](client)
	}

	return client
}

// ConfigRequestExecuter allows you to specify an http request executor for the client
func ConfigRequestExecuter(executer RequestExecuter) ClientConfig {
	return func(c *Client) {
		c.executer = executer
	}
}

// ConfigDefault configures the client with default options to allow quick start
func ConfigDefault() ClientConfig {
	return func(c *Client) {
		c.executer = http.DefaultClient
	}
}
