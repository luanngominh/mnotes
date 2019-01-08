package user

import "net/http"

//Error declaration
var (
	ErrNameEmpty       = errNameEmpty{}
	ErrEmailEmpty      = errEmailEmpty{}
	ErrEmailFormat     = errEmailFormat{}
	ErrEmailExist      = errEmailExist{}
	ErrNameExist       = errNameExist{}
	ErrVerifyCode      = errVerifyCode{}
	ErrUpdateStatus    = errUpdateStatus{}
	ErrVerifyCodeEmpty = errVerifyCodeEmpty{}
	ErrIDInvalid       = errIDInvalid{}
	ErrPasswordEmpty   = errPasswordEmpty{}
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
	return "Name existed"
}

func (errNameExist) StatusCode() int {
	return http.StatusBadRequest
}

//Verify code error
type errVerifyCode struct{}

func (errVerifyCode) Error() string {
	return "Verify code error"
}

func (errVerifyCode) StatusCode() int {
	return http.StatusBadRequest
}

//Update status error
type errUpdateStatus struct{}

func (errUpdateStatus) Error() string {
	return "Update status error"
}

func (errUpdateStatus) StatusCode() int {
	return http.StatusBadRequest
}

//ID empty
type errIDInvalid struct{}

func (errIDInvalid) Error() string {
	return "User ID invalid"
}

func (errIDInvalid) StatusCode() int {
	return http.StatusBadRequest
}

//Verify code empty
type errVerifyCodeEmpty struct{}

func (errVerifyCodeEmpty) Error() string {
	return "Verify code empty"
}

func (errVerifyCodeEmpty) StatusCode() int {
	return http.StatusBadRequest
}

//Password empty
type errPasswordEmpty struct{}

func (errPasswordEmpty) Error() string {
	return "Password empty"
}

func (errPasswordEmpty) StatusCode() int {
	return http.StatusBadRequest
}
