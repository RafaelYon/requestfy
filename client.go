package requestfy

import (
	"context"
	"fmt"
	"io"
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

func (c *Client) newRequest(ctx context.Context, url, method string, body io.Reader) (*http.Request, error) {
	if len(c.baseURL) > 1 {
		url = fmt.Sprintf("%s/%s", c.baseURL, url)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("can't create %s request to '%s': %w", method, url, err)
	}

	return req, nil
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	res, err := c.executer.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't execute %s request to '%s': %w", req.Method, req.URL, err)
	}

	return res, nil
}
