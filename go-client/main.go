package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	proto "github.com/wencan/kit-demo/protocol/google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := proto.NewHealthClient(conn)
	resp, err := client.Check(context.Background(), &proto.HealthCheckRequest{
		Service: "",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Status)
}
