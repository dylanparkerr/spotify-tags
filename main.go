package main

import (
	"fmt"
	"github.com/dylanparkerr/spotify-tags/internal/auth"
)

func main() {
	fmt.Println("starting...")

	token := auth.GetAccessToken()
	auth.GetPlaylists(token)
}
