package endpoints

import "github.com/luanngominh/mnotes/backend/service"

//Endpoints ...
type Endpoints struct {
	//Method endpoint.Endpoint
}

//MakeServerEndpoints create endpoint for service
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{}
}
