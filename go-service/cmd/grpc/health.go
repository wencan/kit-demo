package grpc

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

func NewHealthGRPCServer(service endpoint.HealthService) *HealthGRPCServer {
	return &HealthGRPCServer{
		CheckServer: transport.NewServer(endpoint.NewHealthCheckEndpoint(service), decodeCheckRequest, encodeCheckResponse),
	}
}

func (server *HealthGRPCServer) Check(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	_, resp, err := server.CheckServer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.HealthCheckResponse), nil
}

func (server *HealthGRPCServer) Watch(req *proto.HealthCheckRequest, watcher proto.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "kit not suport streaming request and response")
}

func decodeCheckRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	checkReq := grpcReq.(*proto.HealthCheckRequest)
	return &protocol.HealthCheckRequest{
		Service: checkReq.Service,
	}, nil
}

func encodeCheckResponse(_ context.Context, resp interface{}) (interface{}, error) {
	checkResp := resp.(*protocol.HealthCheckResponse)
	grpcResp := &proto.HealthCheckResponse{}

	var ok bool
	grpcResp.Status, ok = _HealthServiceStatusCodes[checkResp.Status]
	if !ok {
		return nil, status.Errorf(codes.Internal, "unkown health status: %d", checkResp.Status)
	}
	return grpcResp, nil
}
