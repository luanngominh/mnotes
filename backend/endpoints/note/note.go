package note

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/luanngominh/mnotes/backend/service"
)

//MakeCreateEndpoint create note request
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}
