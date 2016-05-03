package middleware

import (
	"net/http"
)

// JSONHandler is set the header to have the content type of json
func JSONHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
