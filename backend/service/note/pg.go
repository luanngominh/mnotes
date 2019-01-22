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

func (s *pgService) Create(ctx context.Context, n *model.Note) (*model.Note, error) {
	return n, s.db.Create(n).Error
}

func (s *pgService) Delete(ctx context.Context, n *model.Note) (*model.Note, error) {
	return nil, nil
}

func (s *pgService) Get(ctx context.Context, query *NoteQuery) ([]*model.Note, error) {
	db := s.db

	if query.ID != "" {
		db.Where("id = ?", query.ID)
	}

	if query.UserID != "" {
		db.Where("user_id = ?", query.UserID)
	}

	if query.Con == 0 && query.Limit == 0 {
		notes := []*model.Note{}
		return notes, db.Order("created_at desc").Offset(query.Con).Limit(query.Limit).
			Select("title, body, id, user_id, created_at").Find(&notes).Error
	}

	notes := []*model.Note{}
	return notes, db.Order("created_at desc").Offset(query.Con).Limit(query.Limit).
		Select("title, body, id, user_id, created_at").Find(&notes).Error
}

func (s *pgService) GetAll(ctx context.Context, userID string) ([]*model.Note, error) {
	notes := []*model.Note{}
	return notes, s.db.Where("id = ?", userID).Order("created_at desc").Find(&notes).Error
}
