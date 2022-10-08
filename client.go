package requestfy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// RequestExecuter abstracts the http client used behind the scenes to perform HTTP requests
type RequestExecuter interface {
	Do(*http.Request) (*http.Response, error)
}

// Client allows you to perform http requests with a simple syntax
type Client struct {
	executer RequestExecuter
	baseURL  string

	// newJsonDecoder stores a function to create a new json decoder
	newJsonDecoder NewDecoder
}

func NewClient(configs ...ClientConfig) *Client {
	client := &Client{}

	for i := range configs {
		configs[i](client)
	}

	return client
}

// Request create a request with using the context.Background
func (c *Client) Request() *Request {
	return c.RequestWithContext(context.Background())
}

// RequestWithContext create a request with using the specified context
func (c *Client) RequestWithContext(ctx context.Context) *Request {
	return &Request{
		client:  c,
		context: ctx,
		headers: make(http.Header),
		params:  make(url.Values),
	}
}

// newRequest creates a request and applies the client's configuration to this request
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

func (c *Client) doRequest(req *http.Request) (*Response, error) {
	res, err := c.executer.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't execute %s request to '%s': %w", req.Method, req.URL, err)
	}

	if res == nil {
		return nil, fmt.Errorf("executer return nil to %s request to '%s'", req.Method, req.URL)
	}

	return &Response{
		Response: res,
		client:   c,
	}, nil
}
