package note

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/luanngominh/mnotes/backend/model"
	"github.com/luanngominh/mnotes/backend/service"
	uuid "github.com/satori/go.uuid"
)

//CreateNoteRequest ...
type CreateNoteRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Title  string    `json:"title"`
	Body   string    `json:"body"`
}

//Response ...
type Response struct {
	Note model.Note `json:"note"`
}

//MakeCreateEndpoint create note request
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateNoteRequest)
		n := &model.Note{
			Title: req.Title,
			Body:  req.Body,
		}

		note, err := s.NoteService.Create(ctx, n)
		if err != nil {
			return nil, err
		}

		return Response{Note: *note}, nil
	}
}
