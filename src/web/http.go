package web

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"niceygy.net/edam/src/eddn"
	"niceygy.net/edam/src/errors"
	"niceygy.net/edam/src/services"
)

var upgrader = websocket.Upgrader{}
var StaticFiles embed.FS

func DoMiddlewareThings(w http.ResponseWriter) {
	w.Header().Set("X-Powered-By", "Go Http.ListenAndServe")
	w.Header().Set("X-Created-By", "Niceygy (Ava Whale) - niceygy@niceygy.net")
	// w.Header().Set("X-Endpoint-Type", "Calculated Live")
}

func Serve() {
	// Serve files from the "static" directory
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		DoMiddlewareThings(w)
		url := "static" + strings.Split(r.RequestURI, "?")[0]
		if strings.HasSuffix(url, "/") {
			url = url + "index.html"
		}
		data, err := StaticFiles.ReadFile(url)
		errors.PanicIfErr(err)
		fmt.Fprintln(w, string(data))
	})

	http.HandleFunc("/data/steamcount", func(w http.ResponseWriter, r *http.Request) {
		DoMiddlewareThings(w)
		fmt.Fprintln(w, services.GetSteamPlayerCount())
	})

	http.HandleFunc("/data/eddncsv", func(w http.ResponseWriter, r *http.Request) {
		data := eddn.CSV_FOR_FTP
		DoMiddlewareThings(w)
		fmt.Fprintln(w, string(data))
	})

	http.HandleFunc("/data/eddncount", func(w http.ResponseWriter, r *http.Request) {
		DoMiddlewareThings(w)
		fmt.Fprintln(w, eddn.GetCurrentEDDNCount())
	})

	http.HandleFunc("/data/activityrating", func(w http.ResponseWriter, r *http.Request) {
		DoMiddlewareThings(w)
		fmt.Fprintln(w, services.OverallActivityRating())
	})

	http.HandleFunc("/data/twitchcount", func(w http.ResponseWriter, r *http.Request) {
		DoMiddlewareThings(w)
		fmt.Fprintln(w, services.GetEliteStreamViewerCount())
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		DoMiddlewareThings(w)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		go func() {
			for {
				messageType := <-eddn.UploaderChannel
				switch messageType {
				case eddn.EDMessage_FSD:
					if err := conn.WriteMessage(websocket.TextMessage, []byte("FSD")); err != nil {
						return
					}
				case eddn.EDMessage_Docked:
					if err := conn.WriteMessage(websocket.TextMessage, []byte("Docked")); err != nil {
						return
					}
				default:
				}

			}
		}()
	})

	log.Println("Started server on :3696")
	if err := http.ListenAndServe(":3696", nil); err != nil {
		log.Fatal(err)
	}
}
