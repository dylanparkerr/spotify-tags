package auth

import (
	"encoding/json"
	"fmt"
	config "github.com/dylanparkerr/spotify-tags/internal/config"
	"io"
	"net/http"
	"strings"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessToken() *AccessToken {
	// creds
	cfg := config.NewConfig()
	id := cfg.ClientID
	secret := cfg.ClientSecret

	// form req
	url := "https://accounts.spotify.com/api/token"
	data := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", id, secret)
	creds := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, creds)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// send req
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// convert to struct
	var accessToken AccessToken
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &accessToken
}

func GetPlaylists(token *AccessToken) {
	user := "dylanparkerrr"

	// form req
	url := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", user)
	// url := fmt.Sprintf("https://api.spotify.com/v1/me/playlists")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	bearer := fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", bearer)

	// send req
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	/*
		right now this is giving me a 404

        if i use my old spotify dylanparkerr i get a resp
        maybe need to make my account public? or get proper auth?
	*/

	fmt.Println(string(body))
}
