package main

/*
 * cli配置
 *
 * wencan
 * 2019-07-07
 */

import (
	"context"
	"log"

	"github.com/wencan/kit-demo/go-cli/transport/grpc"

	"github.com/spf13/cobra"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// HealthClient 健康检查服务接口
type HealthClient interface {
	// Check 检查指定服务的健康状态
	Check(ctx context.Context, serviceName string) (protocol.HealthServiceStatus, error)
}

func main() {
	log.SetFlags(0) // 最简单的日志

	rootCmd := &cobra.Command{
		Use: "kit-demo-cli",
	}

	// 全局flag
	rootCmd.PersistentFlags().String("protobuf", "grpc", "communication protocol, grpc or http")

	// 二级cmd 具体服务
	{
		// health 健康检查
		healthCmd := &cobra.Command{
			Use:   "health",
			Short: "check service health",
		}
		{
			// health check 健康检查check方法
			checkCmd := &cobra.Command{
				Use:     "check",
				Short:   "Check the health of the specified service",
				Example: "kit-demo-cli health check --service kit-demo",
				Run: func(cmd *cobra.Command, args []string) {
					service, err := cmd.LocalFlags().GetString("service")
					if err != nil {
						log.Println(err)
						return
					}
					client, err := grpc.NewHealthGRPCClient(context.Background(), "127.0.0.1:8080")
					if err != nil {
						return
					}
					status, err := client.Check(context.Background(), service)
					if err != nil {
						log.Println(err)
						return
					}
					log.Println("status:", protocol.HealthServiceStatusName(status))
				},
			}
			checkCmd.Flags().String("service", "", "service name")
			checkCmd.MarkFlagRequired("service")
			healthCmd.AddCommand(checkCmd)
		}
		rootCmd.AddCommand(healthCmd)
	}

	rootCmd.Execute()
}
