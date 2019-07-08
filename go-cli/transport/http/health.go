package http

/*
 * 健康检查HTTP客户端实现
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

// NewHealthHTTPClient 创建健康检查HTTP客户端
func NewHealthHTTPClient(ctx context.Context, target string) (*cli_endpoint.HealthEndpoints, error) {
	// url解析必需scheme
	if len(strings.Split(target, "://")) != 2 {
		target = "http://" + target
	}
	addrURL, err := newAddrURL(target)
	if err != nil {
		return nil, err
	}

	checkURL, err := addrURL.Join("/health/check")
	if err != nil {
		return nil, err
	}
	decodeCheckResponse := makeResponseDecoder(func() interface{} { return new(protocol.HealthCheckResponse) })
	checkClient := transport.NewClient(http.MethodGet, checkURL, encodeQueryRequest, decodeCheckResponse)

	return &cli_endpoint.HealthEndpoints{
		CheckEndpoint: checkClient.Endpoint(),
	}, nil
}
