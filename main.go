package main

import (
	"log"
)

func main() {
	log.Println("Loading...")
	// go eRelay()
	go EDDNCsvLoop(&EDDN_CSV_DATA)
	serve()
}
