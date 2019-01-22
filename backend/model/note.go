package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Note store note information
type Note struct {
	//Line below which is called embed struct
	Model

	UserID uuid.UUID  `json:"user_id"`
	Title  string     `json:"title"`
	Body   string     `json:"body"`
	Expire *time.Time `json:"expire,omitempty"`
}
