package note

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	uuid "github.com/satori/go.uuid"

	"github.com/luanngominh/mnotes/backend/model"
	"github.com/luanngominh/mnotes/backend/service"
)

//CreateNoteRequest ...
type CreateNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

//Response ...
type Response struct {
	Note model.Note `json:"note"`
}

//MakeCreateEndpoint create note request
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateNoteRequest)

		//Get authinfo from context, parent by note middleware
		authInfo := (*ctx.Value("auth_info").(*jwt.Claims)).(jwt.MapClaims)

		userID, err := uuid.FromString(authInfo["id"].(string))
		if err != nil {
			return nil, ErrUserIDInvalid
		}

		n := &model.Note{
			UserID: userID,
			Title:  req.Title,
			Body:   req.Body,
		}

		note, err := s.NoteService.Create(ctx, n)
		if err != nil {
			return nil, err
		}

		return Response{Note: *note}, nil
	}
}
