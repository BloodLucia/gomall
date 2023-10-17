package config

import (
	"github.com/spf13/viper"
)

func New() *viper.Viper {
	v := viper.New()

	v.SetConfigType("yaml")
	v.SetConfigName("config.yaml")

	return v
}
