package transport

/*
 * 计算器传输层
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

// CalculatorTransport 传输层传输层客户端接口
type CalculatorTransport interface {
	io.Closer
	NewAddEndpoint() endpoint.Endpoint
	NewSubEndpoint() endpoint.Endpoint
	NewMulEndpoint() endpoint.Endpoint
	NewDivEndpoint() endpoint.Endpoint
}

// CalculatorTransportFactory 健康检查传输层客户端factory
type CalculatorTransportFactory func(ctx context.Context, target string) (CalculatorTransport, error)

// NewCalculatorClient 创建计算器客户端
// 这一步集成服务发现、负载均衡、失败重试
func NewCalculatorClient(transportFactory CalculatorTransportFactory, instancer sd.Instancer, logger kit_log.Logger) *cli_endpoint.CalculatorEndpoints {
	// 加法
	addFactory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		transport, err := transportFactory(context.Background(), instance)
		if err != nil {
			return nil, nil, err
		}
		return transport.NewAddEndpoint(), transport, nil
	}
	addEndpointer := sd.NewEndpointer(instancer, addFactory, logger)
	addBalancer := lb.NewRandom(addEndpointer, time.Now().UnixNano()) // 随机。轮转第一个不是随机的，不适合于cli
	addEndpoint := lb.Retry(3, 3*time.Second, addBalancer)            // Retry返回一个封装的Endpoint

	// 减法
	subFactory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		transport, err := transportFactory(context.Background(), instance)
		if err != nil {
			return nil, nil, err
		}
		return transport.NewSubEndpoint(), transport, nil
	}
	subEndpointer := sd.NewEndpointer(instancer, subFactory, logger)
	subBalancer := lb.NewRandom(subEndpointer, time.Now().UnixNano()) // 随机。轮转第一个不是随机的，不适合于cli
	subEndpoint := lb.Retry(3, 3*time.Second, subBalancer)            // Retry返回一个封装的Endpoint

	// 乘法
	mulFactory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		transport, err := transportFactory(context.Background(), instance)
		if err != nil {
			return nil, nil, err
		}
		return transport.NewMulEndpoint(), transport, nil
	}
	mulEndpointer := sd.NewEndpointer(instancer, mulFactory, logger)
	mulBalancer := lb.NewRandom(mulEndpointer, time.Now().UnixNano()) // 随机。轮转第一个不是随机的，不适合于cli
	mulEndpoint := lb.Retry(3, 3*time.Second, mulBalancer)            // Retry返回一个封装的Endpoint

	// 除法
	divFactory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		transport, err := transportFactory(context.Background(), instance)
		if err != nil {
			return nil, nil, err
		}
		return transport.NewDivEndpoint(), transport, nil
	}
	divEndpointer := sd.NewEndpointer(instancer, divFactory, logger)
	divBalancer := lb.NewRandom(divEndpointer, time.Now().UnixNano()) // 随机。轮转第一个不是随机的，不适合于cli
	divEndpoint := lb.Retry(3, 3*time.Second, divBalancer)            // Retry返回一个封装的Endpoint

	return &cli_endpoint.CalculatorEndpoints{
		AddEndpoint: addEndpoint,
		SubEndpoint: subEndpoint,
		MulEndpoint: mulEndpoint,
		DivEndpoint: divEndpoint,
	}
}
