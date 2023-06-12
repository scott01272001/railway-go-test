package initializers

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func LoadEnvVariables() {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath("./env/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetConfigValue(key string) (string, error) {
	value := viper.GetString(key)
	if value == "" {
		return "", fmt.Errorf("config not exist")
	}
	return value, nil
}
