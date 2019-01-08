package user

import "net/http"

//Define error message and http status code
var (
	ErrUnauthorized = errUnauthorized{}
)

type errUnauthorized struct{}

func (errUnauthorized) Error() string {
	return "Unauthorzied"
}

func (errUnauthorized) StatusCode() int {
	return http.StatusUnauthorized
}
