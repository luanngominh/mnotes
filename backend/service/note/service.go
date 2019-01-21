package note

import (
	"context"

	"github.com/luanngominh/mnotes/backend/model"
)

//Service define note service
type Service interface {
	Create(ctx context.Context, n *model.Note) (*model.Note, error)
	Delete(ctx context.Context, n *model.Note) (*model.Note, error)
	Get(ctx context.Context, query *NoteQuery) ([]*model.Note, error)
}

//NoteQuery ...
type NoteQuery struct {
	ID     string
	UserID string
	Con    int
	Limit  int
}
