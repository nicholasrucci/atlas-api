package middleware

import ()

// JSONError will hold the data that is responded to the client
type JSONError struct {
	Code  int
	Error error
}

func HandleError(status int, err error) JSONError {
	return JSONError{status, err}
}
