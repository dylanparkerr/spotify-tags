package main

import (
	"fmt"
	config "github.com/dylanparkerr/spotify-tags/internal/config"
)

func main() {
	cfg := config.NewConfig()
	key := cfg.ApiKey
	fmt.Println(key)
}
