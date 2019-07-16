package cmd

/*
 * 配置并运行GRPC服务
 *
 * wencan
 * 2019-07-03
 */

import (
	"context"
	"fmt"

	transport "github.com/go-kit/kit/transport/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	grpcsvc "github.com/wencan/kit-demo/go-service/cmd/transport/grpc"
	service "github.com/wencan/kit-demo/go-service/service"
	calculator_proto "github.com/wencan/kit-demo/protocol/github.com/wencan/kit-demo/calculator/grpc_calculator_v1"
	health_proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
)

// NewHandlerOnGRPC 创建基于GRPC的服务处理器
func NewHandlerOnGRPC(ctx context.Context, healthService *service.HealthService, claculatorService *service.CalculatorService, logger *zap.Logger) (fasthttp.RequestHandler, error) {
	// 拦截器要注意顺序
	// 前面的嵌套后面的
	interceptors := []grpc.UnaryServerInterceptor{
		grpc_zap.UnaryServerInterceptor(logger.WithOptions(zap.AddStacktrace(zap.PanicLevel))), // 请求日志，不需要栈信息
		grpcsvc.UnaryErrorInterceptor, // 错误处理拦截器
		grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) (err error) {
			logger.Error("recover a panic", zap.Any("panic", p), zap.Stack("stack"))
			return status.Error(codes.Internal, fmt.Sprint(p))
		})), // recovery panic
	}
	server := grpc.NewServer(grpc_middleware.WithUnaryServerChain(interceptors...))

	//
	options := []transport.ServerOption{
		transport.ServerErrorHandler(NewErrorLogHandler(logger)), // 错误日志输出。不会记录panic
	}
	health_proto.RegisterHealthServer(server, grpcsvc.NewHealthGRPCServer(healthService, options...))
	calculator_proto.RegisterCalculatorServer(server, grpcsvc.NewCalculatorGRPCServer(claculatorService, options...))

	// 服务反射
	reflection.Register(server)

	return fasthttpadaptor.NewFastHTTPHandlerFunc(server.ServeHTTP), nil
}
