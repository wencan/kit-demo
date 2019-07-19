package cmd

import (
	"context"

	errmsg_zap "github.com/wencan/errmsg/logging/zap"
	"go.uber.org/zap"
)

// ErrorLogHandler 错误日志处理器
// 实现kit/transport.ErrorHandler接口
// 用来替换默认的错误日志处理器
type ErrorLogHandler struct {
	logger *zap.Logger
}

// NewErrorLogHandler 创建错误日志处理器
func NewErrorLogHandler(logger *zap.Logger) *ErrorLogHandler {
	return &ErrorLogHandler{
		logger: logger,
	}
}

// Handle 处理错误，写日志
func (handler *ErrorLogHandler) Handle(ctx context.Context, err error) {
	fields := errmsg_zap.Fields(err)
	handler.logger.Error("an error occurred", fields...)
}
