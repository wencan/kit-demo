package grpc

/*
 * 健康检查GRPC客户端实现
 *
 * wencan
 * 2019-07-08
 */

import (
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
	transport "github.com/go-kit/kit/transport/grpc"
	proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
	protocol "github.com/wencan/kit-demo/protocol/model"
	"google.golang.org/grpc"
)

// HealthGRPCClient 健康检查GRPC客户端
type HealthGRPCClient struct {
	CheckEndpoint endpoint.Endpoint
}

// NewHealthGRPCClient 创建健康检查GRPC客户端
func NewHealthGRPCClient(ctx context.Context, target string) (*HealthGRPCClient, error) {
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	service := "grpc.health.v1.Health" // 暂时没找到获取的方法

	encodeCheckRequest := makeRequestEncoder(func() interface{} { return new(proto.HealthCheckRequest) })
	decodeCheckResponse := makeResponseDecoder(func() interface{} { return new(protocol.HealthCheckResponse) })
	checkClient := transport.NewClient(conn, service, "Check", encodeCheckRequest, decodeCheckResponse, new(proto.HealthCheckResponse))

	return &HealthGRPCClient{
		CheckEndpoint: checkClient.Endpoint(),
	}, nil
}

// Check 检查指定服务的健康状态
func (client *HealthGRPCClient) Check(ctx context.Context, service string) (protocol.HealthServiceStatus, error) {
	req := &protocol.HealthCheckRequest{
		Service: service,
	}
	resp, err := client.CheckEndpoint(ctx, req)
	if err != nil {
		return protocol.HealthServiceStatusUnknown, err
	}
	response := resp.(*protocol.HealthCheckResponse)
	return response.Status, nil
}
