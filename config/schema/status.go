package schema

import ()

// DatabaseError will hold the data that is responded to the user when the
// database returns an error
type DatabaseError struct {
	Code  int
	Error error
}

func DBClientError(err error) DatabaseError {
	return DatabaseError{400, err}
}
