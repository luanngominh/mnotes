package note

import "net/http"

//Define error message and status code
var (
	ErrUserIDInvalid   = errUserIDInvalid{}
	ErrContinueInvalid = errContinueInvalid{}
)

type errUserIDInvalid struct{}

func (errUserIDInvalid) Error() string {
	return "User ID invalid"
}

func (errUserIDInvalid) StatusCode() int {
	return http.StatusBadRequest
}

type errContinueInvalid struct{}

func (errContinueInvalid) Error() string {
	return "Missing start offset note"
}

func (errContinueInvalid) StatusCode() int {
	return http.StatusBadRequest
}
