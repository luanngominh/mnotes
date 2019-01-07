package user

import (
	"context"
	"regexp"

	"github.com/luanngominh/mnotes/backend/model"
)

// Declare Regex
const (
	emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

type validationMiddleware struct {
	Service
}

//ValidationMiddleware ...
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, u *model.User) (*model.User, error) {
	if u.Name == "" {
		return nil, ErrNameEmpty
	}

	if u.Email == "" {
		return nil, ErrEmailEmpty
	}

	emailRegex, _ := regexp.Compile(emailRegex)
	if !emailRegex.MatchString(u.Email) {
		return nil, ErrEmailFormat
	}

	//Check email is unique
	//comment for test email
	// users, err := mw.Get(ctx, &userQuery{Email: u.Email})
	// if err != nil {
	// 	return nil, err
	// }

	// if len(users) != 0 {
	// 	return nil, ErrEmailExist
	// }

	//Check name is unique
	users, err := mw.Get(ctx, &userQuery{Name: u.Name})
	if err != nil {
		return nil, err
	}

	if len(users) != 0 {
		return nil, ErrNameExist
	}

	return mw.Service.Create(ctx, u)
}
