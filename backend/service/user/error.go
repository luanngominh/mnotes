package user

import "net/http"

//Error declaration
var (
	ErrNameEmpty   = errNameEmpty{}
	ErrEmailEmpty  = errEmailEmpty{}
	ErrEmailFormat = errEmailFormat{}
	ErrEmailExist  = errEmailExist{}
	ErrNameExist   = errNameExist{}
)

//Define name empty error
type errNameEmpty struct{}

func (errNameEmpty) Error() string {
	return "Name empty error"
}

func (errNameEmpty) StatusCode() int {
	return http.StatusBadRequest
}

//Define email empty error
type errEmailEmpty struct{}

func (errEmailEmpty) Error() string {
	return "Email empty error"
}

func (errEmailEmpty) StatusCode() int {
	return http.StatusBadRequest
}

//Define email format error
type errEmailFormat struct{}

func (errEmailFormat) Error() string {
	return "Email format error"
}

func (errEmailFormat) StatusCode() int {
	return http.StatusBadRequest
}

//Email existed
type errEmailExist struct{}

func (errEmailExist) Error() string {
	return "Email existed"
}

func (errEmailExist) StatusCode() int {
	return http.StatusBadRequest
}

//Name existed
type errNameExist struct{}

func (errNameExist) Error() string {
	return "Email existed"
}

func (errNameExist) StatusCode() int {
	return http.StatusBadRequest
}
