package cmd

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
)

/*
 * endpoint中间件
 *
 * wencan
 * 2019-07-19
 */

// ServerMiddlewares 服务端endpoint中间件
func ServerMiddlewares() endpoint.Middleware {
	return endpoint.Chain(
		ratelimit.NewDelayingLimiter(rate.NewLimiter(rate.Every(time.Millisecond*100), 10)), // 限流，0.1秒一个token，最多10个token
	)
}
