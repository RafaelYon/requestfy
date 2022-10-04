package requestfy

import (
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	*http.Response

	client *Client
}

func (r *Response) Json(v interface{}) error {
	if r.Body == nil {
		return errors.New("can't decode response body is null")
	}

	if r.client.newJsonDecoder == nil {
		return fmt.Errorf("can't decoder JSON, no decoder has been specified to the client")
	}

	defer r.Body.Close()
	if err := r.client.newJsonDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("can't decode JSON in response body: %w", err)
	}

	return nil
}
