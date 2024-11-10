package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ApiKey string
}

func NewConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/repos/spotify-tags/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return Config{
		ApiKey: viper.GetString("spotify.apikey"),
	}
}
