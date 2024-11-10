package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ClientID     string
	ClientSecret string
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
		ClientID:     viper.GetString("spotify.client.id"),
		ClientSecret: viper.GetString("spotify.client.secret"),
	}
}
