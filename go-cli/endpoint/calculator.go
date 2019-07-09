package endpoint

/*
 * 计算器endpoint
 *
 * wencan
 * 2019-07-08
 */

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	protocol "github.com/wencan/kit-demo/protocol/model"
)

// CalculatorEndpoints 计算器客户端endpoint集
// endpoint需要transport赋值
type CalculatorEndpoints struct {
	AddEndpoint endpoint.Endpoint

	SubEndpoint endpoint.Endpoint

	MulEndpoint endpoint.Endpoint

	DivEndpoint endpoint.Endpoint
}

// Add 加法
func (endpoints *CalculatorEndpoints) Add(ctx context.Context, a, b int32) (int32, error) {
	req := &protocol.CalculatorAddRequest{
		A: a,
		B: b,
	}
	resp, err := endpoints.AddEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	response := resp.(*protocol.CalculatorInt32Response)
	return response.Result, nil
}

// Sub 减法
func (endpoints *CalculatorEndpoints) Sub(ctx context.Context, c, d int32) (int32, error) {
	req := &protocol.CalculatorSubRequest{
		C: c,
		D: d,
	}
	resp, err := endpoints.SubEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	response := resp.(*protocol.CalculatorInt32Response)
	return response.Result, nil
}

// Mul 乘法
func (endpoints *CalculatorEndpoints) Mul(ctx context.Context, e, f int32) (int32, error) {
	req := &protocol.CalculatorMulRequest{
		E: e,
		F: f,
	}
	resp, err := endpoints.MulEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	response := resp.(*protocol.CalculatorInt32Response)
	return response.Result, nil
}

// Div 除法
func (endpoints *CalculatorEndpoints) Div(ctx context.Context, m, n int32) (float32, error) {
	req := &protocol.CalculatorDivRequest{
		M: m,
		N: n,
	}
	resp, err := endpoints.DivEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	response := resp.(*protocol.CalculatorFloatResponse)
	return response.Result, nil
}
