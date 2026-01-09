package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const EDDN_CSV_FILEPATH = "static/data/messageCount.csv"

var EDDN_CSV_DATA string

type EDStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Product string `json:"product"`
}

func EDDNCsvLoop(data *string) {
	for {
		downloadEDDNCsv(data)
		time.Sleep(time.Minute * 10)
	}
}

func downloadEDDNCsv(data *string) {
	req, err := http.NewRequest(http.MethodGet, "https://niceygy.net/experiments/edam/data.csv", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("client: bad status code: %d\n", resp.StatusCode)
		return
	}

	_data, err := io.ReadAll(resp.Body)

	*data = string(_data)
}

func getHighestEDDNCount() int {
	stringdata := EDDN_CSV_DATA

	lines := strings.Split(stringdata, "\n")
	largest := 0

	for i := range lines {
		line := lines[i]

		if line == "" {
			break
		}

		count, err := strconv.Atoi(strings.Split(line, ",")[1])
		if err != nil {
			log.Panic(err.Error())
		}

		if count > largest {
			largest = count
		}
	}

	return largest
}

func getCurrentEDDNCount() string {
	stringdata := EDDN_CSV_DATA

	lines := strings.Split(stringdata, "\n")
	line := lines[0]

	return strings.Split(line, ",")[1]
}

func getEDStatus() string {
	req, err := http.NewRequest(http.MethodGet, "https://ed-server-status.orerve.net", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return ""
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return ""
	}
	defer resp.Body.Close()
	var r EDStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		panic(err)
	}
	return r.Status
}
