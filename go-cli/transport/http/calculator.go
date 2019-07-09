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

	transport "github.com/go-kit/kit/transport/http"
	cli_endpoint "github.com/wencan/kit-demo/go-cli/endpoint"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// NewCalculatorHTTPClient 创建计算器HTTP客户端
func NewCalculatorHTTPClient(ctx context.Context, target string) (*cli_endpoint.CalculatorEndpoints, error) {
	// url解析必需scheme
	if len(strings.Split(target, "://")) != 2 {
		target = "http://" + target
	}
	addrURL, err := newAddrURL(target)
	if err != nil {
		return nil, err
	}

	// 两个公共的响应解码器
	decodeInt32Response := makeResponseDecoder(func() interface{} { return new(protocol.CalculatorInt32Response) })
	decodeFloatResponse := makeResponseDecoder(func() interface{} { return new(protocol.CalculatorFloatResponse) })

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

	return &cli_endpoint.CalculatorEndpoints{
		AddEndpoint: addClient.Endpoint(),
		SubEndpoint: subClient.Endpoint(),
		MulEndpoint: mulClient.Endpoint(),
		DivEndpoint: divClient.Endpoint(),
	}, nil
}
