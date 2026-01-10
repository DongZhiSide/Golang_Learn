package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type _RootCmd struct {
	cmd *cobra.Command
}

var (
	roocmd _RootCmd
)

func (rootCmd *_RootCmd) configCmd() {
	rootCmd.cmd = &cobra.Command{
		Use:   `try`,
		Short: `just try`,
		Long:  `just try cobra.`,
		Run:   getRunControl("192.168.1.1"),
	}

	ip := rootCmd.cmd.Flags().String("ip", "127.0.0.1", "input ip")
	fmt.Println("--ip:", *ip)
}

func init() {
	roocmd.configCmd()
}

func Execute() {
	roocmd.cmd.Execute()
}

func getRunControl(ip string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("ip:", ip)
	}
}
