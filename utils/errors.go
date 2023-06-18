package utils

import "errors"

var (
	// ErrRequestStatusCode represents the error when api returned a error code
	ErrRequestStatusCode = errors.New("failed to call http api")
)
