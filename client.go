package requestfy

import (
	"context"
	"fmt"
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

func (c *Client) concatURL(path string) string {
	if len(c.baseURL) < 1 {
		return path
	}

	return fmt.Sprintf("%s%s", c.baseURL, path)
}
