package note

import "net/http"

//define err message and status code
var (
	ErrBodyEmpty  = errBodyEmpty{}
	ErrTitleEmpty = errBodyEmpty{}
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
