package helper

import (
	"encoding/json"
	"net/http"
)

// JSONError will hold the data that is responded to the client
type Response struct {
	Code     int
	Response interface{}
	Error    error
}

func CreateResponse(rw http.ResponseWriter, req *http.Request, status int, response interface{}, incomingError error) error {
	responseError := Response{status, response, incomingError}

	JSONHandler(rw, req)

	rw.WriteHeader(status)

	err := json.NewEncoder(rw).Encode(responseError)
	if err != nil {
		responseError = Response{500, response, err}
	}

	return nil
}
