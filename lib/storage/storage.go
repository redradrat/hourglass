package storage

import (
	"fmt"

	"github.com/redradrat/hourglass/config"
	"github.com/redradrat/hourglass/lib"
	"github.com/spf13/viper"
)

type Storage interface {
	GetLog() (*lib.Log, error)
	WriteToLog(l lib.LogEntry) error
	DeleteFromLog(id string) error
}

func GetBackend() (Storage, error) {
	var be Storage
	configuredStorageBackend := viper.Get(config.DefaultBackendKey)
	switch configuredStorageBackend {
	case config.LocalBackendVal:
		be = LocalStorage{}
	default:
		return nil, fmt.Errorf("unknown storage backend [%s] specified in config", configuredStorageBackend)
	}
	return be, nil
}
