package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/luanngominh/mnotes/backend/model"
	"github.com/luanngominh/mnotes/backend/service"
	"github.com/luanngominh/mnotes/backend/util"
)

// CreateUserRequest ...
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserResponse ...
type CreateUserResponse struct {
	Status string     `json:"status"`
	User   model.User `json:"user"`
}

//MakeRegisterEndpoint ...
func MakeRegisterEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)

		hashed, err := util.HashPassword(req.Password)

		if err != nil {
			return CreateUserResponse{Status: "error"}, err
		}

		verifyToken := util.GenerateToken()

		user := model.User{
			Name:       req.Name,
			Email:      req.Email,
			HashedPass: hashed,
			Verify:     verifyToken,
			Status:     model.UserInactive,
		}

		user, err = s.UserSerivce.Create(ctx, &user)
		if err != nil {
			return CreateUserResponse{Status: "error"}, err
		}

		return CreateUserResponse{Status: "success", User: user}, nil
	}
}
