package main

import (
	"log"
	"strconv"
)

func main() {
	log.Println("Loading...")
	log.Println("PLayerCount = " + strconv.Itoa(getPlayerCount()))
	serve()
}
