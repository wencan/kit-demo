package http

/*
 * 健康检查http服务接口
 * wencan
 * 2019-06-23
 */

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"

	cmd_endpoint "github.com/wencan/kit-demo/go-service/cmd/endpoint"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// RegisterHealthHTTPHandlers 向http router注册健康检查方法处理器
func RegisterHealthHTTPHandlers(router *httprouter.Router, service cmd_endpoint.HealthService, middlewares endpoint.Middleware, options ...transport.ServerOption) error {
	decodeHealthCheckRequest := makeRequestDecoder(func() interface{} { return new(protocol.HealthCheckRequest) })
	router.Handler(http.MethodGet, "/health/check", transport.NewServer(middlewares(cmd_endpoint.NewHealthCheckEndpoint(service)), decodeHealthCheckRequest, encodeResponse, options...))
	return nil
}
