package requestfy

import (
	"context"
	"net/http"
)

// Request stores and facilitates configuration of requests
type Request struct {
	context context.Context
	client  *Client
	headers http.Header
}

// Get performs a request using the GET method
func (r *Request) Get(url string) (*http.Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	for key, headers := range r.headers {
		for _, val := range headers {
			req.Header.Add(key, val)
		}
	}

	return r.client.doRequest(req)
}

func (r *Request) SetHeader(h, v string) *Request {
	r.headers[h] = append(r.headers[h], v)

	return r
}

func (r *Request) GetHeaders() http.Header {
	return r.headers
}
