package note

import (
	"context"

	"github.com/luanngominh/mnotes/backend/model"
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

func (mw validationMiddleware) Create(ctx context.Context, n *model.Note) (*model.Note, error) {
	if n.Title == "" {
		return nil, ErrTitleEmpty
	}

	if n.Body == "" {
		return nil, ErrBodyEmpty
	}

	return mw.Service.Create(ctx, n)
}
