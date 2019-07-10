package http

/*
 * 计算器HTTP客户端实现
 *
 * wencan
 * 2019-07-08
 */

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	transport "github.com/go-kit/kit/transport/http"

	protocol "github.com/wencan/kit-demo/protocol/model"
)

var (
	decodeInt32Response = makeResponseDecoder(func() interface{} { return new(protocol.CalculatorInt32Response) })
	decodeFloatResponse = makeResponseDecoder(func() interface{} { return new(protocol.CalculatorFloatResponse) })
)

// CalculatorHTTPClient 计算器传输层客户端
type CalculatorHTTPClient struct {
	addClient *transport.Client
	subClient *transport.Client
	mulClient *transport.Client
	divClient *transport.Client
}

// NewCalculatorHTTPClient 创建计算器HTTP传输层客户端
func NewCalculatorHTTPClient(ctx context.Context, target string) (*CalculatorHTTPClient, error) {
	// url解析必需scheme
	if len(strings.Split(target, "://")) != 2 {
		target = "http://" + target
	}
	addrURL, err := newAddrURL(target)
	if err != nil {
		return nil, err
	}

	addURL, err := addrURL.Join("/calculator/add")
	if err != nil {
		return nil, err
	}
	addClient := transport.NewClient(http.MethodPost, addURL, encodeFormRequest, decodeInt32Response)

	subURL, err := addrURL.Join("/calculator/sub")
	if err != nil {
		return nil, err
	}
	subClient := transport.NewClient(http.MethodPost, subURL, encodeFormRequest, decodeInt32Response)

	mulURL, err := addrURL.Join("/calculator/mul")
	if err != nil {
		return nil, err
	}
	mulClient := transport.NewClient(http.MethodPost, mulURL, encodeFormRequest, decodeInt32Response)

	divURL, err := addrURL.Join("/calculator/div")
	if err != nil {
		return nil, err
	}
	divClient := transport.NewClient(http.MethodPost, divURL, encodeFormRequest, decodeFloatResponse)

	return &CalculatorHTTPClient{
		addClient: addClient,
		subClient: subClient,
		mulClient: mulClient,
		divClient: divClient,
	}, nil
}

func (client *CalculatorHTTPClient) NewAddEndpoint() endpoint.Endpoint {
	return client.addClient.Endpoint()
}

func (client *CalculatorHTTPClient) NewSubEndpoint() endpoint.Endpoint {
	return client.subClient.Endpoint()
}

func (client *CalculatorHTTPClient) NewMulEndpoint() endpoint.Endpoint {
	return client.mulClient.Endpoint()
}

func (client *CalculatorHTTPClient) NewDivEndpoint() endpoint.Endpoint {
	return client.divClient.Endpoint()
}

func (client *CalculatorHTTPClient) Close() error {
	return nil
}
