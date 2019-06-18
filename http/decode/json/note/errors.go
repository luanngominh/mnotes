package note

import "net/http"

//...
var (
	ErrGetQuery = errGetQuery{}
)

type errGetQuery struct{}

func (errGetQuery) Error() string {
	return "Get query param in valid"
}

func (errGetQuery) StatusCode() int {
	return http.StatusBadRequest
}
