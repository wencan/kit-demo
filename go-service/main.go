package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	grpcsvc "github.com/wencan/kit-demo/go-service/cmd/grpc"
	service "github.com/wencan/kit-demo/go-service/service"
	proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	healthService := service.NewHealthService()

	ln, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	proto.RegisterHealthServer(s, grpcsvc.NewHealthGRPCServer(healthService))
	err = s.Serve(ln)
	if err != nil {
		log.Fatalln(err)
	}
}
