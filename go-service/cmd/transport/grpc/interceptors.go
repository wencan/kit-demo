package grpc

/*
 * 自定义拦截器
 *
 * wencan
 * 2019-07-01
 */

import (
	"context"

	errmsg_grpc "github.com/wencan/errmsg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryErrorInterceptor 错误拦截器。将error转为grpc错误
func UnaryErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		// 优先识别grpc/status错误
		// 然后识别errmsg错误
		// 其它错误将被归类为unknown
		code := status.Code(err)
		if code == codes.Unknown {
			// 返回grpc/status.statusError对象
			// statusError实现了GRPCStatus() *Status方法
			err = errmsg_grpc.Status(err).Err()
		}
	}

	return resp, err
}
