package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/julienschmidt/httprouter"
	"github.com/wencan/multihandler"
	http_zap "github.com/wencan/multihandler/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcsvc "github.com/wencan/kit-demo/go-service/cmd/grpc"
	httpsvc "github.com/wencan/kit-demo/go-service/cmd/http"
	service "github.com/wencan/kit-demo/go-service/service"
	calculator_proto "github.com/wencan/kit-demo/protocol/github.com/wencan/kit-demo/calculator/grpc_calculator_v1"
	health_proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
)

func runGRPCServer(ctx context.Context, network, addr string, healthService *service.HealthService, claculatorService *service.CalculatorService, logger *zap.Logger) error {
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
	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(grpc_zap.UnaryServerInterceptor(logger), grpc_recovery.UnaryServerInterceptor()))
	health_proto.RegisterHealthServer(s, grpcsvc.NewHealthGRPCServer(healthService))
	calculator_proto.RegisterCalculatorServer(s, grpcsvc.NewCalculatorGRPCServer(claculatorService))

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

func runHTTPServer(ctx context.Context, addr string, healthService *service.HealthService, claculatorService *service.CalculatorService, logger *zap.Logger) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	router := httprouter.New()
	httpsvc.RegisterHealthHTTPHandlers(router, healthService)
	httpsvc.RegisterCalculatorHTTPHandlers(router, claculatorService)

	middleware := multihandler.NewMultiMiddleware(http_zap.NewZapLogging(logger))

	s := http.Server{
		Addr:    addr,
		Handler: middleware(router),
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("error", zap.Error(err))
		return err
	}
	return nil
}

func main() {
	healthService := service.NewHealthService()
	claculatorService := service.NewCalculatorService()

	ctx, cancel := context.WithCancel(context.Background())

	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000000Z07:00")) // ISO 8601
	}
	logger, err := logConfig.Build()
	if err != nil {
		log.Println(err)
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		err := runGRPCServer(ctx, "tcp", "127.0.0.1:5051", healthService, claculatorService, logger)
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		time.Sleep(5)
		defer wg.Done()
		defer cancel()
		err := runHTTPServer(ctx, "127.0.0.1:6061", healthService, claculatorService, logger)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
}
