package transport

import (
	"context"

	"github.com/wencan/kit-demo/protocol/model"
)

// HealthClient 健康检查服务接口
type HealthClient interface {
	// Check 检查指定服务的健康状态
	Check(ctx context.Context, serviceName string) (model.HealthServiceStatus, error)
}
