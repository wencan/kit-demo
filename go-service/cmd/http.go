package cmd

/*
 * 配置并运行HTTP服务
 *
 * wencan
 * 2019-07-03
 */

import (
	"context"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wencan/middlewares"
	http_zap "github.com/wencan/middlewares/logging/zap"
	"go.uber.org/zap"

	httpsvc "github.com/wencan/kit-demo/go-service/cmd/transport/http"
	service "github.com/wencan/kit-demo/go-service/service"
)

// NewHTTPHandler 创建HTTP服务处理器
func NewHTTPHandler(ctx context.Context, healthService *service.HealthService, claculatorService *service.CalculatorService, logger *zap.Logger) (http.Handler, error) {
	//
	options := []transport.ServerOption{
		transport.ServerErrorHandler(NewErrorLogHandler(logger)), // 错误日志输出。不会记录panic
	}
	router := httprouter.New()
	httpsvc.RegisterHealthHTTPHandlers(router, healthService, options...)
	httpsvc.RegisterCalculatorHTTPHandlers(router, claculatorService, options...)

	middleware := middlewares.Chain(
		middlewares.LoggingMiddleware(http_zap.NewLogger(logger.WithOptions(zap.AddStacktrace(zap.PanicLevel)))), // 记录请求日志，不需要栈信息
		middlewares.RecoverMiddleware(middlewares.WithRecoveryHandlerOption(func(w http.ResponseWriter, r *http.Request, recovery interface{}) {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("recover a panic", zap.Any("panic", recovery))
		})), // recover panic
	)

	return middleware(router), nil
}
