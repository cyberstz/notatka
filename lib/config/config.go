package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

func Load() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Viper couldn't read in the config file. %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Viper couldn't unmarshal the configuration. %v", err)
	}

	return
}
