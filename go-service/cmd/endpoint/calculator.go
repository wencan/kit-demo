package endpoint

/*
 * 计算器endpoint
 * 服务实现与协议接口的中间层
 *
 * wencan
 * 2019-06-25
 */

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	protocol "github.com/wencan/kit-demo/protocol/model"
)

// CalculatorService 计算器服务接口
type CalculatorService interface {
	Add(ctx context.Context, a, b int32) (int32, error)

	Sub(ctx context.Context, c, d int32) (int32, error)

	Mul(ctx context.Context, e, f int32) (int32, error)

	Div(ctx context.Context, m, n int32) (float32, error)
}

// NewCalculatorAddEndpoint 创建计算器加法方法endpoint
func NewCalculatorAddEndpoint(service CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*protocol.CalculatorAddRequest)
		result, err := service.Add(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return &protocol.CalculatorInt32Response{
			Result: result,
		}, nil
	}
}

// NewCalculatorSubEndpoint 创建计算器减法方法endpoint
func NewCalculatorSubEndpoint(service CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*protocol.CalculatorSubRequest)
		result, err := service.Sub(ctx, req.C, req.D)
		if err != nil {
			return nil, err
		}
		return &protocol.CalculatorInt32Response{
			Result: result,
		}, nil
	}
}

// NewCalculatorMulEndpoint 创建计算器乘法方法endpoint
func NewCalculatorMulEndpoint(service CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*protocol.CalculatorMulRequest)
		result, err := service.Mul(ctx, req.E, req.F)
		if err != nil {
			return nil, err
		}
		return &protocol.CalculatorInt32Response{
			Result: result,
		}, nil
	}
}

// NewCalculatorDivEndpoint 创建计算器除法方法endpoint
func NewCalculatorDivEndpoint(service CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*protocol.CalculatorDivRequest)
		result, err := service.Div(ctx, req.M, req.N)
		if err != nil {
			return nil, err
		}
		return &protocol.CalculatorFloatResponse{
			Result: result,
		}, nil
	}
}
