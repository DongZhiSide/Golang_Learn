package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type RootCmdFlags struct {
	Ip string
}

type RootCmd struct {
	Cmd   *cobra.Command
	Flags RootCmdFlags
}

func NewRootCmd() *RootCmd {
	rootCmd := &RootCmd{}

	rootCmd.Cmd = &cobra.Command{
		Use:   "appName",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files`,
		Run: func(cmd *cobra.Command, args []string) {
			// cmd 就是 rootCmd.cmd, args 是位置参数
			// Do Stuff Here
			fmt.Println("ip is ", rootCmd.Flags.Ip)
		},
	}

	rootCmd.Cmd.PersistentFlags().StringVarP(&rootCmd.Flags.Ip, "ip", "i", "127.0.0.1", "ip address")

	return rootCmd
}
