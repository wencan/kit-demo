package http

/*
 * 计算器http服务接口
 * wencan
 * 2019-06-25
 */

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"

	cmd_endpoint "github.com/wencan/kit-demo/go-service/cmd/endpoint"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// RegisterCalculatorHTTPHandlers 向http router注册计算器方法处理器
func RegisterCalculatorHTTPHandlers(router *httprouter.Router, service cmd_endpoint.CalculatorService, middlewares endpoint.Middleware, options ...transport.ServerOption) error {
	decodeCalculatorAddRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorAddRequest) })
	router.Handler(http.MethodPost, "/calculator/add", transport.NewServer(middlewares(cmd_endpoint.NewCalculatorAddEndpoint(service)), decodeCalculatorAddRequest, encodeResponse, options...))

	decodeCalculatorSubRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorSubRequest) })
	router.Handler(http.MethodPost, "/calculator/sub", transport.NewServer(middlewares(cmd_endpoint.NewCalculatorSubEndpoint(service)), decodeCalculatorSubRequest, encodeResponse, options...))

	decodeCalculatorMulRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorMulRequest) })
	router.Handler(http.MethodPost, "/calculator/mul", transport.NewServer(middlewares(cmd_endpoint.NewCalculatorMulEndpoint(service)), decodeCalculatorMulRequest, encodeResponse, options...))

	decodeCalculatorDivRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorDivRequest) })
	router.Handler(http.MethodPost, "/calculator/div", transport.NewServer(middlewares(cmd_endpoint.NewCalculatorDivEndpoint(service)), decodeCalculatorDivRequest, encodeResponse, options...))
	return nil
}
