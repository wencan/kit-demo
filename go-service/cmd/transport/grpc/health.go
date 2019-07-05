package grpc

/*
 * 健康检查grpc服务接口
 * wencan
 * 2019-06-23
 */

import (
	"context"

	transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/wencan/kit-demo/go-service/cmd/endpoint"
	proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

var (
	// _HealthServiceStatusCodes 健康检查服务状态码映射关联
	_HealthServiceStatusCodes = map[protocol.HealthServiceStatus]proto.HealthCheckResponse_ServingStatus{
		protocol.HealthServiceStatusUnknown:        proto.HealthCheckResponse_UNKNOWN,
		protocol.HealthServiceStatusServing:        proto.HealthCheckResponse_SERVING,
		protocol.HealthServiceStatusNotServing:     proto.HealthCheckResponse_NOT_SERVING,
		protocol.HealthServiceStatusServiceUnknown: proto.HealthCheckResponse_SERVICE_UNKNOWN,
	}
)

// HealthGRPCServer 健康检查GRPC服务
type HealthGRPCServer struct {
	CheckServer transport.Handler
}

// NewHealthGRPCServer 创建健康检查GRPC服务
func NewHealthGRPCServer(service endpoint.HealthService, options ...transport.ServerOption) *HealthGRPCServer {
	decodeCheckRequest := makeRequestDecoder(func() interface{} { return new(protocol.HealthCheckRequest) })
	encodeCheckResponse := makeResponseEncoder(func() interface{} { return new(proto.HealthCheckResponse) })
	return &HealthGRPCServer{
		CheckServer: transport.NewServer(endpoint.NewHealthCheckEndpoint(service), decodeCheckRequest, encodeCheckResponse, options...),
	}
}

// Check 检查指定服务的健康状态
func (server *HealthGRPCServer) Check(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	_, resp, err := server.CheckServer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.HealthCheckResponse), nil
}

// Watch 观察指定服务的健康状态。未实现
func (server *HealthGRPCServer) Watch(req *proto.HealthCheckRequest, watcher proto.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "kit not suport streaming request and response")
}
