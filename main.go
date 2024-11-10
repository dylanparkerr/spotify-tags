package main

import (
	"fmt"
	config "github.com/dylanparkerr/spotify-tags/internal/config"
)

func main() {
	cfg := config.NewConfig()
	id := cfg.ClientID
	secret := cfg.ClientSecret

	fmt.Printf("id: %s\n", id)
	fmt.Printf("secret:%s\n", secret)
}
