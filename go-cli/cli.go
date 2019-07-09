package main

/*
 * cli的逻辑调用部分
 * 连接cobra.Command和业务逻辑
 *
 * 结果cobra的运行逻辑不太符预期，这个写得与实际cmd逻辑不太匹配
 *
 * wencan
 * 2019-07-08
 */

import (
	"context"

	grpc_transport "github.com/wencan/kit-demo/go-cli/transport/grpc"
	http_transport "github.com/wencan/kit-demo/go-cli/transport/http"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// HealthClient 健康检查客户端接口
type HealthClient interface {
	// Check 检查指定服务的健康状态
	Check(ctx context.Context, serviceName string) (protocol.HealthServiceStatus, error)
}

// CalculatorClient 计算器客户端接口
type CalculatorClient interface {
	// Add 加
	Add(ctx context.Context, a, b int32) (int32, error)

	// Sub 减
	Sub(ctx context.Context, c, d int32) (int32, error)

	// Mul 乘
	Mul(ctx context.Context, e, f int32) (int32, error)

	// Div 除
	Div(ctx context.Context, m, n int32) (float32, error)
}

// Cli 根cli
type Cli struct {
	healthClientFactory func(ctx context.Context, target string) (HealthClient, error)

	calcutorClientFactory func(ctx context.Context, target string) (CalculatorClient, error)
}

// NewCliOnGRPC 创建基于GRPC协议的cli
func NewCliOnGRPC() (*Cli, error) {
	return &Cli{
		healthClientFactory: func(ctx context.Context, target string) (HealthClient, error) {
			return grpc_transport.NewHealthGRPCClient(ctx, target)
		},
		calcutorClientFactory: func(ctx context.Context, target string) (CalculatorClient, error) {
			return grpc_transport.NewCalculatorGRPCClient(ctx, target)
		},
	}, nil
}

// NewCliOnHTTP 创建基于HTTP协议的cli
func NewCliOnHTTP() (*Cli, error) {
	return &Cli{
		healthClientFactory: func(ctx context.Context, target string) (HealthClient, error) {
			return http_transport.NewHealthHTTPClient(ctx, target)
		},
		calcutorClientFactory: func(ctx context.Context, target string) (CalculatorClient, error) {
			return http_transport.NewCalculatorHTTPClient(ctx, target)
		},
	}, nil
}

// NewHealthCli 创建健康检查cli
func (cli *Cli) NewHealthCli(ctx context.Context, target string) (*HealthCli, error) {
	client, err := cli.healthClientFactory(ctx, target)
	if err != nil {
		return nil, err
	}
	return &HealthCli{
		HealthClient: client,
	}, nil
}

// NewCalculatorCli 创建计算器cli
func (cli *Cli) NewCalculatorCli(ctx context.Context, target string) (*CalculatorCli, error) {
	client, err := cli.calcutorClientFactory(ctx, target)
	if err != nil {
		return nil, err
	}
	return &CalculatorCli{
		CalculatorClient: client,
	}, nil
}

// HealthCli 健康检查cli
type HealthCli struct {
	HealthClient
}

// CalculatorCli 计算器cli
type CalculatorCli struct {
	CalculatorClient
}
