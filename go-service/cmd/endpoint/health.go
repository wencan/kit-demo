package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/wencan/kit-demo/protocol/model"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

type HealthService interface {
	Check(ctx context.Context, serviceName string) (model.HealthServiceStatus, error)
}

func NewHealthCheckEndpoint(service HealthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*protocol.HealthCheckRequest)
		status, err := service.Check(ctx, req.Service)
		if err != nil {
			return nil, err
		}
		return &protocol.HealthCheckResponse{
			Status: status,
		}, nil
	}
}
