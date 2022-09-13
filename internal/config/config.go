package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func Config() Configurations {
	viper.SetConfigName("proxy")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("[corp] [%s]: Error reading config file: %s\n", time.Now(), err)
	}
	var configuration Configurations
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("[corp] [%s]: Unable to decode configuration: %v\n", time.Now(), err)
	}
	fmt.Printf("[corp] [%s]: Loading configuration from: %s\n", time.Now(), viper.ConfigFileUsed())
	return configuration
}
