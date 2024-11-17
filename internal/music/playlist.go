package music

import (
	// "encoding/json"
	"bytes"
	"fmt"
	// config "github.com/dylanparkerr/spotify-tags/internal/config"
	"io"
	"net/http"
	"net/url"

	"github.com/dylanparkerr/spotify-tags/internal/auth"
)

func GetPlaylists(token *auth.AccessToken) {
	// lmaooo what is this username? why is it like this?
	// need to look at getting my account off of my edu email and and a real user name
	user := "7h0rgmkyhqtd6arkkcyumc36h"

	// form req
	endpoint := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", user)
	// url := fmt.Sprintf("https://api.spotify.com/v1/me/playlists")
	req, err := http.NewRequest("GET", endpoint, nil)
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

	fmt.Println(string(body))
}

func CreatePlaylist(token *auth.AccessToken) {
	user := "7h0rgmkyhqtd6arkkcyumc36h"
	endpoint := url.URL{
		Scheme:     "https",
		Host:       "api.spotify.com",
		Path:       fmt.Sprintf("v1/users/%s/playlists", user),
		ForceQuery: true,
	}
	// idk if this is the way i want to do this
	data := []byte(`
        {
            "name": "test",
            "description": "test desc",
            "public": false
        }
    `)
	req, err := http.NewRequest("POST", endpoint.String(), bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}

	bearer := fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	// TODO:
	// handle the not 200 case

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	println(string(body))
}

func DeletePlaylist(token *auth.AccessToken) {
	// so turns out the API doesnt let you delete playlists
	// best I can do is probably delete the items and
	// update the playlist info like changing the name and
	// description to something like 'deleted'
	//
	// this is going to pose an interesting problem when I
	// need to check the parity betwen the state of spotify
	// and the db.. probably need to have some message in
	// the description of the playlist that indicates we
	// shouldnt try to keep inventory of it
	//
	// that also makes me think maybe we should have something
	// in the description of every playlist the app makes
}
