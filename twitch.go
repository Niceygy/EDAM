package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

/*{
    "data": [
        {
            "id": "80607",
            "name": "Elite: Dangerous",
            "box_art_url": "https://static-cdn.jtvnw.net/ttv-boxart/80607_IGDB-{width}x{height}.jpg",
            "igdb_id": "2955"
        }
    ]
}*/

const ED_TWITCH_ID string = "80607"

var TWITCH_ACCESS_TOKEN_EXPIRY time.Time = time.Now()

var TWITCH_ACCESS_TOKEN string = ""

type TwitchAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type TwitchStreamResponse struct {
	Data []struct {
		ID           string    `json:"id"`
		UserID       string    `json:"user_id"`
		UserLogin    string    `json:"user_login"`
		UserName     string    `json:"user_name"`
		GameID       string    `json:"game_id"`
		GameName     string    `json:"game_name"`
		Type         string    `json:"type"`
		Title        string    `json:"title"`
		ViewerCount  int       `json:"viewer_count"`
		StartedAt    time.Time `json:"started_at"`
		Language     string    `json:"language"`
		ThumbnailURL string    `json:"thumbnail_url"`
		TagIds       []any     `json:"tag_ids"`
		Tags         []string  `json:"tags"`
		IsMature     bool      `json:"is_mature"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

func getTwitchAccessToken() string {
	urlParts := []string{
		"https://id.twitch.tv/oauth2/token?client_id=",
		os.Getenv("EDAM_TWITCH_CLIENTID"),
		"&client_secret=",
		os.Getenv("EDAM_TWITCH_CLIENTSECRET"),
		"&grant_type=client_credentials",
	}
	var url string = strings.Join(urlParts, "")
	resp, err := http.Post(url, "", nil)

	if err != nil {
		log.Panic(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var data TwitchAccessTokenResponse

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	TWITCH_ACCESS_TOKEN = data.AccessToken
	TWITCH_ACCESS_TOKEN_EXPIRY = time.Now().Add(time.Duration(data.ExpiresIn) * time.Second)

	return data.AccessToken

}

func getEliteStreamViewerCount() int {

	if TWITCH_ACCESS_TOKEN_EXPIRY.Unix() < time.Now().Unix() {
		getTwitchAccessToken()
	}

	urlParts := []string{
		"https://api.twitch.tv/helix/streams?game_id=",
		ED_TWITCH_ID,
	}
	var url string = strings.Join(urlParts, "")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err.Error())
	}
	req.Header.Set("Client-ID", os.Getenv("EDAM_TWITCH_CLIENTID"))
	req.Header.Set("Authorization", "Bearer "+TWITCH_ACCESS_TOKEN)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Panic(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var data TwitchStreamResponse

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	var count int = 0

	for _, entry := range data.Data {
		count += entry.ViewerCount
	}

	return count
}
