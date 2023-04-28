package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// LoadConfig loads json configuration file
func Load(path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
