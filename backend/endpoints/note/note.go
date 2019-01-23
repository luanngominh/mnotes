package note

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	uuid "github.com/satori/go.uuid"

	"github.com/luanngominh/mnotes/backend/model"
	"github.com/luanngominh/mnotes/backend/service"
	"github.com/luanngominh/mnotes/backend/service/note"
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

//GetNoteRequest ...
type GetNoteRequest struct {
	Continue int `json:"con"`
	Limit    int `json:"limit"`
}

//GetNoteResponse ...
type GetNoteResponse struct {
	Notes []*model.Note `json:"notes"`
}

//MakeGetNoteEndpoints ...
func MakeGetNoteEndpoints(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNoteRequest)

		//Get authinfo from context, parent by note middleware
		authInfo := (*ctx.Value("auth_info").(*jwt.Claims)).(jwt.MapClaims)

		notes, err := s.NoteService.Get(ctx, &note.NoteQuery{
			UserID: authInfo["id"].(string),
			Con:    req.Continue,
			Limit:  req.Limit,
		})

		if err != nil {
			return nil, err
		}

		return GetNoteResponse{notes}, nil
	}
}

//GetAllNoteResponse ...
type GetAllNoteResponse struct {
	Notes []*model.Note `json:"notes"`
}

//MakeGetAllNoteEndpoints ...
func MakeGetAllNoteEndpoints(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//Get authinfo from context, parent by note middleware
		authInfo := (*ctx.Value("auth_info").(*jwt.Claims)).(jwt.MapClaims)

		notes, err := s.NoteService.GetAll(ctx, authInfo["id"].(string))
		if err != nil {
			return nil, ErrGetAllNote
		}

		return GetAllNoteResponse{notes}, nil
	}
}

//DeleteNoteRequest ...
type DeleteNoteRequest struct {
	ID string `json:"note_id"`
}

//DeleteNoteResponse ...
type DeleteNoteResponse struct {
	Status string `json:"status"`
}

//MakeDeleteNoteEndpoints ...
func MakeDeleteNoteEndpoints(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteNoteRequest)

		// pp.Println(req)
		err := s.NoteService.Delete(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return DeleteNoteResponse{Status: "success"}, nil
	}
}
