package user

import (
	"context"

	"github.com/luanngominh/mnotes/backend/model"
)

//Service interface for user service
type Service interface {
	Create(ctx context.Context, u *model.User) (*model.User, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
}
