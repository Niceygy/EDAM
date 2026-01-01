package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// headers
		w.Header().Set("X-Powered-By", "NightSpeed Connect (Go)")
		w.Header().Set("X-Created-By", "Niceygy (Ava Whale) - niceygy@niceygy.net")

		h.ServeHTTP(w, r)
	})
}

func serve() {

	// Serve files from the "static" directory
	http.Handle("/", middleware(http.FileServer(http.Dir("static"))))

	http.HandleFunc("/data/steamcount", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getSteamPlayerCount())
	})

	http.HandleFunc("/data/eddncsv", func(w http.ResponseWriter, r *http.Request) {
		data := EDDN_CSV_DATA

		// if err != nil {
		// 	log.Println("ERR Open static/data/messageCount.csv: " + err.Error())
		// }
		fmt.Fprintln(w, string(data))
	})

	http.HandleFunc("/data/eddncount", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getCurrentEDDNCount())
	})

	http.HandleFunc("/data/activityrating", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, overallActivityRating())
	})

	log.Println("Starting server on :3696")
	if err := http.ListenAndServe(":3696", nil); err != nil {
		log.Fatal(err)
	}
}
