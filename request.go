package requestfy

import (
	"context"
	"io"
	"net/http"
)

// Request stores and facilitates configuration of requests
type Request struct {
	context context.Context
	client  *Client
}

// Get performs a request using the GET method
func (r *Request) Get(url string) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return r.client.doRequest(req)
}

// Put performs a request using the PUT method
func (r *Request) Put(url string, body io.Reader) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodPut, body)
	if err != nil {
		return nil, err
	}

	return r.client.doRequest(req)
}

func (r *Request) Delete(url string) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}

	return r.client.doRequest(req)
}

func (r *Request) Head(url string) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodHead, nil)
	if err != nil {
		return nil, err
	}

	return r.client.doRequest(req)
}
