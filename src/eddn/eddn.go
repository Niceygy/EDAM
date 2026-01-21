package eddn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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
	if len(UPLOADERS_PAST_HOUR) > 0 {
		//contains something, average what we have

		var total int
		var numEntriesUsed int
		for _, v := range UPLOADERS_PAST_HOUR {
			numEntriesUsed++
			total += v.Messages
		}

		return strconv.Itoa((total / numEntriesUsed) * 60)
	} else if len(UPLOADERS_ALL_TIME) > 0 {
		//get last hour's entry
		return strconv.Itoa(UPLOADERS_ALL_TIME[len(UPLOADERS_ALL_TIME)-1].Messages * 60)
	} else {
		return "0"
	}
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
