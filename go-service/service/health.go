package service

/*
 * 健康检查服务实现
 *
 * wencan
 * 2019-06-23
 */

import (
	"context"

	protocol "github.com/wencan/kit-demo/protocol/model"
)

// HealthService 健康检查服务
type HealthService struct {
}

// NewHealthService 创建健康检查服务
func NewHealthService() *HealthService {
	return &HealthService{}
}

// Check 检查指定服务的健康状况
func (healthService *HealthService) Check(ctx context.Context, serviceName string) (protocol.HealthServiceStatus, error) {
	switch serviceName {
	case "kit-demo", "":
		return protocol.HealthServiceStatusServing, nil
	default:
		return protocol.HealthServiceStatusUnknown, nil
	}
}
