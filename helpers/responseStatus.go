package helper

import (
	"encoding/json"
	"net/http"
)

// JSONError will hold the data that is responded to the client
type JSONError struct {
	Code  int
	Error error
}

func HandleError(rw http.ResponseWriter, req *http.Request, status int, incomingError error) error {
	responseError := JSONError{status, incomingError}

	JSONHandler(rw, req)

	rw.WriteHeader(status)

	err := json.NewEncoder(rw).Encode(responseError)
	if err != nil {
		responseError = JSONError{500, err}
	}

	return nil
}
