package helper

import (
	"encoding/json"
	"net/http"

	"atlas-api/config/schema"
)

// JSONError will hold the data that is responded to the client
type Response struct {
	Code     int
	Response interface{}
	Error    error
}

func OrganizationResponse(rw http.ResponseWriter, req *http.Request, status int, response schema.Organization, incomingError error) error {
	responseError := Response{status, response, incomingError}

	JSONHandler(rw, req)

	rw.WriteHeader(status)

	err := json.NewEncoder(rw).Encode(responseError)
	if err != nil {
		responseError = Response{500, response, err}
	}

	return nil
}

func UserResponse(rw http.ResponseWriter, req *http.Request, status int, response schema.User, incomingError error) error {
	responseError := Response{status, response, incomingError}

	JSONHandler(rw, req)

	rw.WriteHeader(status)

	err := json.NewEncoder(rw).Encode(responseError)
	if err != nil {
		responseError = Response{500, response, err}
	}

	return nil
}

func ProjectResponse(rw http.ResponseWriter, req *http.Request, status int, response schema.Project, incomingError error) error {
	responseError := Response{status, response, incomingError}

	JSONHandler(rw, req)

	rw.WriteHeader(status)

	err := json.NewEncoder(rw).Encode(responseError)
	if err != nil {
		responseError = Response{500, response, err}
	}

	return nil
}
