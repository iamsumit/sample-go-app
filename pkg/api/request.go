package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// Decode decodes a JSON request body into the provided type.
func Decode(r *http.Request, d interface{}) error {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// Decode JSON into a map[string]interface{}.
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	// Use mapstructure to decode the map into the struct.
	err = mapstructure.Decode(data, d)
	if err != nil {
		return err
	}

	return nil
}
