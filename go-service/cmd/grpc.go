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
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	transport "github.com/go-kit/kit/transport/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcsvc "github.com/wencan/kit-demo/go-service/cmd/grpc"
	service "github.com/wencan/kit-demo/go-service/service"
	calculator_proto "github.com/wencan/kit-demo/protocol/github.com/wencan/kit-demo/calculator/grpc_calculator_v1"
	health_proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
)

// RunGRPCServer 配置并运行GRPC服务
func RunGRPCServer(ctx context.Context, network, addr string, healthService *service.HealthService, claculatorService *service.CalculatorService, logger *zap.Logger) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	ln, err := net.Listen(network, addr)
	if err != nil {
		logger.Error("error", zap.Error(err))
		return nil
	}

	// 拦截器要注意顺序
	// 前面的嵌套后面的
	interceptors := []grpc.UnaryServerInterceptor{
		grpc_zap.UnaryServerInterceptor(logger.WithOptions(zap.AddStacktrace(zap.PanicLevel))), // 请求日志，不需要栈信息
		grpcsvc.UnaryErrorInterceptor, // 错误处理拦截器
		grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) (err error) {
			logger.Error("recover a panic", zap.Any("panic", p))
			return status.Error(codes.Internal, fmt.Sprint(p))
		})), // recovery panic
	}
	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(interceptors...))

	//
	options := []transport.ServerOption{
		transport.ServerErrorHandler(NewErrorLogHandler(logger)), // 错误日志输出。不会记录panic
	}
	health_proto.RegisterHealthServer(s, grpcsvc.NewHealthGRPCServer(healthService, options...))
	calculator_proto.RegisterCalculatorServer(s, grpcsvc.NewCalculatorGRPCServer(claculatorService, options...))

	// 服务反射
	reflection.Register(s)

	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()

	err = s.Serve(ln)
	if err != nil && err != grpc.ErrServerStopped {
		logger.Error("error", zap.Error(err))
		return err
	}
	return nil
}
