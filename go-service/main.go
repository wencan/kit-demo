package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/wencan/kit-demo/go-service/cmd"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/wencan/errmsg"
	errmsg_zap "github.com/wencan/errmsg/logging/zap"
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
	logger, err := logConfig.Build(zap.AddStacktrace(zapcore.DPanicLevel))
	if err != nil {
		log.Println(err)
		return
	}

	// errmsg捕获文件位置信息
	// 然后由错误日志处理器统一输出
	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Flongfile)

	grpcHandler, err := cmd.NewGRPCHandler(ctx, healthService, claculatorService, logger.With(zap.String("protocol", "gRPC")))
	if err != nil {
		logger.Error("NewGRPCHandler fail", errmsg_zap.Fields(err)...)
		return
	}
	httpHandler, err := cmd.NewHTTPHandler(ctx, healthService, claculatorService, logger.With(zap.String("protocol", "HTTP")))
	if err != nil {
		logger.Error("NewHTTPHandler fail", errmsg_zap.Fields(err)...)
		return
	}

	// 根据Content-Type判断处理方法
	dispatcher := func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		contentType = strings.Split(contentType, ";")[0]
		switch contentType {
		case "application/grpc":
			grpcHandler.ServeHTTP(w, r)
		default:
			httpHandler.ServeHTTP(w, r)
		}
	}

	// 明文http2服务
	h2s := &http2.Server{}
	handler := h2c.NewHandler(http.HandlerFunc(dispatcher), h2s)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Error("listen failed", zap.Error(err))
	}
	logger.Info("listen on " + ln.Addr().String())
	server := &http.Server{
		Handler: handler,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		// 等待int信号
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, os.Interrupt)

		select {
		case <-sign:
			logger.Info("server shutdown")
			server.Shutdown(ctx)
		case <-ctx.Done():
		}
	}()

	err = server.Serve(ln)
	if err != nil && err != http.ErrServerClosed {
		cancel()
		logger.Error("server failed", zap.Error(err))
	}
	wg.Wait()
}
