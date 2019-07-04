package http

/*
 * 计算器http服务接口
 * wencan
 * 2019-06-25
 */

import (
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"

	"github.com/wencan/kit-demo/go-service/cmd/endpoint"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// RegisterCalculatorHTTPHandlers 向http router注册计算器方法处理器
func RegisterCalculatorHTTPHandlers(router *httprouter.Router, service endpoint.CalculatorService, options ...transport.ServerOption) error {
	decodeCalculatorAddRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorAddRequest) })
	router.Handler(http.MethodPost, "/calculator/add", transport.NewServer(endpoint.NewCalculatorAddEndpoint(service), decodeCalculatorAddRequest, encodeResponse, options...))

	decodeCalculatorSubRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorSubRequest) })
	router.Handler(http.MethodPost, "/calculator/sub", transport.NewServer(endpoint.NewCalculatorSubEndpoint(service), decodeCalculatorSubRequest, encodeResponse, options...))

	decodeCalculatorMulRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorMulRequest) })
	router.Handler(http.MethodPost, "/calculator/mul", transport.NewServer(endpoint.NewCalculatorMulEndpoint(service), decodeCalculatorMulRequest, encodeResponse, options...))

	decodeCalculatorDivRequest := makeRequestDecoder(func() interface{} { return new(protocol.CalculatorDivRequest) })
	router.Handler(http.MethodPost, "/calculator/div", transport.NewServer(endpoint.NewCalculatorDivEndpoint(service), decodeCalculatorDivRequest, encodeResponse, options...))
	return nil
}
