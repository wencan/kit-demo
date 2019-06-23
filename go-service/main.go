package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"

	grpcsvc "github.com/wencan/kit-demo/go-service/cmd/grpc"
	httpsvc "github.com/wencan/kit-demo/go-service/cmd/http"
	service "github.com/wencan/kit-demo/go-service/service"
	proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
)

func runGRPCServer(ctx context.Context, network, addr string, healthService *service.HealthService) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	ln, err := net.Listen(network, addr)
	if err != nil {
		log.Println(err)
		return nil
	}

	s := grpc.NewServer()
	proto.RegisterHealthServer(s, grpcsvc.NewHealthGRPCServer(healthService))

	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()

	err = s.Serve(ln)
	if err != nil && err != grpc.ErrServerStopped {
		log.Println(err)
		return err
	}
	return nil
}

func runHTTPServer(ctx context.Context, addr string, healthService *service.HealthService) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	router := httprouter.New()
	httpsvc.RegisterHealthHTTPHandlers(router, healthService)

	s := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println(err)
		return err
	}
	return nil
}

func main() {
	healthService := service.NewHealthService()

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		err := runGRPCServer(ctx, "tcp", "127.0.0.1:5051", healthService)
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		time.Sleep(5)
		defer wg.Done()
		defer cancel()
		err := runHTTPServer(ctx, "127.0.0.1:6061", healthService)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
}
