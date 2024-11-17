package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	config "github.com/dylanparkerr/spotify-tags/internal/config"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// TODO:
// maybe create an interface for the differnt ways to auth
func GetAccessToken() *AccessToken {
	// creds
	cfg := config.GetInstance()

	fmt.Printf("%+v\n", cfg)
	id := cfg.ClientID
	secret := cfg.ClientSecret
	auth := cfg.Auth
	redURI := cfg.RedirectURI

	// form req
	endpoint := "https://accounts.spotify.com/api/token"
	bodyParams := url.Values{}
	bodyParams.Add("grant_type", "authorization_code")
	bodyParams.Add("code", auth)
	bodyParams.Add("redirect_uri", redURI)
	bodyData := strings.NewReader(bodyParams.Encode())
	req, err := http.NewRequest("POST", endpoint, bodyData)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	creds := fmt.Sprintf("%s:%s", id, secret)
	encodedCreds := base64.StdEncoding.EncodeToString([]byte(creds))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCreds))

	// send req
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	// TODO:
	// handle the not 200 case

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	println(string(body))

	// convert to struct
	var accessToken AccessToken
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &accessToken
}

func PromptAuth() {
	// TODO:
	// when I get around to making the ui, on start up we should do a check
	// that we still have a valid auth code, and only prompt if its not good

	cfg := config.GetInstance()
	id := cfg.ClientID

	//dont know if I want to call this or pass it in
	//if I want to be extra about it for good measure..
	//put genState in state.go with an interface and
	//another implementation that is static for testing
	state := genState(16)
	scope := "playlist-modify-private"

	url := url.URL{
		Scheme:     "https",
		Host:       "accounts.spotify.com",
		Path:       "authorize",
		ForceQuery: true,
	}
	params := url.Query()
	params.Add("client_id", id)
	params.Add("response_type", "code")
	params.Add("redirect_uri", "http://localhost:8080/callback")
	params.Add("state", state)
	params.Add("scope", scope)
	url.RawQuery = params.Encode()

	// TODO:
	// this does not align with my desire to seperate the core of the app from the ui
	// once i start building the interface, this needs to move there
	// should have the ui take this in and then use it for getAccessToken
	println("-------------------------------------------------------------------------")
	println("Go to this link in a browser to authorize the app:")
	println()
	println(url.String())
	println()
	println("-------------------------------------------------------------------------")
}

// generate a random string of [length] characters long
// encouraged by spotify to use in auth request to prevent
// cross-site request forgery
func genState(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	state := make([]rune, length)
	for i := range state {
		state[i] = letters[rand.Intn(len(letters))]
	}
	return string(state)
}
