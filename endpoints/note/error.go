package note

import "net/http"

//Define error message and status code
var (
	ErrUserIDInvalid   = errUserIDInvalid{}
	ErrContinueInvalid = errContinueInvalid{}
	ErrGetAllNote      = errGetAllNote{}
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

type errGetAllNote struct{}

func (errGetAllNote) Error() string {
	return "Get all note error"
}

func (errGetAllNote) StatusCode() int {
	return http.StatusBadRequest
}
