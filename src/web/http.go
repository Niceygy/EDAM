package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/niceygy/edam/eddn"
	"github.com/niceygy/edam/services"
)

// var upgrader = websocket.Upgrader{}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// headers
		w.Header().Set("X-Powered-By", "NightSpeed Connect (Go)")
		w.Header().Set("X-Created-By", "Niceygy (Ava Whale) - niceygy@niceygy.net")

		h.ServeHTTP(w, r)
	})
}

func calcEndpointMiddleware(w http.ResponseWriter) {
	w.Header().Set("X-Powered-By", "NightSpeed Connect (Go)")
	w.Header().Set("X-Created-By", "Niceygy (Ava Whale) - niceygy@niceygy.net")
	w.Header().Set("X-Endpoint-Type", "Calculated Live")

}

func Serve() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	} else if strings.Contains(cwd, "src") {
		cwd = strings.Replace(cwd, "src", "", 1)
	}
	// Serve files from the "static" directory
	http.Handle("/", middleware(http.FileServer(http.Dir(cwd+"/static"))))

	http.HandleFunc("/data/steamcount", func(w http.ResponseWriter, r *http.Request) {
		calcEndpointMiddleware(w)
		fmt.Fprintln(w, services.GetSteamPlayerCount())
	})

	http.HandleFunc("/data/eddncsv", func(w http.ResponseWriter, r *http.Request) {
		data := eddn.CSV_FOR_FTP
		calcEndpointMiddleware(w)
		fmt.Fprintln(w, string(data))
	})

	http.HandleFunc("/data/eddncount", func(w http.ResponseWriter, r *http.Request) {
		calcEndpointMiddleware(w)
		fmt.Fprintln(w, eddn.GetCurrentEDDNCount())
	})

	http.HandleFunc("/data/activityrating", func(w http.ResponseWriter, r *http.Request) {
		calcEndpointMiddleware(w)
		fmt.Fprintln(w, services.OverallActivityRating())
	})

	http.HandleFunc("/data/twitchcount", func(w http.ResponseWriter, r *http.Request) {
		calcEndpointMiddleware(w)
		fmt.Fprintln(w, services.GetEliteStreamViewerCount())
	})

	/*http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		calcEndpointMiddleware(w)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		go func() {
			for {
				uploaderID := <-eddn.UploaderChannel
				if err := conn.WriteMessage(websocket.TextMessage, []byte(uploaderID)); err != nil {
					// log.Println(err)
					return
				}
			}
		}()
	})*/

	log.Println("Started server on :3696")
	if err := http.ListenAndServe(":3696", nil); err != nil {
		log.Fatal(err)
	}
}
