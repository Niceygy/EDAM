package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func serve() {
	// Serve files from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/data/steamcount", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getPlayerCount())
	})

	http.HandleFunc("/data/eddncsv", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("static/data/messageCount.csv")

		if err != nil {
			log.Println("ERR Open static/data/messageCount.csv: " + err.Error())
		}
		fmt.Fprintln(w, string(data))
	})

	http.HandleFunc("/data/eddncount", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("static/data/messageCount.csv")

		if err != nil {
			log.Println("ERR Open static/data/messageCount.csv: " + err.Error())
		}

		stringdata := string(data)

		lines := strings.Split(stringdata, "\n")
		line := lines[len(lines)-2]

		fmt.Fprintln(w, strings.Split(line, ",")[1])
	})

	// Start the server on port 8080
	log.Println("Starting server on :3696")
	if err := http.ListenAndServe(":3696", nil); err != nil {
		log.Fatal(err)
	}
}
