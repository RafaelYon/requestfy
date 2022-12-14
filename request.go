package requestfy

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// Request stores and facilitates configuration of requests
type Request struct {
	context context.Context
	client  *Client
	headers http.Header
	params  url.Values
}

// Get performs a request using the GET method
func (r *Request) Get(url string) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	for key, headers := range r.headers {
		for _, val := range headers {
			req.Header.Add(key, val)
		}
	}

	r.buildQueryParams(req.URL)

	return r.client.doRequest(req)
}

// Put performs a request using the PUT method
func (r *Request) Put(url string, body io.Reader) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodPut, body)
	if err != nil {
		return nil, err
	}

	r.buildQueryParams(req.URL)

	return r.client.doRequest(req)
}

// Post performs a request using the POST method
func (r *Request) Post(url string, body io.Reader) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodPost, body)
	if err != nil {
		return nil, err
	}

	r.buildQueryParams(req.URL)

	return r.client.doRequest(req)
}

func (r *Request) SetHeader(h, v string) *Request {
	r.headers[h] = append(r.headers[h], v)

	return r
}

func (r *Request) SetHeaders(headers http.Header) *Request {
	for key, vals := range headers {
		for _, val := range vals {
			r.SetHeader(key, val)
		}
	}

	return r
}

func (r *Request) GetHeaders() http.Header {
	return r.headers
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

	r.buildQueryParams(req.URL)

	return r.client.doRequest(req)
}

func (r *Request) Patch(url string) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodPatch, nil)
	if err != nil {
		return nil, err
	}

	r.buildQueryParams(req.URL)

	return r.client.doRequest(req)
}

func (r *Request) Options(url string) (*Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodOptions, nil)
	if err != nil {
		return nil, err
	}

	r.buildQueryParams(req.URL)

	return r.client.doRequest(req)
}

func (r *Request) SetParam(k, v string) *Request {
	r.params[k] = append(r.params[k], v)

	return r
}

func (r *Request) buildQueryParams(u *url.URL) {
	u.RawQuery = r.params.Encode()
}
