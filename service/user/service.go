package user

import (
	"context"

	"github.com/luanngominh/mnotes/model"
)

//Service interface for user service
type Service interface {
	Create(ctx context.Context, u *model.User) (*model.User, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Get(ctx context.Context, query *UserQuery) ([]*model.User, error)
	Active(ctx context.Context, userID, verifyCode string) (*model.User, error)
	GetOne(ctx context.Context, query *UserQuery) (*model.User, error)
}

//UserQuery ...
type UserQuery struct {
	ID    string
	Name  string
	Email string
}
