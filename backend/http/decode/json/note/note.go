package note

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/luanngominh/mnotes/backend/endpoints/user"
)

//CreateNoteDecode decode http request
func CreateNoteDecode(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
