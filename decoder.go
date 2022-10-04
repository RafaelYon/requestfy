package requestfy

import (
	"encoding/json"
	"io"
)

// Decoder abstracts the implementation of a decoder from a specific format
type Decoder interface {
	Decode(interface{}) error
}

// NewDecoder generalizes creation of decoder to allow switching between implementations
type NewDecoder func(io.Reader) Decoder

// StdJsonDecoder provides a json.NewDecoder standard decoder
func StdJsonDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}
