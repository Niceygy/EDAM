package main

import "strconv"

func eddnActivityRating() float64 {
	var largest int = getHighestEDDNCount()
	current, _ := strconv.Atoi(getCurrentEDDNCount())

	return (float64(current) / float64(largest)) * 100
}

func overallActivityRating() float64 {
	var eddn float64 = eddnActivityRating()
	return eddn
}
