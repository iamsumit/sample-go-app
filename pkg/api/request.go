package api

import (
	"io/ioutil"
	"net/http"
)

// Decodable is an interface for types that can be decoded from JSON
type Decodable interface {
	DecodeJSON([]byte) error
}

// Decode decodes a JSON request body into the provided type
func Decode(r *http.Request, d Decodable) error {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return d.DecodeJSON(body)
}
