package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/hourglass/config)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in ".config/hourglass".
		viper.AddConfigPath(home + "/.config/hourglass")
		viper.SetConfigName("config")
		fp := home + "/.config/hourglass/config.yaml"
		err = viper.ReadInConfig()
		if err != nil {
			os.Create(fp)
			viper.SafeWriteConfig()
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Invalid config:", err)
		os.Exit(1)
	}
}

func Execute() {
	rootCmd.Execute()
}
