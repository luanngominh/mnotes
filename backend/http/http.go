package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/luanngominh/mnotes/backend/endpoints"
	noteDecode "github.com/luanngominh/mnotes/backend/http/decode/json/note"
	userDecode "github.com/luanngominh/mnotes/backend/http/decode/json/user"
)

//NewHTTPHandler ...
func NewHTTPHandler(endpoints endpoints.Endpoints,
	logger log.Logger,
	userCORS bool) http.Handler {

	r := chi.NewRouter()

	if userCORS {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		})
		r.Use(cors.Handler)
	}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Get("/_warm", httptransport.NewServer(
		endpoint.Nop,
		httptransport.NopRequestDecoder,
		httptransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", httptransport.NewServer(
			endpoints.CreateUser,
			userDecode.RigsterRequestDecode,
			encodeResponse,
			options...,
		).ServeHTTP)

		r.Put("/verify", httptransport.NewServer(
			endpoints.VerifyUser,
			userDecode.VerifyRequestDecode,
			encodeResponse,
			options...,
		).ServeHTTP)

		r.Post("/login", httptransport.NewServer(
			endpoints.Login,
			userDecode.LoginRequestDecode,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	r.Route("/note", func(r chi.Router) {
		r.Use(NoteMiddleware)

		r.Post("/", httptransport.NewServer(
			endpoints.CreateNote,
			noteDecode.CreateNoteDecode,
			encodeResponse,
			options...,
		).ServeHTTP)

		r.Get("/{user_id}", httptransport.NewServer(
			endpoints.GetAllNote,
			httptransport.NopRequestDecoder,
			encodeJSONError,
			options...,
		).ServeHTTP)

		//get_query: con=0&limit=2 encode by base64
		r.Get("/{get_query}", httptransport.NewServer(
			endpoints.GetNote,
			noteDecode.GetNoteDecode,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	return r
}
