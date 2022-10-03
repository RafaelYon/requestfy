package requestfy

import "context"

type Request struct {
	context context.Context
	client  *Client
}
