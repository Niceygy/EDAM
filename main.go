package main

import (
	"embed"

	"niceygy.net/edam/src/eddn"
	"niceygy.net/edam/src/web"

	"log"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	log.Println("Loading...")
	web.StaticFiles = staticFiles
	go eddn.EDDN_RecordHourlyCounts()
	web.Serve()
}
