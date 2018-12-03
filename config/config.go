package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	DefaultBackendKey = "DefaultBackend"
	LocalBackendVal   = "local"
)

func InitConfig(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in ".config/hourglass".
		viper.AddConfigPath(home + "/etc/hourglass")
		viper.AddConfigPath(home + "/.config/hourglass")
		viper.SetConfigName("config")
		setConfigDefaults()
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

func setConfigDefaults() {
	viper.SetDefault(DefaultBackendKey, LocalBackendVal)
}
