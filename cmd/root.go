package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type root_cmd_builder struct{}

func (rcb *root_cmd_builder) build() *cobra.Command {
	root_cmd := rcb.init_cmd()
	rcb.add_flags(root_cmd)
	rcb.add_func(root_cmd)
	return root_cmd
}

func (rcb *root_cmd_builder) init_cmd() *cobra.Command {
	root_cmd := &cobra.Command{
		Use:   "go run main.go",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files`,
	}

	return root_cmd
}

type RootCmdFlags struct {
	Ip string
}

var root_cmd_flags RootCmdFlags

func (rcb *root_cmd_builder) add_flags(root_cmd *cobra.Command) {
	root_cmd.PersistentFlags().StringVarP(&root_cmd_flags.Ip, "ip", "i", "127.0.0.1", "ip address")
}

func (root_cmd_builder) add_func(root_cmd *cobra.Command) {
	root_cmd.RunE = func(cmd *cobra.Command, args []string) error {
		// cmd 就是 rootCmd.Cmd, args 是位置参数
		// Do Stuff Here
		fmt.Println("ip is ", root_cmd_flags.Ip)
		return nil
	}
}

func Execute() {
	root_cmd := (&root_cmd_builder{}).build()
	err := root_cmd.Execute()
	if err != nil {
		fmt.Println(err)
	} else {
		// Do Stuff Here
	}
}
