package transport

/*
 * 健康检查传输层
 * endpoint和特定协议传输层的中间层
 * 实现了服务发现、负载均衡、失败重试
 *
 * wencan
 * 2019-07-10
 */

import (
	"context"
	"io"
	"time"

	"github.com/go-kit/kit/endpoint"
	kit_log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"

	cli_endpoint "github.com/wencan/kit-demo/go-cli/endpoint"
)

// HealthTransport 健康检查传输层客户端接口
type HealthTransport interface {
	io.Closer
	NewCheckEndpoint() endpoint.Endpoint
}

// HealthTransportFactory 健康检查传输层客户端factory
type HealthTransportFactory func(ctx context.Context, target string) (HealthTransport, error)

// NewHealthClient 创建健康检查客户端
// 这一步集成服务发现、负载均衡、失败重试
func NewHealthClient(transportFactory HealthTransportFactory, instancer sd.Instancer, logger kit_log.Logger) *cli_endpoint.HealthEndpoints {
	checkFactory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		transport, err := transportFactory(context.Background(), instance)
		if err != nil {
			return nil, nil, err
		}
		return transport.NewCheckEndpoint(), transport, nil
	}
	checkEndpointer := sd.NewEndpointer(instancer, checkFactory, logger)
	checkBalancer := lb.NewRandom(checkEndpointer, time.Now().UnixNano()) // 随机。轮转第一个不是随机的，不适合于cli
	checkEndpoint := lb.Retry(3, 3*time.Second, checkBalancer)            // Retry返回一个封装的Endpoint

	return &cli_endpoint.HealthEndpoints{
		CheckEndpoint: checkEndpoint,
	}
}
