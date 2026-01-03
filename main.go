package main

import (
	"log"
)

func main() {
	log.Println("Loading...")

	TWITCH_ACCESS_TOKEN = getTwitchAccessToken()
	log.Println(TWITCH_ACCESS_TOKEN)
	count := getEliteStreamViewerCount()
	log.Println(count)

	go EDDNCsvLoop(&EDDN_CSV_DATA)
	serve()
}
