package main

import (
	eddn "github.com/niceygy/edam/eddn"
	"github.com/niceygy/edam/web"

	"log"
)

func main() {
	log.Println("Loading...")
	go eddn.EDDNListener()
	go eddn.EDDNCsvLoop(&eddn.EDDN_CSV_DATA)
	web.Serve()
}
