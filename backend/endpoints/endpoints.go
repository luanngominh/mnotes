package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/luanngominh/mnotes/backend/endpoints/user"
	"github.com/luanngominh/mnotes/backend/service"
)

//Endpoints ...
type Endpoints struct {
	CreateUser endpoint.Endpoint
}

//MakeServerEndpoints create endpoint for service
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateUser: user.MakeRegisterEndpoint(s),
	}
}
