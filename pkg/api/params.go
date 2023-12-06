package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Param returns the web call parameters from the request
func Param(r *http.Request, key string) string {
	m := mux.Vars(r)
	return m[key]
}
