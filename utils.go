package main

import "strconv"

func eddnActivityRating() float64 {
	var largest int = getHighestEDDNCount()
	current, _ := strconv.Atoi(getCurrentEDDNCount())

	return (float64(current) / float64(largest)) * 100
}

func steamActivityRating() float64 {
	var current int = getSteamPlayerCount()
	var largest int = 20000

	return (float64(current) / float64(largest)) * 100
}

func overallActivityRating() float64 {
	//20,000

	var eddn float64 = eddnActivityRating()
	var steam float64 = steamActivityRating()
	return (eddn + steam) / 2
}
