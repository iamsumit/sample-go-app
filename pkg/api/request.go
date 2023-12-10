package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/logger"
	"github.com/mitchellh/mapstructure"
)

// Decode decodes a JSON request body into the provided type.
func Decode(r *http.Request, log logger.Logger, d interface{}) error {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(
			"decoding error while reading the body",
			"error", err.Error(),
			"method", r.Method,
			"endpoint", r.URL.Path,
			"operation", "createUser",
		)

		return ErrDecode
	}
	defer r.Body.Close()

	// Decode JSON into a map[string]interface{}.
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error(
			"decoding error while unmarshalling the body",
			"error", err.Error(),
			"method", r.Method,
			"endpoint", r.URL.Path,
			"operation", "createUser",
		)

		return ErrDecode
	}

	// Use mapstructure to decode the map into the struct.
	err = mapstructure.Decode(data, d)
	if err != nil {
		log.Error(
			"decoding error while decoding the body into the struct",
			"error", err.Error(),
			"method", r.Method,
			"endpoint", r.URL.Path,
			"operation", "createUser",
		)

		return ErrDecode
	}

	return nil
}
