package eddn

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"niceygy.net/edam/errors"
)

// float to string
func ftos(f float64) string {
	return strconv.FormatFloat(f, 'g', 5, 64)
}

/*
Returns the highest number of hourly users
ever seen on the EDDN (by the app)
*/
func GetHighestEDDNCount() int {
	highest := 0
	for _, v := range UPLOADERS_ALL_TIME {
		if v.Messages > highest {
			highest = v.Messages
		}
	}

	return highest
}

/*Returns the last hourly count for EDDN*/
func GetCurrentEDDNCount() string {
	res, err := http.Get("https://eddn.edcd.io:4430/stats/")
	errors.PanicIfErr(err)

	// var body []byte

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err.Error())
	}

	var jdata map[string] /*map[any]*/ any

	err = json.Unmarshal(data, &jdata)
	errors.PanicIfErr(err)
	return ftos(jdata["inbound"].(map[string]any)["60min"].(float64))
}

/*Is ED online?*/
func GetEDStatus() EDState {
	req, err := http.NewRequest(http.MethodGet, "https://ed-server-status.orerve.net", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return EDStateOffline
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return EDStateOffline
	}
	defer resp.Body.Close()
	var r EDStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		panic(err)
	}

	if r.Status == "Good" {
		return EDStateOnline
	} else {
		return EDStateOffline
	}
}
