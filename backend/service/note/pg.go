package note

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/luanngominh/mnotes/backend/model"
)

type pgService struct {
	db *gorm.DB
}

//NewPGService create new pgservice
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func (s *pgService) Create(_ context.Context, n *model.Note) (*model.Note, error) {
	return n, s.db.Create(n).Error
}

func (s *pgService) Delete(_ context.Context, n *model.Note) (*model.Note, error) {
	return nil, nil
}

func (s *pgService) Get(_ context.Context, query *NoteQuery) ([]*model.Note, error) {
	return nil, nil
}
