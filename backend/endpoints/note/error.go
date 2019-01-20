package note

import "net/http"

//Define error message and status code
var (
	ErrUserIDInvalid = errUserIDInvalid{}
)

type errUserIDInvalid struct{}

func (errUserIDInvalid) Error() string {
	return "User ID invalid"
}

func (errUserIDInvalid) StatusCode() int {
	return http.StatusBadRequest
}
