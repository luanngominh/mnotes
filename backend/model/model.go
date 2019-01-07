package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//Model store model template
type Model struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//BeforeCreate generate uuid and created at before create
func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	scope.SetColumn("ID", id)
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
