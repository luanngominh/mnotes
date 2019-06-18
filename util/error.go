package util

import "errors"

//Handle authentication error
var (
	ErrUnauthorize      = errors.New("Unauthorize")
	ErrMissingAuthToken = errors.New("Missing auth token")
	ErrInvalidFormat    = errors.New("Token invalid format")
)
