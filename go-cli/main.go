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
	"os"
	"time"

	kit_log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/etcdv3"
	consul_api "github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/wencan/kit-plugins/sd/mdns"

	"github.com/wencan/kit-demo/go-cli/transport"
	grpc_transport "github.com/wencan/kit-demo/go-cli/transport/grpc"
	http_transport "github.com/wencan/kit-demo/go-cli/transport/http"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

var (
	serviceDirectory = "/services/kit-demo"
)

// HealthClient 健康检查客户端接口
type HealthClient interface {
	// Check 检查指定服务的健康状态
	Check(ctx context.Context, serviceName string) (protocol.HealthServiceStatus, error)
}

// CalculatorClient 计算器客户端接口
type CalculatorClient interface {
	// Add 加
	Add(ctx context.Context, a, b int32) (int32, error)

	// Sub 减
	Sub(ctx context.Context, c, d int32) (int32, error)

	// Mul 乘
	Mul(ctx context.Context, e, f int32) (int32, error)

	// Div 除
	Div(ctx context.Context, m, n int32) (float32, error)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // 最简单的日志

	var instancer sd.Instancer

	var healthClient HealthClient
	var calculatorClient CalculatorClient
	var healthTransportFactory transport.HealthTransportFactory
	var calculatorTransportFactory transport.CalculatorTransportFactory

	// rootCmd.PersistentPreRunE会被二级cmd.PersistentPreRunE覆盖
	// 所以需要二级cmd.PersistentPreRunE主动调研rootCmd.PersistentPreRunE

	rootCmd := &cobra.Command{
		Use: "kit-demo-cli",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// protocol参数
			protocolName, err := cmd.Flags().GetString("protocol")
			if err != nil {
				// cobra 会输出错误信息
				// log.Println(err)
				return err
			}
			// 根据协议
			// 指定传输客户端factory
			switch protocolName {
			case "grpc":
				healthTransportFactory = func(ctx context.Context, target string) (transport.HealthTransport, error) {
					return grpc_transport.NewHealthGRPCClient(ctx, target)
				}
				calculatorTransportFactory = func(ctx context.Context, target string) (transport.CalculatorTransport, error) {
					return grpc_transport.NewCalculatorGRPCClient(ctx, target)
				}
			case "http":
				healthTransportFactory = func(ctx context.Context, target string) (transport.HealthTransport, error) {
					return http_transport.NewHealthHTTPClient(ctx, target)
				}
				calculatorTransportFactory = func(ctx context.Context, target string) (transport.CalculatorTransport, error) {
					return http_transport.NewCalculatorHTTPClient(ctx, target)
				}
			default:
				return errors.New("protocol invalid")
			}

			// 服务发现
			instancer, err = newInstancer(cmd)
			if err != nil {
				return err
			}

			return nil
		},
	}
	// 全局flag
	rootCmd.PersistentFlags().String("protocol", "grpc", "communication protocol, grpc or http")
	rootCmd.PersistentFlags().StringSlice("etcd", []string{}, "etcd servers address")
	rootCmd.PersistentFlags().String("consul", "", "consul server address")

	// 二级cmd 具体服务
	{
		// health 健康检查
		healthCmd := &cobra.Command{
			Use:   "health",
			Short: "check service health",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				// 调用被覆盖的rootCmd.PersistentPreRunE
				// 创建cli
				err := cmd.Parent().PersistentPreRunE(cmd.Parent(), []string{})
				if err != nil {
					return err
				}

				// 健康检查客户端
				healthClient = transport.NewHealthClient(healthTransportFactory, instancer, kit_log.NewLogfmtLogger(os.Stdout))
				return nil
			},
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
					status, err := healthClient.Check(context.Background(), service)
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

		// 计算器
		calculatorCmd := &cobra.Command{
			Use:   "calculator",
			Short: "simple calculator",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				// 调用被覆盖的rootCmd.PersistentPreRunE
				// 创建cli
				err := cmd.Parent().PersistentPreRunE(cmd.Parent(), []string{})
				if err != nil {
					return err
				}

				// 计算器客户端
				calculatorClient = transport.NewCalculatorClient(calculatorTransportFactory, instancer, kit_log.NewLogfmtLogger(os.Stdout))
				return nil
			},
		}
		{
			addCmd := &cobra.Command{
				Use:     "add",
				Example: "kit-demo-cli calcuator add --a 123 --b 456",
				Run: func(cmd *cobra.Command, args []string) {
					a, b, err := parseTwoInt32Flags(cmd.LocalFlags(), "a", "b")
					if err != nil {
						return
					}
					result, err := calculatorClient.Add(context.Background(), a, b)
					if err != nil {
						log.Println(err)
						return
					}
					log.Println("result:", result)
				},
			}
			addCmd.Flags().Int32("a", 0, "")
			addCmd.Flags().Int32("b", 0, "")
			addCmd.MarkFlagRequired("a")
			addCmd.MarkFlagRequired("b")

			subCmd := &cobra.Command{
				Use:     "sub",
				Example: "kit-demo-cli calcuator sub --c 678 --d 345",
				Run: func(cmd *cobra.Command, args []string) {
					c, d, err := parseTwoInt32Flags(cmd.LocalFlags(), "c", "d")
					if err != nil {
						return
					}
					result, err := calculatorClient.Sub(context.Background(), c, d)
					if err != nil {
						log.Println(err)
						return
					}
					log.Println("result:", result)
				},
			}
			subCmd.Flags().Int32("c", 0, "")
			subCmd.Flags().Int32("d", 0, "")
			subCmd.MarkFlagRequired("c")
			subCmd.MarkFlagRequired("d")

			mulCmd := &cobra.Command{
				Use:     "mul",
				Example: "kit-demo-cli calcuator mul --e 123 --f 456",
				Run: func(cmd *cobra.Command, args []string) {
					e, f, err := parseTwoInt32Flags(cmd.LocalFlags(), "e", "f")
					if err != nil {
						return
					}
					result, err := calculatorClient.Mul(context.Background(), e, f)
					if err != nil {
						log.Println(err)
						return
					}
					log.Println("result:", result)
				},
			}
			mulCmd.Flags().Int32("e", 0, "")
			mulCmd.Flags().Int32("f", 0, "")
			mulCmd.MarkFlagRequired("e")
			mulCmd.MarkFlagRequired("f")

			divCmd := &cobra.Command{
				Use:     "div",
				Example: "kit-demo-cli calcuator div --m 123 --n 456",
				Run: func(cmd *cobra.Command, args []string) {
					m, n, err := parseTwoInt32Flags(cmd.LocalFlags(), "m", "n")
					if err != nil {
						return
					}
					result, err := calculatorClient.Div(context.Background(), m, n)
					if err != nil {
						log.Println(err)
						return
					}
					log.Println("result:", result)
				},
			}
			divCmd.Flags().Int32("m", 0, "")
			divCmd.Flags().Int32("n", 0, "")
			divCmd.MarkFlagRequired("m")
			divCmd.MarkFlagRequired("n")

			calculatorCmd.AddCommand(addCmd, subCmd, mulCmd, divCmd)
		}
		rootCmd.AddCommand(calculatorCmd)
	}

	rootCmd.Execute()
}

func newInstancer(cmd *cobra.Command) (sd.Instancer, error) {
	// etcd参数
	etcdServers, err := cmd.Flags().GetStringSlice("etcd")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// consul参数
	consulServer, err := cmd.Flags().GetString("consul")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(etcdServers) > 0 {
		// 如果提供etcd服务器地址参数
		// 使用etcd发现服务
		// etcd客户端
		connCtx, _ := context.WithTimeout(context.Background(), time.Second*10)
		etcdClient, err := etcdv3.NewClient(connCtx, etcdServers, etcdv3.ClientOptions{})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		instancer, err := etcdv3.NewInstancer(etcdClient, serviceDirectory, kit_log.NewLogfmtLogger(os.Stdout))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return instancer, nil
	} else if consulServer != "" {
		config := consul_api.DefaultConfig()
		config.Address = consulServer
		c, err := consul_api.NewClient(config)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		client := consul.NewClient(c)
		instancer := consul.NewInstancer(client, kit_log.NewLogfmtLogger(os.Stdout), "kit-demo", nil, false)
		return instancer, nil
	} else {
		// 如果没提供etcd服务器地址参数
		// 使用mDNS发现服务
		instancer, err := mdns.NewInstancer(serviceDirectory, mdns.InstancerOptions{}, kit_log.NewLogfmtLogger(os.Stdout))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return instancer, nil
	}
}

func parseTwoInt32Flags(flags *pflag.FlagSet, name1, name2 string) (int32, int32, error) {
	int1, err := flags.GetInt32(name1)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	int2, err := flags.GetInt32(name2)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	return int1, int2, nil
}
