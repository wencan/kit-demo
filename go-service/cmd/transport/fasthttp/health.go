package fasthttp

/*
 * 健康检查http服务接口
 * wencan
 * 2019-06-23
 *
 * 已有endpoint和transport/grpc不支持请求/响应对象复用
 * 这里创建兼容Server，不复用响应对象
 * wencan
 * 2019-07-17
 */

import (
	"sync"

	"github.com/fasthttp/router"
	"github.com/wencan/kit-demo/go-service/cmd/endpoint"
	protocol "github.com/wencan/kit-demo/protocol/model"
	transport "github.com/wencan/kit-plugins/transport/fasthttp"
)

var (
	checkRequestPool = sync.Pool{
		New: func() interface{} {
			return new(protocol.HealthCheckRequest)
		},
	}

	acquireCheckRequest = checkRequestPool.Get
	releaseCheckRequest = func(req interface{}) {
		request := req.(*protocol.HealthCheckRequest)
		request.Service = ""
		checkRequestPool.Put(request)
	}
)

// RegisterHealthHTTPHandlers 向http router注册健康检查方法处理器
func RegisterHealthHTTPHandlers(router *router.Router, service endpoint.HealthService, options ...transport.ServerOption) error {
	router.GET("/health/check", transport.NewCompatibleServer(
		endpoint.NewHealthCheckEndpoint(service),
		decodeRequest,
		transport.EncodeJSONResponse,
		acquireCheckRequest,
		releaseCheckRequest,
		options...).ServeFastHTTP)
	return nil
}
