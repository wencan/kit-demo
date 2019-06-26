package endpoint

/*
 * 健康检查endpoint
 * 服务实现与协议接口的中间层
 *
 * wencan
 * 2019-06-23
 */

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/wencan/kit-demo/protocol/model"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// HealthService 健康检查服务接口
type HealthService interface {
	Check(ctx context.Context, serviceName string) (model.HealthServiceStatus, error)
}

// NewHealthCheckEndpoint 创建健康检查endpoint
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
