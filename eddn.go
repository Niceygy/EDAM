package main

import (
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

func EDDNCsvLoop() {
	for {
		downloadEDDNCsv()
		time.Sleep(time.Hour * 1)
	}
}

func downloadEDDNCsv() {
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

	data, err := io.ReadAll(resp.Body)

	EDDN_CSV_DATA = string(data)
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
