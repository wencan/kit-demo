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

	cli_endpoint "github.com/wencan/kit-demo/go-cli/endpoint"

	transport "github.com/go-kit/kit/transport/grpc"
	proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
	protocol "github.com/wencan/kit-demo/protocol/model"
	"google.golang.org/grpc"
)

// NewHealthGRPCClient 创建健康检查GRPC客户端
func NewHealthGRPCClient(ctx context.Context, target string) (*cli_endpoint.HealthEndpoints, error) {
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	service := "grpc.health.v1.Health" // 暂时没找到获取的方法

	encodeCheckRequest := makeRequestEncoder(func() interface{} { return new(proto.HealthCheckRequest) })
	decodeCheckResponse := makeResponseDecoder(func() interface{} { return new(protocol.HealthCheckResponse) })
	checkClient := transport.NewClient(conn, service, "Check", encodeCheckRequest, decodeCheckResponse, new(proto.HealthCheckResponse))

	return &cli_endpoint.HealthEndpoints{
		CheckEndpoint: checkClient.Endpoint(),
	}, nil
}
