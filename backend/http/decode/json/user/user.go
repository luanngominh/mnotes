package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/luanngominh/mnotes/backend/endpoints/user"
)

//RigsterRequestDecode decode http request
func RigsterRequestDecode(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
