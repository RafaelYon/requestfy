package requestfy

import (
	"context"
	"fmt"
	"net/http"
)

type Request struct {
	context context.Context
	client  *Client
}

func (r *Request) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, r.client.concatURL(url), nil)
	if err != nil {
		return nil, fmt.Errorf("can't create GET request: %w", err)
	}

	res, err := r.client.executer.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't execute GET request: %w", err)
	}

	return res, nil
}
