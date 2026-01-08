package main

import (
	"math"
	"strconv"
)

func eddnActivityRating() float64 {
	var largest int = getHighestEDDNCount()
	current, _ := strconv.Atoi(getCurrentEDDNCount())

	return (float64(current) / float64(largest)) * 100
}

func steamActivityRating() float64 {
	var current int = getSteamPlayerCount()
	var largest float64 = 25000

	return (float64(current) / largest) * 100
}

func twitchActivityRating() float64 {
	//use log scale

	var current int = getEliteStreamViewerCount()
	var largest float64 = 10000

	if current <= 0 {
		return 0
	}

	log_min := math.Log(1.0)
	log_max := math.Log(largest)

	log_value := math.Log(float64(current))

	return ((log_value - log_min) / (log_max - log_min)) * 100
}

func overallActivityRating() float64 {
	if getEDStatus() != "Good" {
		return 0
	}

	var eddn float64 = eddnActivityRating()
	var steam float64 = steamActivityRating()
	var twitch float64 = twitchActivityRating()
	return (eddn + steam + twitch) / 3
}
