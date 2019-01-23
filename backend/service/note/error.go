package note

import "net/http"

//define err message and status code
var (
	ErrBodyEmpty     = errBodyEmpty{}
	ErrTitleEmpty    = errBodyEmpty{}
	ErrUserIDInvalid = errUserIDInvalid{}
	ErrNoteIDEmpty   = errNoteIDEmpty{}
	ErrNoteIDInvalid = errNoteIDInvalid{}
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

type errNoteIDEmpty struct{}

func (errNoteIDEmpty) Error() string {
	return "Error note ID empty"
}

func (errNoteIDEmpty) StatusCode() int {
	return http.StatusBadRequest
}

type errNoteIDInvalid struct{}

func (errNoteIDInvalid) Error() string {
	return "Err ID invalid"
}

func (errNoteIDInvalid) StatusCode() int {
	return http.StatusBadRequest
}
