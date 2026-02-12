package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	GeneralConfig `mapstructure:"general"`
	LogConfig     `mapstructure:"log"`
	HttpConfig    `mapstructure:"http"`
}
type GeneralConfig struct {
	Ipv4           string `mapstructure:"ipv4" `
	Port           int    `mapstructure:"port"`
	EnableTLS      bool   `mapstructure:"enable-tls"`
	CrtPath        string `mapstructure:"crt-path"`
	PrivateKeyPath string `mapstructure:"private-key-path"`
}
type HttpConfig struct {
	Root     string `mapstructure:"root"`
	Hostname string `mapstructure:"hostname"`

	CookieSecure   bool   `mapstructure:"cookie-secure"`
	CookieHttponly bool   `mapstructure:"cookie-httpOnly"`
	CookieDomain   string `mapstructure:"cookie-domain"`
}
type LogConfig struct {
	Level string `mapstructure:"level"`
}

func main() {
	cmd := cobra.Command{
		Use:  "viper",
		Long: `viper with cobra`,
		Run: func(cmd *cobra.Command, args []string) {
			c := Config{}
			err := viper.ReadInConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = viper.Unmarshal(&c)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(c)
			fmt.Println(viper.GetInt("general.port"))
		},
	}
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	cmd.Flags().StringP("ipv4", "i", "192.168.1.1", "server's ipv4 address")
	viper.BindPFlag("general.ipv4", cmd.LocalFlags().Lookup("ipv4"))
	cmd.Execute()
	// go run main.go -i 123.123.123.123
	// Output:
	// {{123.123.123.123 12345 true ./a.crt ./a.key} {trace} {/ tiny.mydns.com true true mydns.com}}
	// 12345
}
