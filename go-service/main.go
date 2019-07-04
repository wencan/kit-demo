package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/wencan/errmsg"
	"github.com/wencan/kit-demo/go-service/cmd"
	"github.com/wencan/kit-demo/go-service/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	healthService := service.NewHealthService()
	claculatorService := service.NewCalculatorService()

	ctx, cancel := context.WithCancel(context.Background())

	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := logConfig.Build()
	if err != nil {
		log.Println(err)
		return
	}

	// errmsg捕获文件位置信息
	// 然后由错误日志处理器统一输出
	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Flongfile)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		err := cmd.RunGRPCServer(ctx, "tcp", "127.0.0.1:5051", healthService, claculatorService, logger)
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		time.Sleep(5)
		defer wg.Done()
		defer cancel()
		err := cmd.RunHTTPServer(ctx, "127.0.0.1:6061", healthService, claculatorService, logger)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
}
