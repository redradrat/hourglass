package cmd

import (
	"github.com/redradrat/hourglass/config"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hourglass",
		Short: "Easy and fast time tracker",
		Long:  `A fast, easy-to-use and lightweight time tracking tool, written in Go.`,
	}

	cfgFile string
)

func init() {
	cobra.OnInitialize(func() { config.InitConfig(cfgFile) })
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/hourglass/config)")
}

func Execute() {
	rootCmd.Execute()
}
