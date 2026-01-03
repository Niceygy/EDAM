package main

import (
	"log"
)

func main() {
	log.Println("Loading...")

	go EDDNCsvLoop(&EDDN_CSV_DATA)
	serve()
}
