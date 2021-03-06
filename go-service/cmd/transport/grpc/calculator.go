package grpc

/*
 * 计算器grpc服务接口
 *
 * wencan
 * 2019-07-01
 */

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	transport "github.com/go-kit/kit/transport/grpc"

	cmd_endpoint "github.com/wencan/kit-demo/go-service/cmd/endpoint"
	proto "github.com/wencan/kit-demo/protocol/github.com/wencan/kit-demo/calculator/grpc_calculator_v1"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// CalculatorGRPCServer 计算器GRPC服务
type CalculatorGRPCServer struct {
	AddServer transport.Handler
	SubServer transport.Handler
	MulServer transport.Handler
	DivServer transport.Handler
}

// NewCalculatorGRPCServer 创建计算器GRPC服务
func NewCalculatorGRPCServer(service cmd_endpoint.CalculatorService, middlewares endpoint.Middleware, options ...transport.ServerOption) *CalculatorGRPCServer {
	decodeAddRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorAddRequest) })
	decodeSubRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorSubRequest) })
	decodeMulRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorMulRequest) })
	decodeDivRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorDivRequest) })
	encodeInt32Response := makeResponseEncoder(func() interface{} { return new(proto.CalculatorInt32Response) })
	encodeFloatResponse := makeResponseEncoder(func() interface{} { return new(proto.CalculatorFloatResponse) })
	return &CalculatorGRPCServer{
		AddServer: transport.NewServer(middlewares(cmd_endpoint.NewCalculatorAddEndpoint(service)), decodeAddRequest, encodeInt32Response, options...),
		SubServer: transport.NewServer(middlewares(cmd_endpoint.NewCalculatorSubEndpoint(service)), decodeSubRequest, encodeInt32Response, options...),
		MulServer: transport.NewServer(middlewares(cmd_endpoint.NewCalculatorMulEndpoint(service)), decodeMulRequest, encodeInt32Response, options...),
		DivServer: transport.NewServer(middlewares(cmd_endpoint.NewCalculatorDivEndpoint(service)), decodeDivRequest, encodeFloatResponse, options...),
	}
}

func (server *CalculatorGRPCServer) Add(ctx context.Context, req *proto.CalculatorAddRequest) (*proto.CalculatorInt32Response, error) {
	_, resp, err := server.AddServer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.CalculatorInt32Response), nil
}

func (server *CalculatorGRPCServer) Sub(ctx context.Context, req *proto.CalculatorSubRequest) (*proto.CalculatorInt32Response, error) {
	_, resp, err := server.SubServer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.CalculatorInt32Response), nil
}

func (server *CalculatorGRPCServer) Mul(ctx context.Context, req *proto.CalculatorMulRequest) (*proto.CalculatorInt32Response, error) {
	_, resp, err := server.MulServer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.CalculatorInt32Response), nil
}

func (server *CalculatorGRPCServer) Div(ctx context.Context, req *proto.CalculatorDivRequest) (*proto.CalculatorFloatResponse, error) {
	_, resp, err := server.DivServer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.CalculatorFloatResponse), nil
}
