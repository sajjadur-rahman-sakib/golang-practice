package config

import (
	"log"

	"github.com/spf13/viper"
)

var LocalConfig *Config

type Config struct {
	DBUser string `mapstructure:"DBUSER"`
	DBPass string `mapstructure:"DBPASS"`
	DBIP   string `mapstructure:"DBIP"`
	DbName string `mapstructure:"DBNAME"`
	Port   string `mapstructure:"PORT"`
}

func InitConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading env file: %v", err)
	}

	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error reading env file: %v", err)
	}

	return config
}

func SetConfig() {
	LocalConfig = InitConfig()
}
