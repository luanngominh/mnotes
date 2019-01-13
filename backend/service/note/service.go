package note

import (
	"context"

	"github.com/luanngominh/mnotes/backend/model"
)

//Service define note service
type Service interface {
	Create(_ context.Context, n *model.Note) (*model.Note, error)
	Delete(_ context.Context, n *model.Note) (*model.Note, error)
	Get(_ context.Context, query *NoteQuery) ([]*model.Note, error)
}

//NoteQuery ...
type NoteQuery struct {
	ID     string
	UserID string
}
