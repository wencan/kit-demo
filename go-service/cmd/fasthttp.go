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
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	httpsvc "github.com/wencan/kit-demo/go-service/cmd/transport/fasthttp"
	service "github.com/wencan/kit-demo/go-service/service"
	transport "github.com/wencan/kit-plugins/transport/fasthttp"
	"go.uber.org/zap"
)

func buildZapFields(c context.Context, req *fasthttp.Request, resp *fasthttp.Response) []zap.Field {
	ctx := c.Value(transport.ContextKeyRequestCtx).(*fasthttp.RequestCtx)

	fields := []zap.Field{}
	fields = append(fields, zap.String("remote_addr", ctx.RemoteAddr().String()))
	// user, _, ok := req.BasicAuth()
	// if ok {
	// 	fields = append(fields, zap.String("remote_user", user))
	// }
	fields = append(fields, zap.Time("time_local", ctx.Time()))
	fields = append(fields, zap.ByteString("request_method", req.Header.Method()))
	fields = append(fields, zap.ByteString("request_uri", req.RequestURI()))
	// fields = append(fields, zap.String("server_protocol", req.Proto))
	fields = append(fields, zap.Int("status", resp.StatusCode()))
	fields = append(fields, zap.Int("body_bytes_sent", len(resp.Body())))
	if req.Header.Referer() != nil {
		fields = append(fields, zap.ByteString("http_referer", req.Header.Referer()))
	}
	if req.Header.UserAgent() != nil {
		fields = append(fields, zap.ByteString("http_user_agent", req.Header.UserAgent()))
	}
	fields = append(fields, zap.String("elapsed_time", time.Now().Sub(ctx.Time()).String()))
	return fields
}

// NewHandlerOnHTTP 创建基于HTTP的服务处理器
func NewHandlerOnHTTP(ctx context.Context, healthService *service.HealthService, claculatorService *service.CalculatorService, logger *zap.Logger) (fasthttp.RequestHandler, error) {
	options := []transport.ServerOption{
		transport.ServerErrorHandler(NewErrorLogHandler(logger)), // 错误日志输出。不会记录panic
		transport.ServerFinalizer(func(ctx context.Context, req *fasthttp.Request, resp *fasthttp.Response, err error) {
			fields := buildZapFields(ctx, req, resp)
			if err != nil {
				logger.With(zap.Error(err)).Error(http.StatusText(resp.StatusCode()), fields...)
			} else {
				logger.Info(http.StatusText(resp.StatusCode()), fields...)
			}
		}), // 请求日志
	}
	router := router.New()
	httpsvc.RegisterHealthHTTPHandlers(router, healthService, options...)
	httpsvc.RegisterCalculatorHTTPHandlers(router, claculatorService, options...)

	return router.Handler, nil
}
