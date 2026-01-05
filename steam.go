package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Response struct {
		PlayerCount int `json:"player_count"`
		Result      int `json:"result"`
	} `json:"response"`
}

var LAST_FETCHED time.Time = time.Now()
var CACHED_COUNT int = fetchSteamPlayerCount()

func getSteamPlayerCount() int {
	now := time.Now()
	if now.Sub(LAST_FETCHED).Minutes() > 15 {
		CACHED_COUNT = fetchSteamPlayerCount()
		LAST_FETCHED = now
		// log.Println("Updated steam player numbers")
	}

	return CACHED_COUNT
}

func fetchSteamPlayerCount() int {
	req, err := http.NewRequest(http.MethodGet, "https://api.steampowered.com/ISteamUserStats/GetNumberOfCurrentPlayers/v1?appid=359320", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return 0
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return 0
	}
	defer resp.Body.Close()
	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		panic(err)
	}
	return r.Response.PlayerCount
}
