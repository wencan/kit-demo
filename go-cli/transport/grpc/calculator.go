package grpc

/*
 * 计算器GRPC客户端实现
 *
 * wencan
 * 2019-07-08
 */

import (
	"context"
	"log"

	proto "github.com/wencan/kit-demo/protocol/github.com/wencan/kit-demo/calculator/grpc_calculator_v1"

	cli_endpoint "github.com/wencan/kit-demo/go-cli/endpoint"

	transport "github.com/go-kit/kit/transport/grpc"
	protocol "github.com/wencan/kit-demo/protocol/model"
	"google.golang.org/grpc"
)

// NewCalculatorGRPCClient 创建计算器GRPC客户端
func NewCalculatorGRPCClient(ctx context.Context, target string) (*cli_endpoint.CalculatorEndpoints, error) {
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	service := "grpc.calculator.v1.Calculator"

	// 两个公共解码函数
	decodeInt32Response := makeResponseDecoder(func() interface{} { return new(protocol.CalculatorInt32Response) })
	decodeFloatResponse := makeResponseDecoder(func() interface{} { return new(protocol.CalculatorFloatResponse) })

	encodeAddRequest := makeRequestEncoder(func() interface{} { return new(proto.CalculatorAddRequest) })
	addClient := transport.NewClient(conn, service, "Add", encodeAddRequest, decodeInt32Response, new(proto.CalculatorInt32Response))

	encodeSubRequest := makeRequestEncoder(func() interface{} { return new(proto.CalculatorSubRequest) })
	subClient := transport.NewClient(conn, service, "Sub", encodeSubRequest, decodeInt32Response, new(proto.CalculatorInt32Response))

	encodeMulRequest := makeRequestEncoder(func() interface{} { return new(proto.CalculatorMulRequest) })
	mulClient := transport.NewClient(conn, service, "Mul", encodeMulRequest, decodeInt32Response, new(proto.CalculatorInt32Response))

	encodeDivRequest := makeRequestEncoder(func() interface{} { return new(proto.CalculatorDivRequest) })
	divClient := transport.NewClient(conn, service, "Div", encodeDivRequest, decodeFloatResponse, new(proto.CalculatorFloatResponse))

	return &cli_endpoint.CalculatorEndpoints{
		AddEndpoint: addClient.Endpoint(),
		SubEndpoint: subClient.Endpoint(),
		MulEndpoint: mulClient.Endpoint(),
		DivEndpoint: divClient.Endpoint(),
	}, nil
}
