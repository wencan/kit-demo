package fasthttp

/*
 * 计算器fasthttp服务接口
 * wencan
 * 2019-06-25
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
	addRequestPool = sync.Pool{
		New: func() interface{} {
			return new(protocol.CalculatorAddRequest)
		},
	}
	subRequestPool = sync.Pool{
		New: func() interface{} {
			return new(protocol.CalculatorSubRequest)
		},
	}
	mulRequestPool = sync.Pool{
		New: func() interface{} {
			return new(protocol.CalculatorMulRequest)
		},
	}
	divRequestPool = sync.Pool{
		New: func() interface{} {
			return new(protocol.CalculatorDivRequest)
		},
	}

	acquireAddRequest = addRequestPool.Get
	acquireSubRequest = subRequestPool.Get
	acquireMulRequest = mulRequestPool.Get
	acquireDivRequest = divRequestPool.Get

	releaseAddRequest = func(req interface{}) {
		request := req.(*protocol.CalculatorAddRequest)
		request.A = 0
		request.B = 0
		addRequestPool.Put(request)
	}
	releaseSubRequest = func(req interface{}) {
		request := req.(*protocol.CalculatorSubRequest)
		request.C = 0
		request.D = 0
		subRequestPool.Put(request)
	}
	releaseMulRequest = func(req interface{}) {
		request := req.(*protocol.CalculatorMulRequest)
		request.E = 0
		request.F = 0
		mulRequestPool.Put(request)
	}
	releaseDivRequest = func(req interface{}) {
		request := req.(*protocol.CalculatorDivRequest)
		request.M = 0
		request.N = 0
		divRequestPool.Put(request)
	}
)

// RegisterCalculatorHTTPHandlers 向http router注册计算器方法处理器
func RegisterCalculatorHTTPHandlers(router *router.Router, service endpoint.CalculatorService, options ...transport.ServerOption) error {
	router.POST("/calculator/add", transport.NewCompatibleServer(
		endpoint.NewCalculatorAddEndpoint(service),
		decodeRequest,
		transport.EncodeJSONResponse,
		acquireAddRequest,
		releaseAddRequest,
		options...).ServeFastHTTP)

	router.POST("/calculator/sub", transport.NewCompatibleServer(
		endpoint.NewCalculatorSubEndpoint(service),
		decodeRequest,
		transport.EncodeJSONResponse,
		acquireSubRequest,
		releaseSubRequest,
		options...).ServeFastHTTP)

	router.POST("/calculator/mul", transport.NewCompatibleServer(
		endpoint.NewCalculatorMulEndpoint(service),
		decodeRequest,
		transport.EncodeJSONResponse,
		acquireMulRequest,
		releaseMulRequest,
		options...).ServeFastHTTP)

	router.POST("/calculator/div", transport.NewCompatibleServer(
		endpoint.NewCalculatorDivEndpoint(service),
		decodeRequest,
		transport.EncodeJSONResponse,
		acquireDivRequest,
		releaseDivRequest,
		options...).ServeFastHTTP)
	return nil
}
