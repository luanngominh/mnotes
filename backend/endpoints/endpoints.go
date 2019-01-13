package endpoints

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/luanngominh/mnotes/backend/endpoints/note"
	"github.com/luanngominh/mnotes/backend/endpoints/user"
	"github.com/luanngominh/mnotes/backend/service"
)

//Endpoints ...
type Endpoints struct {
	//User endpoint
	CreateUser endpoint.Endpoint
	VerifyUser endpoint.Endpoint
	Login      endpoint.Endpoint

	//Note endpoint
	CreateNote endpoint.Endpoint
}

//MakeServerEndpoints create endpoint for service
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateUser: user.MakeRegisterEndpoint(s),
		VerifyUser: user.MakeVerifyEndpoint(s),
		Login:      user.MakeLoginEndpoint(s),

		CreateNote: note.MakeCreateEndpoint(s),
	}
}
