package main

import (
	"niceygy.net/edam/web"

	"log"
)

func main() {
	log.Println("Loading...")
	// go eddn.EDDNListener()
	web.Serve()
}
