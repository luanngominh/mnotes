package user

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

func (s *pgService) Create(ctx context.Context, u *model.User) (*model.User, error) {
	return nil, s.db.Create(u).Error
}

func (s *pgService) Update(ctx context.Context, u *model.User) (*model.User, error) {

	return nil, nil
}

func (s *pgService) Get(ctx context.Context, query *userQuery) ([]*model.User, error) {
	db := s.db

	if query.Name != "" {
		db = db.Where("name = ?", query.Name)
	}

	if query.Email != "" {
		db = db.Where("email = ?", query.Email)
	}

	if query.ID != "" {
		db = db.Where("id = ?", query.ID)
	}

	user := []*model.User{}

	return user, db.Find(&user).Error
}
