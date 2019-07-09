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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	protocol "github.com/wencan/kit-demo/protocol/model"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // 最简单的日志

	var cli *Cli

	// rootCmd.PersistentPreRunE会被二级cmd.PersistentPreRunE覆盖
	// 所以需要二级cmd.PersistentPreRunE主动调研rootCmd.PersistentPreRunE

	rootCmd := &cobra.Command{
		Use: "kit-demo-cli",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// protocol参数检查
			protocolName, err := cmd.Flags().GetString("protocol")
			if err != nil {
				// cobra 会输出错误信息
				// log.Println(err)
				return err
			}

			switch protocolName {
			case "grpc":
				cli, err = NewCliOnGRPC()
				if err != nil {
					return err
				}
			case "http":
				cli, err = NewCliOnHTTP()
				if err != nil {
					return err
				}
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
		var healthCli *HealthCli
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

				healthCli, err = cli.NewHealthCli(context.Background(), "127.0.0.1:8080")
				if err != nil {
					return err
				}
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
					status, err := healthCli.Check(context.Background(), service)
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
		var calculatorCli *CalculatorCli
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

				calculatorCli, err = cli.NewCalculatorCli(context.Background(), "127.0.0.1:8080")
				if err != nil {
					return err
				}
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
					result, err := calculatorCli.Add(context.Background(), a, b)
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
					result, err := calculatorCli.Sub(context.Background(), c, d)
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
					result, err := calculatorCli.Mul(context.Background(), e, f)
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
					result, err := calculatorCli.Div(context.Background(), m, n)
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
