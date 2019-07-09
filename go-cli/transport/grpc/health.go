package grpc

/*
 * 健康检查GRPC传输层客户端
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

const (
	healthServiceName = "grpc.health.v1.Health"
)

var (
	encodeCheckRequest  = makeRequestEncoder(func() interface{} { return new(proto.HealthCheckRequest) })
	decodeCheckResponse = makeResponseDecoder(func() interface{} { return new(protocol.HealthCheckResponse) })

	// grpcCheckResponseTemplate kit用这个来推断响应对象类型
	grpcCheckResponseTemplate = new(proto.HealthCheckResponse)
)

// HealthGRPCClient 健康检查传输层客户端
type HealthGRPCClient struct {
	conn *grpc.ClientConn
}

// NewHealthGRPCClient 创建健康检查GRPC传输层客户端
func NewHealthGRPCClient(ctx context.Context, target string) (*HealthGRPCClient, error) {
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &HealthGRPCClient{
		conn: conn,
	}, nil
}

func (client *HealthGRPCClient) NewCheckEndpoint() endpoint.Endpoint {
	checkClient := transport.NewClient(client.conn, healthServiceName, "Check", encodeCheckRequest, decodeCheckResponse, grpcCheckResponseTemplate)
	return checkClient.Endpoint()
}

// Close 关闭连接
// 需要确保客户端对象创建的endpoint全部使用完毕
// 实际由负载均衡逻辑调用Close
func (client *HealthGRPCClient) Close() error {
	return client.conn.Close()
}
