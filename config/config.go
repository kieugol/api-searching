package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	if config == nil {
		var err error
		v := viper.New()
		v.SetConfigType("yaml")
		v.SetConfigName(env)
		v.AddConfigPath("../config/")
		v.AddConfigPath("config/")
		v.AutomaticEnv()
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		err = v.ReadInConfig()
		if err != nil {
			log.Fatal("error on parsing configuration file")
		}
		config = v
	}
}

func GetConfig() *viper.Viper {
	return config
}
