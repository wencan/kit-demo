package main

/*
 * cli配置
 *
 * wencan
 * 2019-07-07
 */

import (
	"context"
	"errors"
	"log"

	"github.com/wencan/errmsg"

	grpc_transport "github.com/wencan/kit-demo/go-cli/transport/grpc"
	http_transport "github.com/wencan/kit-demo/go-cli/transport/http"

	"github.com/spf13/cobra"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

// HealthClient 健康检查服务接口
type HealthClient interface {
	// Check 检查指定服务的健康状态
	Check(ctx context.Context, serviceName string) (protocol.HealthServiceStatus, error)
}

func healthGRPCClientFactory(ctx context.Context, target string) (HealthClient, error) {
	return grpc_transport.NewHealthGRPCClient(ctx, target)
}

func healthHTTPClientFactory(ctx context.Context, target string) (HealthClient, error) {
	return http_transport.NewHealthHTTPClient(ctx, target)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // 最简单的日志

	var healthClientFactory func(ctx context.Context, target string) (HealthClient, error)

	rootCmd := &cobra.Command{
		Use: "kit-demo-cli",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// protocol参数检查

			protocolName, err := cmd.InheritedFlags().GetString("protocol")
			if err != nil {
				// cobra 会输出错误信息
				// log.Println(err)
				return err
			}

			switch protocolName {
			case "grpc":
				healthClientFactory = healthGRPCClientFactory
			case "http":
				healthClientFactory = healthHTTPClientFactory
			default:
				return errors.New("protocol invalid")
			}
			return nil
		},
	}

	// 全局flag
	rootCmd.PersistentFlags().String("protocol", "grpc", "communication protocol, grpc or http")

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
					client, err := healthClientFactory(context.Background(), "127.0.0.1:8080")
					if err != nil {
						return
					}
					status, err := client.Check(context.Background(), service)
					if err != nil {
						errMsg, ok := err.(*errmsg.ErrMsg)
						if ok {
							log.Println(errMsg.String())
						} else {
							log.Println(err)
						}
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
