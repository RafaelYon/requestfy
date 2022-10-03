package requestfy

import (
	"context"
	"net/http"
)

// Request stores and facilitates configuration of requests
type Request struct {
	context context.Context
	client  *Client
}

// Get performs a request using the GET method
func (r *Request) Get(url string) (*http.Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return r.client.doRequest(req)
}
