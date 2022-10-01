package requestfy

type Client struct{}

type ClientConfig func(*Client)

func NewClient(configs ...ClientConfig) *Client {
	client := &Client{}

	for i := range configs {
		configs[i](client)
	}

	return client
}
