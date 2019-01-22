package note

import "net/http"

//define err message and status code
var (
	ErrBodyEmpty     = errBodyEmpty{}
	ErrTitleEmpty    = errBodyEmpty{}
	ErrUserIDInvalid = errUserIDInvalid{}
)

type errTitleEmpty struct{}

func (errTitleEmpty) Error() string {
	return "Error title empty"
}

func (errTitleEmpty) StatusCode() int {
	return http.StatusBadRequest
}

type errBodyEmpty struct{}

func (errBodyEmpty) Error() string {
	return "Error title empty"
}

func (errBodyEmpty) StatusCode() int {
	return http.StatusBadRequest
}

type errUserIDInvalid struct{}

func (errUserIDInvalid) Error() string {
	return "Error userID invalid"
}

func (errUserIDInvalid) StatusCode() int {
	return http.StatusBadRequest
}
