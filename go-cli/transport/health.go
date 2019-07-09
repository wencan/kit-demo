package transport

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/go-kit/kit/endpoint"
	kit_log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"

	protocol "github.com/wencan/kit-demo/protocol/model"
)

// HealthTransport 健康检查传输层客户端接口
type HealthTransport interface {
	io.Closer
	NewCheckEndpoint() endpoint.Endpoint
}

// HealthTransportFactory 健康检查传输层客户端factory
type HealthTransportFactory func(ctx context.Context, target string) (HealthTransport, error)

// HealthClient 健康检查的（终极）客户端——给业务逻辑调用的客户端
// 集成了服务发现、负载均衡、失败重试
type HealthClient struct {
	checkEndpoint endpoint.Endpoint
}

func NewHealthClient(transportFactory HealthTransportFactory, instancer sd.Instancer, logger kit_log.Logger) *HealthClient {
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

	return &HealthClient{
		checkEndpoint: checkEndpoint,
	}
}

func (client *HealthClient) Check(ctx context.Context, serviceName string) (protocol.HealthServiceStatus, error) {
	req := &protocol.HealthCheckRequest{
		Service: serviceName,
	}
	resp, err := client.checkEndpoint(ctx, req)
	if err != nil {
		log.Println(err)
		return protocol.HealthServiceStatusUnknown, nil
	}
	response := resp.(*protocol.HealthCheckResponse)
	return response.Status, nil
}
