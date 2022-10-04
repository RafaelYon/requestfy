package requestfy

import (
	"context"
	"net/http"
)

// Request stores and facilitates configuration of requests
type Request struct {
	context context.Context
	client  *Client
	headers map[string]string
}

// Get performs a request using the GET method
func (r *Request) Get(url string) (*http.Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	if len(r.headers) > 0 {
		for key, val := range r.headers {
			req.Header.Add(key, val)
		}
	}

	return r.client.doRequest(req)
}

func (r *Request) SetHeader(h, v string) *Request {
	r.headers[h] = v

	return r
}

func (r *Request) GetHeaders() map[string]string {
	return r.headers
}

func (r *Request) Delete(url string) (*http.Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}

	return r.client.doRequest(req)
}
