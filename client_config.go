package requestfy

import (
	"net/http"
)

// ClientConfig is a function to configure the client
type ClientConfig func(*Client)

// ConfigRequestExecuter allows you to specify an http request executor for the client
func ConfigRequestExecuter(executer RequestExecuter) ClientConfig {
	return func(c *Client) {
		c.executer = executer
	}
}

// ConfigBaseURL specifies a prefix, a base to apply to the URL of all client requests
func ConfigBaseURL(baseURL string) ClientConfig {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// ConfigJsonDecoder specifies a function to create a new JSON decoder
func ConfigJsonDecoder(newDecoder NewDecoder) ClientConfig {
	return func(c *Client) {
		c.newJsonDecoder = newDecoder
	}
}

// ConfigDefault configures the client with default options to allow quick start
func ConfigDefault() ClientConfig {
	return func(c *Client) {
		c.executer = http.DefaultClient
		c.newJsonDecoder = StdJsonDecoder
	}
}
