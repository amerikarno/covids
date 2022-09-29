package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var App = struct {
	URL string
}{}

func Init() {
	LoadViper()
	LoadAppConfig()
}

func LoadViper() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error when viper read config, error: %v", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func LoadAppConfig() {
	App.URL = viper.GetString("APP.URL")
}
