package grpc

/*
 * 计算器GRPC传输层客户端
 *
 * wencan
 * 2019-07-08
 */

import (
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
	transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	proto "github.com/wencan/kit-demo/protocol/github.com/wencan/kit-demo/calculator/grpc_calculator_v1"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

const (
	calculatorServiceName = "grpc.calculator.v1.Calculator"
)

var (
	decodeInt32Response = makeResponseDecoder(func() interface{} { return new(protocol.CalculatorInt32Response) })
	decodeFloatResponse = makeResponseDecoder(func() interface{} { return new(protocol.CalculatorFloatResponse) })

	encodeAddRequest = makeRequestEncoder(func() interface{} { return new(proto.CalculatorAddRequest) })
	encodeSubRequest = makeRequestEncoder(func() interface{} { return new(proto.CalculatorSubRequest) })
	encodeMulRequest = makeRequestEncoder(func() interface{} { return new(proto.CalculatorMulRequest) })
	encodeDivRequest = makeRequestEncoder(func() interface{} { return new(proto.CalculatorDivRequest) })

	int32ResponseTemplate = new(proto.CalculatorInt32Response)
	floatResponseTemplate = new(proto.CalculatorFloatResponse)
)

// CalculatorGRPCClient 计算器GRPC传输层客户端
type CalculatorGRPCClient struct {
	conn *grpc.ClientConn

	addClient *transport.Client
	subClient *transport.Client
	mulClient *transport.Client
	divClient *transport.Client
}

// NewCalculatorGRPCClient 创建计算器GRPC传输层客户端
func NewCalculatorGRPCClient(ctx context.Context, target string) (*CalculatorGRPCClient, error) {
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	addClient := transport.NewClient(conn, calculatorServiceName, "Add", encodeAddRequest, decodeInt32Response, int32ResponseTemplate)
	subClient := transport.NewClient(conn, calculatorServiceName, "Sub", encodeSubRequest, decodeInt32Response, int32ResponseTemplate)
	mulClient := transport.NewClient(conn, calculatorServiceName, "Mul", encodeMulRequest, decodeInt32Response, int32ResponseTemplate)
	divClient := transport.NewClient(conn, calculatorServiceName, "Div", encodeDivRequest, decodeFloatResponse, floatResponseTemplate)

	return &CalculatorGRPCClient{
		conn:      conn,
		addClient: addClient,
		subClient: subClient,
		mulClient: mulClient,
		divClient: divClient,
	}, nil
}

func (client *CalculatorGRPCClient) NewAddEndpoint() endpoint.Endpoint {
	return client.addClient.Endpoint()
}

func (client *CalculatorGRPCClient) NewSubEndpoint() endpoint.Endpoint {
	return client.subClient.Endpoint()
}

func (client *CalculatorGRPCClient) NewMulEndpoint() endpoint.Endpoint {
	return client.mulClient.Endpoint()
}

func (client *CalculatorGRPCClient) NewDivEndpoint() endpoint.Endpoint {
	return client.divClient.Endpoint()
}

func (client *CalculatorGRPCClient) Close() error {
	return client.conn.Close()
}
