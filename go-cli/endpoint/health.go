package endpoint

/*
 * 健康检查客户端endpoint
 *
 * wencan
 * 2019-07-08
 */

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	protocol "github.com/wencan/kit-demo/protocol/model"
)

// HealthEndpoints 健康检查客户端endpoint集
// endpoint需要transport赋值
type HealthEndpoints struct {
	CheckEndpoint endpoint.Endpoint
}

// Check 检查指定服务的健康状态
func (endpoints *HealthEndpoints) Check(ctx context.Context, service string) (protocol.HealthServiceStatus, error) {
	req := &protocol.HealthCheckRequest{
		Service: service,
	}
	resp, err := endpoints.CheckEndpoint(ctx, req)
	if err != nil {
		return protocol.HealthServiceStatusUnknown, err
	}
	response := resp.(*protocol.HealthCheckResponse)
	return response.Status, nil
}
