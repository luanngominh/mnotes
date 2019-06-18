package endpoints

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/luanngominh/mnotes/endpoints/note"
	"github.com/luanngominh/mnotes/endpoints/user"
	"github.com/luanngominh/mnotes/service"
)

//Endpoints ...
type Endpoints struct {
	//User endpoint
	CreateUser endpoint.Endpoint
	VerifyUser endpoint.Endpoint
	Login      endpoint.Endpoint

	//Note endpoint
	CreateNote endpoint.Endpoint
	GetNote    endpoint.Endpoint
	GetAllNote endpoint.Endpoint
	DeleteNote endpoint.Endpoint
}

//MakeServerEndpoints create endpoint for service
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateUser: user.MakeRegisterEndpoint(s),
		VerifyUser: user.MakeVerifyEndpoint(s),
		Login:      user.MakeLoginEndpoint(s),

		CreateNote: note.MakeCreateEndpoint(s),
		GetNote:    note.MakeGetNoteEndpoints(s),
		GetAllNote: note.MakeGetAllNoteEndpoints(s),
		DeleteNote: note.MakeDeleteNoteEndpoints(s),
	}
}
