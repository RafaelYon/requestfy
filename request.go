package requestfy

import (
	"context"
	"net/http"
)

type Request struct {
	context context.Context
	client  *Client
}

func (r *Request) Get(url string) (*http.Response, error) {
	req, err := r.client.newRequest(r.context, url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return r.client.doRequest(req)
}
