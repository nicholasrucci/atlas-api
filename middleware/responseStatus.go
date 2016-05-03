package middleware

import ()

// DatabaseError will hold the data that is responded to the user when the
// database returns an error
type JSONError struct {
	Code  int
	Error error
}

func HandleError(status int, err error) JSONError {
	return JSONError{status, err}
}
