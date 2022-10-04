package requestfy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	*http.Response
}

func (r *Response) Json(v interface{}) error {
	if r.Body == nil {
		return errors.New("can't decode response body is null")
	}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("can't decode JSON in response body: %w", err)
	}

	return nil
}
