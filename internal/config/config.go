package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

// this may be dumb to have a custom object with all the possible fields
// it makes the code that uses config a little cleaner, but probably
// doesnt scale well
type Config struct {
	ClientID     string
	ClientSecret string
	Auth         string
	RedirectURI  string
}

// TODO:
// implement singleton pattern so we arent creating a bunch of config
// objects everywhere, since I will probably end up spinning up a lot of
// go routines
func GetInstance() *Config {
	return newConfig()
}

func newConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/repos/spotify-tags/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &Config{
		ClientID:     viper.GetString("spotify.client.id"),
		ClientSecret: viper.GetString("spotify.client.secret"),
		Auth:         viper.GetString("spotify.user.auth"),
		RedirectURI:  viper.GetString("spotify.client.redirect"),
	}
}
