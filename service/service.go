package service

import (
	"github.com/luanngominh/mnotes/service/note"
	"github.com/luanngominh/mnotes/service/user"
)

//Service define service
type Service struct {
	UserSerivce user.Service
	NoteService note.Service
}
