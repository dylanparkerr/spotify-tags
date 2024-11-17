package main

import (
	"fmt"

	"github.com/dylanparkerr/spotify-tags/internal/auth"
	"github.com/dylanparkerr/spotify-tags/internal/music"
)

func main() {
	fmt.Println("starting...")

	// auth.PromptAuth()
	// token := auth.GetAccessToken()

	// hydrate and use this until i get db/token refresh working
	token := auth.AccessToken{
		AccessToken: "",
	}
	// music.GetPlaylists(token)
	// fmt.Printf("%+v\n", token)
	auth.PromptAuth()
	music.CreatePlaylist(&token)
}
