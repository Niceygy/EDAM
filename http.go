package main

import (
	"fmt"
	"log"
	"net/http"
)

func serve() {
	// Serve files from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/data/steamcount", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getPlayerCount())
	})

	http.HandleFunc("/data/eddncsv", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "")
	})

	// Start the server on port 8080
	log.Println("Starting server on :3696")
	if err := http.ListenAndServe(":3696", nil); err != nil {
		log.Fatal(err)
	}
}
