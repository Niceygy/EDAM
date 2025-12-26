package main

import (
	"encoding/binary"
	"log"
	"net/http"
)

func serve() {
	// Serve files from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		count := getPlayerCount()
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, count)
		w.Write()
	})

	// Start the server on port 8080
	log.Println("Starting server on :3696")
	if err := http.ListenAndServe(":3696", nil); err != nil {
		log.Fatal(err)
	}
}
