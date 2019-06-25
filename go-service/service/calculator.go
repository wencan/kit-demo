package service

/*
 * 计算器服务实现
 *
 * wencan
 * 2019-06-25
 */

import "context"

// CalculatorService 计算器服务
type CalculatorService struct {
}

// NewCalculatorService 构建计算器服务对象
func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

// Add 加法
func (service *CalculatorService) Add(ctx context.Context, a, b int32) (int32, error) {
	return a + b, nil
}

// Sub 减法
func (service *CalculatorService) Sub(ctx context.Context, c, d int32) (int32, error) {
	return c - d, nil
}

// Mul 乘法
func (service *CalculatorService) Mul(ctx context.Context, e, f int32) (int32, error) {
	return e * f, nil
}

// Div 除法
func (service *CalculatorService) Div(ctx context.Context, m, n int32) (float32, error) {
	return float32(m) / float32(n), nil
}
