package service

import (
	"github.com/luanngominh/mnotes/backend/service/note"
	"github.com/luanngominh/mnotes/backend/service/user"
)

//Service define service
type Service struct {
	UserSerivce user.Service
	NoteService note.Service
}
