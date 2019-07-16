package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/valyala/fasthttp"

	kit_zap "github.com/go-kit/kit/log/zap"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/spf13/pflag"
	"github.com/wencan/errmsg"
	errmsg_zap "github.com/wencan/errmsg/logging/zap"
	"github.com/wencan/kit-plugins/sd/mdns"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wencan/kit-demo/go-service/cmd"
	"github.com/wencan/kit-demo/go-service/service"
)

var (
	etcdServers      = []string{}
	serviceDirectory = "/services/kit-demo"
)

func init() {
	pflag.StringSliceVar(&etcdServers, "etcd", []string{}, "etcd servers address")
	pflag.Parse()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// 日志
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

	// 服务逻辑
	healthService := service.NewHealthService()
	claculatorService := service.NewCalculatorService()

	// 服务接口处理
	// grpcHandler, err := cmd.NewHandlerOnGRPC(ctx, healthService, claculatorService, logger.With(zap.String("protocol", "gRPC")))
	// if err != nil {
	// 	logger.Error("NewHandlerOnGRPC fail", errmsg_zap.Fields(err)...)
	// 	return
	// }
	httpHandler, err := cmd.NewHandlerOnHTTP(ctx, healthService, claculatorService, logger.With(zap.String("protocol", "HTTP")))
	if err != nil {
		logger.Error("NewHandlerOnHTTP fail", errmsg_zap.Fields(err)...)
		return
	}

	// // 根据Content-Type判断处理方法
	// dispatcher := func(ctx *fasthttp.RequestCtx) {
	// 	contentType := string(ctx.Request.Header.ContentType())
	// 	switch contentType {
	// 	case "application/grpc":
	// 		grpcHandler(ctx)
	// 	default:
	// 		httpHandler(ctx)
	// 	}
	// }

	// // 明文http2服务
	// h2s := &http2.Server{}
	// handler := h2c.NewHandler(http.HandlerFunc(dispatcher), h2s)

	// 监听
	ln, err := net.Listen("tcp", ":")
	if err != nil {
		logger.Error("listen failed", zap.Error(err))
		return
	}
	logger.Info("listen on " + ln.Addr().String())

	// 服务地址
	ip := getOutboundIP().String()
	_, sport, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		log.Println(err)
		return
	}
	port, err := strconv.Atoi(sport)
	if err != nil {
		log.Println(err)
		return
	}
	instance := fmt.Sprintf("%s:%d", ip, port)

	// 服务注册
	var registrar sd.Registrar
	if len(etcdServers) > 0 {
		// etcd
		etcdClient, err := etcdv3.NewClient(ctx, etcdServers, etcdv3.ClientOptions{})
		if err != nil {
			log.Println(err)
			ln.Close()
			return
		}
		registrar = etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
			Key:   serviceDirectory + "/" + instance,
			Value: instance,
		}, kit_zap.NewZapSugarLogger(logger.With(zap.String("sd", "etcd")), zap.InfoLevel))
	} else {
		// mDNS
		service := mdns.Service{
			Instance: instance,
			Service:  serviceDirectory,
			Port:     port,
		}
		registrar, err = mdns.NewRegistrar(service, kit_zap.NewZapSugarLogger(logger.With(zap.String("sd", "mDNS")), zap.InfoLevel))
		if err != nil {
			log.Println(err)
			ln.Close()
			return
		}
	}
	registrar.Register()
	defer registrar.Deregister()

	// 服务
	server := &fasthttp.Server{
		Name:    "kit-demo", // UserAgent
		Handler: httpHandler,
	}

	// 优雅退出
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		// 等待退出信号
		// 用户取消为int信号
		// docker stop为term信号
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, os.Interrupt, syscall.SIGTERM)

		select {
		case <-sign:
			logger.Info("server shutdown")
			server.Shutdown()
		case <-ctx.Done():
		}
	}()

	// 开始服务
	err = server.Serve(ln)
	if err != nil {
		cancel()
		logger.Error("server failed", zap.Error(err))
	}

	// 等待子goroutine退出
	wg.Wait()
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
