package service

import (
	"context"

	protocol "github.com/wencan/kit-demo/protocol/model"
)

// HealthService 健康检查服务
type HealthService struct {
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (healthService *HealthService) Check(ctx context.Context, serviceName string) (protocol.HealthServiceStatus, error) {
	switch serviceName {
	case "kit-demo", "":
		return protocol.HealthServiceStatusServing, nil
	default:
		return protocol.HealthServiceStatusUnknown, nil
	}
}
