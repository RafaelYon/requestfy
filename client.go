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

func NewClient(configs ...ClientConfig) *Client {
	client := &Client{}

	for i := range configs {
		configs[i](client)
	}

	return client
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
