package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Load proxy configuration file from configs folder,
// relative to current working directory.
func Config() (Configurations, error) {
	viper.SetConfigName("proxy")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yml")
	var configuration Configurations
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("[corp] [%s]: Error reading config file: %v\n", time.Now(), err)
		return configuration, err
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("[corp] [%s]: Unable to decode configuration: %v\n", time.Now(), err)
	} else {
		fmt.Printf("[corp] [%s]: Loading configuration from: %s\n", time.Now(), viper.ConfigFileUsed())
	}
	return configuration, err
}
