package requestfy

import (
	"context"
	"net/http"
)

// RequestExecuter abstracts the http client used behind the scenes to perform HTTP requests
type RequestExecuter interface {
	Do(*http.Request) (*http.Response, error)
}

// Client Allows you to perform http requests with a simple syntax
type Client struct {
	executer RequestExecuter
	baseURL  string
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

// ConfigBaseURL specifies a prefix, a base to apply to the URL of all client requests
func ConfigBaseURL(baseURL string) ClientConfig {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// ConfigDefault configures the client with default options to allow quick start
func ConfigDefault() ClientConfig {
	return func(c *Client) {
		c.executer = http.DefaultClient
	}
}

func (c *Client) Request() *Request {
	return c.RequestWithContext(context.Background())
}

func (c *Client) RequestWithContext(ctx context.Context) *Request {
	return &Request{
		client:  c,
		context: ctx,
	}
}
