package main

import (
	"log"
)

func main() {
	log.Println("Loading...")
	go EDDNCsvLoop()
	serve()
}
