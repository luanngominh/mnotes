package user

import (
	"context"

	"github.com/luanngominh/mnotes/backend/model"
)

//Service interface for user service
type Service interface {
	Create(ctx context.Context, u *model.User) (*model.User, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Get(ctx context.Context, query *userQuery) ([]*model.User, error)
}

//UserQuery ...
type userQuery struct {
	ID    string
	Name  string
	Email string
}
