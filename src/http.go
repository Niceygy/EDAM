package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// headers
		w.Header().Set("X-Powered-By", "NightSpeed Connect (Go)")
		w.Header().Set("X-Created-By", "Niceygy (Ava Whale) - niceygy@niceygy.net")

		h.ServeHTTP(w, r)
	})
}

func serve() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	} else if strings.Contains(cwd, "src") {
		cwd = strings.Replace(cwd, "src", "", 1)
	}
	// Serve files from the "static" directory
	http.Handle("/", middleware(http.FileServer(http.Dir(cwd+"/static"))))

	http.HandleFunc("/data/steamcount", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getSteamPlayerCount())
	})

	http.HandleFunc("/data/eddncsv", func(w http.ResponseWriter, r *http.Request) {
		data := EDDN_CSV_DATA

		fmt.Fprintln(w, string(data))
	})

	http.HandleFunc("/data/eddncount", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getCurrentEDDNCount())
	})

	http.HandleFunc("/data/activityrating", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, overallActivityRating())
	})

	http.HandleFunc("/data/twitchcount", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getEliteStreamViewerCount())
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		go func() {
			for {
				uploaderID := <-uploaderChannel
				if err := conn.WriteMessage(websocket.TextMessage, []byte(uploaderID)); err != nil {
					// log.Println(err)
					return
				}
			}
		}()
	})

	log.Println("Starting server on :3696")
	if err := http.ListenAndServe(":3696", nil); err != nil {
		log.Fatal(err)
	}
}
