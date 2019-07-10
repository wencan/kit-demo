package http

/*
 * 健康检查HTTP传输层
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
	decodeCheckResponse = makeResponseDecoder(func() interface{} { return new(protocol.HealthCheckResponse) })
)

// HealthHTTPClient 健康检查传输层客户端
type HealthHTTPClient struct {
	checkClient *transport.Client
}

// NewHealthHTTPClient 创建健康检查HTTP传输层客户端
func NewHealthHTTPClient(ctx context.Context, target string) (*HealthHTTPClient, error) {
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
	checkClient := transport.NewClient(http.MethodGet, checkURL, encodeQueryRequest, decodeCheckResponse)

	return &HealthHTTPClient{
		checkClient: checkClient,
	}, nil
}

// NewCheckEndpoint 创建check方法endpoint
func (client *HealthHTTPClient) NewCheckEndpoint() endpoint.Endpoint {
	return client.checkClient.Endpoint()
}

// Close 关闭传输层客户端
// 实现什么都没干。负载均衡的Factory需要这个
func (client *HealthHTTPClient) Close() error {
	return nil
}
