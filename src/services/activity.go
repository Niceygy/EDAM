package services

import (
	"math"
	"strconv"

	"github.com/niceygy/edam/eddn"
)

func eddnActivityRating() float64 {
	var largest int = eddn.GetHighestEDDNCount()
	current, _ := strconv.Atoi(eddn.GetCurrentEDDNCount())

	return (float64(current) / float64(largest)) * 100
}

func steamActivityRating() float64 {
	var current int = GetSteamPlayerCount()
	var largest float64 = 25000

	return (float64(current) / largest) * 100
}

func twitchActivityRating() float64 {
	//use log scale

	var current int = GetEliteStreamViewerCount()
	var largest float64 = 10000

	if current <= 0 {
		return 0
	}

	log_min := math.Log(1.0)
	log_max := math.Log(largest)

	log_value := math.Log(float64(current))

	return ((log_value - log_min) / (log_max - log_min)) * 100
}

func OverallActivityRating() float64 {
	if eddn.GetEDStatus() != "Good" {
		return 0
	}

	var eddn float64 = eddnActivityRating()
	var steam float64 = steamActivityRating()
	var twitch float64 = twitchActivityRating()
	return (eddn + steam + twitch) / 3
}
